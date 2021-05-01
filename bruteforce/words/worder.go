package words

type Worder interface {
	Next() string
	GetCharsetIfAvailable() Alphabet
}