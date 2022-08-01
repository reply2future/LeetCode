package main

func myPow(x float64, n int) float64 {
	if x == 1 || x == 0 {
		return x
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}
	var _sum float64 = 1
	_neg := false
	if n < 0 {
		_neg = true
		n = -n
	}
	for i := 0; i < n; i++ {
		_sum *= x
	}
	if _neg {
		return 1 / _sum
	}
	return _sum
}
