package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (myatomic *NoPad) IncreaseAll() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}

type Pad struct {
	a uint64
	_ [7]uint64
	b uint64
	_ [7]uint64
	c uint64
	_ [7]uint64
}

func (myatomic *Pad) IncreaseAll() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}

type MyAtomic interface {
	IncreaseAll()
}

func testAtomicIncrease(myatomic MyAtomic, n int) {
	addTimes := 100
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				myatomic.IncreaseAll()
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

func BenchmarkNoPad(b *testing.B) {
	myatomic := &NoPad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic, b.N)
}

func BenchmarkPad(b *testing.B) {
	myatomic := &Pad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic, b.N)
}
