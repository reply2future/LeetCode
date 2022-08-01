/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */
package main

import (
	utils "reply2future.com/utils"
	"sort"
)

// @lc code=start
func hIndex(citations []int) int {
    sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	cLen := len(citations)
	for i := 0; i < cLen; i++ {
		if citations[i] < i + 1 {
			return i
		}
	}
	return cLen
}
// @lc code=end
func main() {
	utils.AssertEqual1[int,[]int]("example 1",3,hIndex,[]int{3,0,6,1,5})
	utils.AssertEqual1[int,[]int]("example 2",1,hIndex,[]int{1,3,1})
	utils.AssertEqual1[int,[]int]("example 3",1,hIndex,[]int{3})
}
