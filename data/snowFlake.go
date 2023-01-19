package main

import (
	"errors"
	"sync"
	"time"
)

/**
雪花算法
	 41bit timestamp | 10 bit machineID : 5bit workerID 5bit dataCenterID ｜ 12 bit sequenceBits
最多使用69年
*/

const (
	workerIDBits     = uint64(5) // 10bit 工作机器ID中的 5bit workerID
	dataCenterIDBits = uint64(5) // 10 bit 工作机器ID中的 5bit dataCenterID
	sequenceBits     = uint64(12)

	maxWorkerID     = int64(-1) ^ (int64(-1) << workerIDBits) //节点ID的最大值 用于防止溢出
	maxDataCenterID = int64(-1) ^ (int64(-1) << dataCenterIDBits)
	maxSequence     = int64(-1) ^ (int64(-1) << sequenceBits)

	timeLeft = uint8(22) // timeLeft = workerIDBits + sequenceBits // 时间戳向左偏移量
	dataLeft = uint8(17) // dataLeft = dataCenterIDBits + sequenceBits
	workLeft = uint8(12) // workLeft = sequenceBits // 节点IDx向左偏移量
	// 2020-05-20 08:00:00 +0800 CST
	twepoch = int64(1589923200000) // 常量时间戳(毫秒)
)

type IDGenerator struct {
	mu           sync.Mutex
	LastStamp    int64 // 记录上一次ID的时间戳
	WorkerID     int64 // 该节点的ID
	DataCenterID int64 // 该节点的 数据中心ID
	Sequence     int64 // 当前毫秒已经生成的ID序列号(从0 开始累加) 1毫秒内最多生成4096个ID
}

// 分布式情况下,我们应通过外部配置文件或其他方式为每台机器分配独立的id
func NewSnowFlake(workerID, dataCenterID int64) *IDGenerator {
	return &IDGenerator{
		WorkerID:     workerID,
		LastStamp:    0,
		Sequence:     0,
		DataCenterID: dataCenterID,
	}
}

func (g *IDGenerator) getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func (g *IDGenerator) Acquire() (uint64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.acquire()
}

func (g *IDGenerator) acquire() (uint64, error) {
	curTime := g.getMilliSeconds()

	if g.LastStamp == curTime {

		g.Sequence = (g.Sequence + 1) & maxSequence

		if g.Sequence == 0 {
			for curTime <= g.LastStamp {
				curTime = g.getMilliSeconds()
			}
		}
	} else {
		if curTime < g.LastStamp {
			return 0, errors.New("time is moving backwards,waiting until")
		} else {
			g.Sequence = 0
		}
		g.LastStamp = curTime
	}

	id := ((curTime - twepoch) << timeLeft) |
		(g.DataCenterID << dataLeft) |
		(g.WorkerID << workLeft) |
		g.Sequence

	return uint64(id), nil
}
