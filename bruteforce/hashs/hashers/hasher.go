package hashers

type Hasher interface {
	Example() string
	DecodeInput(data string) []byte
	Hash(datas []string) [][]byte
	Compare(transformedData []byte, referenceData []byte) bool
	ProcessWithWildcard(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string
	IsValid(data string) bool
}

