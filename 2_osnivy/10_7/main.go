//А теперь напишите функцию findMaxNegative(),
//которая вернёт максимальное отрицательное число (а в случае отсутствия отрицательных чисел — ошибку)

package main

import "fmt"

func main() {
	arr := []int64{-1, -12, -5}
	var min int64
	min, _ = findMaxNegative(arr)
	fmt.Println(min)
}

func findMaxNegative(arr []int64) (int64, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("slice is null len")
	}
	min := arr[0]
	for _, i := range arr[1:] {
		//fmt.Println(i)
		if i < min {
			min = i
		}
	}
	return min, nil
}
