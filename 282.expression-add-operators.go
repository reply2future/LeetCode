/*
 * @lc app=leetcode id=282 lang=golang
 *
 * [282] Expression Add Operators
 */
package main
import (
	utils "reply2future.com/utils"
	"strconv"
)
// @lc code=start
// +,-,*
func addOperators(num string, target int) []string {
	return addOperatorsRecursive(num, target)
}

// #1 recursive =================================================================
func addOperatorsRecursive(num string, target int) []string {
	nLen := len(num)
	for i := 1; i <= nLen; i++ {
		n, e := strconv.Atoi(num[:i])
		if e != nil {
			panic(e)
		}

		
	}
}
// #1 =================================================================
// @lc code=end
func main() {
	utils.AssertEqual2[[]string,string,int]("example 1",[]string{"1*2*3","1+2+3"},addOperators,"123",6)
	utils.AssertEqual2[[]string,string,int]("example 2",[]string{"2*3+2","2+3*2"},addOperators,"232",8)
	utils.AssertEqual2[[]string,string,int]("example 3",[]string{},addOperators,"3456237490",9191)
}
