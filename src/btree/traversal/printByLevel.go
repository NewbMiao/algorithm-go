package traversal

import (
	. "btree"
	"linkedlist"
	"fmt"
	"container/list"
)

func PrintByLevel(bt *Node) {
	var last, nlast *Node //当前层最后一个节点，下一层最后一个节点
	queue := linkedlist.NewList()
	queue.Push(bt)
	last = bt
	level := 1
	fmt.Printf("level %d: ", level)
	for !queue.IsEmpty() {
		tmp := queue.Pop()
		if tmp == nil {
			return
		}
		node := tmp.(*Node)
		if node == nil {
			return
		}

		if node.Left != nil {
			queue.Push(node.Left)
			nlast = node.Left
		}
		if node.Right != nil {
			queue.Push(node.Right)
			nlast = node.Right
		}
		fmt.Printf("%d ", node.Value)

		if last == node && !queue.IsEmpty() {
			level += 1
			fmt.Printf("\nlevel %d: ", level)
			last = nlast
		}
	}
	fmt.Println()
}

func PrintByZigZag(bt *Node) {
	deQueue := list.New()
	var node, last, nlast *Node //当前层最后一个节点，下一层最后一个节点
	last = bt
	level := 1
	str := "from left to right"
	leftToRight := true
	fmt.Printf("level %d %s: ", level, str)

	deQueue.PushFront(bt)

	for deQueue.Len() != 0 {
		if leftToRight {
			tmp := deQueue.Front()
			if tmp == nil {
				continue
			}
			deQueue.Remove(tmp)
			node = tmp.Value.(*Node)
			if node == nil {
				continue
			}
			if node.Left != nil {
				if nlast == nil {
					nlast = node.Left
				}
				deQueue.PushBack(node.Left)

			}

			if node.Right != nil {
				if nlast == nil {
					nlast = node.Right
				}
				deQueue.PushBack(node.Right)

			}


		} else {
			tmp := deQueue.Back()
			if tmp == nil {
				continue
			}
			deQueue.Remove(tmp)
			node = tmp.Value.(*Node)
			if node == nil {
				continue
			}
			if node.Right != nil {
				if nlast == nil {
					nlast = node.Right
				}
				deQueue.PushFront(node.Right)

			}
			if node.Left != nil {
				if nlast == nil {
					nlast = node.Left
				}
				deQueue.PushFront(node.Left)

			}

		}

		fmt.Printf("%d ", node.Value)
		if last == node && deQueue.Len() > 0 {
			level += 1
			str = "from left to right"
			if leftToRight {
				str = "from right to left"
			}
			fmt.Printf("\nlevel %d %s: ", level, str)

			leftToRight = !leftToRight
			last = nlast
			nlast = nil
		}
	}

}
