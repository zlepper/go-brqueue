package brqueue

import (
	"net"
	"sync"
	"sync/atomic"
)

type factoryFunc func() (net.Conn, error)

// A connection pool for reusing of connections
type connectionPool struct {
	nextIndex int32
	// A map of all connection. true if active, false if not
	connections []*connection
}

// Creates a new pool with the specified number of connections
func newConnectionPool(count int, factory factoryFunc) (*connectionPool, error) {
	pool := &connectionPool{connections: make([]*connection, count)}
	var wg sync.WaitGroup
	wg.Add(count)
	var outerError error
	for i := 0; i < count; i++ {
		index := i
		go func() {
			conn, err := factory()
			if err != nil {
				outerError = err
				return
			}
			pool.connections[index] = &connection{
				conn:conn,
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if outerError != nil {
		pool.close()
		return nil, outerError
	}

	return pool, nil
}

// Closes all the connections. Returns the first error encountered
func (p *connectionPool) close() error {
	var erro error
	for _, conn := range p.connections {
		if conn != nil {
			err := conn.close()
			if err != nil && erro != nil {
				erro = err
			}
		}
	}
	return erro
}

func (p *connectionPool) get() *connection {

	index := atomic.AddInt32(&p.nextIndex, 1) % int32(len(p.connections))
	return p.connections[index]

}
