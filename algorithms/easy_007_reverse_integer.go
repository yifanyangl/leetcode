package algorithms

func reverse(x int) int {
	var digits []int
	for divisor := 1; ; divisor *= 10 {
		resultDivide := x / divisor
		if resultDivide == 0 {
			break
		}
		digits = append(digits, resultDivide%10)
	}

	reversed := 0
	denom := 1
	for i := len(digits) - 1; i >= 0; i-- {
		reversed += denom * digits[i]
		denom *= 10
	}
	if reversed > 1<<31-1 || reversed < -1<<31 {
		reversed = 0
	}
	return reversed
}
