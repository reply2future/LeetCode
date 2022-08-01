/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
package main

import (
	assert "reply2future.com/utils"
	"fmt"
	"strings"
	"time"
)

// @lc code=start
func minLen(p string) int {
	sum := 0
	for _, c := range p {
		if c != '*' {
			sum += 1
		}
	}
	return sum
}

func isMatch(s string, p string) bool {
	// p simplify
	// pMinLen > sLen return false
	// tail match first
	p = simplifyPattern(p)

	if minLen(p) > len(s) {
		return false
	}

	if !splitSubMatch(s, p) {
		return false
	}

	if !tailMatch(s, p) {
		return false
	}

	// return isMatchRecursive(s, p)
	return isMatchDp(s, p)
}

func tailMatch(s string, p string) bool {
	pLen := len(p)
	sLen := len(s)
	for i := 1; i <= pLen; i++ {
		if p[pLen-i] == '*' {
			return true
		}

		if p[pLen-i] != s[sLen-i] && p[pLen-i] != '?' {
			return false
		}
	}
	return true
}

func splitSubMatch(s string, p string) bool {
	ps := strings.FieldsFunc(p, func(r rune) bool {
		return r == '*' || r == '?'
	})
	for _, c := range ps {
		if len(c) != 0 && !strings.Contains(s, c) {
			return false
		}
	}
	return true
}

func simplifyPattern(p string) string {
	var ret []rune
	for i, c := range p {
		if c != '*' {
			ret = append(ret, c)
		} else {
			if i == 0 || p[i-1] != '*' {
				ret = append(ret, c)
			}
		}
	}
	return (string)(ret)
}

// #1 recursive
func isMatchRecursive(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	if sLen == 0 {
		return minLen(p) == 0
	}

	if pLen == 0 {
		return false
	}

	if p[0] == '*' {
		return isMatchRecursive(s, p[1:]) || isMatchRecursive(s[1:], p)
	} else {
		return (s[0] == p[0] || p[0] == '?') && isMatchRecursive(s[1:], p[1:])
	}
}

// #2 dp
func isMatchDp(s string, p string) bool {
	memo := make(map[string]map[string]bool)
	var _isMatchDp func(s string, p string) bool
	_isMatchDp = func(s string, p string) bool {
		sLen := len(s)
		pLen := len(p)
		if sLen == 0 {
			return minLen(p) == 0
		}

		if pLen == 0 {
			return false
		}

		match, ok := memo[s]

		if ok {
			m, pOk := match[p]
			if pOk {
				return m
			}
		} else {
			memo[s] = make(map[string]bool)
		}

		if p[0] == '*' {
			memo[s][p] = _isMatchDp(s, p[1:]) || _isMatchDp(s[1:], p)
		} else {
			memo[s][p] = (s[0] == p[0] || p[0] == '?') && _isMatchDp(s[1:], p[1:])
		}
		return memo[s][p]
	}
	return _isMatchDp(s, p)
}

// @lc code=end
func main() {
	fmt.Println(isMatch("aa", "a"), "false")
	fmt.Println(isMatch("aa", "*"), "true")
	fmt.Println(isMatch("ab", "?*"), "true")
	fmt.Println(isMatch("cb", "?a"), "false")
	fmt.Println(isMatch("adceb", "*a*b"), "true")
	st := time.Now()
	fmt.Println(isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"), "false")
	et := time.Now()
	fmt.Println("use time:", et.Sub(st))
	st = time.Now()
	// *aa*ba*a*bb*aa*ab*a*aaaaaa*a*aaaa*bbabb*b*b*aaaaaaaaa*a*ba*bbb*a*ba*bb*bb*a*b*bb
	fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
		"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"), "false")
	et = time.Now()
	fmt.Println("use time:", et.Sub(st))
	st = time.Now()
	// *bbb*a*abb*b*a*bbbbaab*b*aaba*a*b*a*abb*aa*b*bb*abbbb*b*bbbabaa*b*ba*a*ba*b*a*a*aaa
	fmt.Println(isMatch("aaaaaabbaabaaaaabababbabbaababbaabaababaaaaabaaaabaaaabababbbabbbbaabbababbbbababbaaababbbabbbaaaaaaabbaabbbbababbabbaaabababaaaabaaabaaabbbbbabaaabbbaabbbbbbbaabaaababaaaababbbbbaabaaabbabaabbaabbaaaaba",
		"*bbb**a*******abb*b**a**bbbbaab*b*aaba*a*b**a*abb*aa****b*bb**abbbb*b**bbbabaa*b**ba**a**ba**b*a*a**aaa"), "false")
	et = time.Now()
	fmt.Println("use time:", et.Sub(st))
	st = time.Now()
	// *a*b*b*b*b*bb*b*babaab*ba*a*aaa*baa*b*bbbb*bbaa*a*a*a*a*b*a*a*ba*aa*a*a*
	fmt.Println(isMatch("aaaabaabaabbbabaabaabbbbaabaaabaaabbabbbaaabbbbbbabababbaabbabbbbaababaaabbbababbbaabbbaabbaaabbbaabbbbbaaaabaaabaabbabbbaabababbaabbbabababbaabaaababbbbbabaababbbabbabaaaaaababbbbaabbbbaaababbbbaabbbbb",
		"**a*b*b**b*b****bb******b***babaab*ba*a*aaa***baa****b***bbbb*bbaa*a***a*a*****a*b*a*a**ba***aa*a**a*"), "false")
	et = time.Now()
	fmt.Println("use time:", et.Sub(st))
	assert.AssertEqual2[bool, string, string]("aa -> a", false, isMatch, "aa", "a")
}
