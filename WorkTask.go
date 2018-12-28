package brqueue

// A task that can be worked on.
// Do not attempt to create a new instance yourself, it will fail.
type WorkTask struct {
	id      string
	message []byte
	valid   bool
}

func (w *WorkTask) GetId() string {
	return w.id
}

func (w *WorkTask) GetMessage() []byte {
	return w.message
}

func newWorkTask(id string, message []byte) *WorkTask {
	return &WorkTask{
		id: id, message: message,
		valid: true,
	}
}
