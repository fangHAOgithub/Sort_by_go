package main

import (
	"fmt"
	"sort_test/About_sort"
)

func main() {

	arr := [...]int{99, 51, 41, 2, 31, 55, -8, 999, 564, 12, 1, 0, 33, 28, 556}
	fmt.Println("原始的数据：", arr)
	//res_arr := About_sort.BubbleSort(arr[:])	// 冒泡排序
	//fmt.Println("冒泡排序：",res_arr) // 从小到大[2 31 41 51 99]

	//res_arr := About_sort.SelectionSort(arr[:])
	//fmt.Println("选择排序：",res_arr)

	res_arr := About_sort.InsertionSort(arr[:])
	fmt.Println("插入排序：", res_arr)

}
