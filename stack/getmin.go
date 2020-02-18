package stack

import . "github.com/NewbMiao/AlgorithmCodeNote/kit/stack"
//用栈获取最小值
type MyStack struct {
	SData *S
	SMin  *S //维护最小值
}

func NewMyStack() *MyStack {
	return &MyStack{
		New(), New(),
	}
}

func (s *MyStack) Push(d int) {
	s.SData.Push(d)
	if s.SMin.IsEmpty() || s.SMin.Peek().(int) >= d {
		s.SMin.Push(d)
	}
}

func (s *MyStack) Pop() int {
	tmp, ok := s.SData.Pop()
	if !ok {
		return -1
	}
	if s.SData.IsEmpty() || (s.SMin.IsEmpty() || s.SMin.Peek().(int) >= tmp.(int)) {
		s.SMin.Pop()
	}
	return tmp.(int)

}
func (s *MyStack) GetMin() int {
	if s.SMin.IsEmpty() {
		return -1
	}

	return s.SMin.Peek().(int)
}
