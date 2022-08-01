/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	// return isPowerOfTwoForLoop(n)
	// return isPowerOfTwoRecursive(n)
	return isPowerOfTwoBitOp(n)
}

// #1 for-loop
func isPowerOfTwoForLoop(n int) bool {
	for {
		remainder := n % 2
		if remainder != 0 {
			return false
		}
		quotient := n / 2
		if quotient == 1 || quotient == 2 {
			return true
		}
		n = quotient
	}
}

// #2 recursive
func isPowerOfTwoRecursive(n int) bool {
	if n == 1 || n == 2 {
		return true
	}

	if n%2 != 0 {
		return false
	}
	return isPowerOfTwoRecursive(n / 2)
}

// #3 bit op
func isPowerOfTwoBitOp(n int) bool {
	return n > 0 && (n & (n - 1) == 0)
}

// @lc code=end
func main() {
	utils.AssertEqual1[bool, int]("n = 1", true, isPowerOfTwo, 1)
	utils.AssertEqual1[bool, int]("n = 16", true, isPowerOfTwo, 16)
	utils.AssertEqual1[bool, int]("n = 3", false, isPowerOfTwo, 3)
}
