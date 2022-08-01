/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
package main

import "fmt"

// @lc code=start
func nextPermutation(nums []int) {
	// # recursion
	// recursion(nums)

	// # direction
	direction(nums)
}

func recursion(nums []int) {
	reverse := next(nums)
	if !reverse {
		reverseArray(nums)
	}
}

func direction(nums []int) {
	nLen := len(nums)
	if nLen == 0 || nLen == 1 {
		return
	}

	needReverse := true
	for i := nLen - 1; i > 0; i-- {
		// 1, 3, 5, 4, 2
		if nums[i] > nums[i-1] {
			for j := nLen - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					swap := nums[j]
					nums[j] = nums[i-1]
					nums[i-1] = swap
					reverseArray(nums[i:])
					break
				}
			}

			needReverse = false
			break
		}
	}

	if needReverse {
		reverseArray(nums)
	}
}

func reverseArray(nums []int) {
	nLen := len(nums)
	var swap int
	for i := 0; i < nLen/2; i++ {
		swap = nums[i]
		nums[i] = nums[nLen-1-i]
		nums[nLen-1-i] = swap
	}
}

func next(restNum []int) bool {
	rLen := len(restNum)
	if rLen == 0 || rLen == 1 {
		return false
	}
	if rLen == 2 && restNum[0] >= restNum[1] {
		return false
	}

	reverse := next(restNum[1:])
	if !reverse {
		var swap int
		match := false
		for i := rLen - 1; i > 0; i-- {
			if restNum[0] >= restNum[i] {
				continue
			}
			swap = restNum[0]
			restNum[0] = restNum[i]
			restNum[i] = swap
			reverseArray(restNum[1:])
			match = true
			break
		}
		return match
	}

	return true
}

// @lc code=end

func main() {
	a := []int{1, 2, 3}
	nextPermutation(a)
	fmt.Println(a) // []int{1, 3, 2}
	nextPermutation(a)
	fmt.Println(a) // []int{2, 1, 3}, "The two words should be the same.")
	nextPermutation(a)
	fmt.Println(a) // []int{2, 3, 1}, "The two words should be the same.")
	nextPermutation(a)
	fmt.Println(a) // []int{3, 1, 2}, "The two words should be the same.")
	nextPermutation(a)
	fmt.Println(a) // []int{3, 2, 1}, "The two words should be the same.")
	nextPermutation(a)
	fmt.Println(a) // []int{1, 2, 3}, "The two words should be the same.")

	b := []int{1, 2, 3, 4, 5}
	nextPermutation(b)
	fmt.Println(b) // []int{1, 2, 3, 5, 4}
	nextPermutation(b)
	fmt.Println(b) // []int{1, 2, 4, 3, 5}
	nextPermutation(b)
	fmt.Println(b) // []int{1, 2, 4, 5, 3}
	nextPermutation(b)
	fmt.Println(b) // []int{1, 2, 5, 3, 4}
	nextPermutation(b)
	fmt.Println(b) // []int{1, 2, 5, 4, 3}
	nextPermutation(b)
	fmt.Println(b) // []int{1, 3, 2, 4, 5}

	c := []int{1}
	nextPermutation(c)
	fmt.Println(c) // []int{1}
}
