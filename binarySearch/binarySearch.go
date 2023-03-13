package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, el int) (int, bool) {
	sort.Ints(arr)
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == el {
			return arr[mid], true
		}
		if arr[mid] < el {
			start = mid + 1
		}
		if arr[mid] > el {
			end = mid - 1
		}
	}
	return 0, false
}
func main() {
	var N int

	fmt.Printf("Введите количество элементов массива:\n")
	fmt.Scan(&N)

	arr := make([]int, N)
	fmt.Printf("Введите  элементы массива:\n")
	for i := 0; i < N; i++ {
		var n int
		fmt.Scan(&n)
		arr[i] = n
	}
	fmt.Printf("Введите число которое вы ищите:\n")
	var el int
	fmt.Scan(&el)

	el, flag := binarySearch(arr, el)
	if flag == false {
		fmt.Println("Число отсутствует в данном массиве!")
	} else if flag == true {
		fmt.Printf("Данное число найдено: %d", el)
	}
}
