/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */
package main
import (
	assert "reply2future.com/utils"
)
// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	return insertBruteForce(intervals, newInterval)
}

// #1 brute force
func insertBruteForce(intervals [][]int, newInterval []int) [][]int {
	// insert situations:
	sOverlaps, startIndex := binarySearchIndex(intervals, 0, newInterval[0])
	eOverlaps, endIndex := binarySearchIndex(intervals[startIndex:], startIndex, newInterval[1])

	ret := make([][]int, 0)
	var rightValue int
	var leftValue int
	var lastStartIndex int
	if eOverlaps {
		rightValue = intervals[endIndex][1]
		lastStartIndex = endIndex + 1
	} else {
		rightValue = newInterval[1]
		lastStartIndex = endIndex
	}
	if sOverlaps {
		leftValue = intervals[startIndex][0]
	} else {
		leftValue = newInterval[0]
	}

	ret = append(ret, intervals[0:startIndex]...)
	ret = append(ret, []int{leftValue, rightValue})
	ret = append(ret, intervals[lastStartIndex:]...)

	return ret
}

func binarySearchIndex(intervals [][]int, offset int, n int) (bool, int) {
	iLen := len(intervals)
	if iLen == 0 {
		return false, offset
	}

	mid := iLen / 2
	
	midValue := intervals[mid]
	// insert situations:
	// 1. no overlaps, new elements:
	//   mid: e(i) < s(n), e(n) < s(i+1)
	//   front: e(n) < s(0)
	//   end: s(n) > e(len - 1)
	// 2. overlaps
	//   overlapsIndices -> remove elements
	//   startIndex: e(i) < s(n) < s(i+1) || s(n) < s(0) || s(i) < s(n) < e(i)
	//   endIndex: same above
	if midValue[0] <= n && midValue[1] >= n {
		return true, mid + offset
	}
	if midValue[0] > n {
		return binarySearchIndex(intervals[0:mid], offset, n)
	}
	if midValue[1] < n {
		return binarySearchIndex(intervals[mid + 1:], offset + mid + 1, n)
	}
	// un-arrived
	return false, -1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
func main() {
	assert.AssertEqual2[[][]int, [][]int, []int]("[[1,3],[6,9]] + [2,5]", [][]int{{1,5},{6,9}}, insert, [][]int{{1,3},{6,9}}, []int{2,5})
	assert.AssertEqual2[[][]int, [][]int, []int]("[[1,2],[3,5],[6,7],[8,10],[12,16]] + [4,8]", [][]int{{1,2},{3,10},{12,16}}, insert, [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}, []int{4,8})
	assert.AssertEqual2[[][]int, [][]int, []int]("[[3,5],[12,15]] + [6,6]", [][]int{{3,5},{6,6},{12,15}}, insert, [][]int{{3,5},{12,15}}, []int{6,6})
	assert.AssertEqual2[[][]int, [][]int, []int]("[[1,5]] + [6,8]", [][]int{{1,5},{6,8}}, insert, [][]int{{1,5}}, []int{6,8})
}
