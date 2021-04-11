package main

import (
	"calc"
	"fmt"
	"strconv"
)

func main() {
	var a, act, b string
	var res float64
	fmt.Println("Введите (через enter) a +(-*/) b")
	_, err1 := fmt.Scanln(&a)
	_, err2 := fmt.Scanln(&act)
	_, err3 := fmt.Scanln(&b)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Не удалось прочитать", err1, err2, err3)
	} else {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		res = calc.NewCalculator(af, bf, act).Calculate()
		fmt.Printf("%v %v %v = %.2f\n", a, act, b, res)
	}

}
