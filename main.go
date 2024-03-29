package main

import (
	"flag"
	"fmt"
	"github.com/ngirot/BruteForce/bruteforce"
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"os"
)

func main() {
	var value = flag.String("value", "", "The value to be tested")
	var alphabet = flag.String("alphabet", "", "The file containing all characters")
	var dictionary = flag.String("dictionary", "", "The file containing all words to be tested")
	var hashType = flag.String("type", "sha256", "The hash type")
	var saltBefore = flag.String("salt-before", "", "The salt added to the end of the generated word")
	var saltAfter = flag.String("salt-after", "", "The salt added to the beginning of the generated word")
	var gpu = flag.Bool("gpu", false, "Use GPU instead of CPU for hash computation")
	var size = flag.String("size", "0", "The password range size, or unique size. Examples: '5', '3-7'")

	flag.Parse()

	var processingUnitConfiguration = conf.NewProcessingUnitConfiguration(*gpu)
	var processingUnitAvailability = processingUnitConfiguration.CheckAvailability()
	if processingUnitAvailability != nil {
		fmt.Printf("%s\n", processingUnitAvailability.Error())
		return
	}

	var wordSizeConf, err = conf.NewWordSizeLimitConf(*size)
	if err != nil {
		fmt.Printf("Parameter '%s' is not a valid size limit. Example of valid limit: '5' or '8-10'", *size)
		return
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
	if result, error := bruteforce.Launch(hashConf, wordConf, processingUnitConfiguration, wordSizeConf); error == nil {
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
