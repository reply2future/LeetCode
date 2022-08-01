/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func numIslands(grid [][]byte) int {
	sum := 0
	m := len(grid)
	n := len(grid[0])

	var visit func(*[][]byte, int, int)
	visit = func(grid *[][]byte, i int, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || (*grid)[i][j] == '0' {
			return
		}

		(*grid)[i][j] = '0'

		visit(grid, i - 1, j)
		visit(grid, i, j - 1)
		visit(grid, i, j + 1)
		visit(grid, i + 1, j)
	}

	for i, row := range grid {
		for j, val := range row {
			if val == '1' {
				visit(&grid, i, j)
				sum += 1
			}
		}
	}

	return sum
}

// @lc code=end
func main() {
	utils.AssertEqual1[int, [][]byte]("example 1", 1, numIslands, [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	})
	utils.AssertEqual1[int, [][]byte]("example 2", 3, numIslands, [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	})
	utils.AssertEqual1[int, [][]byte]("example 3", 0, numIslands, [][]byte{
		{'0'},
	})
	utils.AssertEqual1[int, [][]byte]("example 4", 1, numIslands, [][]byte{
		{'1'},
	})
	utils.AssertEqual1[int, [][]byte]("example 5", 1, numIslands, [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	})
	utils.AssertEqual1[int, [][]byte]("example 6", 1, numIslands, [][]byte{
		{'1', '1', '1'},
		{'0', '0', '1'},
		{'1', '0', '1'},
		{'1', '1', '1'},
	})
}
