/*
 * @lc app=leetcode id=313 lang=golang
 *
 * [313] Super Ugly Number
 */
package main
import (
	utils "reply2future.com/utils"
	"math"
)
// @lc code=start
func nthSuperUglyNumber(n int, primes []int) int {
	pLen := len(primes)
	indices := make([]int, pLen)
    out := make([]int, n)
	out[0] = 1

	for j := 1; j < n; j++ {
		minVal := math.MaxInt32
		minIdx := make([]int, 0)
		for i := 0; i < pLen; i++ {
			current := primes[i] * out[indices[i]]
			if current < minVal {
				minIdx = []int{i}
				minVal = current
			} else if current == minVal {
				minIdx = append(minIdx, i)
			}
		}

		out[j] = minVal
		for _, v := range minIdx {
			indices[v] += 1
		}
	}

	return out[n-1]
}
// @lc code=end
func main() {
	utils.AssertEqual2[int,int,[]int]("example 1",32,nthSuperUglyNumber,12,[]int{2,7,13,19})
	utils.AssertEqual2[int,int,[]int]("example 2",1,nthSuperUglyNumber,1,[]int{2,3,5})
	utils.AssertEqual2[int,int,[]int]("example 3",2,nthSuperUglyNumber,2,[]int{2,3,5})
}
