/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */
package main

import (
	utils "reply2future.com/utils"
	"strings"
)

// @lc code=start
func checkRecord(s string) bool {
	// A 2 total
	// L >=3 consecutive
	// return forceBrute551(s)
	return strings.Count(s, "A") < 2 && strings.Index(s, "LLL") == -1
}

// #1 force brute =================================================================
func forceBrute551(s string) bool {
	A := 0
	L := 0

	for _, ch := range s {
		if ch == 'A' {
			if A == 1 {
				return false
			}

			A++
			L = 0
		} else if ch != 'L' {
			L = 0
		} else {
			if L == 2 {
				return false
			}

			L++
		}
	}

	return true

}

// #1 force brute =================================================================

// @lc code=end
func main() {
	utils.AssertEqual1[bool, string]("example 1", true, checkRecord, "PPALLP")
	utils.AssertEqual1[bool, string]("example 2", false, checkRecord, "PPALLL")
}
