package buff

import (
	"sync"
)

// BuffObjPool manages buff object pooling
type BuffObjPool struct {
	pool sync.Pool
}

// Singleton instance
var (
	buffPoolInstance *BuffObjPool
	buffPoolOnce     sync.Once
)

// GetInstance returns the singleton BuffObjPool instance
func GetBuffPoolInstance() *BuffObjPool {
	buffPoolOnce.Do(func() {
		buffPoolInstance = &BuffObjPool{
			pool: sync.Pool{
				New: func() interface{} {
					return &FYBuff{}
				},
			},
		}
	})
	return buffPoolInstance
}

// NewBuff creates a new buff with the given UID
func (p *BuffObjPool) NewBuff(bufUID int64) *FYBuff {
	buff := p.pool.Get().(*FYBuff)
	buff.UID = bufUID
	buff.init()
	return buff
}

// DelBuff returns a buff to the pool
func (p *BuffObjPool) DelBuff(buff *FYBuff) {
	if buff == nil {
		return
	}
	// Reset buff state before returning to pool
	buff.reset()
	p.pool.Put(buff)
}
