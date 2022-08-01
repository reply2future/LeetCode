/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */
package main

import (
	"fmt"
	utils "reply2future.com/utils"
)

// @lc code=start
type LRUNode struct {
	Prev  *LRUNode
	Next  *LRUNode
	Key   int
	Value int
}
// using head and tail node will be easier
type LRUCache struct {
	HashMap   map[int]*LRUNode
	LatestKey int
	LeastKey  int
	Capacity  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		HashMap:  make(map[int]*LRUNode),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	record, ok := this.HashMap[key]
	if !ok {
		return -1
	}

	this.UpdateLatestKey(key)
	
	return record.Value
}

func (this *LRUCache) UpdateLatestKey(key int) {
	if key == this.LatestKey {
		return
	}

	record, ok := this.HashMap[key]
	if !ok {
		return
	}

	oLatestNode := this.HashMap[this.LatestKey]
	oLatestNode.Next = record

	if key == this.LeastKey {
		this.LeastKey = record.Next.Key
		record.Next.Prev = nil
	} else {
		record.Next.Prev = record.Prev
		record.Prev.Next = record.Next
	}
	record.Prev = oLatestNode
	record.Next = nil
	this.LatestKey = key
}

func (this *LRUCache) Put(key int, value int) {
	record, ok := this.HashMap[key]
	if !ok {
		hLen := len(this.HashMap)
		
		oLatestNode := this.HashMap[this.LatestKey]
		newNode := &LRUNode{Value: value, Key: key, Prev: oLatestNode}
		this.HashMap[key] = newNode
		if hLen == 0 {
			this.LatestKey = key
			this.LeastKey = key
			return
		}

		oLatestNode.Next = newNode
		this.LatestKey = key

		if hLen >= this.Capacity {
			oLeastNode := this.HashMap[this.LeastKey]
			this.LeastKey = oLeastNode.Next.Key
			oLeastNode.Next.Prev = nil
			oLeastNode.Next = nil
			delete(this.HashMap, oLeastNode.Key)
		}
		return
	}

	record.Value = value
	this.UpdateLatestKey(key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
func main() {
	fmt.Println("first example ================================")
	obj1 := Constructor(2)
	obj1.Put(1, 1)
	obj1.Put(2, 2)
	utils.AssertEqual1[int, int]("get 1", 1, obj1.Get, 1)
	obj1.Put(3, 3)
	utils.AssertEqual1[int, int]("get 2", -1, obj1.Get, 2)
	obj1.Put(4, 4)
	utils.AssertEqual1[int, int]("get 1", -1, obj1.Get, 1)
	utils.AssertEqual1[int, int]("get 3", 3, obj1.Get, 3)
	utils.AssertEqual1[int, int]("get 4", 4, obj1.Get, 4)

	fmt.Println("second example ================================")
	obj2 := Constructor(1)
	obj2.Put(2, 1)
	utils.AssertEqual1[int, int]("get 2", 1, obj2.Get, 2)
	obj2.Put(3, 2)
	utils.AssertEqual1[int, int]("get 2", -1, obj2.Get, 2)
	utils.AssertEqual1[int, int]("get 3", 2, obj2.Get, 3)

	fmt.Println("third example ================================")
	obj3 := Constructor(2)
	obj3.Put(2, 1)
	obj3.Put(2, 2)
	utils.AssertEqual1[int, int]("get 2", 2, obj3.Get, 2)
	obj3.Put(1, 1)
	obj3.Put(4, 1)
	utils.AssertEqual1[int, int]("get 2", -1, obj3.Get, 2)
}
