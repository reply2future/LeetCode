/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 */
package main

import (
	"math"
	utils "reply2future.com/utils"
	"sort"
)

// @lc code=start
func wiggleSort(nums []int) {
	// bruteForce324(nums)
	sort324(nums)
}

// stackoverflow
// #1 bruteForce =================================================================
func bruteForce324(nums []int) {
	var travel func(rest []int, before int, idx int) ([]int, bool)
	travel = func(rest []int, before int, idx int) ([]int, bool) {
		rLen := len(rest)
		if rLen == 0 {
			return []int{}, true
		}

		for i := 0; i < rLen; i++ {
			if (idx%2 == 0 && rest[i] >= before) || (idx%2 == 1 && rest[i] <= before) {
				continue
			}

			newRest := make([]int, 0)
			newRest = append(newRest, rest[:i]...)
			newRest = append(newRest, rest[i+1:]...)

			if val, ok := travel(newRest, rest[i], idx+1); ok {
				return append([]int{rest[i]}, val...), true
			}
		}

		return nil, false
	}

	ret, ok := travel(nums, math.MaxInt, 0)
	if !ok {
		panic("there is no valid result")
	}

	copy(nums, ret)
}

// #1 bruteForce =================================================================

// #2 sort =================================================================
// 1. sort the slice by asc order
// 2. split two piece
//   1) like [1, 2] and [3, 4]
//   2) like [1, 2, 3] and [4, 5]
// 3. let right slice into left slice, like [1, 3, 2, 4]
// 4. if there is duplication key like [1, 2, 2, 4] => merge [2, 4] [1, 2] => [2,4,1,2]
func sort324(nums []int) {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	nLen := len(nums)
	ret := make([]int, nLen)

	lIdx := 0
	rIdx := 0
	remainder := nLen % 2
	quotient := nLen / 2
	mid := quotient + remainder
	left := nums[:mid]
	right := nums[mid:]

	duplicateIdx := -1

	for i := 0; i < nLen; i++ {
		if i%2 == 0 {
			ret[i] = left[lIdx]
			lIdx++
		} else {
			ret[i] = right[rIdx]
			rIdx++
		}
		if i > 0 && ret[i] == ret[i-1] {
			duplicateIdx = i
		}
	}

	if duplicateIdx != -1 {
		copy(nums, ret[duplicateIdx:])
		copy(nums[nLen-duplicateIdx:], ret[:duplicateIdx])
	} else {
		copy(nums, ret)
	}
}

// #2 sort =================================================================

// @lc code=end
func main() {
	compareFn := func(nums []int) []int {
		wiggleSort(nums)
		return nums
	}

	// for #1
	// utils.AssertEqual1[[]int, []int]("example 1", []int{1, 5, 1, 6, 1, 4}, compareFn, []int{1, 5, 1, 1, 6, 4})
	// utils.AssertEqual1[[]int, []int]("example 2", []int{1, 3, 2, 3, 1, 2}, compareFn, []int{1, 3, 2, 2, 3, 1})

	// for #2
	utils.AssertEqual1[[]int, []int]("example 1", []int{1, 4, 1, 5, 1, 6}, compareFn, []int{1, 5, 1, 1, 6, 4})
	utils.AssertEqual1[[]int, []int]("example 2", []int{1, 2, 1, 3, 2, 3}, compareFn, []int{1, 3, 2, 2, 3, 1})
	utils.AssertEqual1[[]int, []int]("example 3", []int{5, 6, 5, 6, 5, 6, 4, 5}, compareFn, []int{4, 5, 5, 5, 5, 6, 6, 6})
}
