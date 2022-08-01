/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 */
package main
import (
	utils "reply2future.com/utils"
	"sort"
)
// @lc code=start
// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000
// words[i] consists only of lowercase English letters.
// 
// two words dont share common letters
// max of length(word[i]) * length(word[j])
func maxProduct(words []string) int {
	// return maxProductDirect(words)    
	return maxProductBitOpt(words)    
}

// #1 direct =================================================================
func maxProductDirect(words []string) (ret int) {
	cache := make([]map[byte]interface{}, len(words))

	for i, word := range words {
		cache[i] = make(map[byte]interface{})
		wbs := []byte(word)
		for _, wb := range wbs {
			cache[i][wb] = struct{}{}
		}
	}

	for i := 0; i < len(words) - 1; i++ {
		iLen := len(words[i])
		iwbs := cache[i]
		for j := i + 1; j < len(words); j++ {
			jLen := len(words[j])
			nMax := iLen * jLen
			if nMax <= ret {
				continue
			}

			jwbs := cache[j]
			match := true 
			for key, _ := range iwbs {
				if _, ok := jwbs[key]; ok {
					match = false
					break
				}
			}

			if !match {
				continue
			}

			ret = nMax
		}
	}

	return
}
// #1 ========================================================================
// #2 bit operation ==========================================================
func maxProductBitOpt(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	alpha := make([]int, len(words))
	for i, w := range words {
		for _, c := range w {
			alpha[i] = alpha[i] | (0x1 << uint(c-'a'))
		}
	}
	max := 0
	for i := 0; i < len(words); i++ {
		if max >= len(words[i])*len(words[i]) {
			break
		}
		for j := i + 1; j < len(words); j++ {
			if max >= len(words[i])*len(words[j]) {
				break
			}
			if alpha[i]&alpha[j] != 0 {
				continue
			}
			max = len(words[i]) * len(words[j])
			break
		}
	}
	return max
}
// #2 ========================================================================

// @lc code=end
func main() {
	utils.AssertEqual1[int,[]string]("example 1",16,maxProduct,[]string{"abcw","baz","foo","bar","xtfn","abcdef"})
	utils.AssertEqual1[int,[]string]("example 2",4,maxProduct,[]string{"a","ab","abc","d","cd","bcd","abcd"})
	utils.AssertEqual1[int,[]string]("example 3",0,maxProduct,[]string{"a","aa","aaa","aaaa"})
}
