/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func findTheDifference(s string, t string) byte {
	ss := make([]int, 26)
	bytes := []byte(s)
	for _, char := range bytes {
		ss[char-'a'] += 1
	}

	bytes = []byte(t)
	for _, char := range bytes {
		if ss[char-'a'] == 0 {
			return char
		} else {
			ss[char-'a'] -= 1
		}
	}
	panic("unreachable")
}

// @lc code=end
func main() {
	utils.AssertEqual2[byte, string, string]("example 1", 'e', findTheDifference, "abcd", "abcde")
	utils.AssertEqual2[byte, string, string]("example 2", 'y', findTheDifference, "", "y")
}
