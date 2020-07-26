package traversal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/NewbMiao/algorithm-go/kit/btree"
	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

//层级遍历序列化
func SerialByLevel(bt *btree.Node) (res string) {
	queue := linkedlist.NewList()
	queue.Push(bt)

	res += fmt.Sprint(bt.Value, "!")
	for !queue.IsEmpty() {
		tmp := queue.Pop()
		if tmp == nil {
			return
		}

		cur := tmp.(*btree.Node)
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
func ReconByLevel(res string) (bt *btree.Node) {
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
		cur := node.(*btree.Node)
		index++
		cur.Left = genNode(strs[index])
		index++
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

func genNode(s string) *btree.Node {
	if s == "#" {
		return nil
	}
	if tmp, err := strconv.Atoi(s); err == nil {
		return &btree.Node{Value: tmp}
	}
	return nil
}

func SerialByPre(bt *btree.Node) (s string) {
	if bt == nil {
		return "#!"
	}
	s = fmt.Sprint(bt.Value, "!")
	s += SerialByPre(bt.Left)
	s += SerialByPre(bt.Right)
	return
}

func ReconByPre(s string) (bt *btree.Node) {
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

func reconPre(queue *linkedlist.LList) (bt *btree.Node) {
	if queue == nil {
		return
	}
	node := queue.Pop()
	if node == nil {
		return
	}
	if tmp, ok := node.(*btree.Node); ok && tmp != nil {
		bt = tmp
		bt.Left = reconPre(queue)
		bt.Right = reconPre(queue)
	}

	return
}
