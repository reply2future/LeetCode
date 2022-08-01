/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */
package main
import (
	utils "reply2future.com/utils"
	"strings"
)
// @lc code=start
func detectCapitalUse(word string) bool {
	return strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}
// @lc code=end
func main() {
	utils.AssertEqual1[bool,string]("example 1",true,detectCapitalUse,"USA")
	utils.AssertEqual1[bool,string]("example 2",false,detectCapitalUse,"FlaG")
	utils.AssertEqual1[bool,string]("example 3",true,detectCapitalUse,"g")
}
