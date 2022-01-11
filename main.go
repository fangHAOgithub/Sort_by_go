package main

import (
	"fmt"
	"sort_test/About_sort"
)

func main() {

	arr := [...]int{99, 51, 41, 2, 31, 55, -8, 999, 564, 12, 1, 0, 33, 28, 556}

	res_arr := About_sort.BubbleSort(arr[:])
	fmt.Println(res_arr) // 从小到大[2 31 41 51 99]
	fmt.Println(arr)
}
