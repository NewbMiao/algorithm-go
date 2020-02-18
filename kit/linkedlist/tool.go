package linkedlist

func Arr2LList(arr []interface{}) *LList {
	if len(arr) == 0 {
		return nil
	}
	l := NewList()
	for i := 0; i < len(arr); i++ {
		l.Push(arr[i])
	}
	return l
}

func CvtLNode2arr(l *LNode) (res []interface{}) {
	cur := l
	for cur != nil {
		res = append(res, cur.Value)
		cur = cur.Next
	}
	return
}
