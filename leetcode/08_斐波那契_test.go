package main

import "testing"

func Test8(t *testing.T) {
	t.Log(fib(45))
}

// 斐波那契描述
// 0 1 1 2 3 5 8 13 21 34 55 89 144
func fib(i int) int {
	const mod int = 1e9 + 7
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		dp := make([]int, i+1)
		dp[0] = 0
		dp[1] = 1
		for j := 2; j <= i; j++ {
			dp[j] = (dp[j-1] + dp[j-2]) % mod
		}
		return dp[i]
	}
}
