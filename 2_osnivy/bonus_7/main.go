//Напишите программу, выводящую первые 20 простых чисел.
//Простое число — натуральное число больше 1, которое делится без остатка только на себя и на 1.
package main

import "fmt"

func main() {
	natur1()
}

func natur1() {
	n := 0
	for i := 1; n < 20; i++ {
		//fmt.Println(n, i)
		if check(i) {
			n++
			fmt.Println("Bingo", n, i)
		}
		if n == 20 {
			break
		}
	}
	fmt.Println("ok")
}

func check(number int) bool {
	for i := 2; i <= number; i++ {
		if (number%i == 0) && (i != number) {
			return false
		}

	}
	return true
}
