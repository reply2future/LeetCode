/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */
package main
import (
	utils "reply2future.com/utils"
)
// @lc code=start
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	return _longestConsecutive(nums)
}

// #1 O(n)
func _longestConsecutive(nums []int) int {
	maps := convertToMap(nums)

	lc := 1
	for {
		if len(maps) == 0 {
			break
		}

		cc := 1
		key := getKey(maps)
		delete(maps, key)

		// asc direction
		ascKey := key + 1
		for {
			_, ok := maps[ascKey]
			if !ok {
				break
			}
			delete(maps, ascKey)
			cc += 1
			ascKey += 1
		}

		// desc direction
		descKey := key - 1
		for {
			_, ok := maps[descKey]
			if !ok {
				break
			}
			delete(maps, ascKey)
			cc += 1
			ascKey -= 1
		}

		if cc > lc {
			lc = cc
		}
	}

	return lc
}

func convertToMap(nums []int) map[int]struct{} {
	maps := make(map[int]struct{})
	for _, num := range nums {
		maps[num] = struct{}{}
	}
	return maps
}

func getKey(maps map[int]struct{}) int {
	for key := range maps {
		return key
	}
	return -1
}

// @lc code=end
func main() {
	utils.AssertEqual1[int,[]int]("[100,4,200,1,3,2]",4,longestConsecutive,[]int{100,4,200,1,3,2})
	utils.AssertEqual1[int,[]int]("[0,3,7,2,5,8,4,6,0,1]",9,longestConsecutive,[]int{0,3,7,2,5,8,4,6,0,1})
}
