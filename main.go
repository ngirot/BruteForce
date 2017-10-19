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
	if hash == "--benchmark" {
		var hasherBenchResult = BenchHasher()
		fmt.Printf("One CPU hasher : %d h/s", hasherBenchResult)
		os.Exit(0)
	}


	fmt.Printf("Start brute-forcing...\n")

	var result = Launch(hash)

	if result != "" {
		fmt.Printf("Found : %s\n", result)
	} else {
		fmt.Printf("Not found\n")
	}
}