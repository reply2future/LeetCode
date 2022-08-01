/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func reverseVowels(s string) string {
	vowels := map[rune]interface{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},
	}
	var indices []int
	for i, v := range s {
		if _, ok := vowels[v]; ok {
			indices = append(indices, i)
		}
	}

	i, j := 0, len(indices) - 1

	runes := []rune(s)
	for j > i {
		si := indices[i]
		ei := indices[j]
		runes[si], runes[ei] = runes[ei], runes[si]
		i++
		j--
	}

	return string(runes)
}

// @lc code=end
func main() {
	utils.AssertEqual1[string, string]("example 1", "holle", reverseVowels, "hello")
	utils.AssertEqual1[string, string]("example 2", "leotcede", reverseVowels, "leetcode")
	utils.AssertEqual1[string, string]("example 3", "Aa", reverseVowels, "aA")
	utils.AssertEqual1[string, string]("example 4", "race car", reverseVowels, "race car")
}
