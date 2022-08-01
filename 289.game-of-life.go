/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func gameOfLife(board [][]int) {
	gameOfLifeDirection(board)
}

// #1 direct =================================================================
type Position struct {
	i int
	j int
}

func gameOfLifeDirection(board [][]int) {
	memo := make(map[Position]int)

	rLen := len(board)
	cLen := len(board[0])
	for i, rows := range board {
		for j, v := range rows {
			nl := 0

			for m := i - 1; m <= i+1; m++ {
				for n := j - 1; n <= j+1; n++ {
					if m < 0 || m >= rLen || n < 0 || n >= cLen {
						continue
					}

					if board[m][n] == 1 {
						nl++
					}
				}
			}

			if v == 1 {
				nl--

				if nl < 2 || nl > 3 {
					memo[Position{i:i, j:j}] = 0
				}
			} else if nl == 3 {
				memo[Position{i:i, j:j}] = 1
			}
		}
	}

	for position, v := range memo {
		board[position.i][position.j] = v
	}
}

// ===========================================================================
// @lc code=end

func main() {
	e1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(e1)
	utils.AssertEqual0[[][]int]("example 1", [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}, func() [][]int { return e1 })

	e2 := [][]int{{1, 1}, {1, 0}}
	gameOfLife(e2)
	utils.AssertEqual0[[][]int]("example 2", [][]int{{1, 1}, {1, 1}}, func() [][]int { return e2 })
}
