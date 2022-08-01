/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */
package main
import (
	utils "reply2future.com/utils"
	"strings"
	"unicode"
)
// @lc code=start
func licenseKeyFormatting(s string, k int) string {
	s = strings.Replace(s, "-", "", -1)
	sLen := len(s)
	if sLen < k {
		return strings.ToUpper(s)
	}

	remainder := sLen % k
	quotient := sLen / k
	if remainder == 0 {
		quotient--
	}
	ret := make([]rune, sLen + quotient)
	for i := 0; i < remainder; i++ {
		ret[i] = unicode.ToUpper(rune(s[i]))
	}
	j := 0
	if remainder > 0 {
		ret[remainder] = '-'
		j = remainder + 1
	}

	s = s[remainder:]
	sLen = len(s)

	for i := 0; i < sLen; i++ {
		ret[j] = unicode.ToUpper(rune(s[i]))
		if (i + 1) % k == 0 && i != sLen - 1 {
			ret[j+1] = '-'
			j++
		}
		j++
	}
	return string(ret)
}
// @lc code=end
func main() {
	utils.AssertEqual2[string,string,int]("example 1","5F3Z-2E9W",licenseKeyFormatting,"5F3Z-2e-9-w",4)
	utils.AssertEqual2[string,string,int]("example 2","2-5G-3J",licenseKeyFormatting,"2-5g-3-J",2)
	utils.AssertEqual2[string,string,int]("example 3","A-A-A-A",licenseKeyFormatting,"a-a-a-a-",1)
}
