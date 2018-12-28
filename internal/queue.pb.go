// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Priority int32

const (
	Priority_LOW  Priority = 0
	Priority_HIGH Priority = 1
)

var Priority_name = map[int32]string{
	0: "LOW",
	1: "HIGH",
}
var Priority_value = map[string]int32{
	"LOW":  0,
	"HIGH": 1,
}

func (x Priority) String() string {
	return proto.EnumName(Priority_name, int32(x))
}
func (Priority) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{0}
}

// Puts a new message in the queue
type EnqueueRequest struct {
	// The actual task to enqueue
	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The priority of the task
	Priority Priority `protobuf:"varint,2,opt,name=priority,proto3,enum=Priority" json:"priority,omitempty"`
	// What capabilities are required to handle the task
	RequiredCapabilities []string `protobuf:"bytes,3,rep,name=requiredCapabilities,proto3" json:"requiredCapabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnqueueRequest) Reset()         { *m = EnqueueRequest{} }
func (m *EnqueueRequest) String() string { return proto.CompactTextString(m) }
func (*EnqueueRequest) ProtoMessage()    {}
func (*EnqueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{0}
}
func (m *EnqueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnqueueRequest.Unmarshal(m, b)
}
func (m *EnqueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnqueueRequest.Marshal(b, m, deterministic)
}
func (dst *EnqueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnqueueRequest.Merge(dst, src)
}
func (m *EnqueueRequest) XXX_Size() int {
	return xxx_messageInfo_EnqueueRequest.Size(m)
}
func (m *EnqueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnqueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnqueueRequest proto.InternalMessageInfo

func (m *EnqueueRequest) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *EnqueueRequest) GetPriority() Priority {
	if m != nil {
		return m.Priority
	}
	return Priority_LOW
}

func (m *EnqueueRequest) GetRequiredCapabilities() []string {
	if m != nil {
		return m.RequiredCapabilities
	}
	return nil
}

type EnqueueResponse struct {
	// The id of the created task
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnqueueResponse) Reset()         { *m = EnqueueResponse{} }
func (m *EnqueueResponse) String() string { return proto.CompactTextString(m) }
func (*EnqueueResponse) ProtoMessage()    {}
func (*EnqueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{1}
}
func (m *EnqueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnqueueResponse.Unmarshal(m, b)
}
func (m *EnqueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnqueueResponse.Marshal(b, m, deterministic)
}
func (dst *EnqueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnqueueResponse.Merge(dst, src)
}
func (m *EnqueueResponse) XXX_Size() int {
	return xxx_messageInfo_EnqueueResponse.Size(m)
}
func (m *EnqueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnqueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnqueueResponse proto.InternalMessageInfo

func (m *EnqueueResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Removes a single message from the queue
type PopRequest struct {
	// What capabilities the client has available
	AvailableCapabilities []string `protobuf:"bytes,1,rep,name=availableCapabilities,proto3" json:"availableCapabilities,omitempty"`
	// If the server should wait for a message to come in, before returning
	// with a response
	WaitForMessage       bool     `protobuf:"varint,2,opt,name=waitForMessage,proto3" json:"waitForMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopRequest) Reset()         { *m = PopRequest{} }
func (m *PopRequest) String() string { return proto.CompactTextString(m) }
func (*PopRequest) ProtoMessage()    {}
func (*PopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{2}
}
func (m *PopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopRequest.Unmarshal(m, b)
}
func (m *PopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopRequest.Marshal(b, m, deterministic)
}
func (dst *PopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopRequest.Merge(dst, src)
}
func (m *PopRequest) XXX_Size() int {
	return xxx_messageInfo_PopRequest.Size(m)
}
func (m *PopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PopRequest proto.InternalMessageInfo

func (m *PopRequest) GetAvailableCapabilities() []string {
	if m != nil {
		return m.AvailableCapabilities
	}
	return nil
}

func (m *PopRequest) GetWaitForMessage() bool {
	if m != nil {
		return m.WaitForMessage
	}
	return false
}

type PopResponse struct {
	// True if there was a message available
	HadResult bool `protobuf:"varint,3,opt,name=hadResult,proto3" json:"hadResult,omitempty"`
	// The actual message to process
	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The id of the message
	// Should be returned with the acknowledge request
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopResponse) Reset()         { *m = PopResponse{} }
func (m *PopResponse) String() string { return proto.CompactTextString(m) }
func (*PopResponse) ProtoMessage()    {}
func (*PopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{3}
}
func (m *PopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopResponse.Unmarshal(m, b)
}
func (m *PopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopResponse.Marshal(b, m, deterministic)
}
func (dst *PopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopResponse.Merge(dst, src)
}
func (m *PopResponse) XXX_Size() int {
	return xxx_messageInfo_PopResponse.Size(m)
}
func (m *PopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PopResponse proto.InternalMessageInfo

func (m *PopResponse) GetHadResult() bool {
	if m != nil {
		return m.HadResult
	}
	return false
}

func (m *PopResponse) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AcknowledgeRequest struct {
	// The id of the message to acknowledge
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcknowledgeRequest) Reset()         { *m = AcknowledgeRequest{} }
func (m *AcknowledgeRequest) String() string { return proto.CompactTextString(m) }
func (*AcknowledgeRequest) ProtoMessage()    {}
func (*AcknowledgeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{4}
}
func (m *AcknowledgeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcknowledgeRequest.Unmarshal(m, b)
}
func (m *AcknowledgeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcknowledgeRequest.Marshal(b, m, deterministic)
}
func (dst *AcknowledgeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcknowledgeRequest.Merge(dst, src)
}
func (m *AcknowledgeRequest) XXX_Size() int {
	return xxx_messageInfo_AcknowledgeRequest.Size(m)
}
func (m *AcknowledgeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AcknowledgeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AcknowledgeRequest proto.InternalMessageInfo

func (m *AcknowledgeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Returned when a message has been acknowledged
type AcknowledgeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcknowledgeResponse) Reset()         { *m = AcknowledgeResponse{} }
func (m *AcknowledgeResponse) String() string { return proto.CompactTextString(m) }
func (*AcknowledgeResponse) ProtoMessage()    {}
func (*AcknowledgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{5}
}
func (m *AcknowledgeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcknowledgeResponse.Unmarshal(m, b)
}
func (m *AcknowledgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcknowledgeResponse.Marshal(b, m, deterministic)
}
func (dst *AcknowledgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcknowledgeResponse.Merge(dst, src)
}
func (m *AcknowledgeResponse) XXX_Size() int {
	return xxx_messageInfo_AcknowledgeResponse.Size(m)
}
func (m *AcknowledgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AcknowledgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AcknowledgeResponse proto.InternalMessageInfo

// Write the full current queue back to the tpc stream
// As the response is not instant, it doesn't gaurantee that some
// of the messages hasn't been completed in the meantime, or
// that new messages has arrived
type ListFullQueueRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFullQueueRequest) Reset()         { *m = ListFullQueueRequest{} }
func (m *ListFullQueueRequest) String() string { return proto.CompactTextString(m) }
func (*ListFullQueueRequest) ProtoMessage()    {}
func (*ListFullQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{6}
}
func (m *ListFullQueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFullQueueRequest.Unmarshal(m, b)
}
func (m *ListFullQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFullQueueRequest.Marshal(b, m, deterministic)
}
func (dst *ListFullQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFullQueueRequest.Merge(dst, src)
}
func (m *ListFullQueueRequest) XXX_Size() int {
	return xxx_messageInfo_ListFullQueueRequest.Size(m)
}
func (m *ListFullQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFullQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFullQueueRequest proto.InternalMessageInfo

// A partial response about a single message in the queue
type ListFullQueueResponse struct {
	// The index of this message, useful if the client wants to order messages
	SegmentIndex int32 `protobuf:"varint,1,opt,name=segmentIndex,proto3" json:"segmentIndex,omitempty"`
	// The actual message
	Message []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// The id of the message
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// If the listing has finished. When this is received, the
	// rest of the fields doesn't make sense
	Finished             bool     `protobuf:"varint,5,opt,name=finished,proto3" json:"finished,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFullQueueResponse) Reset()         { *m = ListFullQueueResponse{} }
func (m *ListFullQueueResponse) String() string { return proto.CompactTextString(m) }
func (*ListFullQueueResponse) ProtoMessage()    {}
func (*ListFullQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{7}
}
func (m *ListFullQueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFullQueueResponse.Unmarshal(m, b)
}
func (m *ListFullQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFullQueueResponse.Marshal(b, m, deterministic)
}
func (dst *ListFullQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFullQueueResponse.Merge(dst, src)
}
func (m *ListFullQueueResponse) XXX_Size() int {
	return xxx_messageInfo_ListFullQueueResponse.Size(m)
}
func (m *ListFullQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFullQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFullQueueResponse proto.InternalMessageInfo

func (m *ListFullQueueResponse) GetSegmentIndex() int32 {
	if m != nil {
		return m.SegmentIndex
	}
	return 0
}

func (m *ListFullQueueResponse) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ListFullQueueResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListFullQueueResponse) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

type ErrorResponse struct {
	// What went wrong
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{8}
}
func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (dst *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(dst, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Wraps the requests for easier parsing on the socket layer
type RequestWrapper struct {
	RefId int32 `protobuf:"varint,10,opt,name=refId,proto3" json:"refId,omitempty"`
	// Types that are valid to be assigned to Message:
	//	*RequestWrapper_Enqueue
	//	*RequestWrapper_Pop
	//	*RequestWrapper_Acknowledge
	//	*RequestWrapper_ListFullQueue
	Message              isRequestWrapper_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RequestWrapper) Reset()         { *m = RequestWrapper{} }
func (m *RequestWrapper) String() string { return proto.CompactTextString(m) }
func (*RequestWrapper) ProtoMessage()    {}
func (*RequestWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{9}
}
func (m *RequestWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestWrapper.Unmarshal(m, b)
}
func (m *RequestWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestWrapper.Marshal(b, m, deterministic)
}
func (dst *RequestWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestWrapper.Merge(dst, src)
}
func (m *RequestWrapper) XXX_Size() int {
	return xxx_messageInfo_RequestWrapper.Size(m)
}
func (m *RequestWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_RequestWrapper proto.InternalMessageInfo

type isRequestWrapper_Message interface {
	isRequestWrapper_Message()
}

type RequestWrapper_Enqueue struct {
	Enqueue *EnqueueRequest `protobuf:"bytes,1,opt,name=enqueue,proto3,oneof"`
}
type RequestWrapper_Pop struct {
	Pop *PopRequest `protobuf:"bytes,2,opt,name=pop,proto3,oneof"`
}
type RequestWrapper_Acknowledge struct {
	Acknowledge *AcknowledgeRequest `protobuf:"bytes,3,opt,name=acknowledge,proto3,oneof"`
}
type RequestWrapper_ListFullQueue struct {
	ListFullQueue *ListFullQueueRequest `protobuf:"bytes,4,opt,name=listFullQueue,proto3,oneof"`
}

func (*RequestWrapper_Enqueue) isRequestWrapper_Message()       {}
func (*RequestWrapper_Pop) isRequestWrapper_Message()           {}
func (*RequestWrapper_Acknowledge) isRequestWrapper_Message()   {}
func (*RequestWrapper_ListFullQueue) isRequestWrapper_Message() {}

func (m *RequestWrapper) GetMessage() isRequestWrapper_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RequestWrapper) GetRefId() int32 {
	if m != nil {
		return m.RefId
	}
	return 0
}

func (m *RequestWrapper) GetEnqueue() *EnqueueRequest {
	if x, ok := m.GetMessage().(*RequestWrapper_Enqueue); ok {
		return x.Enqueue
	}
	return nil
}

func (m *RequestWrapper) GetPop() *PopRequest {
	if x, ok := m.GetMessage().(*RequestWrapper_Pop); ok {
		return x.Pop
	}
	return nil
}

func (m *RequestWrapper) GetAcknowledge() *AcknowledgeRequest {
	if x, ok := m.GetMessage().(*RequestWrapper_Acknowledge); ok {
		return x.Acknowledge
	}
	return nil
}

func (m *RequestWrapper) GetListFullQueue() *ListFullQueueRequest {
	if x, ok := m.GetMessage().(*RequestWrapper_ListFullQueue); ok {
		return x.ListFullQueue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RequestWrapper) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RequestWrapper_OneofMarshaler, _RequestWrapper_OneofUnmarshaler, _RequestWrapper_OneofSizer, []interface{}{
		(*RequestWrapper_Enqueue)(nil),
		(*RequestWrapper_Pop)(nil),
		(*RequestWrapper_Acknowledge)(nil),
		(*RequestWrapper_ListFullQueue)(nil),
	}
}

func _RequestWrapper_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RequestWrapper)
	// message
	switch x := m.Message.(type) {
	case *RequestWrapper_Enqueue:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Enqueue); err != nil {
			return err
		}
	case *RequestWrapper_Pop:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pop); err != nil {
			return err
		}
	case *RequestWrapper_Acknowledge:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Acknowledge); err != nil {
			return err
		}
	case *RequestWrapper_ListFullQueue:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListFullQueue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RequestWrapper.Message has unexpected type %T", x)
	}
	return nil
}

func _RequestWrapper_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RequestWrapper)
	switch tag {
	case 1: // message.enqueue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EnqueueRequest)
		err := b.DecodeMessage(msg)
		m.Message = &RequestWrapper_Enqueue{msg}
		return true, err
	case 2: // message.pop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PopRequest)
		err := b.DecodeMessage(msg)
		m.Message = &RequestWrapper_Pop{msg}
		return true, err
	case 3: // message.acknowledge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AcknowledgeRequest)
		err := b.DecodeMessage(msg)
		m.Message = &RequestWrapper_Acknowledge{msg}
		return true, err
	case 4: // message.listFullQueue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListFullQueueRequest)
		err := b.DecodeMessage(msg)
		m.Message = &RequestWrapper_ListFullQueue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RequestWrapper_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RequestWrapper)
	// message
	switch x := m.Message.(type) {
	case *RequestWrapper_Enqueue:
		s := proto.Size(x.Enqueue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestWrapper_Pop:
		s := proto.Size(x.Pop)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestWrapper_Acknowledge:
		s := proto.Size(x.Acknowledge)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestWrapper_ListFullQueue:
		s := proto.Size(x.ListFullQueue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ResponseWrapper struct {
	RefId int32 `protobuf:"varint,10,opt,name=refId,proto3" json:"refId,omitempty"`
	// Types that are valid to be assigned to Message:
	//	*ResponseWrapper_Enqueue
	//	*ResponseWrapper_Pop
	//	*ResponseWrapper_Acknowledge
	//	*ResponseWrapper_Error
	//	*ResponseWrapper_ListFullQueue
	Message              isResponseWrapper_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ResponseWrapper) Reset()         { *m = ResponseWrapper{} }
func (m *ResponseWrapper) String() string { return proto.CompactTextString(m) }
func (*ResponseWrapper) ProtoMessage()    {}
func (*ResponseWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_queue_919abe04e1fc201f, []int{10}
}
func (m *ResponseWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseWrapper.Unmarshal(m, b)
}
func (m *ResponseWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseWrapper.Marshal(b, m, deterministic)
}
func (dst *ResponseWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseWrapper.Merge(dst, src)
}
func (m *ResponseWrapper) XXX_Size() int {
	return xxx_messageInfo_ResponseWrapper.Size(m)
}
func (m *ResponseWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseWrapper proto.InternalMessageInfo

type isResponseWrapper_Message interface {
	isResponseWrapper_Message()
}

type ResponseWrapper_Enqueue struct {
	Enqueue *EnqueueResponse `protobuf:"bytes,1,opt,name=enqueue,proto3,oneof"`
}
type ResponseWrapper_Pop struct {
	Pop *PopResponse `protobuf:"bytes,2,opt,name=pop,proto3,oneof"`
}
type ResponseWrapper_Acknowledge struct {
	Acknowledge *AcknowledgeResponse `protobuf:"bytes,3,opt,name=acknowledge,proto3,oneof"`
}
type ResponseWrapper_Error struct {
	Error *ErrorResponse `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}
type ResponseWrapper_ListFullQueue struct {
	ListFullQueue *ListFullQueueResponse `protobuf:"bytes,5,opt,name=listFullQueue,proto3,oneof"`
}

func (*ResponseWrapper_Enqueue) isResponseWrapper_Message()       {}
func (*ResponseWrapper_Pop) isResponseWrapper_Message()           {}
func (*ResponseWrapper_Acknowledge) isResponseWrapper_Message()   {}
func (*ResponseWrapper_Error) isResponseWrapper_Message()         {}
func (*ResponseWrapper_ListFullQueue) isResponseWrapper_Message() {}

func (m *ResponseWrapper) GetMessage() isResponseWrapper_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ResponseWrapper) GetRefId() int32 {
	if m != nil {
		return m.RefId
	}
	return 0
}

func (m *ResponseWrapper) GetEnqueue() *EnqueueResponse {
	if x, ok := m.GetMessage().(*ResponseWrapper_Enqueue); ok {
		return x.Enqueue
	}
	return nil
}

func (m *ResponseWrapper) GetPop() *PopResponse {
	if x, ok := m.GetMessage().(*ResponseWrapper_Pop); ok {
		return x.Pop
	}
	return nil
}

func (m *ResponseWrapper) GetAcknowledge() *AcknowledgeResponse {
	if x, ok := m.GetMessage().(*ResponseWrapper_Acknowledge); ok {
		return x.Acknowledge
	}
	return nil
}

func (m *ResponseWrapper) GetError() *ErrorResponse {
	if x, ok := m.GetMessage().(*ResponseWrapper_Error); ok {
		return x.Error
	}
	return nil
}

func (m *ResponseWrapper) GetListFullQueue() *ListFullQueueResponse {
	if x, ok := m.GetMessage().(*ResponseWrapper_ListFullQueue); ok {
		return x.ListFullQueue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ResponseWrapper) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ResponseWrapper_OneofMarshaler, _ResponseWrapper_OneofUnmarshaler, _ResponseWrapper_OneofSizer, []interface{}{
		(*ResponseWrapper_Enqueue)(nil),
		(*ResponseWrapper_Pop)(nil),
		(*ResponseWrapper_Acknowledge)(nil),
		(*ResponseWrapper_Error)(nil),
		(*ResponseWrapper_ListFullQueue)(nil),
	}
}

func _ResponseWrapper_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ResponseWrapper)
	// message
	switch x := m.Message.(type) {
	case *ResponseWrapper_Enqueue:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Enqueue); err != nil {
			return err
		}
	case *ResponseWrapper_Pop:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pop); err != nil {
			return err
		}
	case *ResponseWrapper_Acknowledge:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Acknowledge); err != nil {
			return err
		}
	case *ResponseWrapper_Error:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *ResponseWrapper_ListFullQueue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListFullQueue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ResponseWrapper.Message has unexpected type %T", x)
	}
	return nil
}

