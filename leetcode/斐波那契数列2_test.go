package leetcode

import "testing"

func Test2(t *testing.T) {
	t.Log("结果: ", k2(4))
	t.Log("结果: ", k2(5))
	t.Log("结果: ", k2(10))
}

func k2(n int) int {
	if n == 1 || n == 0 {
		return 0
	}
	if n == 2 || n == 3 {
		return 1
	}
	a, b := 1, 1

	for i := 0; i < n-3; i++ {
		j := a
		a = b
		b = j + b
	}
	return b
}
