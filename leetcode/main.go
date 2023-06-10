package main

import "fmt"

func main() {
	fmt.Println(isStraight([]int{4, 7, 5, 9, 2}))
}

func isStraight(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums)-1; j++ {
			if nums[i] > nums[j+1] {
				nums[i], nums[j+1] = nums[j+1], nums[i]
			}
		}
	}
	count := 0
	v := 1
	countMap := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1]-1 {
			v++
		}
		if nums[i] == 0 {
			count++
		} else {
			countMap[nums[i]]++
			if countMap[nums[i]] >= 2 {
				return false
			}
		}
	}
	if v == 3 && count == 2 || v == 5 {
		return true
	}
	if count == 2 {
		card := nums[count:]
		for i := 0; i < len(card)-1; i++ {
			if card[i+1]-card[i] > 3 {
				return false
			}
		}
		return true
	} else if count == 1 && v == 3 {
		card := nums[count:]
		for i := 0; i < len(card)-1; i++ {
			if card[i+1]-card[i] > 2 {
				return false
			}
		}
		return true
	} else if count == 3 {
		card := nums[count:]
		return card[count+1]-card[count] < 4
	}

	return v == 5 || count >= 4
}
