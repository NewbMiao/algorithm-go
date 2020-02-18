package stack

import . "github.com/NewbMiao/AlgorithmCodeNote/kit/stack"

//用栈来排序栈
func sortStackByStack(s *S) {
	help := New()
	if s.IsEmpty() {
		return
	}
	for !s.IsEmpty() {
		cur, _ := s.Pop()
		for !help.IsEmpty() && help.Peek().(int) < cur.(int) {
			tmp, _ := help.Pop()
			s.Push(tmp)
		}
		help.Push(cur)
	}

	for !help.IsEmpty() {
		tmp, _ := help.Pop()
		s.Push(tmp)
	}
}
