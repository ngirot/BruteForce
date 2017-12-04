package hashs

type Hasher interface {
	Hash(data string) []byte
}