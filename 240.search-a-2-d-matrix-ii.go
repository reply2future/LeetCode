/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	rLen := len(matrix)
	cLen := len(matrix[0])
	if matrix[0][0] > target || matrix[rLen-1][cLen-1] < target {
		return false
	}

	// return searchMatrixRecursive(matrix, target)
	return searchMatrixDirections(matrix, target)
}

// #1 recursive =================================================================
type MatrixKey struct {
	i int
	j int
}

func searchMatrixRecursive(matrix [][]int, target int) bool {
	memo := make(map[MatrixKey]bool)
	rLen := len(matrix)
	cLen := len(matrix[0])
	var visit func([][]int, int, int, int) bool
	visit = func(matrix [][]int, target int, i int, j int) bool {
		if i > rLen-1 || j > cLen-1 {
			return false
		}

		res, ok := memo[MatrixKey{i: i, j: j}]
		if ok {
			return res
		}

		if matrix[i][j] == target {
			memo[MatrixKey{i: i, j: j}] = true
			return true
		} else {
			memo[MatrixKey{i: i, j: j}] = false 
			if matrix[i][j] > target {
				return visit(matrix, target, i+1, j)
			} else {
				return visit(matrix, target, i, j+1) || visit(matrix, target, i+1, j)
			}
		}
	}

	return visit(matrix, target, 0, 0)
}

// #1 ===========================================================================

// #2 direction ===========================
func searchMatrixDirections(matrix [][]int, target int) bool {
	rLen := len(matrix)
	cLen := len(matrix[0])

	i := 0
	j := cLen - 1

	for i < rLen && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}
// #2 =====================================
// @lc code=end
func main() {
	utils.AssertEqual2[bool, [][]int, int]("example 1", true, searchMatrix, [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5)
	utils.AssertEqual2[bool, [][]int, int]("example 2", true, searchMatrix, [][]int{
		{1, 3, 4, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{4, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 6)
	utils.AssertEqual2[bool, [][]int, int]("example 3", false, searchMatrix, [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 20)
}
