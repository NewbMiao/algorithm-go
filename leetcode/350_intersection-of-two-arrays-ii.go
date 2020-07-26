package leetcode

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m1 := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m1[v]++
	}
	k := 0
	for _, v := range nums2 {
		if m1[v] > 0 {
			m1[v]--
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}

func intersectDoublepointer(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var i, j, k int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
