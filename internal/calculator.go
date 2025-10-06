package calculator

type Calculator struct{}

func (* Calculator) Add(a,b int) int {
	return a + b
}

func (*Calculator) Subtract(a,b int) int {
	return a - b
}

func (*Calculator) Multiply(a,b int) float32 {
	return float32(a) * float32(b)
}

func (*Calculator) Divide(a,b int) float32 {
	if b == 0 {
		panic("division by zero")
	}
	return float32(a) / float32(b)
}	

func NewCalculator() *Calculator {
	return &Calculator{}
}