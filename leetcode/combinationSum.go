package leetcode

import (
	"sort"
)

/*
https://leetcode-cn.com/problems/combination-sum

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。

说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 .

*/
func combinationSum(candidates []int, target int) [][]int {
	l := len(candidates)

	tmp := candidates
	if l == 0 {
		return nil
	}
	for _, v := range candidates {
		if v != 0 && target >= 2*v {
			ts := target/v - 1
			for i := 0; i < ts; i++ {
				tmp = append(tmp, v)
			}
		}
	}
	sort.Ints(tmp)
	res := []int{}
	return findCombinationSum(tmp, res, len(tmp)-1, target)
}

func findCombinationSum(nums, res []int, index, target int) (out [][]int) {
	if index < 0 || target <= 0 {
		return
	}

	if target == nums[index] {
		tmp := make([]int, 0, len(res)+1)
		tmp = append(tmp, res...)
		tmp = append(tmp, nums[index])
		sort.Ints(tmp)
		out = append(out, tmp)
	}

	res = append(res, nums[index])
	out = append(out, findCombinationSum(nums, res, index-1, target-nums[index])...)

	res = res[:len(res)-1]
	offset := 1
	for index > 0 && nums[index] == nums[index-1] {
		index--
	}
	out = append(out, findCombinationSum(nums, res, index-offset, target)...)

	return
}
