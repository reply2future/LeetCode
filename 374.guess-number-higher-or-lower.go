/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */
package main

import (
	utils "reply2future.com/utils"
)

var picked int

func guess(num int) int {
	if num > picked {
		return -1
	} else if num < picked {
		return 1
	} else {
		return 0
	}
}

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	start := 1
	end := n

	for {
		if start == end {
			return start
		}
		mid := (start + end) / 2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1:
			start = mid + 1
		case 0:
			return mid
		}
	}
}

// @lc code=end
func main() {
	picked = 6
	utils.AssertEqual1[int, int]("example 1", 6, guessNumber, 10)
	picked = 1
	utils.AssertEqual1[int, int]("example 2", 1, guessNumber, 1)
	picked = 1
	utils.AssertEqual1[int, int]("example 3", 1, guessNumber, 2)
}
