/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
	nm := make(map[int]interface{})
	for _, num := range nums {
		nm[num] = struct{}{}
	}
	
	ret := make([]int, 0)
	for i := 1; i <= n; i++ {
		if _, ok := nm[i]; !ok {
			ret = append(ret, i)
		}
	}
	return ret
}
// @lc code=end
func main() {
	utils.AssertEqual1[[]int,[]int]("example 1",[]int{5,6},findDisappearedNumbers,[]int{4,3,2,7,8,2,3,1})
	utils.AssertEqual1[[]int,[]int]("example 2",[]int{2},findDisappearedNumbers,[]int{1,1})
}
