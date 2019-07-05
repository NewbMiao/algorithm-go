package leetcode

import (
	"stack"
	"math"
)

/*
https://leetcode-cn.com/problems/trapping-rain-water/
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

*/
//双指针，当前可接雨水取决左右较低的柱子
func trap(height []int) int {
	l := len(height)
	if l <= 1 {
		return 0
	}
	left := 0
	right := l - 1
	leftMax := 0
	rightMax := 0
	ans := 0
	for left < right {

		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			//fmt.Println("left", left, leftMax, ans)
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			//fmt.Println("right", right, rightMax, ans)

			right--
		}
	}
	return ans
}

//栈
func trap1(height []int) int {
	l := len(height)
	if l <= 1 {
		return 0
	}
	ans := 0
	st := stack.New()
	for i := 0; i < l; i++ {
		for !st.IsEmpty() && height[st.Peek().(int)] < height[i] {
			top, _ := st.Peek().(int)
			st.Pop()
			if st.IsEmpty() {
				break
			}
			pre, _ := st.Peek().(int)
			distance := i - pre - 1
			bound_height := int(math.Min(float64(height[pre]), float64(height[i]))) - height[top]
			ans += distance * bound_height
			//fmt.Println(i, top, pre, distance, bound_height)
		}
		st.Push(i)
	}
	return ans
}
