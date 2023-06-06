package main

import (
	"fmt"
	"testing"
)

// 编写测试
func Test5(t *testing.T) {
	fmt.Println(hammingWeight(11)) // Output: 3
}


func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		// 按位与运算符: 两个数都为1时，结果才为1，否则为0
		// 按位运算例如: 11 = 1011, 1 = 0001, 1011 & 0001 = 0001 = 1
		count += int(num & 1)
		// 右移一位例如: 1010 >> 0101 , 右位运算相当于除以2
		// 再次右移一位例如: 0101 >> 0010 ,移除最后1位
		// 再次右移一位例如: 0010 >> 0001 ,移除最后1位
		num >>= 1
		// 打压二进制字符串
		//fmt.Printf("%b\n", num) // Output: 1011
	}
	return count
}
