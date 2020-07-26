package stack

import "github.com/NewbMiao/algorithm-go/kit/stack"

//递归原地逆序栈
func Reverse(s *stack.S) {
	if s.IsEmpty() {
		return
	}
	i := getAndRemoveLastElement(s)
	Reverse(s)
	s.Push(i)
}

//获取并移除最后一个元素
func getAndRemoveLastElement(s *stack.S) interface{} {
	res, _ := s.Pop()
	if s.IsEmpty() {
		return res
	}
	last := getAndRemoveLastElement(s)
	s.Push(res)
	return last
}
