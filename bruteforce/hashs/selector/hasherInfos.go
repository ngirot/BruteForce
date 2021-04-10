package selector

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
)

type HasherInfos struct {
	benchDescription string
	builder          func() hashers.Hasher
}

func NewHasherInfos(benchDescription string, builder func() hashers.Hasher) HasherInfos {
	return HasherInfos{benchDescription: benchDescription, builder: builder}
}

func (hi *HasherInfos) Description() string {
	return hi.benchDescription
}

func (hi *HasherInfos) Build() hashers.Hasher {
	return hi.builder()
}
