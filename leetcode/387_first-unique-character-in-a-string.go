package leetcode

func firstUniqChar(s string) int {
	records := [26]int{}
	for k, v := range s {
		records[v-'a'] = k
	}

	for k, v := range s {
		if records[v-'a'] == k {
			return k
		}
		records[v-'a'] = -1
	}
	return -1
}
