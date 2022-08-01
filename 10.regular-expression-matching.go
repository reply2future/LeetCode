/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
package main 
import (
	"fmt"
	"time"
)
// @lc code=start
type RegExpType int

const (
	// ab, cd
	NORMAL_STR 			RegExpType = 0
	// .*
	QUANTIFIERS_ANY 	RegExpType = 1
	// a*, b*
	QUANTIFIERS_ALP 	RegExpType = 2
	// .
	SINGLE_DOT			RegExpType = 3
)

type RegExpObj struct {
	Value				byte
	Type 				RegExpType
	Indexes				[]int
	isFixed 			bool
}

func (regExpObj RegExpObj) isFull() bool {
	return (regExpObj.Type == NORMAL_STR || regExpObj.Type == SINGLE_DOT) && len(regExpObj.Indexes) > 0
}

func (regExpObj RegExpObj) isMatch(r byte) bool {
	switch (regExpObj.Type) {
	case NORMAL_STR:
		return r == regExpObj.Value
	case QUANTIFIERS_ALP:
		return r == regExpObj.Value
	case SINGLE_DOT:
		return true
	case QUANTIFIERS_ANY:
		return true
	default:
		return false
	}
}

func parse(p string) (error, []RegExpObj) {
	var ret []RegExpObj
	for i := 0; i < len(p); i++ {
		if i + 1 < len(p) && p[i + 1] == '*' {
			if p[i] == '.' {
				ret = append(ret, RegExpObj{
					Value: p[i],
					Type: QUANTIFIERS_ANY,
					isFixed: false,
				})
			} else {
				ret = append(ret, RegExpObj{
					Value: p[i],
					Type: QUANTIFIERS_ALP,
					isFixed: false,
				})
			}
			i++
		} else if p[i] == '.' {
			ret = append(ret, RegExpObj{
				Value: p[i],
				Type: SINGLE_DOT,
				isFixed: true,
			})
		} else {
			ret = append(ret, RegExpObj{
				Value: p[i],
				Type: NORMAL_STR,
				isFixed: true,
			})
		}
	}
	return nil, ret
}

func simplify(rs []RegExpObj) (ret []RegExpObj) {
	// merge duplication regexp, such a*b*c*.*d* => .*
	for _, r := range rs {
		switch (r.Type) {
		case SINGLE_DOT:
			ret = append(ret, r)
		case NORMAL_STR:
			ret = append(ret, r)
		case QUANTIFIERS_ALP:
			retLen := len(ret)
			if retLen > 0 && ((ret[retLen - 1].Type == QUANTIFIERS_ALP && ret[retLen - 1].Value == r.Value) || (ret[retLen - 1].Type == QUANTIFIERS_ANY)) {
				continue
			}
			ret = append(ret, r)
		case QUANTIFIERS_ANY:
			for i := len(ret) - 1; i >= 0; i-- {
				if ret[i].Type == SINGLE_DOT || ret[i].Type == NORMAL_STR {
					break	
				}
				ret = ret[:i]
			}
			ret = append(ret, r)
		default:
			panic("RegExpObj.Type error")
		}
	}
	return ret
}

