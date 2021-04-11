package calc

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

type calculator struct {
	a, b float64
	act  string
}

func NewCalculator(a, b float64, act string) calculator {
	return calculator{a, b, act}
}

func (c calculator) Calculate() float64 {
	var ret float64
	switch c.act {
	case "+":
		ret = c.calc_sum()
	case "-":
		ret = c.calc_min()
	case "*":
		ret = c.calc_multi()
	case "/":
		ret = c.calc_div()
	}
	return ret
}

func (c *calculator) calc_sum() float64 {
	return c.a + c.b
}
func (c *calculator) calc_min() float64 {
	return c.a - c.b
}
func (c *calculator) calc_multi() float64 {
	return c.a * c.b
}
func (c *calculator) calc_div() float64 {
	return c.a / c.b
}
