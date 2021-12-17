package pool

import (
	"sync"
)

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// new pool, size is its maximum
func New(size int) *pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Use coroutine, delta is the amount used
func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// one Coroutines complete
func (p *pool) Done() {
	<-p.queue
	p.wg.Done()
}

//Wait for all coroutines to complete
func (p *pool) Wait() {
	p.wg.Wait()
}
