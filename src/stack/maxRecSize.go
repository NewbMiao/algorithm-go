package stack

import (
	"math"
	. "kit/stack"
)

//获取矩阵中连续为1的最大矩形长度
func maxRecSize(m [][]int) int {
	l := len(m)
	if l == 0 {
		return 0
	}
	max := 0
	height := make([]int, len(m[0]))
	for i := 0; i < l; i++ {
		for j := range m[i] {
			if m[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		max = int(math.Max(float64(max), float64(getMaxRec(height))))
	}
	return max
}

func getMaxRec(height []int) int {
	if len(height) == 0 {
		return 0
	}
	st := New()
	max := 0
	for i, v := range height {
		//栈内index对应size值大小递减，遇到大值，取栈顶左右能扩展的最大矩形大小
		for !st.IsEmpty() && height[st.Peek().(int)] >= v {
			top, _ := st.Pop()
			j := top.(int)
			k := -1
			if !st.IsEmpty() {
				k = st.Peek().(int)
			}
			//扩展最左是k+1 到右 i-1
			distance := i - k - 1
			max = int(math.Max(float64(max), float64(distance*height[j])))
		}
		st.Push(i)
	}

	for !st.IsEmpty() {
		top, _ := st.Pop()
		j := top.(int)
		k := -1
		if !st.IsEmpty() {
			k = st.Peek().(int)
		}
		distance := len(height) - k - 1

		max = int(math.Max(float64(max), float64(distance*height[j])))
	}

	return max
}
