package main

import (
	"fmt"
)

func linSearchEl(arr []int, el int) {
	for _, v := range arr {
		if v == el {
			fmt.Println(v)
			break
		}
	}
	fmt.Println("This number doesn't exists in this array")
}

func main() {
	var N int
	fmt.Scan(&N) // обозначаем длинну массива через консоль

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		var n int
		fmt.Scan(&n)
		arr[i] = n
	}
	linSearchEl(arr, 100)
}
