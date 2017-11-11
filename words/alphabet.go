package words

import (
	"io/ioutil"
)

type Alphabet struct {
	letters []string
}

func BuildAlphabet(letters []string) Alphabet {
	return &alphabet{letters}
}

func LoadAlphabet(file string) Alphabet {
	var data,_ = ioutil.ReadFile(file)
	var letters = make([]string, len(data), len(data))

	for i, c := range data {
		letters[i] = string(c)
	}

	return Alphabet{letters}
}

func (a *Alphabet) length() int {
	return len(a.letters)
}

func (a *Alphabet) letter(position int) string {
	return a.letters[position]
}
