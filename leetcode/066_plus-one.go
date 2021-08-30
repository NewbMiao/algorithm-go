package leetcode

func plusOne(digits []int) []int {
	var sum int
	l := len(digits)
	if l < 1 {
		return nil
	}
	sum = 1
	for i := l - 1; i >= 0; i-- {
		if sum > 0 {
			sum += digits[i]
			digits[i] = sum % 10
			sum /= 10
		}
	}
	if sum > 0 {
		return append([]int{sum}, digits...)
	}
	return digits
}
