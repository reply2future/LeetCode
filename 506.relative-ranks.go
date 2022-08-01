/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */
package main
import (
	utils "reply2future.com/utils"
	"sort"
	"strconv"
)
// @lc code=start
func findRelativeRanks(score []int) []string {
	return bruteForce506(score)    
}

// #1 brute force =================================================================
func bruteForce506(score []int) []string {
	sLen := len(score)
	cloneScore := make([]int, sLen)
	copy(cloneScore, score)

	sort.Slice(cloneScore, func(i, j int) bool {return cloneScore[i] > cloneScore[j]})
	mp := make(map[int]string)
	for i, v := range cloneScore {
		switch i {
		case 0:
			mp[v] = "Gold Medal"
		case 1:
			mp[v] = "Silver Medal"
		case 2:
			mp[v] = "Bronze Medal"
		default:
			mp[v] = strconv.Itoa(i+1)
		}
	}
	ret := make([]string, sLen)
	for i, v := range score {
		ret[i] = mp[v]
	}

	return ret
}
// #1 brute force =================================================================
// @lc code=end
func main() {
	utils.AssertEqual1[[]string,[]int]("example 1",[]string{"Gold Medal","Silver Medal","Bronze Medal","4","5"},findRelativeRanks,[]int{5,4,3,2,1})
	utils.AssertEqual1[[]string,[]int]("example 2",[]string{"Gold Medal","5","Bronze Medal","Silver Medal","4"},findRelativeRanks,[]int{10,3,8,9,4})
}
