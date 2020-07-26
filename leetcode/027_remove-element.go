package leetcode

func removeElement(nums []int, val int) int {
	cur := 0
	for _, v := range nums {
		if v != val {
			nums[cur] = v
			cur++
		}
	}
	return cur
}
