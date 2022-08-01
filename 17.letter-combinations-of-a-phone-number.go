/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package main

import "fmt"

// @lc code=start
var ALP_MAPS = map[byte][]string{
	'2': {"a", "b", "c"},
	'3' : {"d", "e", "f"},
	'4' : {"g", "h", "i"},
	'5' : {"j", "k", "l"},
	'6' : {"m", "n", "o"},
	'7' : {"p", "q", "r", "s"},
	'8'  : {"t", "u", "v"},
	'9'  : {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	ret := []string{}
	digitsLen := len(digits)
	if digitsLen == 0 {
		return ret
	}

	alps, ok := ALP_MAPS[digits[0]]
	if !ok {
		panic("Input error digits")
	}

	if digitsLen == 1 {
		return alps
	}

	for _, r := range alps {
		rest := letterCombinations(digits[1:])
		for _, _r := range rest {
			ret = append(ret, r + _r)
		}
	}
	return ret
}

// @lc code=end

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
}