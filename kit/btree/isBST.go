package btree

// 是否为搜索二叉树：中序遍历为升序.
func IsBST(bt *Node) (res bool) {
	if bt == nil {
		return
	}
	res = true
	cur1 := bt
	var cur2, pre *Node
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				cur2.Right = nil
			}
		}

		if pre != nil && pre.Value > cur1.Value {
			res = false // 记录，等恢复更改的右指针指向再返回
		}
		pre = cur1
		cur1 = cur1.Right
	}
	return
}
