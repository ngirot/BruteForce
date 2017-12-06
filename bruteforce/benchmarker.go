package bruteforce

import (
	"math"
	"github.com/ngirot/BruteForce/bruteforce/hashs"
	"github.com/ngirot/BruteForce/bruteforce/words"
)

const hashTobench = 10 * 1000 * 1000

func BenchHasher() int {
	var hasher = hashs.NewHasherSha256()

	var chrono = NewChrono()
	chrono.Start()

	for i := 0; i < hashTobench; i++ {
		hasher.Hash("1234567890")
	}

	chrono.End()

	return int(math.Floor(hashTobench / chrono.DurationInSeconds()))
}

func BenchBruter() int {
	var alphabet = words.LoadAlphabet("alphabet.default.data")
	var worder = words.NewWorder(alphabet, 1, 0)

	var chrono = NewChrono()
	chrono.Start()

	for i := 0; i < hashTobench; i++ {
		worder.Next()
	}

	chrono.End()

	return int(math.Floor(hashTobench / chrono.DurationInSeconds()))
}
