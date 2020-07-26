package leetcode

import "math"

/*
https://leetcode-cn.com/problems/container-with-most-water/
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。.
*/
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l := 0
	r := len(height) - 1
	max := 0
	for r > l {
		tmp := int(math.Min(float64(height[r]), float64(height[l])))
		max = int(math.Max(float64(max), float64(tmp*(r-l))))
		//移动短的那个指针
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return max
}
