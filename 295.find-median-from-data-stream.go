/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */
package main

import (
	"container/heap"
	utils "reply2future.com/utils"
)

// minimum heap
// small, large heap
// 1. k, k => small.push(n), small.remove(k), large.push(k)
// 2. k, k + 1 => large.push(n), large.pop(), small.push()
// @lc code=start

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	// largeHeap: asc order
	// smallHeap: negative number of asc order
	smallHeap, largeHeap *IntHeap
	isOdd bool
}

func Constructor() MedianFinder {
	smallHeap := &IntHeap{}
	heap.Init(smallHeap)
	largeHeap := &IntHeap{}
	heap.Init(largeHeap)
	return MedianFinder{smallHeap: smallHeap, largeHeap: largeHeap}
}

func (this *MedianFinder) AddNum(num int) {
	this.isOdd = !this.isOdd
	sLen := len(*this.smallHeap)
	lLen := len(*this.largeHeap)
	if len(*this.largeHeap) == 0 {
		heap.Push(this.largeHeap, num)
		return
	}
	if sLen == lLen {
		heap.Push(this.smallHeap, -num)
		heap.Push(this.largeHeap, -heap.Pop(this.smallHeap).(int))
	} else {
		heap.Push(this.largeHeap, num)
		heap.Push(this.smallHeap, -heap.Pop(this.largeHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.isOdd {
		return float64((*this.largeHeap)[0])
	} else {
		return float64(-(*this.smallHeap)[0] + (*this.largeHeap)[0]) / 2.0
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	utils.AssertEqual0[float64]("example 1", 1.5, mf.FindMedian)
	mf.AddNum(3)
	utils.AssertEqual0[float64]("example 2", 2.0, mf.FindMedian)

	mf = Constructor()
	mf.AddNum(-1)
	mf.AddNum(-2)
	mf.AddNum(-3)
	mf.AddNum(-4)
	mf.AddNum(-5)
	utils.AssertEqual0[float64]("example 3", -3.0, mf.FindMedian)
}
