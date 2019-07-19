package stack

import "fmt"

type S struct {
	Data []interface{}
}

func New() *S {
	return &S{}
}

func (s *S) Pop() (r interface{}, ok bool) {
	if s == nil || len(s.Data) == 0 {
		return
	}
	r = s.Data[0]
	s.Data = s.Data[1:]
	ok = true
	return
}

func (s *S) Push(v interface{}) {
	if s == nil {
		s = new(S)
	}
	s.Data = append([]interface{}{v}, s.Data...)
}

func (s *S) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *S) Clear() {
	if !s.IsEmpty() {
		s.Data = []interface{}{}
	}
}

func (s *S) Peek() interface{} {
	if s == nil || len(s.Data) == 0 {
		return nil
	}
	return s.Data[0]
}

func (s *S) String() string {
	return fmt.Sprintf("%v", s.Data)
}
