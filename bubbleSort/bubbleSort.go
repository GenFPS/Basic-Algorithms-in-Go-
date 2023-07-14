package main

import (
	"fmt"
)

func bubbleSort(slice []int) []int {
	for i := len(slice) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			var tmp int
			// если предыдущий элемент больше следующего
			if slice[j] > slice[j+1] {
				// то сохраняем значение текущего элемента
				tmp = slice[j]
				// меняем значение текущего элемента на следующий
				slice[j] = slice[j+1]
				// а значение следущего элемента присвивается значением предыдущего элемента, который был больше
				slice[j+1] = tmp
			}
		}
	}
	return slice
}

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2}
	fmt.Println(bubbleSort(arr)) // [1 2 3 5 6 7 8]
}
