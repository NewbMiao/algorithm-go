package leetcode

/*
https://leetcode-cn.com/problems/permutations/
给定一个没有重复数字的序列，返回其所有可能的全排列。

https://leetcode-cn.com/problems/permutations/solution/quan-pai-lie-by-leetcode/
*/
func permute(nums []int) [][]int {
	out := make([][]int, 0)
	backtrackPermute(len(nums), 0, nums, &out)
	return out
}

func backtrackPermute(n, first int, nums []int, out *[][]int) {
	if n == first {
		// 拷贝nums，防止回溯结束后nums重置为1，2，3
		tmp := make([]int, n)
		copy(tmp, nums)
		*out = append(*out, tmp)
		return
	}
	for i := first; i < n; i++ {
		// 尝试位置first和i替换的排列
		if first != i {
			nums[first], nums[i] = nums[i], nums[first]
		}
		// 尝试first+1的全排列
		backtrackPermute(n, first+1, nums, out)
		// 恢复当前替换，使不影响后续回溯
		if first != i {
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
}
