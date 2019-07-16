package linkedlist

func RemoveLastKthNode(l *LList, k int) {
	if l.IsEmpty() || k < 1 {
		return
	}
	node := l.Head
	for node != nil {
		k--
		node = node.Next
	}
	if k > 0 { //无倒数kth node
		return
	}
	if k == 0 { //倒数kth node为头结点
		l.Head = l.Head.Next
		return
	}
	//此时k-N,在遍历一次，k==0时，找到N-k,即删除节点的前一个节点
	node = l.Head
	for node != nil {
		k++
		if k == 0 {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
}

func RemoveLastKthDNode(l *DList, k int) {
	if l.IsEmpty() || k < 1 {
		return
	}
	node := l.Head
	for node != nil {
		k--
		node = node.Next
	}
	if k > 0 { //无倒数kth node
		return
	}
	if k == 0 { //倒数kth node为头结点
		head := l.Head
		l.Head = head.Next
		l.Head.Last = head.Last
		return
	}
	//此时k-N,在遍历一次，k==0时，找到N-k,即删除节点的前一个节点
	node = l.Head
	for node != nil {
		k++
		if k == 0 {
			next := node.Next
			node.Next = next.Next
			next.Last = node.Last
			break
		}
		node = node.Next
	}
}
