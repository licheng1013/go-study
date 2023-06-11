package main

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 8, 10}
	fmt.Println(AssertNumber(SearchTarget(list, 1), 0))
	fmt.Println(AssertNumber(SearchTarget(list, 2), 1))
	fmt.Println(AssertNumber(SearchTarget(list, 3), 2))
	fmt.Println(AssertNumber(SearchTarget(list, 4), 3))
	fmt.Println(AssertNumber(SearchTarget(list, 5), 4))
	fmt.Println(AssertNumber(SearchTarget(list, 8), 5))
	fmt.Println(AssertNumber(SearchTarget(list, 10), 6))
	fmt.Println(AssertNumber(SearchTarget(list, 11), -1))
}

func SearchTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		centerIndex := (left + right) / 2
		if nums[centerIndex] < target {
			left = centerIndex + 1
		} else if target < nums[centerIndex] {
			right = centerIndex - 1
		} else {
			return centerIndex
		}
	}
	return -1
}
