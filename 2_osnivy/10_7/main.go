//А теперь напишите функцию findMaxNegative(),
//которая вернёт максимальное отрицательное число (а в случае отсутствия отрицательных чисел — ошибку)

package main

import "fmt"

func main() {
	arr := []int64{-1, -12, -5, 4, 5, -12, 3, 99}
	//arr := []int64{}
	//var res int64
	//res, _ = findMaxNegative(arr)
	//res, _ = findMostOftenRepeated(arr)
	//var res map[int64]int64
	//res, _ = findMostOftenRepeatedWithMap(arr)
	var res []int64
	//res = trimNegative(arr)
	res = trimLessAverage(arr)
	fmt.Println(res)
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

//Эту функцию можно оптимизировать.
//Нет необходимости сравнивать текущий элемент с элементами,
// которые стоят перед ним. Если там будет число равное текущему,
// то количество его повторений вы уже посчитали раньше, в предыдущих итерациях.
// Если же там не будет равного числа, то, очевидно, нет смысла рассматривать.
//Оптимизируйте функцию findMostOftenRepeatedOptimized() таким образом, чтобы текущий элемент не сравнивался с предыдущими.
func findMostOftenRepeated(array []int64) (mostOften int64, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		//fmt.Println("+")
		currentCount := 0
		for _, numberToCompare := range array[i:] {
			//fmt.Println("+")
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

//На самом деле, эта задача, наверное, лучше всего решается с использованием хэш-мапы (структура map в Go).
// И работать так будет быстрее, всего O(n). Правда, увеличатся затраты по памяти, но память обычно дешевле процессора.

func findMostOftenRepeatedWithMap(array []int64) (mostOften map[int64]int64, err error) {
	ret := make(map[int64]int64)
	if len(array) == 0 {
		return ret, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	for _, number := range array {

		if _, ok := ret[number]; ok {
			ret[number] += 1
		} else {
			ret[number] = 1
		}
	}
	maxm := make(map[int64]int64)
	var count, num int64
	for k, v := range ret {
		//fmt.Println(k, v)
		if v > count {
			count = v
			num = k

		}
	}
	maxm[num] = count
	return maxm, nil
}

//Написать функцию trimNegative() (в качестве аргумента принимает слайс интов и возвращает слайс интов),
// которая удаляет все отрицательные элементы.
func trimNegative(array []int64) []int64 {
	var ret []int64
	for _, i := range array {
		if i >= 0 {
			ret = append(ret, i)
		}
	}

	return ret
}

//Найти среднее арифметическое всех чисел массива.
//Как в прошлой задаче, добавить в новый массив только те числа,
// которые меньше найденного в первом пункте.
func trimLessAverage(array []int64) []int64 {
	var ret []int64
	var sum int64 = 0
	var mean int64 = 0
	if len(array) == 0 {
		return ret
	}
	for _, i := range array {
		sum += i
	}
	mean = sum / int64(len(array))
	fmt.Println(mean)
	for _, i := range array {
		if i < mean {
			ret = append(ret, i)
		}
	}

	return ret
}
