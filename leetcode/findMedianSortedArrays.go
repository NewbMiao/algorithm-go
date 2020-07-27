package leetcode

import "math"

/*
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-shu-b/

          left_part          |        right_part
    A[0], A[1], ..., A[i-1]  |  A[i], A[i+1], ..., A[m-1]
    B[0], B[1], ..., B[j-1]  |  B[j], B[j+1], ..., B[n-1]

len(left_part)=len(right_part)
max(left_part) <= min(right_part)
即：
i+j=m - i + n - j + 1
B[j−1]≤A[i] 以及 A[i−1]≤B[j].
*/
func findMedianSortedArrays(A, B []int) float64 {
	m := len(A)
	n := len(B)
	if m > n { // to ensure m<=n
		return findMedianSortedArrays(B, A)
	}
	//在A中二分查找maxLeft
	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && B[j-1] > A[i] {
			iMin = i + 1 // i is too small
		} else if i > iMin && A[i-1] > B[j] {
			iMax = i - 1 // i is too big
		} else { // i is perfect
			maxLeft := 0.0
			if i == 0 { //A[i-1]不存在，只有B[j−1]≤A[i]
				maxLeft = float64(B[j-1])
			} else if j == 0 { //B[j-1]不存在，只有A[i−1]≤B[j]
				maxLeft = float64(A[i-1])
			} else {
				maxLeft = math.Max(float64(A[i-1]), float64(B[j-1]))
			}
			if (m+n)%2 == 1 { //奇数个
				return maxLeft
			}

			minRight := 0.0
			if i == m {
				minRight = float64(B[j])
			} else if j == n {
				minRight = float64(A[i])
			} else {
				minRight = math.Min(float64(B[j]), float64(A[i]))
			}

			return (maxLeft + minRight) / 2.0
		}
	}
	return 0.0
}
