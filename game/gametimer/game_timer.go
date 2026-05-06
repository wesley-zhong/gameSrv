package gametimer

import (
	"gameSrv/pkg/timer"
	"time"
)

var sharedTimer *timer.ShardedWheelTimer = timer.NewShardedWheelTimer(time.Millisecond, 2048)

func Start() {
	sharedTimer.Start()
}

func Stop() {
	sharedTimer.Stop()
}

func AddTask(key int64, delay time.Duration, job func()) {
	sharedTimer.AddTask(key, delay, job)
}

func RemoveTask(key int64) bool {
	return sharedTimer.RemoveTask(key)
}
