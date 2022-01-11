package About_sort


func BubbleSort(arr []int) []int {
	// 冒泡排序
	arr_len := len(arr)
	for i := 0; i <= arr_len; i++ {
		for j := 1; j < arr_len-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
