/*
 * @lc app=leetcode id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 */
package main
import (
	utils "reply2future.com/utils"
	"strconv"
)
// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var res []rune
	switch {
	case numerator < 0 && denominator < 0:
		numerator, denominator = -numerator, -denominator
	case numerator < 0 && denominator > 0:
		numerator = -numerator
		res = append(res, '-')
	case numerator > 0 && denominator < 0:
		denominator = -denominator
		res = append(res, '-')
	}

	// key: numerator, value: start index of res
	memo := make(map[int]int)
	quotient := numerator / denominator 
	remaider := numerator % denominator

	res = append(res, []rune(strconv.Itoa(quotient))...)

	if remaider == 0 {
		return string(res)
	}

	res = append(res, '.')

	for remaider != 0 {
		numerator = remaider * 10	
		
		sIndex, exists := memo[numerator]
		if exists {
			res = append(res, ')', ' ')
			copy(res[sIndex+1:], res[sIndex:])
			res[sIndex] = '('
			break
		}

		quotient = numerator / denominator 
		remaider = numerator % denominator	

		memo[numerator] = len(res)
		res = append(res, []rune(strconv.Itoa(quotient))...)
	}

	return string(res)
}
// @lc code=end
func main() {
	utils.AssertEqual2[string,int,int]("1/2","0.5",fractionToDecimal,1,2)
	utils.AssertEqual2[string,int,int]("2/1","2",fractionToDecimal,2,1)
	utils.AssertEqual2[string,int,int]("4/333","0.(012)",fractionToDecimal,4,333)
	utils.AssertEqual2[string,int,int]("400/333","1.(201)",fractionToDecimal,400,333)
	utils.AssertEqual2[string,int,int]("4/330","0.0(12)",fractionToDecimal,4,330)
	utils.AssertEqual2[string,int,int]("10/3","3.(3)",fractionToDecimal,10,3)
	utils.AssertEqual2[string,int,int]("220/3","73.(3)",fractionToDecimal,220,3)
	utils.AssertEqual2[string,int,int]("-50/8","-6.25",fractionToDecimal,-50,8)
	utils.AssertEqual2[string,int,int]("-22/-2","11",fractionToDecimal,-22,-2)
	utils.AssertEqual2[string,int,int]("0/-2","0",fractionToDecimal,0,-2)
}
