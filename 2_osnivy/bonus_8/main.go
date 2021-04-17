//Напишите функцию, которая ищет факториал натурального числа в итеративном и рекурсивном виде.
//Факториал n — обозначается как n! — произведение чисел от 1 до n.
//Например, факториал 4 — это 4! = 1 * 2 * 3 * 4 = 24
package main

import "fmt"

func main() {
	iter(4)
	fmt.Println(recur(4))
}

func iter(n int64) {
	fak := int64(1)
	for i := int64(1); i <= n; i++ {
		fak *= i
	}
	fmt.Println("ok", fak)
}

func recur(n int64) int64 {
	if n < 2 {
		return 1
	}
	return n * recur(n-1)
}
