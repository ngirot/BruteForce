package bruteforce

type TesterBuilder struct {
	Build func() Tester
}

type Tester struct {
	Notify func(data string)
	Test   func(data []string) int
}
