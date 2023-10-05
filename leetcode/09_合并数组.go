package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 先把nums2的值放到nums1的后面，然后排序
	for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[i-m]
	}
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
}
