/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */
package main

import (
	utils "reply2future.com/utils"
	"strings"
)

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseStr(s string, k int) string {
	sLen := len(s)
	var sb strings.Builder
	quotient := sLen / k
	remainder := sLen % k
	var counter int
	if remainder != 0 {
		counter = quotient + 1
	} else {
		counter = quotient
	}
	for i := 0; i < counter; i++ {
		startIdx := i * k
		endIdx := min(startIdx+k, sLen)
		if i%2 == 0 {
			sb.WriteString(reverse541(s[startIdx:endIdx]))
		} else {
			sb.WriteString(s[startIdx:endIdx])
		}
	}

	return sb.String()
}

func reverse541(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

// @lc code=end
func main() {
	utils.AssertEqual2[string, string, int]("example 1", "bacdfeg", reverseStr, "abcdefg", 2)
	utils.AssertEqual2[string, string, int]("example 2", "bacd", reverseStr, "abcd", 2)
}
