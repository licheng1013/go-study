package main

import (
	"fmt"
	"testing"
)

// 编写测试案例
// 1. 1,2,3,4,5  4,5,3,2,1
// 2. 1,2,3,4,5  4,3,5,1,2
// 测试方法
func Test6(t *testing.T) {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 1, 2}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	var nums = len(popped)
	if len(pushed) != nums {
		return false
	}
	stack := make([]int, 0)
	for _, k := range pushed {
		// 添加到栈里面
		stack = append(stack, k)
		//  在每次添加时,对比出栈数组, 如果一致继续对比继续弹出,不一致则继续添加
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
	}
	return len(stack) == 0
}
