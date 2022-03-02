package puppet

import (
	"math"
	"sync"
	"time"
)

type sig struct {
}
type Pool struct {
	// capacity of the pool.
	capacity int32
	// running is the number of the currently running goroutines.
	running int32
	// expiryDuration set the expired time (second) of every worker.
	expiryDuration time.Duration
	// workers is a slice that store the available workers.
	workers []*Worker
	// release is used to notice the pool to closed itself.
	release chan sig
	// lock for synchronous operation.
	lock sync.Mutex
	once sync.Once
}

// NewPool generates a instance of ants pool
func NewPool(size int) (*Pool, error) {
	return NewTimingPool(size, DefaultCleanIntervalTime)
}

// NewTimingPool generates a instance of ants pool with a custom timed task
func NewTimingPool(size, expiry int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	if expiry <= 0 {
		return nil, ErrInvalidPoolExpiry
	}
	p := &Pool{
		capacity:       int32(size),
		freeSignal:     make(chan sig, math.MaxInt32),
		release:        make(chan sig, 1),
		expiryDuration: time.Duration(expiry) * time.Second,
	}
	// 启动定期清理过期worker任务，独立goroutine运行，
	// 进一步节省系统资源
	p.monitorAndClear()
	return p, nil
}
