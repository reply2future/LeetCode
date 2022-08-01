/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */
package main

import (
	utils "reply2future.com/utils"
	"strconv"
)

// @lc code=start
func summaryRanges(nums []int) []string {
	var ret []string
	cursor := 0
	nLen := len(nums)
	for cursor < nLen {
		if cursor == nLen-1 || nums[cursor]+1 != nums[cursor+1] {
			ret = append(ret, strconv.Itoa(nums[cursor]))
			cursor += 1
			continue
		}

		i := 1
		for ; cursor+i < nLen-1; i++ {
			if nums[i+cursor]+1 != nums[i+cursor+1] {
				break
			}
		}

		ret = append(ret, strconv.Itoa(nums[cursor])+"->"+strconv.Itoa(nums[cursor+i]))

		cursor += i + 1
	}

	return ret
}

// @lc code=end
func main() {
	utils.AssertEqual1[[]string, []int]("example 1", []string{"0->2", "4->5", "7"}, summaryRanges, []int{0, 1, 2, 4, 5, 7})
	utils.AssertEqual1[[]string, []int]("example 2", []string{"0", "2->4", "6", "8->9"}, summaryRanges, []int{0, 2, 3, 4, 6, 8, 9})
}
