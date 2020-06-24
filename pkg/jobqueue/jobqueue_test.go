package jobqueue

import (
	"testing"
)

const test_string = "test_answer"

type testJob struct {
	answer string
}

func (tj *testJob) Process() {
	tj.answer = test_string
}

func TestJobQueue(t *testing.T) {
	jQueue, _ := NewJobQueue(5)
	tj := &testJob{}
	jQueue.PushJob(tj)
	jQueue.Close()
	if tj.answer != test_string {
		t.Errorf(
			"job was not done properly, get result: %s, expected: %s",
			tj.answer,
			test_string)
	}
}
