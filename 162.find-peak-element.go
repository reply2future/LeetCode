/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
// nums[-1] = nums[n] = -∞
func findPeakElement(nums []int) int {
	return binaryFindPeakElement(nums, 0, len(nums) - 1)
	// return directFindPeakElement(nums)
}

// #1 binary search
func binaryFindPeakElement(nums []int, startIndex int, endIndex int) int {

	nLen := len(nums)
	if startIndex > endIndex || startIndex < 0 || endIndex >= nLen {
		return -1
	}

	midIndex := (endIndex + startIndex) / 2
	if (midIndex-1 == -1 || nums[midIndex-1] < nums[midIndex]) && (midIndex+1 == nLen || nums[midIndex+1] < nums[midIndex]) {
		return midIndex
	}

	if endIndex == startIndex {
		return -1
	}

	left := binaryFindPeakElement(nums, startIndex, midIndex-1)
	if left == -1 {
		return binaryFindPeakElement(nums, midIndex+1, endIndex)
	}
	return left
}

// #2 direct
func directFindPeakElement(nums []int) int {
	nLen := len(nums)
	if nLen == 1 {
		return 0
	}

	left, right := 0, nLen-1
	for {
		mid := (left + right) / 2

		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			}
			left = mid + 1
			continue
		}

		if mid == nLen-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			}
			right = mid - 1
			continue
		}

		if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
			return mid
		}

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// @lc code=end
func main() {
	utils.AssertEqual1[int, []int]("[1,2,3,1]", 2, findPeakElement, []int{1, 2, 3, 1})
	utils.AssertEqual1[int, []int]("[1,2,1,3,5,6,4]", 1, findPeakElement, []int{1, 2, 1, 3, 5, 6, 4})
	utils.AssertEqual1[int, []int]("[1,2,3,4,5]", 4, findPeakElement, []int{1, 2, 3, 4, 5})
}
