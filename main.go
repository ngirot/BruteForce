package main

import (
	"flag"
	"fmt"
	"github.com/ngirot/BruteForce/bruteforce"
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/hashs"
	"github.com/ngirot/BruteForce/bruteforce/maths"
	"os"
)

func main() {
	var bench = flag.Bool("benchmark", false, "Launch a benchmark")
	var value = flag.String("value", "", "The value to be tested")
	var alphabet = flag.String("alphabet", "", "The file containing all characters")
	var dictionary = flag.String("dictionary", "", "The file containing all words to be tested")
	var hashType = flag.String("type", "sha256", "The hash type")
	var saltBefore = flag.String("salt-before", "", "The salt added to the end of the generated word")
	var saltAfter = flag.String("salt-after", "", "The salt added to the beginning of the generated word")
	var gpu = flag.Bool("gpu", false, "Use GPU instead of CPU for hash computation")

	flag.Parse()

	var processingUnitConfiguration = conf.NewProcessingUnitConfiguration(*gpu)
	var processingUnitAvailability = processingUnitConfiguration.CheckAvailability()
	if processingUnitAvailability != nil {
		fmt.Printf("%s\n", processingUnitAvailability.Error())
		return
	}

	if *bench {
		var types = hashs.AllHasherTypes()
		for _, t := range types {
			var hasherCreator, _ = hashs.HasherCreator(t, processingUnitConfiguration)
			var description = hashs.HasherBenchmarkDescription(t)

			fmt.Printf("=== %s ===\n", description)
			fmt.Printf("One CPU   : ")
			var timeOneCpu = bruteforce.BenchHasherOneCpu(hasherCreator)
			fmt.Printf("%s\n", maths.FormatNumber(timeOneCpu, "h/s"))

			fmt.Printf("Multi CPU : ")
			var timeMultiCpu = bruteforce.BenchHasherMultiCpu(hasherCreator)
			fmt.Printf("%s\n", maths.FormatNumber(timeMultiCpu, "h/s"))

			fmt.Print("\n")
		}

		fmt.Printf("=== Word generator ===\n")
		fmt.Printf("One CPU   : ")
		fmt.Printf("%s\n", maths.FormatNumber(bruteforce.BenchWorderOneCpu()/1000, "w/s"))
		fmt.Printf("Multi CPU : ")
		fmt.Printf("%s\n", maths.FormatNumber(bruteforce.BenchWorderMultiCpu()/1000, "w/s"))
		os.Exit(0)
	}

	if *value == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Start brute-forcing '%s' (%s)...\n", *value, *hashType)

	var hashConf = conf.NewHash(*value, *hashType)
	var wordConf = conf.NewWordConf(*dictionary, *alphabet, *saltBefore, *saltAfter)

	var chrono = bruteforce.NewChrono()
	chrono.Start()
	if result, error := bruteforce.Launch(hashConf, wordConf, processingUnitConfiguration); error == nil {
		chrono.End()

		if result != "" {
			fmt.Printf("\rFound: %s in %d s\n", result, chrono.DurationInRoundedSeconds())
		} else {
			fmt.Printf("\rNothing found\n")
		}
	} else {
		fmt.Printf("Error initializing Hasher '%s':\n", *hashType)
		fmt.Printf("%s\n", error)
	}
}


