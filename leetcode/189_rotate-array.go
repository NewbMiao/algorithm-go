package leetcode

func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	k %= len(nums)
	reverseArray(nums)
	reverseArray(nums[:k])
	reverseArray(nums[k:])
}

func rotateAppend(nums []int, k int) {
	l := len(nums)
	if l < 2 {
		return
	}
	k %= l
	target := append(nums[l-k:l], nums[:l-k]...)
	copy(nums, target)
}
