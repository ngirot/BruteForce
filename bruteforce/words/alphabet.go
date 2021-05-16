package words

import (
	"io/ioutil"
)

type Alphabet interface {
	Length() int
	Letter(position int) string
	AsCharset() []string
}

type alphabet struct {
	letters []string
}

func BuildAlphabet(letters []string) Alphabet {
	return &alphabet{letters}
}

func DefaultAlphabet() Alphabet {
	return &alphabet{defaultAlphabet}
}

func LoadAlphabet(file string) Alphabet {
	var data, _ = ioutil.ReadFile(file)
	var letters = make([]string, len(data), len(data))

	for i, c := range data {
		letters[i] = string(c)
	}

	return &alphabet{letters}
}

func (a *alphabet) Length() int {
	return len(a.letters)
}

func (a *alphabet) Letter(position int) string {
	return a.letters[position]
}

func (a *alphabet) AsCharset() []string {
	return a.letters
}
