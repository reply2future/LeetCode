/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func wordBreak(s string, wordDict []string) []string {
    if len(s) == 0 {
		return []string{}
	}
	return dpWordBreak2(s, wordDict)
}
type Record struct {
	IsMatch bool
	Matches []string
}
// #1 dp
func dpWordBreak2(s string, wordDict []string) []string {
	memo := make(map[string]Record)
	var r func(string, []string) Record
	r = func(s string, wordDict []string) Record {
		sLen := len(s)
		if sLen == 0 {
			return Record{IsMatch: true, Matches: []string{}}
		}

		record, exist := memo[s]
		if exist {
			return record 
		}

		var matches []string
		for _, v := range wordDict {
			vLen := len(v)
			if vLen > sLen || s[:vLen] != v {
				continue
			}

			record := r(s[vLen:], wordDict)
			if !record.IsMatch {
				continue
			}

			if len(record.Matches) == 0 {
				matches = append(matches, v)
				continue
			}

			for _, m := range record.Matches {
				matches = append(matches, v + " " + m)
			}
		}
		memo[s] = Record{IsMatch: len(matches) > 0, Matches: matches}
		return memo[s]
	}
	record := r(s, wordDict)
	if !record.IsMatch {
		return []string{}
	}
	return record.Matches
}
// @lc code=end
func main() {
	utils.AssertEqual2[[]string,string,[]string]("catsanddog",[]string{"cat sand dog","cats and dog"},wordBreak,"catsanddog",[]string{"cat","cats","and","sand","dog"})
	utils.AssertEqual2[[]string,string,[]string]("pineapplepenapple",[]string{"pine apple pen apple","pine applepen apple","pineapple pen apple"},wordBreak,"pineapplepenapple",[]string{"apple","pen","applepen","pine","pineapple"})
	utils.AssertEqual2[[]string,string,[]string]("catsandog",[]string{},wordBreak,"catsandog",[]string{"cats","dog","sand","and","cat"})
}
