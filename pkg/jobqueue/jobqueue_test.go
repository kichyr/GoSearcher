package jobqueue

import (
	"testing"
	"time"
)

type testJob struct {
	jobId int
	res   chan int
}

func (tj *testJob) Process() {
	if tj.jobId == 0 {
		time.Sleep(1 * time.Second)
		tj.res <- tj.jobId
	} else {
		tj.res <- tj.jobId
	}
}

func TestJobQueueOneWorker(t *testing.T) {
	jQueue, _ := NewJobQueue(1)
	res := make(chan int, 2)
	jQueue.PushJob(&testJob{0, res})
	jQueue.PushJob(&testJob{1, res})
	jQueue.Close()
	res1 := <-res
	res2 := <-res

	if !(res1 == 0 && res2 == 1) {
		t.Errorf(
			"For worker number = 1 JobQueue got the result in an inconsistent sequence with push order.",
		)
	}
}
func TestJobQueueTwoWorkers(t *testing.T) {
	jQueue, _ := NewJobQueue(2)
	res := make(chan int, 2)
	jQueue.PushJob(&testJob{0, res})
	jQueue.PushJob(&testJob{1, res})
	jQueue.Close()
	res1 := <-res
	res2 := <-res

	if !(res1 == 1 && res2 == 0) {
		t.Errorf(
			"For worker number = 2 JobQueue should have done the faster job first but that didn't happen",
		)
	}
}
