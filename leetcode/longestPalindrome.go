package leetcode

import (
	"math"
	"strings"
)

func generateDivided(s string, d rune) string {
	if len(s) == 0 {
		return ""
	}
	res := make([]rune, 0, 2*len(s)+1)
	for _, v := range s {
		res = append(res, d, v)
	}
	res = append(res, d)
	return string(res)
}

// Manacher
//https://leetcode-cn.com/problems/two-sum/solution/zhong-xin-kuo-san-dong-tai-gui-hua-by-liweiwei1419/
func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}
	sDivided := generateDivided(s, '#')
	slen := len(sDivided)
	p := make([]int, slen) //回文半径，p[i]-1 实际回文最大长度

	mx := 0 //从开始到现在使用“中心扩散法”能得到的“最长回文子串”能延伸到的最右端的位置。 mx = id + p[id]
	id := 0 //从开始到现在使用“中心扩散法”能得到的“最长回文子串”的中心的位置
	lpdLen := 1
	lpdStr := string(s[0])
	for i := 0; i < slen; i++ {
		//p[j] < mx-i，i 和 j 关于 id 中心对称：即 j = 2 * id - i
		if i < mx {
			p[i] = int(math.Min(float64(p[2*id-i]), float64(mx-i)))
		} else {
			// 走到这里，只可能是因为 i = mx
			//if (i > mx) {
			//	panic("errInfo: i>mx")
			//}
			p[i] = 1
		}
		//p[j] >= mx-i，中心扩展去匹配
		for i-p[i] >= 0 && i+p[i] < slen && sDivided[i-p[i]] == sDivided[i+p[i]] {
			p[i]++
		}
		if i+p[i] > mx {
			mx = i + p[i]
			id = i
		}

		if p[i]-1 > lpdLen {
			lpdLen = p[i] - 1
			lpdStr = sDivided[i-p[i]+1 : i+p[i]]
		}
	}
	return strings.ReplaceAll(lpdStr, "#", "")
}

//s[j,i] 是否为回文：  dp[i, j] = s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1]).
func longestPalindrome1(s string) (res string) {
	if len(s) <= 1 {
		return s
	}
	res = string(s[0])
	llen := 1
	dp := make([][]bool, len(s))
	for k := range dp {
		dp[k] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if i-j+1 > llen {
					llen = i - j + 1
					res = s[j : i+1]
				}
			}
		}
	}
	return res
}

//中心扩展.
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var start, end int

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   //奇数回文序列
		len2 := expandAroundCenter(s, i, i+1) //偶数回文序列
		lenMax := int(math.Max(float64(len1), float64(len2)))
		if lenMax > end-start {
			start = i - (lenMax-1)/2
			end = i + lenMax/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (l int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1 // rightPre-leftPre+1 上一次循环后right+1,left-1,多加了2，所以-1
}
