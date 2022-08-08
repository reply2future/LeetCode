/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */
package main

import (
	"math"
	utils "reply2future.com/utils"
	"strconv"
)

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 || n == 1 {
		return int(math.Pow10(n))
	}

	// return bruteForce357(n)
	return arthemetic357(n)
}

// #1 brute force ==============================
func bruteForce357(n int) int {
	nLen := int(math.Pow10(n))
	uniques := 0

	for i := 0; i < nLen; i++ {
		bytes := []byte(strconv.Itoa(i))
		bLen := len(bytes)
		memo := make(map[byte]int)
		for m, b := range bytes {
			if memo[b] == 1 {
				rLen := bLen-1-m
				rUniques := int(math.Pow10(rLen))
				uniques += rUniques
				
				i += rUniques - 1
				break
			}
			memo[b] = 1
		}
	}

	return nLen - uniques
}
// #1 brute force ==============================

// @lc code=end
func main() {
	utils.AssertEqual1[int, int]("example 1", 91, countNumbersWithUniqueDigits, 2)
	utils.AssertEqual1[int, int]("example 2", 1, countNumbersWithUniqueDigits, 0)
	utils.AssertEqual1[int, int]("example 3", 2345851, countNumbersWithUniqueDigits, 8)
}
