package main
// 6.8 Напишите программу, которая считывает первое число, затем арифметический оператор (+, -, *, /),
// затем второе число, после чего, в зависимости от арифметического оператора, производит нужное действие и выводит строку с результатом.
import (
	"fmt"
	"strconv"
)

func main() {
	var a, act, b string
	_, err1 := fmt.Scanln(&a)
	_, err2 := fmt.Scanln(&act)
	_, err3 := fmt.Scanln(&b)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Не удалось прочитать", err1, err2, err3)
	} else {
		doAct(a, b, act)
	}
}

func doAct(a string, b string, act string) {
	var ret float64
	//aa, _ := strconv.Atoi(a)
	//bb, _ := strconv.Atoi(b)
	aa, _ := strconv.ParseFloat(a, 64)
	bb, _ := strconv.ParseFloat(b, 64)

	switch act {
	case "+":
		ret = aa + bb
	case "-":
		ret = aa - bb
	case "*":
		ret = aa * bb
	case "/":
		ret = aa / bb
	}
	fmt.Printf("%v %v %v = %.2f\n", a, act, b, ret)

}
