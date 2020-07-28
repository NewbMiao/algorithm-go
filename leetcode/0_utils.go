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

func generateLinkedListFromArray(arr []int) *ListNode {
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

// link end node to node at `pos`
func generateCycleLinkedList(arr []int, pos int) *ListNode {
	if pos < 0 {
		return generateLinkedListFromArray(arr)
	}
	head := &ListNode{}
	cur := head
	cycleNode := &ListNode{}
	for i := 0; i < len(arr); i++ {
		cur.Val = arr[i]
		cur.Next = &ListNode{}
		if i == pos {
			cycleNode = cur
		}
		if i == len(arr)-1 {
			cur.Next = cycleNode
			break
		}
		cur = cur.Next
	}
	return head
}
