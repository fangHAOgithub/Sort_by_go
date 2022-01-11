package About_sort

//排序算法汇总，https://github.com/hustcc/JS-Sorting-Algorithm

func BubbleSort(arr []int) []int {
	// 冒泡排序， 时间复杂度 平均和最坏是O(n^2)， 最好是O(n)， 稳定排序
	// 原理概述，通过两两比较，前者大于后者则进行交换，外层循环第一次循环走完，最大的数会在末尾
	// 依次进行循环，直到数组有序
	arr_len := len(arr)
	for i := 0; i <= arr_len; i++ {
		for j := 1; j < arr_len-i; j++ {
			if arr[j] < arr[j-1] { // 两两比较
				arr[j], arr[j-1] = arr[j-1], arr[j] // 交换操作
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	// 选择排序，时间复杂度都是 O(n^2)，不稳定
	// 原理概述
	// 1、先找到最大/小的元素，放在初始位置
	// 2、再从剩余元素继续寻找最大/小的元素，放在已排序的末尾
	// 3、重复第二步，直到元素排好序
	len_arr := len(arr)
	for i := 0; i < len_arr-1; i++ {
		min_num := i
		for j := i + 1; j < len_arr; j++ {
			if arr[min_num] > arr[j] {
				min_num = j
			}
		}
		arr[i], arr[min_num] = arr[min_num], arr[i]
	}
	return arr
}

func InsertionSort(arr []int) []int {
	// 插入排序，时间复杂度最坏 O(n^2)，最好 O(n)，稳定的排序算法
	// 1、将第一个元素看成有序，后面的元素是未排序
	// 2、从头到尾扫描未排序的序列，扫描每个元素，插入到有序序列的合适位置（如未排序和已排序的相等，则放在其后面）
	len_arr := len(arr)
	for i := 0; i < len_arr; i++ {
		order := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > order {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				arr[j+1] = order
				break
			}
		}
	}
	return arr
}
