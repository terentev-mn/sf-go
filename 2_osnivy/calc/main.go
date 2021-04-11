//Создайте пакет calc, содержащий неэкспортируемую структуру calculator.
// Напишите экспортируемую функцию-конструктор для создания экземпляра структуры calculator.

//Добавьте экспортируемый (публичный) метод Calculate для структуры calculator.
//В качестве параметров метод должен принимать 2 числа типа float64 и строчный оператор (+, -, *, /).
//Возвращать метод должен значение типа float64.

//Добавьте неэкспортируемые (приватные) методы для каждой операции (сложения, вычитания, умножения, деления).
// Каждый приватный метод должен принимать на вход 2 числа типа float64 и возвращать значение типа float64.
// (В методе деления чисел должны быть проверка делителя на равенство 0).
//Метод Calculate должен в зависимости от переданного оператора вызывать один из приватных методов.
// (Если в качестве оператора передан +, то должен быть вызван приватный метод для сложения чисел).

// Создайте пакет main c функцией main().
// Внутри функции main необходимо из консоли считать 2 числа и оператор (как вы это делали в задании предыдущего модуля).
// Затем создать экземпляр структуры calculator из пакета calc и вызвать метод Calculate, передав ему полученные из консоли значения.
// Полученный из метода Calculate результат нужно распечатать в консоль.
package main

import (
	"calc"
	"fmt"
	"strconv"
)

func main() {
	var a, act, b string
	fmt.Println("Введите (через enter) a +(-*/) b")
	_, err1 := fmt.Scanln(&a)
	_, err2 := fmt.Scanln(&act)
	_, err3 := fmt.Scanln(&b)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Не удалось прочитать", err1, err2, err3)
	} else {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		calc.NewCalculator(af, bf, act).Calculate()
	}

}