package gopool

type WorkerPool struct {
	workerCount   int
	goHashFuc     func() uint
	workerQueSize int
	workers       []*Worker
}

// worker count
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

func (pool *WorkerPool) SubmitTask(task func()) error {
	//hashCode := pool.goHashFuc()
	//index := hashCode % uint(pool.workerCount)
	pool.workers[0].AsyExecute(task)
	return nil
}

func (pool *WorkerPool) HashWorker(hashCode int) *Worker {
	size := len(pool.workers)
	return pool.workers[hashCode%size]

}
