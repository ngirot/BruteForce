package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Start brute-forcing...\n")

	var result = TestAllStrings(testValue, displayValue)

	if result != "" {
		fmt.Printf("Found : %s\n", result)
	} else {
		fmt.Printf("Not found\n")
	}
}

func displayValue(data string)  {
	fmt.Printf("Done: %s\n", data)
}

func testValue(data string) bool {
	return Hash(data) == "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
}