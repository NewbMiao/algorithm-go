package traversal

import (
	. "btree"
	"linkedlist"
	"fmt"
	"strings"
	"strconv"
)

//层级遍历序列化
func SerialByLevel(bt *Node) (res string) {
	queue := linkedlist.NewList()
	queue.Push(bt)

	res += fmt.Sprint(bt.Value, "!")
	for !queue.IsEmpty() {
		tmp := queue.Pop()
		if tmp == nil {
			return
		}

		cur := tmp.(*Node)
		if cur.Left != nil {
			res += fmt.Sprint(cur.Left.Value, "!")
			queue.Push(cur.Left)
		} else {
			res += "#!"
		}
		if cur.Right != nil {
			res += fmt.Sprint(cur.Right.Value, "!")
			queue.Push(cur.Right)
		} else {
			res += "#!"
		}
	}
	return
}

//层级遍历反序列化
func ReconByLevel(res string) (bt *Node) {
	strs := strings.Split(res, "!")
	index := 0
	queue := linkedlist.NewList()
	bt = genNode(strs[index])
	queue.Push(bt)
	for !queue.IsEmpty() {
		node := queue.Pop()
		if node == nil {
			return
		}
		cur := node.(*Node)
		index += 1
		cur.Left = genNode(strs[index])
		index += 1
		cur.Right = genNode(strs[index])

		if cur.Left != nil {
			queue.Push(cur.Left)
		}

		if cur.Right != nil {
			queue.Push(cur.Right)
		}
	}
	return
}

func genNode(s string) *Node {
	if s == "#" {
		return nil
	}
	if tmp, err := strconv.Atoi(s); err == nil {
		return &Node{Value: tmp}
	}
	return nil
}

func SerialByPre(bt *Node) (s string) {
	if bt == nil {
		return "#!"
	}
	s = fmt.Sprint(bt.Value, "!")
	s += SerialByPre(bt.Left)
	s += SerialByPre(bt.Right)
	return
}

func ReconByPre(s string) (bt *Node) {
	if s == "" {
		return
	}
	strs := strings.Split(s, "!")
	queue := linkedlist.NewList()
	for _, v := range strs {
		queue.Push(genNode(v))
	}
	bt = reconPre(queue)
	return
}

func reconPre(queue *linkedlist.LList) (bt *Node) {
	if queue == nil {
		return
	}
	node := queue.Pop()
	if node == nil {
		return
	}
	if tmp, ok := node.(*Node); ok && tmp != nil {
		bt = tmp
		bt.Left = reconPre(queue)
		bt.Right = reconPre(queue)
	}

	return
}
