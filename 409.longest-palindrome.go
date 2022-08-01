/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func longestPalindrome(s string) int {
    // <= 1 single character
	mp := make(map[byte]int)
	for _, c := range s {
		char := byte(c)
		if _, ok := mp[char]; ok {
			mp[char] += 1
		} else {
			mp[char] = 1
		}
	}

	ret := 0
	singled := false
	for _, val := range mp {
		if val % 2 == 0 {
			ret += val
		} else {
			if !singled {
				ret += val
				singled = true
			} else {
				ret += val - 1
			}
		}
	}
	return ret
}
// @lc code=end
func main() {
	utils.AssertEqual1[int,string]("example 1",7,longestPalindrome,"abccccdd")
	utils.AssertEqual1[int,string]("example 2",1,longestPalindrome,"a")
	utils.AssertEqual1[int,string]("example 3",3,longestPalindrome,"ccc")
}
