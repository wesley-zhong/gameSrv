package gopool

import (
	"sync/atomic"
)

type Worker struct {
	taskChan  chan func()
	taskCount atomic.Int64
}

func newWorker(taskCount int) *Worker {
	return &Worker{
		taskChan:  make(chan func(), taskCount),
	}
}

func (w *Worker) AsyExecute(task func()) {
	w.taskCount.Add(1)
	w.taskChan <- task
}

func (w *Worker) Start() {
	for task := range w.taskChan {
		task()
		w.taskCount.Add(-1)
	}
}

func (w *Worker) TaskCount() int64 {
	return w.taskCount.Load()
}
