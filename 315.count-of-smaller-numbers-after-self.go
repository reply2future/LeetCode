/*
 * @lc app=leetcode id=315 lang=golang
 *
 * [315] Count of Smaller Numbers After Self
 */
package main

import (
	utils "reply2future.com/utils"
	"math"
)

// @lc code=start
func countSmaller(nums []int) []int {
	return binaryInsertCountSmaller(nums)
}

// #1 binary insert sort =================================================================
// time limite
func binaryInsertCountSmaller(nums []int) (ret []int) {
	nLen := len(nums)
	ret = make([]int, nLen)

	sNums := []int{math.MinInt32, nums[nLen-1], math.MaxInt32}
	ret[nLen-1] = 0

	for i := nLen - 2; i >= 0; i-- {
		nNums, idx := binaryInsertNum(sNums, nums[i])
		sNums = nNums
		ret[i] = idx
	}

	return
}

// asc order
// return sorted list and insert index, from 0 to n
// if equal then idx - 1
func binaryInsertNum(nums []int, num int) (ret []int, idx int) {
	tmp := nums[:]
	// ret = make([]int, len(tmp)+1)
	
	for {
		nLen := len(tmp)
		if nLen == 0 {
			// copy(ret, nums[:idx])
			// ret[idx] = num
			// copy(ret[idx+1:], nums[idx:])
			ret = append(ret, nums[:idx]...)
			ret = append(ret, num)
			ret = append(ret, nums[idx:]...)
			// remove the first min value
			idx--
			return
		}

		mid := nLen / 2
		if tmp[mid] > num {
			tmp = tmp[:mid]
		} else if tmp[mid] < num {
			idx += mid + 1
			tmp = tmp[mid+1:]
		} else {
			// find first not equal
			for i := mid - 1; i >= 0; i-- {
				if tmp[i] != num {
					idx += i + 1
					break
				}
			}
			tmp = []int{}
		}
	}
}
// #1 ====================================================================================

// @lc code=end
func main() {
	utils.AssertEqual1[[]int, []int]("example 1", []int{2, 1, 1, 0}, countSmaller, []int{5, 2, 6, 1})
	utils.AssertEqual1[[]int, []int]("example 2", []int{0}, countSmaller, []int{-1})
	utils.AssertEqual1[[]int, []int]("example 3", []int{0, 0}, countSmaller, []int{-1, -1})
	utils.AssertEqual1[[]int, []int]("example 4", []int{10,27,10,35,12,22,28,8,19,2,12,2,9,6,12,5,17,9,19,12,14,6,12,5,12,3,0,10,0,7,8,4,0,0,4,3,2,0,1,0}, countSmaller, []int{26,78,27,100,33,67,90,23,66,5,38,7,35,23,52,22,83,51,98,69,81,32,78,28,94,13,2,97,3,76,99,51,9,21,84,66,65,36,100,41})
	utils.AssertEqual1[[]int, []int]("example 5", []int{0, 0, 0, 0}, countSmaller, []int{1, 2, 3, 4})
	// time limit
	expected := utils.ReadJsonFromFile[[]int]("./data/315.expected.json")
	inputData := utils.ReadJsonFromFile[[]int]("./data/315.input.json")
	utils.AssertEqual1[[]int, []int]("example 6", expected, countSmaller, inputData)
}