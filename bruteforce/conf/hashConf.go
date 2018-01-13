package conf

type HashConf struct  {
	Value string
	HashType string
}

func NewHash(value string, hashType string) HashConf {
	return HashConf{value, hashType}
}
