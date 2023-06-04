package main

import (
	"fmt"
)

func main() {
	fmt.Println([]int{0}[1:])
	fmt.Println(5 / 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	length := len(nums)
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
	return nums
}
