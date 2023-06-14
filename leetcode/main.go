package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
func pivotIndex(nums []int) int {
	i := len(nums) / 2
	a := nums[0:i]
	b := nums[i+1:]
	aLength := len(a)
	bLength := len(b)
	sumA := make([]int, aLength)
	sumB := make([]int, bLength)
	for index, k := range a {
		if index == 0 {
			sumA[index] = k
		} else {
			sumA[index] = sumA[index-1] + k
		}
	}
	for index, k := range b {
		if index == 0 {
			sumB[index] = k
		} else {
			sumB[index] = sumB[index-1] + k
		}
	}
	if aLength > bLength {
		for index, _ := range sumA {
			if sumA[index] == sumB[bLength-1] {
				return i
			}
		}
	} else {
		for index, _ := range sumB {
			if sumB[index] == sumA[bLength-1] {
				return i
			}
		}
	}
	return -1
}
