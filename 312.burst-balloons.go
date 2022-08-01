/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 */
package main

import (
	"strconv"

	utils "reply2future.com/utils"
)

// @lc code=start
// num.length [1,300]
func maxCoins(nums []int) int {
	// return recursiveMaxCoins(nums)
	return dcMaxCoins(nums)
}

// #1 recursive =================================================================
// time limited
func numsToString(nums []int) string {
	ret := ""
	for i, v := range nums {
		if i != 0 {
			ret += ("," + strconv.Itoa(v))
		} else {
			ret += strconv.Itoa(v)
		}
	}
	return ret
}

func recursiveMaxCoins(nums []int) int {
	memo := make(map[string]int)

	var fn func([]int) int
	fn = func(nums []int) int {
		nLen := len(nums)
		if nLen == 0 {
			return 0
		}

		if v, ok := memo[numsToString(nums)]; ok {
			return v
		}

		maxCoins := 0
		for i := 0; i < nLen; i++ {
			coins := nums[i]
			nextNums := make([]int, nLen-1)
			if i > 0 {
				coins *= nums[i-1]
				copy(nextNums, nums[0:i])
			}
			if i < nLen-1 {
				coins *= nums[i+1]
				copy(nextNums[i:], nums[i+1:])
			}

			nextStr := numsToString(nextNums)
			if v, ok := memo[nextStr]; ok {
				coins += v
			} else {
				memo[nextStr] = recursiveMaxCoins(nextNums)
				coins += memo[nextStr]
			}

			if coins > maxCoins {
				maxCoins = coins
			}
		}

		return maxCoins
	}
	return fn(nums)
}

// #1 ===========================================================================

// #2 divide and conquer ========================================================
// init the array with i...k...j, to add head and tail with `1`, and it doesn't need to care about the burst balloons
type DcPosition struct {
	Left  int
	Right int
}

func dcMaxCoins(nums []int) int {
	nLen := len(nums)
	models := make([]int, nLen+2)
	models[0] = 1
	copy(models[1:nLen+1], nums)
	models[nLen+1] = 1

	memo := make(map[DcPosition]int)
	var fn func(int, int) int
	fn = func(left int, right int) int {
		if left+1 >= right {
			return 0
		}
		pos := DcPosition{Left: left, Right: right}
		if v, ok := memo[pos]; ok {
			return v
		}

		maxCoins := 0
		for i := left + 1; i < right; i++ {
			maxCoins = max(maxCoins, models[left]*models[i]*models[right]+fn(left, i)+fn(i, right))
		}
		memo[pos] = maxCoins
		return maxCoins
	}

	return fn(0, nLen+1)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// #2 ===========================================================================

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
func main() {
	utils.AssertEqual1[int, []int]("example 1", 167, maxCoins, []int{3, 1, 5, 8})
	utils.AssertEqual1[int, []int]("example 2", 10, maxCoins, []int{1, 5})
	utils.AssertEqual1[int, []int]("example 3", 1717, maxCoins, []int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3, 3})
}
