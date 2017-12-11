package bruteforce

import (
	"math"
	"github.com/ngirot/BruteForce/bruteforce/hashs"
	"github.com/ngirot/BruteForce/bruteforce/words"
)

const hashTobench = 10 * 1000 * 1000

func BenchHasherOneCpu(hasher hashs.Hasher) int {
	var chrono = NewChrono()
	chrono.Start()

	for i := 0; i < hashTobench; i++ {
		hasher.Hash("1234567890")
	}

	chrono.End()

	return int(math.Floor(hashTobench / chrono.DurationInSeconds()))
}

func BenchBruter() int {
	var alphabet = words.DefaultAlphabet()
	var worder = words.NewWorder(alphabet, 1, 0)

	var chrono = NewChrono()
	chrono.Start()

	for i := 0; i < hashTobench; i++ {
		worder.Next()
	}

	chrono.End()

	return int(math.Floor(hashTobench / chrono.DurationInSeconds()))
}
