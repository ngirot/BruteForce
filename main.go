package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s HASH\n", os.Args[0])
		os.Exit(1)
	}

	var hash = os.Args[1]
	if hash == "--benchmark" {
		fmt.Printf("One CPU hasher : %d kh/s\n", BenchHasher()/1000)
		fmt.Printf("One CPU word generator : %d kw/s\n", BenchBruter()/1000)
		os.Exit(0)
	}

	fmt.Printf("Start brute-forcing...\n")

	var chrono = NewChrono()
	chrono.Start()
	var result = Launch(hash)
	chrono.End()

	if result != "" {
		fmt.Printf("Found : %s\n", result)
		fmt.Printf("In %f s", chrono.DurationInSeconds())
	} else {
		fmt.Printf("Not found\n")
	}
}
