/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
package main

import "fmt"

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	// return bruteForce(height)
	return optimizeBruteForce(height)
}

// brute force
func bruteForce(height []int) int {
	// two border side cannot trap
	// 1. the barrel effect
	// 2. use variable to record the match height, matchHeight <= height
	// 3. result = height x distance
	// 4. how to match: height <= cursorHeight
	// 5. how to next bar: matchHeight == height
	// 6. last bar will be ignored
	hLen := len(height)
	sum := 0

	// hLen - 1: last bar will be ignored
	for i := 0; i < hLen-1; i++ {
		matchHeight := 0
		for j := i + 1; j < hLen; j++ {
			if height[j] <= matchHeight {
				continue
			}

			useHeight := min(height[j], height[i])
			distance := (j - i - 1)
			sum += (useHeight - matchHeight) * distance
			matchHeight = useHeight

			if matchHeight == height[i] {
				break
			}
		}
	}

	return sum
}

func optimizeBruteForce(height []int) int {
	// two border side cannot trap
	// 1. the barrel effect
	// 2. use variable to record the match height, matchHeight <= height
	// 3. result = height x distance
	// 4. how to match: height <= cursorHeight
	// 5. how to next bar: matchHeight == height
	// 6. last bar will be ignored
	hLen := len(height)
	sum := 0

	// hLen - 1: last bar will be ignored
	for i := 0; i < hLen-1; i++ {
		if height[i] == 0 {
			continue
		}
		
		var maxIndex int
		matchHeight := 0
		for j := i + 1; j < hLen; j++ {
			if height[j] <= matchHeight {
				continue
			}

			maxIndex = j
			matchHeight = min(height[j], height[i])

			if height[j] >= height[i] {
				break
			}
		}

		if maxIndex == 0 {
			break	
		}

		middleHolder := 0

		for m := i + 1; m < maxIndex; m++ {
			middleHolder += height[m]
		}

		distance := (maxIndex - i - 1)
		sum += matchHeight * distance - middleHolder
		i = maxIndex - 1
	}

	return sum
}

func dp(height []int) int {
	sum := 0

	return sum
}

// @lc code=end

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), "6")
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}), "9")
	fmt.Println(trap([]int{0, 2, 0}), "0")
}
