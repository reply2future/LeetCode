/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func isPowerOfThree(n int) bool {
    if n < 0 {
		return false
	} else if n == 0 {
		return false 
	} else if n == 1 {
		return true
	}

	for {
		quotient := n / 3
		remainder := n % 3
		if remainder != 0 {
			return false
		}

		if remainder == 0 && quotient == 1 {
			return true
		}

		n = quotient
	}
}
// @lc code=end
func main() {
	utils.AssertEqual1[bool,int]("example 1",true,isPowerOfThree,27)
	utils.AssertEqual1[bool,int]("example 2",false,isPowerOfThree,0)
	utils.AssertEqual1[bool,int]("example 3",true,isPowerOfThree,9)
	utils.AssertEqual1[bool,int]("example 4",true,isPowerOfThree,1)
}
