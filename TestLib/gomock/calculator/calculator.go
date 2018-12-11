package calculator

// Calculator 是计算器结构体
type Calculator struct {
	math Math
}

func NewCalculator(m Math) *Calculator {

}

func (c *Calculator) add(a, b int) int {
	return c.math.Add(a, b)
}
