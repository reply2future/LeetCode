package main

/**
 * https://leetcode.com/explore/learn/card/recursion-i/253/conclusion/1675/
 */
import "math"

func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}
	if K == 2 {
		return 1
	}
	// odd contract value
	_count := 0
	// find mid position of K
	_len := int(math.Pow(2, (float64)(N-1)))
	// f(1) = 0; f(2) = 1
	for {
		_len = _len / 2
		if K <= _len {
			continue
		}
		K -= _len
		_count++
		if (K == 1 && _count%2 != 1) || (K == 2 && _count%2 == 1) {
			return 0
		}
		if (K == 2 && _count%2 != 1) || (K == 1 && _count%2 == 1) {
			return 1
		}
	}
}

/**
 * 遞歸做法
 */
func _kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}
	if K == 2 {
		return 1
	}
	if K%2 == 1 {
		return _kthGrammar(N-1, (K+1)/2)
	}
	return (_kthGrammar(N-1, K/2) + 1) % 2
}
