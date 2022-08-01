/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */
package main

import (
	"math"

	utils "reply2future.com/utils"
)

// @lc code=start

func numSquares(n int) int {
	// return numSquaresDp(n)
	return numSquaresBfs(n)
}

// #1 dp =================================================================
func getMaxRoot(n int) int {
	v := math.Sqrt(float64(n))
	return int(v)
}

type NKey struct {
	n       int
	maxRoot int
}

func numSquaresDp(n int) int {
	memo := make(map[NKey]int)

	var numSquaresRecursive func(int) int
	numSquaresRecursive = func(n int) int {
		if n == 0 {
			return 0
		}
		maxRoot := getMaxRoot(n)
		if maxRoot*maxRoot == n {
			return 1
		}
		if maxRoot == 1 {
			return n
		}

		nkey := NKey{n, maxRoot}
		v, ok := memo[nkey]
		if ok {
			return v
		}

		min := -1
		var count int
		for i := maxRoot; i >= 1; i-- {
			i2 := i * i
			if i2 > n {
				continue
			}
			count = numSquaresRecursive(n - i2)
			if min == -1 || count+1 < min {
				min = count + 1
			}
		}

		memo[nkey] = min
		return min
	}

	return numSquaresRecursive(n)
}

// #1 =================================================================

// #2 bfs =============================================================
func numSquaresBfs(n int) int {
	var numPerfectSquares int
	var squares []int

	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	queue := []int{n}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		numPerfectSquares++
		currentLength := len(queue)

		for _, popped := range queue {
			for _, square := range squares {
				if popped == square {
					return numPerfectSquares
				}

				remainder := popped - square
				if remainder > 0 && !visited[remainder] {
					visited[remainder] = true
					queue = append(queue, remainder)
				}
			}
		}

		queue = queue[currentLength:]
	}

	return -1
}

// #2 =================================================================

// @lc code=end
func main() {
	utils.AssertEqual1[int, int]("example 1", 3, numSquares, 12)
	utils.AssertEqual1[int, int]("example 2", 2, numSquares, 13)
	utils.AssertEqual1[int, int]("example 3", 3, numSquares, 329)
	utils.AssertEqual1[int, int]("example 4", 3, numSquares, 7115)
}
