/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func addStrings(num1 string, num2 string) string {
    var ret []byte
	var runes1, runes2 []byte
	if len(num1) > len(num2) {
		runes1 = []byte(num1)
		runes2 = []byte(num2)
		ret = make([]byte, len(num1)+1)
	} else {
		runes2 = []byte(num1)
		runes1 = []byte(num2)
		ret = make([]byte, len(num2)+1)
	}

	rLen1 := len(runes1)
	rLen2 := len(runes2)
	quotient := 0
	for i := 1; i <= rLen1; i++ {
		v1 := int(runes1[rLen1-i] - '0')
		v2 := 0
		if i <= rLen2 {
			v2 = int(runes2[rLen2-i] - '0')
		}
		sum := v1 + v2 + quotient
		quotient = sum / 10
		remainder := sum % 10
		
		ret[rLen1+1-i] = byte(remainder + '0')
	}

	if quotient == 1 {
		ret[0] = '1'
	} else {
		ret = ret[1:]
	}

	return string(ret)
}
// @lc code=end
func main() {
	utils.AssertEqual2[string,string,string]("example 1","134",addStrings,"11","123")
	utils.AssertEqual2[string,string,string]("example 2","533",addStrings,"456","77")
	utils.AssertEqual2[string,string,string]("example 3","0",addStrings,"0","0")
	utils.AssertEqual2[string,string,string]("example 4","13",addStrings,"4","9")
}