func deriveLen(seq []RegExpObj) (bool, int) {
	isFixed := true
	len := 0

	for _, seqItem := range seq {
		if !seqItem.isFixed {
			isFixed = false 
		} else {
			len++
		}
	}
	return isFixed, len
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clearIndexes(seq []RegExpObj) {
	for i, seqItem := range seq {
		if len(seqItem.Indexes) > 0 {
			seq[i].Indexes = []int{}
		}
	}
}

func isSeqMatch(s string, seq []RegExpObj) bool {
	sCursor := 0
	seqCursor := 0
	sLen := len(s)
	seqLen := len(seq)

	if sLen == 0 && seqLen == 0 {
		return true
	}

	if (sLen == 0 && seqLen != 0) || (sLen != 0 && seqLen == 0) {
		return false
	}

	// TODO: .* at the end of seq, return true
	
	for {

		// s to the end by seq has something not match yet
		if sCursor >= sLen {
			_, dLen := deriveLen(seq[min(seqCursor, seqLen):])
			return dLen <= 0
		}

		if seqCursor < seqLen {
			currentRegExpObj := seq[seqCursor]
			if currentRegExpObj.Type == QUANTIFIERS_ALP || currentRegExpObj.Type == QUANTIFIERS_ANY {
				seqCursor++
				continue
			}

			if currentRegExpObj.isMatch(s[sCursor]) {
				seq[seqCursor].Indexes = append(currentRegExpObj.Indexes, sCursor)
				sCursor++
				seqCursor++
				continue
			}
		}

		match := false
		coveredValue := false
		for i := seqCursor - 1; i >= 0; i-- {
			previousRegExpObj := seq[i]
			if previousRegExpObj.isMatch(s[sCursor]) {
				if previousRegExpObj.isFixed {
					if !coveredValue && previousRegExpObj.isFull() {
						coveredValue = true
					}
					clearIndexes(seq[:i])
					previousMatch := isSeqMatch(s[:sCursor], seq[:i])
					if !previousMatch {
						continue
					}
					seq[i].Indexes = []int{}
				} else if coveredValue {
					clearIndexes(seq[:i + 1])
					previousMatch := isSeqMatch(s[:sCursor], seq[:i + 1])
					if !previousMatch {
						continue
					}
				}
				
				seq[i].Indexes = append(seq[i].Indexes, sCursor)
				
				match = true
				sCursor++
				seqCursor = i + 1
				break
			} else {
				if !coveredValue && len(previousRegExpObj.Indexes) > 0 {
					coveredValue = true
				}	
				seq[i].Indexes = []int{}
			}
		}
		if !match {
			return false
		}
	}
}

func isMatch(s string, p string) bool {
	// parse the p string to object
	// type: normal string, quantifiers, single dot
	// merge quantifiers
	// find sub match string
	e, seq := parse(p)
	if e != nil {
		panic(e)
	}
	// seq = simplify(seq)

	isFixed, seqLen := deriveLen(seq)
	sLen := len(s)
	if isFixed && seqLen != sLen {
		return false
	}
	if !isFixed && sLen < seqLen {
		return false
	}

	return isSeqMatch(s, seq)
}
// @lc code=end

func isMatchTest(s string, p string, expected bool) {
	t1 := time.Now()
	// ret := isMatch(s, p)
	// ret := isMatchRecursion(s, p)
	ret := isMatchDp(s, p)
	t2 := time.Now()
	fmt.Println("use time: ", t2.Sub(t1))
	if ret != expected {
		panic(fmt.Sprintf("s: %s, p: %s, expected: %v, got: %v", s, p, expected, ret))
	}
}

// recursion
func isMatchRecursion(s string, p string) bool {
	pLen := len(p)
	sLen := len(s)
	if pLen == 0 {
		return sLen == 0
	}

	firstMatch := sLen > 0 && (p[0] == s[0] || p[0] == '.')
	if pLen >= 2 && p[1] == '*' {
		return isMatchRecursion(s, p[2:]) || (firstMatch && isMatchRecursion(s[1:], p))
	} else {
		return firstMatch && isMatchRecursion(s[1:], p[1:])
	}
}

// dp
func isMatchDp(s string, p string) bool {
	memo := make(map[string]map[string]struct{})

	var _isMatchRecursion func(s string, p string) bool
	_isMatchRecursion = func (s string, p string) bool {
		pLen := len(p)
		sLen := len(s)
		if pLen == 0 {
			return sLen == 0
		}

		matchMap, ok := memo[s]

		if ok {
			_, ok := matchMap[p]
			if ok {
				return true
			}
		}
	
		firstMatch := sLen > 0 && (p[0] == s[0] || p[0] == '.')
		_, pOk := memo[s]
		if !pOk {
			memo[s] = make(map[string]struct{})
		}
		if pLen >= 2 && p[1] == '*' {
			if _isMatchRecursion(s, p[2:]) || (firstMatch && _isMatchRecursion(s[1:], p)) {
				memo[s][p] = struct{}{}
				return true
			}
			return false
		} else {
			if firstMatch && _isMatchRecursion(s[1:], p[1:]) {
				memo[s][p] = struct{}{}
				return true
			}
			return false
		}
	}

	return _isMatchRecursion(s, p)
}

func main() {
	isMatchTest("aa", "a*", true)
	isMatchTest("aa", "a", false)
	isMatchTest("aa", "a*b", false)
	isMatchTest("aa", "a..", false)
	isMatchTest("aab", "a*.", true)
	isMatchTest("b", "a*.", true)
	isMatchTest("abcdbcd", "a.*bcd", true)
	isMatchTest("abcdbcd", "a.*bcd.*bcd", true)
	isMatchTest("abcdbcd", "a.*bcd.*", true)
	isMatchTest("aa", "a*", true)
	isMatchTest("aa", ".*", true)
	isMatchTest("aabcd", "a*b*cd", true)
	isMatchTest("cd", "a*b*cd", true)
	isMatchTest("mcd", "a*b*cd", false)
	isMatchTest("cf", "a*b*cd*e*f*", true)
	isMatchTest("af", "a*b*cd*e*f*", false)
	isMatchTest("aaae", "a*b*ad*e", true)
	isMatchTest("aeae", "a*b*ad*e", false)

	isMatchTest("aaba", "ab*a*c*a", false)
	isMatchTest("aabcbcbcaccbcaabc", ".*a*aa*.*b*.c*.*a*", true)

	isMatchTest("bcbabcaaca", "a*c*a*b*.*aa*c*", true)
	isMatchTest("bcbabcaacacbcabac", "a*c*a*b*.*aa*c*a*a*", true)

	isMatchTest("aab", "a.*ac*", false)
	isMatchTest("aabcc", "a.*ac*.", false)
	isMatchTest("acbccccacccaabcc", "a.*ac*.", false)
	isMatchTest("acbccccacccaabcc", "ab*.*c*b*ac*c*c*.", false)

	isMatchTest("bab", "b*a*", false)
	isMatchTest("bbab", "b*a*", false)
}
