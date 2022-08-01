/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package main

import (
    "fmt"
    "math"
)

 // @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    _mergedSortedArrays := mergeSortedArrays(nums1, nums2)

	return calMedian(_mergedSortedArrays)
}

func calMedian(nums []int) float64 {
	_len := len(nums)
	if !isOddNum(_len) {
		return (float64)(nums[_len / 2 - 1] + nums[_len / 2]) / 2
	}

	return (float64)(nums[_len / 2])
}

func isOddNum(num int) bool {
	return num % 2 != 0
}

func classifyArrays(nums1 []int, nums2 []int) ([]int, []int) {
	if len(nums1) > len(nums2) {
		return nums1, nums2
	}

	return nums2, nums1
}

func binarySearch(nums []int, n int) int {
	_len := len(nums)
	if _len == 0 {
		return _len
	}

	compareNum := nums[_len / 2]
	if compareNum == n {
		return _len / 2 + 1
	} else if compareNum > n {
		return binarySearch(nums[:_len / 2], n)
	} else {
		return _len / 2 + 1 + binarySearch(nums[_len / 2 + 1:], n)
	}
}

func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}

	if len(nums2) == 0 {
		return nums1
	}
	
	var _retArrays []int

	_biggerNums, _smallerNums := classifyArrays(nums1, nums2)
	_index := len(_biggerNums) / 2
	_compareNum := _biggerNums[_index]

	if _compareNum < _smallerNums[0] {
		_retArrays = append(_retArrays, _biggerNums[:_index + 1]...)

		_mergedSortedArrays := mergeSortedArrays(_biggerNums[_index + 1:], _smallerNums)
		return append(_retArrays, _mergedSortedArrays...)
	} else if _compareNum == _smallerNums[0] {
		_retArrays = append(_retArrays, _biggerNums[:_index + 1]...)
		_retArrays = append(_retArrays, _smallerNums[0])

		_mergedSortedArrays := mergeSortedArrays(_biggerNums[_index + 1:], _smallerNums[1:])
		return append(_retArrays, _mergedSortedArrays...)	
	} else {
		_index = binarySearch(_biggerNums[:_index], _smallerNums[0])

		_retArrays = append(_retArrays, _biggerNums[:_index]...)
		_retArrays = append(_retArrays, _smallerNums[0])

		_mergedSortedArrays := mergeSortedArrays(_biggerNums[_index:], _smallerNums[1:])
		return append(_retArrays, _mergedSortedArrays...)	
	}
}

// @lc code=end

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
// https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2471/Very-concise-O(log(min(MN)))-iterative-solution-with-detailed-explanation
func _findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    N1 := len(nums1)
    N2 := len(nums2)
    if N1 < N2 {
		return _findMedianSortedArrays(nums2, nums1)	// Make sure A2 is the shorter one.
	}     
    lo := 0
	hi := N2 * 2
    for {
		if lo > hi {
			break
		}
        mid2 := (lo + hi) / 2   // Try Cut 2 
        mid1 := N1 + N2 - mid2  // Calculate Cut 1 accordingly
  
		L1 := math.MinInt
		L2 := math.MinInt
		R1 := math.MaxInt
		R2 := math.MaxInt
		if mid1 != 0 {
			L1 = nums1[(mid1-1)/2]
		}
		if mid2 != 0 {
			L2 = nums2[(mid2-1)/2]
		}
		if mid1 != N1 * 2 {
			R1 = nums1[(mid1)/2]
		}
		if mid2 != N2 * 2 {
        	R2 = nums2[(mid2)/2];
		}
        
        if L1 > R2 {
			lo = mid2 + 1
		} else if L2 > R1 {
			hi = mid2 - 1
		} else {
			return (float64)(max(L1,L2) + min(R1, R2)) / 2
		}
    }
    return -1
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(_findMedianSortedArrays(nums1, nums2))
}

