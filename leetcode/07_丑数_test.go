package main

import (
	"fmt"
	"testing"
)

func Test7(t *testing.T) {
	fmt.Println(nthUglyNumber(1352))
}

// 丑数是指只包含质因数 2,3,5 的正整数。
// 什么是质因数？质因数就是一个数的因数中的质数
func nthUglyNumber(n int) int {
	k := 1
	for i := 2; k != n; i++ {
		v := i
		for v%2 == 0 {
			v /= 2
		}
		for v%3 == 0 {
			v /= 3
		}
		for v%5 == 0 {
			v /= 5
		}
		if v == 1 {
			k++
		}
		if k == n {
			return i
		}
	}
	return k
}
