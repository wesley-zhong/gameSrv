package actor

import (
	"gameSvr/pkg/log"
	"time"
)

type Actor struct {
	id       uint32
	mq       chan *CallMethod
	cbActors map[int64]*Actor
}
type CallMethod struct {
	id   int64
	call func()
	cb   func()
}

func NewActor(msgCap int) *Actor {
	return &Actor{
		id:       1,
		mq:       make(chan *CallMethod, msgCap),
		cbActors: make(map[int64]*Actor, 0),
	}
}

func (actor *Actor) Call(call func()) {
	callmethod := &CallMethod{
		id:   int64(0),
		call: call,
	}
	actor.mq <- callmethod
}

func (actor *Actor) CallWithBack(call func(), callBack func(), calledActor *Actor) {
	id := 444 //gen 64
	actor.cbActors[int64(id)] = calledActor
	callMethod := &CallMethod{
		id:   int64(id),
		call: call,
		cb:   callBack,
	}
	actor.mq <- callMethod

}

func (actor *Actor) Run() {
	for {
		select {
		case callMethod, ok := <-actor.mq:
			if ok {
				callMethod.call()
				if callMethod.id == 0 {
					return
				}
				actor.cbActors[callMethod.id].Call(callMethod.cb)

			} else {
				log.Infof("------ some error")
			}
		default:
			time.Sleep(1 * time.Millisecond)
			//log.Infof("-------- default")
		}
	}
}
