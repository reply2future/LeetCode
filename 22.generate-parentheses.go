/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package main
import "fmt"
// @lc code=start
func generateParenthesis(n int) []string {
	retMaps := make(map[string]struct{})
	if n == 1 {
		return []string{"()"}
	}

	parenthesis := generateParenthesis(n - 1)
	for _, p := range parenthesis {
		for i := 0; i <= len(p); i++ {
			inner := p[:i] + "()" + p[i:]
			retMaps[inner] = struct{}{}
		}
	}

	return mapKeys(retMaps)
}

func mapKeys(mymap map[string]struct{}) []string {
    keys := make([]string, 0, len(mymap))
    for k := range mymap {
        keys = append(keys, k)
    }
	return keys
}

// @lc code=end
func main() {
	fmt.Println(generateParenthesis(6))
	fmt.Println(generateParenthesis(5))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
