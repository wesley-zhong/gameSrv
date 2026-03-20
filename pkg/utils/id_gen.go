package utils

import (
	"fmt"
	"sync"
	"time"
)

type SnowFlakeIdWorker struct {
	twepoch          int64
	workerIdBits     int64
	dataCenterIdBits int64
	maxWorkerId      int64
	maxDataCenterId int64
	sequenceBits    int64
	workerIdShift   int64
	dataCenterIdShift int64
	timestampLeftShift int64
	sequenceMask    int64
	workerId        int64
	dataCenterId    int64
	sequence        int64
	lastTimestamp   int64
	lock            sync.Mutex
}

var idWorker *SnowFlakeIdWorker

func IdGenInit(dataCenterId int64, workerId int64) {
	idWorker = new(SnowFlakeIdWorker)
	idWorker.init(dataCenterId, workerId)
}

func NextId() int64 {
	return idWorker.nextId()
}

func (p *SnowFlakeIdWorker) init(dataCenterId int64, workerId int64) {
	p.twepoch = 1622476800000
	p.workerIdBits = 5
	p.dataCenterIdBits = 5
	p.maxWorkerId = -1 ^ (-1 << p.workerIdBits)
	p.maxDataCenterId = -1 ^ (-1 << p.dataCenterIdBits)
	p.sequenceBits = 12
	p.workerIdShift = p.sequenceBits
	p.dataCenterIdShift = p.sequenceBits + p.workerIdBits
	p.timestampLeftShift = p.sequenceBits + p.workerIdBits + p.dataCenterIdBits
	p.sequenceMask = -1 ^ (-1 << p.sequenceBits)

	if workerId > p.maxWorkerId || workerId < 0 {
		panic(fmt.Errorf("Worker ID cannot be greater than %d or less than 0", p.maxWorkerId))
	}
	if dataCenterId > p.maxDataCenterId || dataCenterId < 0 {
		panic(fmt.Errorf("DataCenter ID cannot be greater than %d or less than 0", p.maxDataCenterId))
	}

	p.workerId = workerId
	p.dataCenterId = dataCenterId
	p.sequence = 0
	p.lastTimestamp = -1
}

func (p *SnowFlakeIdWorker) nextId() int64 {
	p.lock.Lock()
	defer p.lock.Unlock()

	timestamp := p.timeGen()
	if timestamp < p.lastTimestamp {
		panic(fmt.Errorf("Clock moved backwards. Refusing to generate id for %d milliseconds", p.lastTimestamp-timestamp))
	}

	if p.lastTimestamp == timestamp {
		p.sequence = (p.sequence + 1) & p.sequenceMask
		if p.sequence == 0 {
			timestamp = p.tilNextMillis(p.lastTimestamp)
		}
	} else {
		p.sequence = 0
	}

	p.lastTimestamp = timestamp

	return ((timestamp - p.twepoch) << p.timestampLeftShift) |
		(p.dataCenterId << p.dataCenterIdShift) |
		(p.workerId << p.workerIdShift) | p.sequence
}

func (p *SnowFlakeIdWorker) tilNextMillis(lastTimestamp int64) int64 {
	timestamp := p.timeGen()
	for timestamp <= lastTimestamp {
		timestamp = p.timeGen()
	}
	return timestamp
}

func (p *SnowFlakeIdWorker) timeGen() int64 {
	return time.Now().UnixNano() / 1e6
}