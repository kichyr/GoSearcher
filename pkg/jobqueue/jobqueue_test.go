package jobqueue

import (
	"testing"
)

type testJob struct {
	answer string
}

func (tj *testJob) Process() {
	tj.answer = "test_answer"
}

func TestJobQueue(t *testing.T) {
	jQueue := NewJobQueue(5)
	tj := &testJob{}
	jQueue.PushJob()
	jQueue.Close()
	assert.Equal(t, tj.answer, "test_answer", "The two words should be the same.")

}
