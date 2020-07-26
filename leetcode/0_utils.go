package leetcode

func reverseArray(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (l *ListNode) convertListNodesToArray() (res []int) {
	curNode := l
	for curNode != nil {
		res = append(res, curNode.Val)
		curNode = curNode.Next
	}
	return
}

func generateListNodesFromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	l := &ListNode{}
	curNode := l
	for i := 0; i < len(arr); i++ {
		curNode.Val = arr[i]
		if i == len(arr)-1 {
			break
		}
		curNode.Next = &ListNode{}
		curNode = curNode.Next
	}
	return l
}
