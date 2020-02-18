package traversal

import (
	. "github.com/NewbMiao/AlgorithmCodeNote/kit/btree"
)

func MorrisIn(bt *Node) (res []int) {
	if bt == nil {
		return
	}

	cur1 := bt     //当前根节点
	var cur2 *Node //当前根节点最右节点
	for cur1 != nil {
		cur2 = cur1.Left //先处理左子树
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				//寻找当前子树最右节点
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				//最右节点的右节点指向根节点
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				//恢复最右节点的右节点指向空节点
				cur2.Right = nil
			}
		}

		res = append(res, cur1.Value)
		//左子树处理完，切到右节点继续
		cur1 = cur1.Right
	}
	return
}

func MorrisPre(bt *Node) (res []int) {
	if bt == nil {
		return
	}

	cur1 := bt     //当前根节点
	var cur2 *Node //当前根节点最右节点
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				//寻找当前子树最右节点
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				//最右节点的右节点指向根节点
				cur2.Right = cur1
				//打印根节点
				res = append(res, cur1.Value)
				cur1 = cur1.Left
				continue
			} else {
				//恢复最右节点的右节点指向空节点
				cur2.Right = nil
			}
		} else {
			//打印左节点或右节点
			res = append(res, cur1.Value)
		}
		//左子树处理完，切到又节点继续
		cur1 = cur1.Right
	}
	return
}

func MorrisLast(bt *Node) (res []int) {
	if bt == nil {
		return
	}

	cur1 := bt     //当前根节点
	var cur2 *Node //当前根节点最右节点
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				//寻找当前子树最右节点
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				//最右节点的右节点指向根节点
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				//恢复最右节点的右节点指向空节点
				cur2.Right = nil
				//左右根获取左子树
				res = append(res, getEdgeNode(cur1.Left)...)
			}
		}
		//左子树处理完，切到又节点继续
		cur1 = cur1.Right
	}
	//获取根节点右边界
	res = append(res, getEdgeNode(bt)...)

	return
}

func getEdgeNode(bt *Node) (res []int) {
	//逆序右边界
	tail := reverseEdge(bt)
	cur := tail
	for cur != nil {
		res = append(res, cur.Value)
		cur = cur.Right
	}
	reverseEdge(tail)

	return
}

func reverseEdge(from *Node) *Node {
	var pre, next *Node
	for from != nil {
		next = from.Right
		from.Right = pre
		pre = from
		from = next
	}
	return pre
}
