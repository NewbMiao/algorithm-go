package misc

import (
	"fmt"
)

func findSumArr(arr, res []int, sum, index int) {
	if sum <= 0 || index < 0 {
		return
	}
	if sum == arr[index] {
		for _, v := range res {
			fmt.Printf("%d,", v)
		}
		fmt.Printf("%d\n", arr[index])
	}
	//res = append([]int{arr[index]}, res...) //放res头部
	res = append(res, arr[index]) //尾部
	findSumArr(arr, res, sum-arr[index], index-1)
	//res = res[1:] //res头部移除
	res = res[:len(res)-1] //尾部移除
	findSumArr(arr, res, sum, index-1)
}
