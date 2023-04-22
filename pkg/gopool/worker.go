package gopool

import (
	"gameSvr/pkg/log"
	"sync/atomic"
	"time"
)

type Worker struct {
	taskChan  chan func()
	taskCount int64
}

func newWorker(taskCount int) *Worker {
	return &Worker{
		taskChan:  make(chan func(), taskCount),
		taskCount: 0,
	}
}

func (worker *Worker) AsyExcute(task func()) {
	atomic.AddInt64(&worker.taskCount, 1)
	worker.taskChan <- task
}

func (worker *Worker) Start() {
	for {
		select {
		case task, ok := <-worker.taskChan:
			if ok {
				task()
				atomic.AddInt64(&worker.taskCount, -1)
			} else {
				log.Infof("------ some error")
			}
		default:
			time.Sleep(1 * time.Millisecond)
			//log.Infof("-------- default")
		}
	}
}

func (worker *Worker) TaskCount() int64 {
	return worker.taskCount
}
