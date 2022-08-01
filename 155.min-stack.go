/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */
package main
import (
	utils "reply2future.com/utils"
	"fmt"
)
// @lc code=start
type MinStackElement struct {
	Value int
	CurrentMin int
}

type MinStack struct {
    Stack []MinStackElement
}


func Constructor() MinStack {
    return MinStack{
		Stack: make([]MinStackElement, 0),
	}
}


func (this *MinStack) Push(val int)  {
	sLen := len(this.Stack)
	var min int
	if sLen == 0 {
		min = val
	} else {
		prev := this.Stack[sLen - 1].CurrentMin
		if prev > val {
			min = val
		} else {
			min = prev
		}
	}
    this.Stack = append(this.Stack, MinStackElement{Value: val, CurrentMin: min})
}


func (this *MinStack) Pop()  {
	sLen := len(this.Stack)
	if sLen == 0 {
		return
	}
    this.Stack = this.Stack[:sLen - 1]
}


func (this *MinStack) Top() int {
	sLen := len(this.Stack)
	if sLen == 0 {
		panic("Empty stack cannot call Top()")
	}
    return this.Stack[sLen - 1].Value
}


func (this *MinStack) GetMin() int {
	sLen := len(this.Stack)
	if sLen == 0 {
		panic("Empty stack cannot call GetMin()")
	}
    return this.Stack[sLen - 1].CurrentMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
func main() {
	fmt.Println("================ example 1 =====================")	
	obj1 := Constructor()
	obj1.Push(-2)
	obj1.Push(0)
	obj1.Push(-3)
	utils.AssertEqual0[int]("getMin",-3,obj1.GetMin)
	obj1.Pop()
	utils.AssertEqual0[int]("top",0,obj1.Top)
	utils.AssertEqual0[int]("getMin",-2,obj1.GetMin)
}
