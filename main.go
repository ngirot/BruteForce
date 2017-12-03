package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	var bench = flag.Bool("benchmark", false, "Launch a benchmark")
	var value = flag.String("value", "", "The value to be tested");
	flag.Parse()

	if *bench {
		fmt.Printf("One CPU hasher : %d kh/s\n", BenchHasher()/1000)
		fmt.Printf("One CPU word generator : %d kw/s\n", BenchBruter()/1000)
		os.Exit(0)
	}

	if *value == "" {
		flag.Usage();
		os.Exit(1)
	}

	fmt.Printf("Start brute-forcing...\n")

	var chrono = NewChrono()
	chrono.Start()
	var result = Launch(*value)
	chrono.End()

	if result != "" {
		fmt.Printf("Found : %s\n", result)
		fmt.Printf("In %f s", chrono.DurationInSeconds())
	} else {
		fmt.Printf("Not found\n")
	}
}
