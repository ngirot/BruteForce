package bruteforce

import "github.com/ngirot/BruteForce/bruteforce/hashs/hashers"

type TesterBuilder struct {
	Build func() Tester
}

type Tester struct {
	Notify func(data string)
	Test   func(data []string) int
	Target func() string
	Hasher func() hashers.Hasher
}
