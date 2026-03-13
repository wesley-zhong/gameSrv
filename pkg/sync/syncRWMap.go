package sync

import (
	"sync"
)

// SyncRWMap 分段锁的并发安全Map，基于泛型实现
// 将数据分成多个分片(shard)，每个分片有独立的RWMutex
// 相比单一锁，多个不同key的操作可以并发执行，大幅减少锁竞争

type SyncRWMap[K comparable, V any] struct {
	shards    []*shard[K, V]
	shardMask uint64
}

type shard[K comparable, V any] struct {
	data   map[K]V
	rwLock sync.RWMutex
}

// DefaultShardCount 默认分片数量
const DefaultShardCount = 64

// NewSyncRWMap 创建一个新的SyncRWMap，使用默认分片数量(64)
func NewSyncRWMap[K comparable, V any]() *SyncRWMap[K, V] {
	return NewSyncRWMapWithShards[K, V](DefaultShardCount)
}

// NewSyncRWMapWithShards 创建指定分片数量的SyncRWMap
// shardCount 会被调整为2的幂次方以优化位运算
func NewSyncRWMapWithShards[K comparable, V any](shardCount int) *SyncRWMap[K, V] {
	if shardCount <= 0 {
		shardCount = 1
	}
	shardCount = roundToPowerOfTwo(shardCount)

	shards := make([]*shard[K, V], shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = &shard[K, V]{
			data:   make(map[K]V),
			rwLock: sync.RWMutex{},
		}
	}

	return &SyncRWMap[K, V]{
		shards:    shards,
		shardMask: uint64(shardCount - 1),
	}
}

// roundToPowerOfTwo 计算大于等于n的最小2的幂次方
func roundToPowerOfTwo(n int) int {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return n
}

// getShard 根据key的哈希值获取对应的分片
func (m *SyncRWMap[K, V]) getShard(key K) *shard[K, V] {
	// 使用内置的哈希函数
	hash := uint64(defaultHash(key))
	idx := hash & m.shardMask
	return m.shards[idx]
}

// defaultHash 默认的哈希函数
func defaultHash[K comparable](k K) uint64 {
	switch v := any(k).(type) {
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	case uintptr:
		return uint64(v)
	case string:
		return fnvHash64(v)
	default:
		// 对于其他类型，使用Go内置的hash函数
		return uint64(comparableHash(k))
	}
}

// fnvHash64 FNV-1a哈希算法，用于字符串
func fnvHash64(s string) uint64 {
	const offset64 uint64 = 14695981039346656037
	const prime64 uint64 = 1099511628211
	hash := offset64
	for i := 0; i < len(s); i++ {
		hash ^= uint64(s[i])
		hash *= prime64
	}
	return hash
}

// comparableHash 用于其他comparable类型的简单哈希
func comparableHash[K comparable](k K) uint64 {
	// 简单的实现，实际使用中可以考虑更复杂的哈希
	var h uint64
	// Go不允许直接遍历comparable类型的字节
	// 这里只是一个占位，实际使用时可以根据具体类型优化
	return h
}

// Store 设置key对应的value
func (m *SyncRWMap[K, V]) Store(key K, value V) {
	shard := m.getShard(key)
	shard.rwLock.Lock()
	defer shard.rwLock.Unlock()
	shard.data[key] = value
}

// Load 获取key对应的value，如果不存在则返回false
func (m *SyncRWMap[K, V]) Load(key K) (V, bool) {
	shard := m.getShard(key)
	shard.rwLock.RLock()
	defer shard.rwLock.RUnlock()
	value, ok := shard.data[key]
	return value, ok
}

// LoadOrStore 获取key对应的value，如果不存在则存储并返回false
// 返回的loaded为true表示值已存在，false表示是新插入的
func (m *SyncRWMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	shard := m.getShard(key)
	shard.rwLock.Lock()
	defer shard.rwLock.Unlock()
	actual, loaded = shard.data[key]
	if loaded {
		return
	}
	shard.data[key] = value
	actual = value
	loaded = false
	return
}

// LoadAndDelete 获取并删除key对应的value
func (m *SyncRWMap[K, V]) LoadAndDelete(key K) (V, bool) {
	shard := m.getShard(key)
	shard.rwLock.Lock()
	defer shard.rwLock.Unlock()
	value, ok := shard.data[key]
	if ok {
		delete(shard.data, key)
	}
	return value, ok
}

// Delete 删除key对应的value
func (m *SyncRWMap[K, V]) Delete(key K) {
	shard := m.getShard(key)
	shard.rwLock.Lock()
	defer shard.rwLock.Unlock()
	delete(shard.data, key)
}

// Range 遍历所有键值对
// fn返回false时停止遍历
// 注意：遍历期间持有所有分片的读锁，避免长时间遍历
func (m *SyncRWMap[K, V]) Range(fn func(key K, value V) bool) {
	for _, shard := range m.shards {
		shard.rwLock.RLock()
		for key, value := range shard.data {
			if !fn(key, value) {
				shard.rwLock.RUnlock()
				return
			}
		}
		shard.rwLock.RUnlock()
	}
}

// Size 返回map中的键值对总数
func (m *SyncRWMap[K, V]) Size() int {
	size := 0
	for _, shard := range m.shards {
		shard.rwLock.RLock()
		size += len(shard.data)
		shard.rwLock.RUnlock()
	}
	return size
}

// Len 返回map中的键值对总数（Size的别名）
func (m *SyncRWMap[K, V]) Len() int {
	return m.Size()
}

// Clear 清空所有键值对
func (m *SyncRWMap[K, V]) Clear() {
	for _, shard := range m.shards {
		shard.rwLock.Lock()
		shard.data = make(map[K]V)
		shard.rwLock.Unlock()
	}
}

// Keys 返回所有key的切片（快照）
func (m *SyncRWMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Size())
	m.Range(func(key K, value V) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}

// Values 返回所有value的切片（快照）
func (m *SyncRWMap[K, V]) Values() []V {
	values := make([]V, 0, m.Size())
	m.Range(func(key K, value V) bool {
		values = append(values, value)
		return true
	})
	return values
}