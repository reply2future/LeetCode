/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func judgeCircle(moves string) bool {
    bytes := []byte(moves)
	mp := make(map[byte]int)

	for _, char := range bytes {
		mp[char] += 1
	}

	if mp['U'] != mp['D'] {
		return false
	}
	if mp['L'] != mp['R'] {
		return false
	}
	return true
}
// @lc code=end
func main() {
	utils.AssertEqual1[bool,string]("example 1",true,judgeCircle,"UD")
	utils.AssertEqual1[bool,string]("example 2",false,judgeCircle,"LL")
}
