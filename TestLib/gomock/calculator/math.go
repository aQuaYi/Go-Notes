package calculator

//go:generate mockgen -source math.go -destination math_mock.go -package calculator

// Math interface define add and sub operation
type Math interface {
	Add(int, int) int
	Sub(int, int) int
}
