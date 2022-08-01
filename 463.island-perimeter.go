/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func islandPerimeter(grid [][]int) int {
	return bruteForce463(grid)    
}

// #1 brute force =================================================================
func bruteForce463(grid [][]int) int {
	ret := 0
	rLen := len(grid)
	cLen := len(grid[0])
	for i, rows := range grid {
		for j, col := range rows {
			if col == 0 {
				continue
			}

			if j == 0 || rows[j-1] == 0 {
				ret++
			}
			if j == cLen-1 || rows[j+1] == 0 {
				ret++
			}
			if i == 0 || grid[i-1][j] == 0 {
				ret++
			}
			if i == rLen-1 || grid[i+1][j] == 0 {
				ret++
			}
		}
	}
	return ret
}
// #1 =============================================================================
// @lc code=end
func main() {
	utils.AssertEqual1[int,[][]int]("example 1",16,islandPerimeter,[][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}})
	utils.AssertEqual1[int,[][]int]("example 2",4,islandPerimeter,[][]int{{1}})
	utils.AssertEqual1[int,[][]int]("example 3",4,islandPerimeter,[][]int{{1,0}})
}
