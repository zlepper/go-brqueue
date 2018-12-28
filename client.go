package brqueue

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/zlepper/go-brqueue/internal"
	"net"
	"sync"
)

type Priority int

const (
	HighPriority Priority = iota
	LowPriority  Priority = iota
)

var (
	ErrInvalidPriority = errors.New("invalid priority")
	ErrInvalidResponse = errors.New("invalid response")
	ErrInvalidTask     = errors.New("invalid task")
)

// A raw wrapper around the TCP connection to the brqueue server.
// You probably want to add another layer, so you aren't dealing with raw byte arrays.
// This client is safe to concurrent usage.
type Client struct {
	// The underlying TCP connection
	pool *connectionPool
	// Ref id generation
	nextRefId      int32
	nextRefIdMutex sync.Mutex
	callbacks      map[int32]chan *internal.ResponseWrapper
	callbacksMutex sync.Mutex
}

// Creates a new client and connects to the given hostname:port
// An error is returned if the TCP connection cannot be created
func NewClient(hostname string, port int) (*Client, error) {
	pool, err := newConnectionPool(12, func() (net.Conn, error) {
		return net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port))
	})
	if err != nil {
		return nil, err
	}

	c := &Client{
		pool:		pool,
		callbacks: make(map[int32]chan *internal.ResponseWrapper),
	}
	return c, nil
}

// Closes the underlying socket
func (c *Client) Close() error {
	c.callbacksMutex.Lock()
	for _, channel := range c.callbacks {
		close(channel)
	}
	c.callbacksMutex.Unlock()

	return c.pool.close()
}

func (c *Client) getNextRefId() int32 {
	c.nextRefIdMutex.Lock()
	defer c.nextRefIdMutex.Unlock()
	c.nextRefId += 1
	return c.nextRefId
}

// Actually sends the message across the wire
func (c *Client) sendMessage(message *internal.RequestWrapper, connection *connection) (refId int32, err error) {
	refId = c.getNextRefId()
	message.RefId = refId

	data, err := proto.Marshal(message)
	if err != nil {
		return 0, err
	}

	size := intToByteArray(int32(len(data)))

	_, err = connection.conn.Write(append(size, data...))
	if err != nil {
		return 0, err
	}

	return refId, nil
}

// Reads the next message of the socket
func (c *Client) readNextMessage(connection *connection) (*internal.ResponseWrapper, error) {
	conn := connection.conn
	connection.readLock.Lock()
	defer connection.readLock.Unlock()
	sizeBytes := make([]byte, 4)
	_, err := conn.Read(sizeBytes)
	if err != nil {
		return nil, err
	}

	data := make([]byte, byteArrayToInt(sizeBytes))

	_, err = conn.Read(data)
	if err != nil {
		return nil, err
	}

	var wrapper internal.ResponseWrapper
	err = proto.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return &wrapper, nil
}

// Waits for a response for the given refId
func (c *Client) waitForResponse(refId int32, connection *connection) (*internal.ResponseWrapper, error) {
	c.callbacksMutex.Lock()
	res, ok := c.callbacks[refId]
	if !ok {
		res = make(chan *internal.ResponseWrapper, 1)
		c.callbacks[refId] = res
	}
	c.callbacksMutex.Unlock()
	defer func() {
		c.callbacksMutex.Lock()
		delete(c.callbacks, refId)
		c.callbacksMutex.Unlock()
	}()

	message, err := c.readNextMessage(connection)
	if err != nil {
		return nil, err
	}

	if message.RefId == refId {
		errorMessage := message.GetError()
		if errorMessage != nil {
			return nil, errors.New("response error: " + errorMessage.Message)
		}

		return message, nil
	}
	// Get the actual callback waiting for a message.
	c.callbacksMutex.Lock()
	cha, ok := c.callbacks[message.RefId]
	// If there is no such callback, create it, and pass the message to that listener instead
	if !ok {
		cha = make(chan *internal.ResponseWrapper, 1)
		c.callbacks[message.RefId] = cha
	}
	cha <- message
	c.callbacksMutex.Unlock()


	// Wait for a message to appear for us
	message = <-res
	errorMessage := message.GetError()
	if errorMessage != nil {
		return nil, errors.New("response error: " + errorMessage.Message)
	}
	return message, nil
}

// executes a request, and waits for a response
func (c *Client) executeRequest(message *internal.RequestWrapper) (*internal.ResponseWrapper, error) {
	connection := c.pool.get()
	refId, err := c.sendMessage(message, connection)
	if err != nil {
		return nil, err
	}

	return c.waitForResponse(refId, connection)
}

// Enqueues a new message
func (c *Client) EnqueueRequest(message []byte, priority Priority, requiredCapabilities []string) (string, error) {
	if priority != HighPriority && priority != LowPriority {
		return "", ErrInvalidPriority
	}

	var prio internal.Priority

	if priority == HighPriority {
		prio = internal.Priority_HIGH
	} else {
		prio = internal.Priority_LOW
	}

	requestWrapper := &internal.RequestWrapper{
		Message: &internal.RequestWrapper_Enqueue{
			Enqueue: &internal.EnqueueRequest{
				Message:              message,
				Priority:             prio,
				RequiredCapabilities: requiredCapabilities,
			},
		},
	}

	responseWrapper, err := c.executeRequest(requestWrapper)
	if err != nil {
		return "", err
	}

	response := responseWrapper.GetEnqueue()
	if response == nil {
		return "", ErrInvalidPriority
	}
	return response.Id, nil
}

// Pops a single message off the queue
func (c *Client) Pop(availableCapabilities []string, waitForMessages bool) (*WorkTask, error) {
	requestWrapper := &internal.RequestWrapper{
		Message: &internal.RequestWrapper_Pop{
			Pop: &internal.PopRequest{
				AvailableCapabilities: availableCapabilities,
				WaitForMessage:        waitForMessages,
			},
		},
	}

	responseWrapper, err := c.executeRequest(requestWrapper)
	if err != nil {
		return nil, err
	}
	response := responseWrapper.GetPop()
	if response == nil {
		return nil, ErrInvalidResponse
	}
	if response.HadResult {
		return newWorkTask(response.Id, response.Message), nil
	}
	return nil, nil
}

// Acknowledges the completion of the given task
func (c *Client) Acknowledge(task *WorkTask) error {
	if !task.valid {
		return ErrInvalidTask
	}

	requestWrapper := &internal.RequestWrapper{
		Message: &internal.RequestWrapper_Acknowledge{
			Acknowledge: &internal.AcknowledgeRequest{
				Id: task.id,
			},
		},
	}

	responseWrapper, err := c.executeRequest(requestWrapper)
	if err != nil {
		return err
	}
	response := responseWrapper.GetAcknowledge()
	if response == nil {
		return ErrInvalidResponse
	}
	return nil
}
