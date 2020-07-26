package array

import (
	"sort"
)

//获取arr中和为sum的所有子数组.
func FindSumArr(arr []int, sum int) (sub [][]int) {
	sort.Ints(arr)                         //升序arr
	index := len(arr) - 1                  //倒着计算子数组
	res := []int{}                         //中间子数组结果集
	return getSumArr(arr, res, sum, index) //res尾部操作
	//return getSumArr2(arr, res, sum, index)  //res头部操作
}
func getSumArr(arr, res []int, sum, index int) (sub [][]int) {
	if sum <= 0 || index < 0 {
		return
	}
	if sum == arr[index] {
		tmp := []int{}
		tmp = append(tmp, res...)
		tmp = append(tmp, arr[index])
		sub = append(sub, tmp)
	}
	res = append(res, arr[index]) //尾部
	//有arr[index]组成的子数组
	sub = append(sub, getSumArr(arr, res, sum-arr[index], index-1)...)
	res = res[:len(res)-1] //尾部移除
	//无arr[index]组成的子数组
	sub = append(sub, getSumArr(arr, res, sum, index-1)...)
	return
}

func getSumArr2(arr, res []int, sum, index int) (sub [][]int) {
	if sum <= 0 || index < 0 {
		return
	}
	if sum == arr[index] {
		tmp := []int{}
		tmp = append(tmp, res...)
		tmp = append(tmp, arr[index])
		sub = append(sub, tmp)
	}
	res = append([]int{arr[index]}, res...) //放res头部
	//有arr[index]组成的子数组
	sub = append(sub, getSumArr(arr, res, sum-arr[index], index-1)...)
	res = res[1:] //res头部移除
	//无arr[index]组成的子数组
	sub = append(sub, getSumArr(arr, res, sum, index-1)...)
	return
}
