/*
 * @lc app=leetcode id=521 lang=golang
 *
 * [521] Longest Uncommon Subsequence I
 */
package main
import (
	utils "reply2future.com/utils"
	"fmt"
)
// @lc code=start
func findLUSlength(a string, b string) int {
    // return bruteForce521(a, b)
    return clear521(a, b)
}
// #1 brute force =================================================================
func bruteForce521(a string, b string) int {
	if a == b {
		return -1
	}

	aLen := len(a)
	bLen := len(b)

	if aLen > bLen {
		return aLen
	} else if aLen < bLen {
		return bLen
	}

	for i := aLen; i > 0; i-- {
		bigSubsequence := getAllSubsequences(a, i)
		smallSubsequence := getAllSubsequences(b, i)

		if containUncommonSubsequence(bigSubsequence, smallSubsequence) {
			return i
		}
	}

	return -1
}

func containUncommonSubsequence(big []string, small []string) bool {
	for _, b := range big {
		found := false 
		for _, s := range small {
			if b == s {
				found = true
				break
			}
		}
		if !found {
			return true
		}
	}

	for _, s := range small {
		found := false 
		for _, b := range big {
			if b == s {
				found = true
				break
			}
		}
		if !found {
			return true
		}
	}

	return false 
}

func getAllSubsequences(s string, l int) (ret []string) {
	ret = make([]string, 0)

	sLen := len(s)
	if sLen < l {
		return
	}

	if l == 0 {
		return append(ret, "")
	}

	for i := 0; i < sLen && sLen - i >= l; i++ {
		subRet := getAllSubsequences(s[i+1:], l-1)
		for _, v := range subRet {
			ret = append(ret, fmt.Sprintf("%c%v", s[i], v))
		}
	}

	return
}
// #1 brute force =================================================================

// #2 clear =================================================================
func clear521(a string, b string) int {
	if a == b {
		return -1
	}
	aLen := len(a)
	bLen := len(b)
	if aLen > bLen {
		return aLen
	}
	return bLen
}
// #2 clear =================================================================

// @lc code=end
func main() {
	utils.AssertEqual2[int,string,string]("example 1",3,findLUSlength,"aba","cdc")
	utils.AssertEqual2[int,string,string]("example 2",3,findLUSlength,"aaa","bbb")
	utils.AssertEqual2[int,string,string]("example 3",-1,findLUSlength,"aaa","aaa")
	utils.AssertEqual2[int,string,string]("example 4",17,findLUSlength,"aefawfawfawfaw","aefawfeawfwafwaef")
}
