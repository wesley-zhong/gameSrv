package timer

import (
	"sync"
	"time"
)

// Task 定义任务节点
type Task struct {
	key    int64
	rounds int
	job    func()
	next   *Task // 手动链表，避免 container/list 的额外的内存分配
}

// bucket 槽位，使用独立的互斥锁减少竞争
type bucket struct {
	mu   sync.Mutex
	head *Task
}

// ShardedWheelTimer 高性能时间轮
type ShardedWheelTimer struct {
	interval    time.Duration
	slotNum     int
	slots       []*bucket
	currentSlot int

	taskPool    sync.Pool
	stopChannel chan struct{}
}

// NewShardedWheelTimer 初始化
func NewShardedWheelTimer(interval time.Duration, slotNum int) *ShardedWheelTimer {
	sw := &ShardedWheelTimer{
		interval:    interval,
		slotNum:     slotNum,
		slots:       make([]*bucket, slotNum),
		stopChannel: make(chan struct{}),
		taskPool: sync.Pool{
			New: func() interface{} {
				return &Task{}
			},
		},
	}
	for i := 0; i < slotNum; i++ {
		sw.slots[i] = &bucket{}
	}
	return sw
}

// Start 启动时间轮转动
func (sw *ShardedWheelTimer) Start() {
	ticker := time.NewTicker(sw.interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				sw.tick()
			case <-sw.stopChannel:
				ticker.Stop()
				return
			}
		}
	}()
}

// AddTask 添加任务
func (sw *ShardedWheelTimer) AddTask(key int64, delay time.Duration, job func()) {
	steps := int(delay / sw.interval)
	rounds := steps / sw.slotNum
	slotIdx := (sw.currentSlot + steps) % sw.slotNum

	b := sw.slots[slotIdx]

	// 从池中获取任务对象
	t := sw.taskPool.Get().(*Task)
	t.key = key
	t.rounds = rounds
	t.job = job

	b.mu.Lock()
	// 插入链表头部 (O(1))
	t.next = b.head
	b.head = t
	b.mu.Unlock()
}

// tick 核心执行逻辑
func (sw *ShardedWheelTimer) tick() {
	b := sw.slots[sw.currentSlot]

	b.mu.Lock()
	defer b.mu.Unlock()

	var prev *Task
	curr := b.head

	for curr != nil {
		if curr.rounds <= 0 {
			// 1. 执行任务 (使用 goroutine 异步执行，避免阻塞时间轮)
			go curr.job()

			// 2. 从链表中移除
			next := curr.next
			if prev == nil {
				b.head = next
			} else {
				prev.next = next
			}

			// 3. 重置并回收对象到池中
			target := curr
			curr = next

			target.key = 0
			target.job = nil
			target.next = nil
			sw.taskPool.Put(target)
		} else {
			curr.rounds--
			prev = curr
			curr = curr.next
		}
	}

	// 移动指针
	sw.currentSlot = (sw.currentSlot + 1) % sw.slotNum
}

// RemoveTask 移除指定 key 的任务
func (sw *ShardedWheelTimer) RemoveTask(key int64) bool {
	// 遍历所有槽位查找任务
	for _, b := range sw.slots {
		b.mu.Lock()
		var prev *Task
		curr := b.head

		for curr != nil {
			if curr.key == key {
				// 找到目标任务，从链表中移除
				next := curr.next
				if prev == nil {
					b.head = next
				} else {
					prev.next = next
				}

				// 重置并回收对象到池中
				curr.key = 0
				curr.rounds = 0
				curr.job = nil
				curr.next = nil
				sw.taskPool.Put(curr)

				b.mu.Unlock()
				return true
			}
			prev = curr
			curr = curr.next
		}
		b.mu.Unlock()
	}
	return false
}

// Stop 关闭时间轮
func (sw *ShardedWheelTimer) Stop() {
	close(sw.stopChannel)
}
