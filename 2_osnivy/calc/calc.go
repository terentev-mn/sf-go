package calc

import "fmt"

type calculator struct {
	a, b float64
	act  string
}

func NewCalculator(a, b float64, act string) calculator {
	return calculator{a, b, act}
}

func (c calculator) Calculate() {
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
	fmt.Printf("%v %v %v = %.2f\n", c.a, c.act, c.b, ret)
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
