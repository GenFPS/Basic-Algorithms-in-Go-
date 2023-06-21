package main

import "fmt"

// реализация префиксной сумммы
func prefixSumArr(arr []int) []int {
	n := len(arr)
	prefic := make([]int, n)
	prefic[0] = arr[0]
	for i := 1; i < n; i++ {
		prefic[i] = prefic[i-1] + arr[i]
	}
	return prefic
}

// нахождение суммы на заданном интервале
func subSumArr(arr []int, l, r int) int {
	if l == 0 {
		return arr[r]
	}
	return arr[r] - arr[l-1]
}

func main() {
	// массив целых чисел
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	// создадим массив префиксной суммы на основе массива выше
	prefixArr := prefixSumArr(arr)
	fmt.Println(prefixArr)

	// 0 <= l, r < len(arr)
	var l, r int = 0, 6
	fmt.Println(subSumArr(prefixArr, l, r))
}
