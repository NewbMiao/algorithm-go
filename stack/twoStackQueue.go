package stack

import "github.com/NewbMiao/algorithm-go/kit/stack"

//双栈实现队列
type TwoStackQueue struct {
	SPush *stack.S
	SPop  *stack.S
}

func NewTwoStackQueue() *TwoStackQueue {
	return &TwoStackQueue{
		stack.New(), stack.New(),
	}
}

func (t *TwoStackQueue) Add(i int) {
	t.SPush.Push(i)
}

func (t *TwoStackQueue) Poll() int {
	t.loadQueue()
	i, _ := t.SPop.Pop()
	tmp, _ := i.(int)
	return tmp
}

func (t *TwoStackQueue) Peek() int {
	t.loadQueue()
	tmp, _ := t.SPop.Peek().(int)
	return tmp
}

func (t *TwoStackQueue) loadQueue() {
	if t.SPush.IsEmpty() || !t.SPop.IsEmpty() {
		return
	}
	for !t.SPush.IsEmpty() {
		i, _ := t.SPush.Pop()
		t.SPop.Push(i)
	}
}
