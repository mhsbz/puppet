package puppet

import (
	"time"
)

const (
	DefaultCleanIntervalTime = 5
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
}

func NewPool(cap int32) *Pool {
	return NewPoolWithExpire(cap, DefaultCleanIntervalTime)
}

func NewPoolWithExpire(cap int32, expireTime int32) *Pool {
	workers := make([]*Worker, cap)
	release := make(chan sig)
	p := &Pool{
		capacity:       cap,
		running:        0,
		expiryDuration: time.Duration(expireTime),
		workers:        workers,
		release:        release,
	}
	return p
}
