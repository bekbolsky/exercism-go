package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sumOfSquares := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}
	return sumOfSquares
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
