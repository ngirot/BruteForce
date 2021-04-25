package maths

func PowInt(base int, exponent int) int {
	var result = 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
