package leetcode

import (
	"sort"
)

/*
https://leetcode-cn.com/problems/3sum/
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。.

*/
func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	if nums[0] > 0 || nums[l-1] < 0 {
		return nil
	}
	for i := 0; i < l-2; i++ {
		// 最左为正数则无解
		if nums[i] > 0 {
			break
		}
		// 后边的相同值计算过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		first := i + 1
		last := l - 1

		for last > first {
			sum := nums[i] + nums[first] + nums[last]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[first], nums[last]})
			}
			if sum <= 0 {
				for first < last && nums[first] == nums[first+1] {
					first++
				}
				first++
			}
			if sum >= 0 {
				for first < last && nums[last] == nums[last-1] {
					last--
				}
				last--
			}
		}
	}

	return res
}
