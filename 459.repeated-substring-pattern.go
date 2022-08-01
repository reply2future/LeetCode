/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */
package main
import (
	utils "reply2future.com/utils"
	"strings"
)
// @lc code=start
func repeatedSubstringPattern(s string) bool {
	// return bruteForce459(s)
	// return kmp459(s)
	return crazy459(s)
}

// #1 brute force =================================================================
func bruteForce459(s string) bool {
	sLen := len(s)
	stack := make([]byte, 0)
	for _, char := range s {
		stack = append(stack, byte(char))
		skLen := len(stack)
		if sLen % skLen != 0 {
			continue
		}

		if sLen < skLen * 2 {
			return false
		}
		repeated := true 
		for j := skLen; j < sLen; j++ {
			if stack[j % skLen] != s[j] {
				repeated = false	
				break	
			}
		}
		if repeated {
			return true
		}
	}

	return false

}
// #1 =============================================================================

// #2 kmp =========================================================================
func kmp459(s string) bool {
	nxt := make([]int, len(s))
    j := 0
    for i := 1; i < len(s); i++ {
        for j > 0 && s[i] != s[j] {
            j = nxt[j-1]
        }
        if s[i] == s[j] {
            j++
        }
        nxt[i] = j
    }
    n := len(s)
    return nxt[n-1] > 0 && n % (n-nxt[n-1]) == 0
}
// #2 =============================================================================

// #3 crazy =======================================================================
func crazy459(s string) bool {
	return strings.Contains(s[1:]+s[:len(s)-1], s)
}
// #3 =============================================================================
// @lc code=end
func main() {
	utils.AssertEqual1[bool,string]("example 1",true,repeatedSubstringPattern,"abab")
	utils.AssertEqual1[bool,string]("example 2",false,repeatedSubstringPattern,"aba")
	utils.AssertEqual1[bool,string]("example 3",true,repeatedSubstringPattern,"abcabcabcabc")
	utils.AssertEqual1[bool,string]("example 4",true,repeatedSubstringPattern,"aaaaa")
	utils.AssertEqual1[bool,string]("example 5",false,repeatedSubstringPattern,"a")
}
