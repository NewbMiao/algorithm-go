package linkedlist

import (
	"fmt"

	"github.com/NewbMiao/algorithm-go/kit/linkedlist"
)

//环形链表模拟约瑟夫环问题
func JosephusKill1(l *linkedlist.LList, m int) (res int) {
	if l == nil || l.Head == nil || m < 1 {
		return
	}
	if l.Head.Next == l.Head {
		return l.Head.Value.(int)
	}
	head := l.Head
	//获取尾节点
	last := head
	for last.Next != head {
		last = last.Next
	}

	cnt := 0

	//依次后移last->head,没遇m删head节点
	for head != last {
		cnt++
		if cnt == m {
			last.Next = head.Next
			fmt.Println(head.Value, "out")
			cnt = 0
		} else {
			last = last.Next
		}
		head = head.Next
	}
	return head.Value.(int)
}

//递归计算最后报数位置
func JosephusKill2(l *linkedlist.LList, m int) (res int) {
	if l == nil || l.Head == nil || m < 1 {
		return
	}
	if l.Head.Next == l.Head {
		return l.Head.Value.(int)
	}
	n := 1
	cur := l.Head.Next
	for cur != l.Head {
		n++
		cur = cur.Next
	}
	tmp := getLive(n, m)
	cur = l.Head
	for {
		tmp--
		if tmp == 0 {
			break
		}
		cur = cur.Next
	}
	return cur.Value.(int)
}

//报数编号对应关系 B=(A-1)%i+1
//移除报号为m的节点，新老环关系 old=(new+m -1)%i+1
func getLive(i, m int) int {
	if i == 1 {
		return 1
	}
	return (getLive(i-1, m)+m-1)%i + 1
}
