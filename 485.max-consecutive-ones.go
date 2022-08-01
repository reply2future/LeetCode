/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
    max := 0

	count := 0
	for _, v := range nums {
		if v == 0 {
			if count > max {
				max = count
			}
			count = 0
		} else {
			count++
		}
	}
	if count > max {
		max = count
	}
	return max
}
// @lc code=end
func main() {
	utils.AssertEqual1[int,[]int]("example 1",3,findMaxConsecutiveOnes,[]int{1,1,0,1,1,1})
	utils.AssertEqual1[int,[]int]("example 2",2,findMaxConsecutiveOnes,[]int{1,0,1,1,0,1})
}
