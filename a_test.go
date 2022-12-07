package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

type NoPad struct {
	a [2]uint64
	b [2]uint64
	c [2]uint64
}

func (myatomic *NoPad) IncreaseAll() {
	atomic.AddUint64(&myatomic.a[0], 1)
	atomic.AddUint64(&myatomic.b[0], 1)
	atomic.AddUint64(&myatomic.c[0], 1)
}

type Pad struct {
	a [8]uint64
	b [8]uint64
	c [8]uint64
}

func (myatomic *Pad) IncreaseAll() {
	atomic.AddUint64(&myatomic.a[0], 1)
	atomic.AddUint64(&myatomic.b[0], 1)
	atomic.AddUint64(&myatomic.c[0], 1)
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
	myatomic.a = [2]uint64{0, 1}
	b.ResetTimer()
	testAtomicIncrease(myatomic, b.N)
}

func BenchmarkPad(b *testing.B) {
	myatomic := &Pad{}
	myatomic.a = [8]uint64{0}
	myatomic.b = [8]uint64{0}
	myatomic.c = [8]uint64{0}
	b.ResetTimer()
	testAtomicIncrease(myatomic, b.N)
}
