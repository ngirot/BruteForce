package conf

type WordConf struct {
	Dictionary string
	Alphabet string
	Salt string
}

func NewWordConf(dictionary string, alphabet string, salt string) WordConf {
	return WordConf{dictionary, alphabet, salt}
}
