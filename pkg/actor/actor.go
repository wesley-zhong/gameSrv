package actor

import (
	"gameSrv/pkg/gopool"
	"runtime"
)

type Actor struct {
	id           uint32
	routerWorker *gopool.Worker
}

var routerPool = gopool.StartNewWorkerPool(runtime.NumCPU(), 1024)

func NewActor(id int) *Actor {
	return &Actor{
		id:           1,
		routerWorker: routerPool.HashWorker(id),
	}
}

func (actor *Actor) Call(call func()) {
	actor.routerWorker.AsyExecute(call)
}

func (actor *Actor) CallWithBack(call func(), callBack func(), calledActor *Actor) {
	actor.routerWorker.AsyExecute(func() {
		call()
		calledActor.routerWorker.AsyExecute(callBack)
	})
}
