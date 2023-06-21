package main

import (
	"fmt"
	"sort"
)

// Дан целочисленный массив, найдите в нем непрерывный подмассив с наибольшей суммой
// (алгоримт Кадане)

func kadane(arr []int) int {
	maxSoFar := 0
	maxEndHere := 0

	for i := 0; i < len(arr); i++ {
		maxEndHere = maxEndHere + arr[i]

		// если максимальная сумма отрицательна, то обнуляем ее
		maxEndHere = maxOfSum(maxEndHere, 0)

		// обновляем результат, если текущая сумма подмассива окажется больше
		maxSoFar = maxOfSum(maxSoFar, maxEndHere)
	}
	return maxSoFar
}

func maxOfSum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{5, 4, -1, 7, 8}

	// Данные строчки кода нужны в случае, если все элементы массива отрицательны
	cloneArr := make([]int, len(arr))
	numbOfElCop := copy(cloneArr, arr)

	fmt.Println(numbOfElCop) // выводит кол-во скопированных элементов
	sort.Ints(cloneArr)

	if cloneArr[len(cloneArr)-1] < 0 {
		fmt.Println(cloneArr[len(cloneArr)-1])
	} else {
		fmt.Println(kadane(arr))
	}
}

/*
P.S. Если в массиве содержатся все отрицательные элементы, то ответом
	 будет являтся максимальное значение элемента в данном массиве (45 строчка кода).
*/
