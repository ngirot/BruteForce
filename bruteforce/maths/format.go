package maths

import "fmt"

func FormatNumber(number int, unit string) string {
	if number < 1000 {
		return fmt.Sprintf("%d %s", number, unit)
	}
	if number < 1000000 {
		return fmt.Sprintf("%.2f k%s", float64(number)/1000, unit)
	}
	return fmt.Sprintf("%.2f M%s", float64(number)/1000000, unit)
}
