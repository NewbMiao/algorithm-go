package leetcode

/*
https://leetcode-cn.com/problems/next-permutation/submissions/
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

*/
func nextPermutation(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	//寻找右侧第一个非降序数字
	i := l - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i -= 1
	}
	if i >= 0 {
		//交换右侧只比nums[i]大一点的数
		j := l - 1
		for j >= 0 && nums[i] >= nums[j] {
			j -= 1
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums = reverseNums(nums, i+1)
	return nums
}

func reverseNums(nums []int, start int) []int {
	end := len(nums) - 1
	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
	return nums
}
