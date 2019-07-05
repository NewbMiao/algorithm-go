package leetcode

func extremeInsertionIndex(nums []int, target int, left bool) int {
	lo := 0
	hi := len(nums) //!important mid后移，右侧边界 lo-1
	for lo < hi {
		mid := (hi + lo) / 2
		if nums[mid] > target ||
			(left && nums[mid] == target) { //left==true：获取range左边界，false则为右边界（lo-1）
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}
	if l == 1 && target == nums[0] {
		return []int{0, 0}
	}
	leftIdx := extremeInsertionIndex(nums, target, true)
	if leftIdx == l || nums[leftIdx] != target {
		return []int{-1, -1}
	}
	return []int{
		leftIdx,
		extremeInsertionIndex(nums, target, false) - 1,
	}

}
