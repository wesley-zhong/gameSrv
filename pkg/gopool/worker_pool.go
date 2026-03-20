package gopool

type WorkerPool struct {
	workerCount   int
	workerQueSize int
	workers       []*Worker
}

func StartNewWorkerPool(workerCount int, workQueSize int) *WorkerPool {
	pool := &WorkerPool{
		workerCount:   workerCount,
		workerQueSize: workQueSize,
		workers:       make([]*Worker, workerCount),
	}
	for i := 0; i < workerCount; i++ {
		pool.workers[i] = newWorker(workQueSize)
		go pool.workers[i].Start()
	}
	return pool
}

func (pool *WorkerPool) SubmitTask(hashCode int64, task func()) error {
	index := int(hashCode&0x7FFFFFFF) % pool.workerCount
	pool.workers[index].AsyExecute(task)
	return nil
}

func (pool *WorkerPool) HashWorker(hashCode int64) *Worker {
	size := len(pool.workers)
	return pool.workers[int(hashCode&0x7FFFFFFF)%size]
}
