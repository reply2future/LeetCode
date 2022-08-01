/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */
package main

import (
	utils "reply2future.com/utils"
	"strconv"
	"strings"
)

// @lc code=start
func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")

	s = clearBrackets(s)

	sum := 0
	i := 0
	var j int
	for i < len(s) {
		switch s[i] {
		case '+':
			for j = i + 1; j < len(s); j++ {
				if s[j] == '+' || s[j] == '-' {
					break
				}
			}
			num, err := strconv.Atoi(s[i+1 : j])
			if err != nil {
				panic(err)
			}
			sum += num
		case '-':
			for j = i + 1; j < len(s); j++ {
				if s[j] == '+' || s[j] == '-' {
					break
				}
			}
			num, err := strconv.Atoi(s[i+1 : j])
			if err != nil {
				panic(err)
			}
			sum -= num
		default:
			for j = i + 1; j < len(s); j++ {
				if s[j] == '+' || s[j] == '-' {
					break
				}
			}
			num, err := strconv.Atoi(s[i:j])
			if err != nil {
				panic(err)
			}
			sum += num
		}

		i = j
	}

	return sum
}

func clearBrackets(s string) string {
	var sb strings.Builder
	var b []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if i == 0 || s[i-1] == '+' {
				b = append(b, '+')
			} else {
				b = append(b, '-')
			}
		} else if s[i] == ')' {
			b = b[:len(b)-1]
		} else {
			// -(12) || 12
			if (i-1 > 0 && s[i-1] == '(') || (s[i] != '-' && s[i] != '+') {
				sb.WriteByte(s[i])
			} else {
				sb.WriteString(string(b))
				sb.WriteByte(s[i])
			}
		}
	}
	ret := sb.String()
	r := strings.NewReplacer("--", "+", "++", "+", "+-", "-", "-+", "-")
	for {
		bLen := len(ret)
		ret = r.Replace(ret)
		aLen := len(ret)
		if aLen == bLen {
			break
		}
	}
	return ret
}

// @lc code=end
func main() {
	utils.AssertEqual1[int, string]("1 + 1", 2, calculate, "1 + 1")
	utils.AssertEqual1[int, string](" 2-1 + 2 ", 3, calculate, " 2-1 + 2 ")
	utils.AssertEqual1[int, string]("(1+(4+5+2)-3)+(6+8)", 23, calculate, "(1+(4+5+2)-3)+(6+8)")
	utils.AssertEqual1[int, string](" -(2-1) + 2 ", 1, calculate, " -(2-1) + 2 ")
	utils.AssertEqual1[int, string]("1-(-1-(2+1-(3+2)-1))", -1, calculate, "1-(-1-(2+1-(3+2)-1))")
	utils.AssertEqual1[int, string]("- (3 - (- (4 + 5) ) )", -12, calculate, "- (3 - (- (4 + 5) ) )")
	utils.AssertEqual1[int, string]("2-4-(8+2-6+(8+4-(1)+8-10))", -15, calculate,"2-4-(8+2-6+(8+4-(1)+8-10))")
}
