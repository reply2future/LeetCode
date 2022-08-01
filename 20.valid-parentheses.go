/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package main 
import "fmt"
// @lc code=start
var pairBracketMap = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func isValid(s string) bool {
	var queue []rune
    runes := []rune(s)

	if len(runes) % 2 == 1 {
		return false
	}

	for _, r := range runes {
		closeBracket, ok := pairBracketMap[r]

		if ok {
			queue = append(queue, closeBracket)
			continue
		}
		qLen := len(queue)
		if qLen == 0 || queue[qLen - 1] != r {
			return false
		}

		queue = queue[:qLen - 1]
	}

	return len(queue) == 0
}
// @lc code=end

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("["))
	fmt.Println(isValid("(("))
	fmt.Println(isValid("{[]}"))
}