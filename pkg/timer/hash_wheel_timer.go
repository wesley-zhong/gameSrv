package timer

import (
	"sync"
	"time"
)

// TaskNode 优化后的任务节点，避免使用 container/list 的指针开销
type TaskNode struct {
	key    int64
	rounds int
	job    func()
	next   *TaskNode
}

type bucket struct {
	sync.RWMutex
	head *TaskNode
}

type ShardedWheelTimer struct {
	interval    time.Duration
	slots       []*bucket
	slotNum     int
	currentSlot int

	// 使用任务池减少 GC
	nodePool sync.Pool
}

func NewShardedWheelTimer(interval time.Duration, slotNum int) *ShardedWheelTimer {
	sw := &ShardedWheelTimer{
		interval: interval,
		slotNum:  slotNum,
		slots:    make([]*bucket, slotNum),
		nodePool: sync.Pool{
			New: func() interface{} { return &TaskNode{} },
		},
	}
	for i := 0; i < slotNum; i++ {
		sw.slots[i] = &bucket{}
	}
	return sw
}

// AddTask 极致优化：直接定位 Slot 插入，无需全局 Map
func (sw *ShardedWheelTimer) AddTask(key int64, delay time.Duration, job func()) {
	steps := int(delay / sw.interval)
	rounds := steps / sw.slotNum
	slotIdx := (sw.currentSlot + steps) % sw.slotNum

	b := sw.slots[slotIdx]

	// 从池中获取节点
	node := sw.nodePool.Get().(*TaskNode)
	node.key = key
	node.rounds = rounds
	node.job = job

	b.Lock()
	node.next = b.head
	b.head = node
	b.Unlock()
}

// tick 扫描当前 Slot
func (sw *ShardedWheelTimer) tick() {
	b := sw.slots[sw.currentSlot]

	b.Lock()
	prev := (*TaskNode)(nil)
	curr := b.head

	for curr != nil {
		if curr.rounds <= 0 {
			// 1. 移除节点
			next := curr.next
			if prev == nil {
				b.head = next
			} else {
				prev.next = next
			}

			// 2. 执行任务 (建议移交给协程池)
			go curr.job()

			// 3. 回收节点
			tmp := curr
			curr = next
			sw.nodePool.Put(tmp)
		} else {
			curr.rounds--
			prev = curr
			curr = curr.next
		}
	}
	b.Unlock()

	sw.currentSlot = (sw.currentSlot + 1) % sw.slotNum
}
