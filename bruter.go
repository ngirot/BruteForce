package main

type tester func(data string) bool
type status func(data string)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}

func TestAllStrings(fn tester, fi status) string {

	for _, value := range alphabet {
		fi(value)
		if fn(value) {
			return value
		}
	}

	return ""
}
