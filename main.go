package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s HASH", os.Args[0])
		os.Exit(1)
	}

	var hash = os.Args[1]

	fmt.Printf("Start brute-forcing...\n")

	var result = TestAllStrings(testValue(hash), displayValue)

	if result != "" {
		fmt.Printf("Found : %s\n", result)
	} else {
		fmt.Printf("Not found\n")
	}
}

var parsed = 0
func displayValue(data string)  {
	parsed++
	if(parsed%1000000==0) {
		fmt.Printf("Done: %s\n", data)
	}
}

func testValue(hash string) func(string) bool {
	return func(data string) bool {
		return Hash(data) == hash
	}

}