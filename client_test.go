package brqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_EnqueueRequest_Pop(t *testing.T) {
	client, err := NewClient("localhost", 6431)
	if assert.NoError(t, err) {

		id, err := client.EnqueueRequest([]byte("Hello"), HighPriority, []string{"foo"})
		if assert.NoError(t, err) {

			task, err := client.Pop([]string{"foo"}, true)
			if assert.NoError(t, err) {
				assert.Equal(t, id, task.id)

				assert.NoError(t, client.Acknowledge(task))
			}
		}
	}
}
