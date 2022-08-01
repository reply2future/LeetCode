package main

/**
 *You are climbing a stair case. It takes n steps to reach to the top.
 *Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
 */
var _cache map[int]int = make(map[int]int)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	_v, _ok := _cache[n]
	if _ok {
		return _v
	}
	_r := climbStairs(n-1) + climbStairs(n-2)
	_cache[n] = _r
	return _r
}