func _ResponseWrapper_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ResponseWrapper)
	switch tag {
	case 1: // message.enqueue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EnqueueResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ResponseWrapper_Enqueue{msg}
		return true, err
	case 2: // message.pop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PopResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ResponseWrapper_Pop{msg}
		return true, err
	case 3: // message.acknowledge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AcknowledgeResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ResponseWrapper_Acknowledge{msg}
		return true, err
	case 4: // message.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ErrorResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ResponseWrapper_Error{msg}
		return true, err
	case 5: // message.listFullQueue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListFullQueueResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ResponseWrapper_ListFullQueue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ResponseWrapper_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ResponseWrapper)
	// message
	switch x := m.Message.(type) {
	case *ResponseWrapper_Enqueue:
		s := proto.Size(x.Enqueue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResponseWrapper_Pop:
		s := proto.Size(x.Pop)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResponseWrapper_Acknowledge:
		s := proto.Size(x.Acknowledge)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResponseWrapper_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResponseWrapper_ListFullQueue:
		s := proto.Size(x.ListFullQueue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*EnqueueRequest)(nil), "EnqueueRequest")
	proto.RegisterType((*EnqueueResponse)(nil), "EnqueueResponse")
	proto.RegisterType((*PopRequest)(nil), "PopRequest")
	proto.RegisterType((*PopResponse)(nil), "PopResponse")
	proto.RegisterType((*AcknowledgeRequest)(nil), "AcknowledgeRequest")
	proto.RegisterType((*AcknowledgeResponse)(nil), "AcknowledgeResponse")
	proto.RegisterType((*ListFullQueueRequest)(nil), "ListFullQueueRequest")
	proto.RegisterType((*ListFullQueueResponse)(nil), "ListFullQueueResponse")
	proto.RegisterType((*ErrorResponse)(nil), "ErrorResponse")
	proto.RegisterType((*RequestWrapper)(nil), "RequestWrapper")
	proto.RegisterType((*ResponseWrapper)(nil), "ResponseWrapper")
	proto.RegisterEnum("Priority", Priority_name, Priority_value)
}

func init() { proto.RegisterFile("queue.proto", fileDescriptor_queue_919abe04e1fc201f) }

var fileDescriptor_queue_919abe04e1fc201f = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0xfd, 0xd1, 0x90, 0x78, 0x9c, 0x3a, 0xd1, 0x26, 0xa9, 0x2c, 0x04, 0x22, 0x58, 0x50,
	0x85, 0x0f, 0xf9, 0x60, 0x90, 0xe0, 0x02, 0x12, 0x45, 0x2d, 0x8e, 0x54, 0x44, 0x59, 0x09, 0x55,
	0xe2, 0xe6, 0xe0, 0x69, 0xba, 0xe0, 0xda, 0xce, 0xda, 0xa6, 0x70, 0xe4, 0xc0, 0x3f, 0xe1, 0xc4,
	0x1f, 0xe4, 0x8a, 0xb2, 0xb1, 0x1d, 0xdb, 0xb5, 0x72, 0xcb, 0xce, 0x4c, 0x76, 0xde, 0xe7, 0x9d,
	0xf1, 0x82, 0xbe, 0xca, 0x30, 0x43, 0x3b, 0xe6, 0x51, 0x1a, 0x59, 0xbf, 0x65, 0x30, 0x8e, 0x43,
	0x11, 0xa1, 0xb8, 0xca, 0x30, 0x49, 0x89, 0x09, 0xdd, 0x2b, 0x4c, 0x12, 0x6f, 0x89, 0xa6, 0x3c,
	0x95, 0x67, 0x7d, 0x5a, 0x1c, 0xc9, 0x43, 0xe8, 0xc5, 0x9c, 0x45, 0x9c, 0xa5, 0x3f, 0x4d, 0x65,
	0x2a, 0xcf, 0x0c, 0x47, 0xb3, 0xcf, 0xf2, 0x00, 0x2d, 0x53, 0xc4, 0x81, 0x31, 0xc7, 0x55, 0xc6,
	0x38, 0xfa, 0x6f, 0xbd, 0xd8, 0x5b, 0xb0, 0x80, 0xa5, 0x0c, 0x13, 0x53, 0x9d, 0xaa, 0x33, 0x8d,
	0xb6, 0xe6, 0xac, 0xfb, 0x30, 0x28, 0x65, 0x24, 0x71, 0x14, 0x26, 0x48, 0x0c, 0x50, 0x98, 0x2f,
	0x24, 0x68, 0x54, 0x61, 0xbe, 0xf5, 0x15, 0xe0, 0x2c, 0x8a, 0x0b, 0x95, 0xcf, 0x61, 0xe2, 0x7d,
	0xf7, 0x58, 0xe0, 0x2d, 0x02, 0xac, 0x75, 0x91, 0x45, 0x97, 0xf6, 0x24, 0x39, 0x04, 0xe3, 0xda,
	0x63, 0xe9, 0x49, 0xc4, 0xdf, 0xe7, 0x88, 0x6b, 0x8e, 0x1e, 0x6d, 0x44, 0xad, 0x4f, 0xa0, 0x8b,
	0x5e, 0xb9, 0x94, 0x3b, 0xa0, 0x5d, 0x7a, 0x3e, 0xc5, 0x24, 0x0b, 0x52, 0x53, 0x15, 0xff, 0xd8,
	0x06, 0x76, 0x18, 0xb6, 0x41, 0x50, 0x4a, 0x84, 0x07, 0x40, 0xde, 0x7c, 0xf9, 0x16, 0x46, 0xd7,
	0x01, 0xfa, 0xcb, 0xd2, 0xf0, 0x26, 0xe8, 0x04, 0x46, 0xb5, 0xaa, 0x8d, 0x08, 0xeb, 0x00, 0xc6,
	0xa7, 0x2c, 0x49, 0x4f, 0xb2, 0x20, 0xf8, 0x58, 0x99, 0x97, 0xf5, 0x4b, 0x86, 0x49, 0x23, 0x91,
	0xcb, 0xb6, 0xa0, 0x9f, 0xe0, 0xf2, 0x0a, 0xc3, 0x74, 0x1e, 0xfa, 0xf8, 0x43, 0xb4, 0xe8, 0xd0,
	0x5a, 0xac, 0x2a, 0x5e, 0x6d, 0x13, 0xbf, 0x57, 0xc8, 0x22, 0xb7, 0xa1, 0x77, 0xc1, 0x42, 0x96,
	0x5c, 0xa2, 0x6f, 0x76, 0x84, 0x07, 0xe5, 0xd9, 0x7a, 0x04, 0xfb, 0xc7, 0x9c, 0x47, 0xbc, 0x6c,
	0xdd, 0xf0, 0x44, 0x2b, 0xaf, 0xb5, 0xfe, 0xc9, 0x60, 0xe4, 0xd2, 0xcf, 0xb9, 0x17, 0xc7, 0xc8,
	0xc9, 0x18, 0x3a, 0x1c, 0x2f, 0xe6, 0xbe, 0x09, 0x42, 0xe0, 0xe6, 0x40, 0x9e, 0x40, 0x17, 0x37,
	0x2b, 0x21, 0xae, 0xd0, 0x9d, 0x81, 0x5d, 0xdf, 0x54, 0x57, 0xa2, 0x45, 0x05, 0xb9, 0x07, 0x6a,
	0x1c, 0xc5, 0xc2, 0x6a, 0xdd, 0xd1, 0xed, 0xed, 0xa2, 0xb8, 0x12, 0x5d, 0x67, 0xc8, 0x0b, 0xd0,
	0xbd, 0xad, 0xa9, 0x82, 0x55, 0x77, 0x46, 0xf6, 0xcd, 0x71, 0xb8, 0x12, 0xad, 0x56, 0x92, 0x57,
	0xb0, 0x1f, 0x54, 0xdd, 0x15, 0x8e, 0xe8, 0xce, 0xc4, 0x6e, 0x1b, 0x86, 0x2b, 0xd1, 0x7a, 0xf5,
	0x91, 0x56, 0x1a, 0x61, 0xfd, 0x51, 0x60, 0x50, 0x18, 0xb4, 0x1b, 0xfd, 0x69, 0x13, 0x7d, 0x68,
	0x37, 0xbe, 0x8e, 0x2a, 0xfb, 0xb4, 0xca, 0xde, 0xb7, 0x2b, 0x8b, 0x5b, 0xc0, 0xbf, 0x6c, 0x83,
	0x1f, 0xdb, 0x2d, 0x5b, 0xd6, 0xa4, 0x3f, 0x84, 0x0e, 0xae, 0x07, 0x9b, 0x53, 0x1b, 0x76, 0x6d,
	0xcc, 0xae, 0x44, 0x37, 0x69, 0xf2, 0xba, 0xe9, 0x52, 0x47, 0xd4, 0x1f, 0xd8, 0xad, 0x9b, 0xb9,
	0xcb, 0xa6, 0xc7, 0x77, 0xa1, 0x57, 0x3c, 0x2a, 0xa4, 0x0b, 0xea, 0xe9, 0x87, 0xf3, 0xa1, 0x44,
	0x7a, 0xb0, 0xe7, 0xce, 0xdf, 0xb9, 0x43, 0xf9, 0x68, 0xf4, 0xb9, 0xbb, 0xe0, 0x02, 0xfc, 0xaf,
	0x52, 0xfc, 0x5a, 0xdc, 0x12, 0xaf, 0xd9, 0xb3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xac, 0xf7,
	0xd4, 0xb9, 0xdc, 0x04, 0x00, 0x00,
}
