/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */
package main

import (
	"fmt"
	utils "reply2future.com/utils"
	"strconv"
)

// @lc code=start
// 0 <= turned on <= 10
// 1:00 and 10:02
// hours: 8,4,2,1,0
// minutes: 32,16,8,4,2,1,0
func readBinaryWatch(turnedOn int) []string {
	ret := []string{}
	if turnedOn >= 9 {
		return ret
	}

	var loop func(int, int, int, int)
	loop = func(n int, hours int, minutes int, pos int) {
		if hours >= 12 || minutes > 59 {
			return
		}
		if n == 0 {
			ret = append(ret, strconv.Itoa(hours)+":"+fmt.Sprintf("%02d", minutes))
			return
		}

		// hours + minutes => array [10]int{1,2,4,8,16,32,   1,2,4,8}
		for i := pos; i < 10; i++ {
			if i < 6 {
				// minutes
				loop(n-1, hours, minutes+1<<i, i+1)
			} else {
				// hours
				loop(n-1, hours+1<<(i-6), minutes, i+1)
			}
		}
	}
	loop(turnedOn, 0, 0, 0)

	return ret
}

// @lc code=end
func main() {
	readBinaryWatch(2)
	utils.AssertEqual1[[]string, int]("example 1", []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}, readBinaryWatch, 1)
	utils.AssertEqual1[[]string, int]("example 2", []string{}, readBinaryWatch, 9)
}
