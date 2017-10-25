package main

import (
	"fmt"
	"os"
	"time"
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

	var start = time.Now().UnixNano()
	var result = Launch(hash)
	var end = time.Now().UnixNano()

	if result != "" {
		fmt.Printf("Found : %s\n", result)
		fmt.Printf("In %f s", (float64(end-start) / 1000000000))
	} else {
		fmt.Printf("Not found\n")
	}
}
