package conf

type WordConf struct {
	Dictionary string
	Alphabet string
	SaltBefore string
	SaltAfter string
}

func NewWordConf(dictionary string, alphabet string, saltBefore string, saltAfter string) WordConf {
	return WordConf{dictionary, alphabet, saltBefore, saltAfter}
}

func (conf *WordConf) IsForAlphabet() bool {
	return conf.Dictionary == ""
}
