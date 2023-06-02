package leetcode

import "testing"

func Test3(t *testing.T) {
	// 必须有序,如果没有找到则返回-1

	nums := []int{2, 7, 11, 13, 15, 21, 24}
	target := 45

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if target == sum {
				t.Logf("a: %v, b: %v, sum: %v", i, j, sum)
				return
			}
		}
	}
	t.Log("没有找到: -1")
}
