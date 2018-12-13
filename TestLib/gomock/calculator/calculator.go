package calculator

// Calculator 是计算器结构体
type Calculator struct {
	math Math
}

// Add return sum of a and b
func (c *Calculator) Add(a, b int) int {
	return c.math.Add(a, b)
}

// Sub return
func (c *Calculator) Sub(a []int, b *int) int {
	return c.math.Sub(a, b)
}
