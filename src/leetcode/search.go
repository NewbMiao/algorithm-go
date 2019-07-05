package leetcode

/*
https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。

*/
func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left := 0
	right := l - 1
	if nums[left] < nums[right] { //无旋转点
		if target >= nums[0] && target <= nums[l-1] {
			left = 0
			right = l - 1
			return binSearch(nums, left, right, target)
		}
	} else {
		//	二分查找旋转点
		pivot := 0
		for left <= right {
			pivot = (left + right) / 2
			if nums[pivot] > nums[pivot+1] {
				pivot = pivot + 1
				break
			} else {
				//!important pivot==left时,防止right越界成为-1，先处理right
				if nums[pivot] < nums[left] { //左侧降序，旋转点在左侧
					right = pivot - 1
				} else { //左侧升序，旋转点在右侧
					left = pivot + 1
				}
			}
		}

		if target == nums[pivot] {
			return pivot
		}
		if target <= nums[pivot-1] && target >= nums[0] {
			left = 0
			right = pivot - 1
			return binSearch(nums, left, right, target)
		} else if target >= nums[pivot] && target <= nums[l-1] {
			left = pivot
			right = l - 1
			return binSearch(nums, left, right, target)
		}
	}

	return -1

}

func binSearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
