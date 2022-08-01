/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 */
package main

import (
	utils "reply2future.com/utils"
)

type Iterator struct {
	cursor   int
	storages []int
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return len(this.storages) > this.cursor
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	this.cursor += 1
	return this.storages[this.cursor-1]
}

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter      *Iterator
	peekValue int
	isPeeking bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.isPeeking || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.isPeeking {
		this.isPeeking = false
		return this.peekValue
	}

	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if !this.isPeeking {
		this.isPeeking = true
		this.peekValue = this.iter.next()
	}

	return this.peekValue
}

// @lc code=end
func main() {
	pi := Constructor(&Iterator{storages: []int{1, 2, 3}})
	utils.AssertEqual0[int]("next()",1,pi.next)
	utils.AssertEqual0[int]("peek()",2,pi.peek)
	utils.AssertEqual0[int]("next()",2,pi.next)
	utils.AssertEqual0[int]("next()",3,pi.next)
	utils.AssertEqual0[bool]("hasNext()",false,pi.hasNext)
}
