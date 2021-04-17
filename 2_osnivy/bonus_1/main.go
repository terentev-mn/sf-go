//Напишите программу, которая выводит на экран числа от 1 до 100.
//При этом вместо чисел, кратных трём, программа должна выводить слово Fizz, а вместо чисел, кратных пяти — слово Buzz.
//Если число кратно и трём, и пяти, то программа должна выводить слово FizzBuzz.
// почему-то  c case у меня не получилось. понял дело в switch VAR.. не надо указывать переменную
package main

import "fmt"

func main() {
	//ifCase()
	swCase()
}

func swCase() {
	for i := 1; i <= 50; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz", i)
		case i%3 == 0:
			fmt.Println("Fizz", i)
		case i%5 == 0:
			fmt.Println("Buzz", i)
		default:
			fmt.Println(i)
		}
	}
}

func ifCase() {
	for i := 1; i <= 50; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else {
			fmt.Println(i)
		}
	}
}
