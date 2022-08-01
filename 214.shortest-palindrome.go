/*
 * @lc app=leetcode id=214 lang=golang
 *
 * [214] Shortest Palindrome
 */
package main

import (
	utils "reply2future.com/utils"
	"strings"
)

// @lc code=start
func shortestPalindrome(s string) string {
	// return shortestPalindromeRecursive(s)
	return shortestPalindromeDirection(s)
}

// #1 recursive
func shortestPalindromeRecursive(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return ""
	}

	midIndex := sLen / 2
	remainder := sLen % 2

	if isReverseString(s[:midIndex], s[midIndex+remainder:]) {
		return s
	}
	
	var sb strings.Builder
	letter := s[sLen-1]
	sb.WriteByte(letter)
	sb.WriteString(shortestPalindrome(s[:sLen-1]))
	sb.WriteByte(letter)
	return sb.String()
}

func shortestPalindromeDirection(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return ""
	}

	cLen := sLen
	for {
		midIndex := cLen / 2
		remainder := cLen % 2

		if cLen == 1 || isReverseString(s[:midIndex], s[midIndex+remainder:cLen]) {
			break
		}
		cLen -= 1
	}

	return reverseString(s[cLen:]) + s
}

func reverseString(s string) string {
	runes := []rune(s)
	sLen := len(runes)
	for i, j := 0, sLen-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isReverseString(s1 string, s2 string) bool {
	sLen := len(s1)
	if sLen != len(s2) {
		return false
	}

	for i, char := range s1 {
		if char != rune(s2[sLen-i-1]) {
			return false
		}
	}
	return true
}

// @lc code=end
func main() {
	utils.AssertEqual1[string, string]("aacecaaa", "aaacecaaa", shortestPalindrome, "aacecaaa")
	utils.AssertEqual1[string, string]("abcd", "dcbabcd", shortestPalindrome, "abcd")
}
