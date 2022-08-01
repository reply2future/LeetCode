package main

func getRow(rowIndex int) []int {
	_output := make([]int, rowIndex+1)
	_mid := rowIndex / 2
	for i := 0; i <= rowIndex; i++ {
		if i <= _mid {
			_output[i] = _cal(rowIndex, i)
		} else {
			_output[i] = _output[rowIndex-i]
		}
	}
	return _output
}

func _cal(x int, y int) int {
	if y == 0 || x == y {
		return 1
	}

	if y == 1 {
		return x
	}

	if (y - 1) == (x - 1 - y) {
		return 2 * _cal(x-1, y-1)
	}
	return _cal(x-1, y-1) + _cal(x-1, y)
}

/**
 * 使用二項式方法直接計算本身這一列的規則
 */
func _directGetRow(rowIndex int) []int {
	_output := make([]int, rowIndex+1)
	_output[0] = 1
	for i := 1; i <= rowIndex; i++ {
		_output[i] = _output[i-1] * (rowIndex - i + 1) / i
	}
	return _output
}
