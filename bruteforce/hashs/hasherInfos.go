package hashs

type hasherInfos struct {
	benchDescription string
	builder          func() Hasher
}

func NewHasherInfos(benchDescription string, builder func() Hasher) hasherInfos {
	return hasherInfos{benchDescription: benchDescription, builder: builder}
}
