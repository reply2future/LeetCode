/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */
package main
import (
	utils "reply2future.com/utils"
	"math"
)
// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	// return bruteForce643(nums, k)
	return sliceWindow643(nums, k)
}

// #1 brute force =================================================================
func bruteForce643(nums []int, k int) float64 {
	nLen := len(nums)
	maxAverage := -math.MaxFloat64

	for i := 0; i <= nLen - k; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum += nums[i + j]
		}
		result := float64(sum) / float64(k)
		if result > maxAverage {
			maxAverage = result
		}
	}

	return maxAverage
}
// #1 brute force =================================================================

// #2 slice windows =================================================================
func sliceWindow643(nums []int, k int) float64 {
	nLen := len(nums)
	maxSum := 0
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	
	maxSum = sum

	for pos := 0; pos < nLen - k; pos++ {
		sum -= nums[pos]
		sum += nums[pos+k]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)	
}
// #2 slice windows =================================================================

// @lc code=end
func main() {
	utils.AssertEqual2[float64,[]int,int]("example 1",12.75000,findMaxAverage,[]int{1,12,-5,-6,50,3},4)
	utils.AssertEqual2[float64,[]int,int]("example 2",5.00000,findMaxAverage,[]int{5},1)
	utils.AssertEqual2[float64,[]int,int]("example 3",-1.00000,findMaxAverage,[]int{-1},1)
}
