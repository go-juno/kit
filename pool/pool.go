package pool

import (
	"sync"
)

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// new pool, size is its maximum
func New(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Use coroutine, delta is the amount used
func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// one Coroutines complete
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

//Wait for all coroutines to complete
func (p *Pool) Wait() {
	p.wg.Wait()
}
