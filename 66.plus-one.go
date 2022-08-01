/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
package main

import (
	assert "reply2future.com/utils"
)

// @lc code=start
func plusOne(digits []int) []int {
	dLen := len(digits)
	for i := dLen - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}

	ret := make([]int, dLen+1)
	ret[0] = 1
	return ret
}

// @lc code=end
func main() {
	assert.AssertEqual1[[]int, []int]("[1,2,3]", []int{1, 2, 4}, plusOne, []int{1, 2, 3})
	assert.AssertEqual1[[]int, []int]("[4,3,2,1]", []int{4, 3, 2, 2}, plusOne, []int{4, 3, 2, 1})
	assert.AssertEqual1[[]int, []int]("[9]", []int{1, 0}, plusOne, []int{9})
	assert.AssertEqual1[[]int, []int]("[9,9]", []int{1, 0, 0}, plusOne, []int{9, 9})
}
