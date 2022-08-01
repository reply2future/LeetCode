/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
// sell -> cooldown
// buy -> sell -> cooldown -> buy
// x buy -> buy
// period: [buy -> n1 * cooldown -> sell -> cooldown] && buy < sell
import "math"

func maxProfit(prices []int) int {
	// return dpRecursiveMaxProfit(prices)
	return bottomUpDpMaxProfit(prices)
}

// #1 dp recursive =================================================
// time limited
func dpRecursiveMaxProfit(prices []int) int {
	memo := make(map[int]int)
	var dpMaxProfit func(prices []int) int
	dpMaxProfit = func(prices []int) int {
		pLen := len(prices)
		if v, ok := memo[pLen]; ok {
			return v
		}

		maxProfit := 0
		for i := 0; i < pLen-1; i++ {
			buy := prices[i]
			for j := i + 1; j < pLen; j++ {
				if prices[j] <= buy {
					continue
				}

				sum := 0
				sum += (prices[j] - buy)
				// cooldown
				if j+2 <= pLen {
					mLen := pLen - j - 2
					if v, ok := memo[mLen]; ok {
						sum += v
					} else {
						result := dpMaxProfit(prices[j+2:])
						memo[mLen] = result
						sum += result
					}
				}

				if sum > maxProfit {
					maxProfit = sum
				}
			}
		}

		return maxProfit
	}

	return dpMaxProfit(prices)
}

// #1 ==============================================================
// #2 bottom up dp =================================================

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func bottomUpDpMaxProfit(prices []int) int {
	// initialization
	cool_down, sell, hold := 0, 0, math.MinInt64

	for _, stock_price_of_Day_i := range prices {

		prev_cool_down, prev_sell, prev_hold := cool_down, sell, hold

		// Max profit of cooldown on Day i comes from either cool down of Day_i-1, or sell out of Day_i-1 and today Day_i is cooling day
		cool_down = Max(prev_cool_down, prev_sell)

		// Max profit of sell on Day_i comes from hold of Day_i-1 and sell on Day_i
		sell = prev_hold + stock_price_of_Day_i

		// Max profit of hold on Day_i comes from either hold of Day_i-1, or cool down on Day_i-1 and buy on Day_i
		hold = Max(prev_hold, prev_cool_down-stock_price_of_Day_i)

	}

	// The state of max profit must be either sell or cool down
	return Max(sell, cool_down)
}

// #2 ==============================================================

// @lc code=end
func main() {
	utils.AssertEqual1[int, []int]("exmaple 1", 3, maxProfit, []int{1, 2, 3, 0, 2})
	utils.AssertEqual1[int, []int]("exmaple 2", 0, maxProfit, []int{1})
	utils.AssertEqual1[int, []int]("exmaple 3", 3, maxProfit, []int{1, 2, 4})
	utils.AssertEqual1[int, []int]("exmaple 4", 0, maxProfit, []int{5, 4, 3, 2, 1})
}
