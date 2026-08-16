package differenceofsquares

// so turns out ^ is used for bitwise XOR, not exponents
//
// it also turns out range starts at 0, not 1, oops
func SquareOfSum(n int) int {
	if n == 1 {
		return n
	}
	squareSums := 0
	for int := range n + 1 { // i think this is O(n)? idk Big O notation
		squareSums += int
	}
	return squareSums * squareSums
}

func SumOfSquares(n int) int {
	if n == 1 {
		return n
	}
	sumSquares := 0
	for int := range n + 1 {
		sumSquares += int * int
	}
	return sumSquares
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
