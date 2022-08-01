/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// return recursionWordBreak(s, wordDict)    // timeout
	return dpWordBreak(s, wordDict)
}
// #1 recursion - timeout
func recursionWordBreak(s string, wordDict []string) bool {
	sLen := len(s)
	if sLen == 0 {
		return true
	}

	for _, word := range wordDict {
		wLen := len(word)
		if sLen < wLen {
			continue
		}

		isMatch := s[:wLen] == word && recursionWordBreak(s[wLen:], wordDict)
		if isMatch {
			return true
		}
	}
	return false
}
// #2 dp
func dpWordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)
	var r func(string, []string) bool
	r = func(s string, wordDict []string) bool {
		sLen := len(s)
		if sLen == 0 {
			return true
		}

		isMatch, exist := memo[s]
		if exist {
			return isMatch
		}

		memo[s] = false
	
		for _, word := range wordDict {
			wLen := len(word)
			if sLen < wLen {
				continue
			}
	
			isMatch := s[:wLen] == word && r(s[wLen:], wordDict)
			if isMatch {
				memo[s] = true
				break
			}
		}

		return memo[s] 
	}

	return r(s, wordDict)
}
// @lc code=end
func main() {
	utils.AssertEqual2[bool,string,[]string]("leetcode",true,wordBreak,"leetcode",[]string{"leet","code"})
	utils.AssertEqual2[bool,string,[]string]("applepenapple",true,wordBreak,"applepenapple",[]string{"apple","pen"})
	utils.AssertEqual2[bool,string,[]string]("catsandog",false,wordBreak,"catsandog",[]string{"cats","dog","sand","and","cat"})
}
