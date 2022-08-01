/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
package main

import (
	assert "reply2future.com/utils"
)

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var ret []int

	rLen := len(matrix)

	if rLen == 0 {
		return []int{}
	}
	if rLen == 1 {
		return matrix[0]
	}

	cLen := len(matrix[0])

	if cLen == 1 {
		for i := 0; i < rLen; i++ {
			ret = append(ret, matrix[i][0])
		}
		return ret
	}

	// top, right, bottom, left
	ret = append(ret, matrix[0][:cLen - 1]...)
	for i := 0; i < rLen - 1; i++ {
		ret = append(ret, matrix[i][cLen - 1])
	}
	for i := cLen - 1; i > 0; i-- {
		ret = append(ret, matrix[rLen - 1][i])
	}
	for i := rLen - 1; i > 0; i-- {
		ret = append(ret, matrix[i][0])
	}

	var inners [][]int
	for i := 1; i < rLen - 1 && cLen > 2; i++ {
		inners = append(inners, matrix[i][1:cLen - 1])
	}

	return append(ret, spiralOrder(inners)...)
}

// @lc code=end
func main() {
	assert.AssertEqual1[[]int, [][]int]("[[1,2,3],[4,5,6],[7,8,9]]", []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	assert.AssertEqual1[[]int, [][]int]("[[1,2,3]]", []int{1, 2, 3}, spiralOrder, [][]int{{1, 2, 3}})
	assert.AssertEqual1[[]int, [][]int]("[[1]]", []int{1}, spiralOrder, [][]int{{1}})
	assert.AssertEqual1[[]int, [][]int]("[[1], [2], [3]]", []int{1, 2, 3}, spiralOrder, [][]int{{1}, {2}, {3}})
	assert.AssertEqual1[[]int, [][]int]("[[1,2], [3,4]]", []int{1, 2, 4, 3}, spiralOrder, [][]int{{1, 2}, {3, 4}})
	assert.AssertEqual1[[]int, [][]int]("[[1,2], [3,4], [5,6]]", []int{1, 2, 4, 6, 5, 3}, spiralOrder, [][]int{{1, 2}, {3, 4}, {5, 6}})
}
