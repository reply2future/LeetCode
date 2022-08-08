/*
 * @lc app=leetcode id=327 lang=golang
 *
 * [327] Count of Range Sum
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {
	// return bruteForce327(nums, lower, upper)
	// return binaryIndexedTree327(nums, lower, upper)
	return mergeSort327(nums, lower, upper)
}

// TLE
// #1 brute force =================================================================
type Range327 struct {
	Start int
	End   int
}

func bruteForce327(nums []int, lower int, upper int) int {
	nLen := len(nums)
	sum := make([]int, nLen+1)
	for i, val := range nums {
		sum[i+1] = val + sum[i]
	}

	count := 0
	for i := 0; i < nLen; i++ {
		for j := nLen; j > i; j-- {
			result := sum[j] - sum[i]
			if result < lower || result > upper {
				continue
			}
			count++
		}
	}

	return count
}

// #1 brute force =================================================================

// TLE
// #2 binary indexed tree =================================================================
type BinaryIndexedTree327 []int

func (this BinaryIndexedTree327) lowBit(i int) int { return i & -i }

func build(nums []int) BinaryIndexedTree327 {
	bLen := len(nums) + 1
	bit := make(BinaryIndexedTree327, bLen)
	copy(bit[1:], nums)

	for i := 1; i < bLen; i++ {
		j := i + bit.lowBit(i)
		if j < bLen {
			bit[j] += bit[i]
		}
	}

	return bit
}

func (this BinaryIndexedTree327) prefixSum(idx int) int {
	idx += 1
	result := 0
	for idx > 0 {
		result += this[idx]
		idx -= this.lowBit(idx)
	}

	return result
}

func (this BinaryIndexedTree327) rangeSum(fromIdx int, toIdx int) int {
	return this.prefixSum(toIdx) - this.prefixSum(fromIdx-1)
}

func binaryIndexedTree327(nums []int, lower int, upper int) int {
	bit := build(nums)

	nLen := len(nums)
	count := 0
	for i := 0; i < nLen; i++ {
		for j := nLen - 1; j >= i; j-- {
			result := bit.rangeSum(i, j)
			if result < lower || result > upper {
				continue
			}
			count++
		}
	}

	return count
}

// #2 binary indexed tree =================================================================

// https://leetcode.com/problems/count-of-range-sum/discuss/77990/Share-my-solution
// #3 merge sort =================================================================
func mergeSort327(nums []int, lower int, upper int) int {
	nLen := len(nums)
	sums := make([]int, nLen+1)
	for i, v := range nums {
		sums[i+1] = v + sums[i]
	}
	return countWhileMergeSort(sums, 0, nLen+1, lower, upper)
}

func countWhileMergeSort(sums []int, start int, end int, lower int, upper int) int {
	if end-start <= 1 {
		return 0
	}
	mid := (start + end) / 2
	count := countWhileMergeSort(sums, start, mid, lower, upper) + countWhileMergeSort(sums, mid, end, lower, upper)

	// three elements as the minimum unit: left, mid, right
	j, k, t := mid, mid, mid
	cache := make([]int, end-start)
	for i, r := start, 0; i < mid; i++ {
		for k < end && sums[k]-sums[i] < lower {
			k++
		}
		for j < end && sums[j]-sums[i] <= upper {
			j++
		}
		// like BST, make the sum which is less than sum[i] to the left	
		// because the merge will ignore the ordered of the unit's internal
		for t < end && sums[t] < sums[i] {
			cache[r] = sums[t]
			r++
			t++
		}
		cache[r] = sums[i]
		r++
		
		count += j - k
	}

	copy(sums[start:t], cache)

	return count
}

// #3 merge sort =================================================================

// @lc code=end
func main() {
	utils.AssertEqual3[int, []int, int, int]("example 1", 3, countRangeSum, []int{-2, 5, -1}, -2, 2)
	utils.AssertEqual3[int, []int, int, int]("example 2", 1, countRangeSum, []int{0}, 0, 0)
}
