package cpu

import "github.com/ngirot/BruteForce/bruteforce/hashs/hashers"

func genericProcessWithWildCard(hasher hashers.Hasher, charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	var expanded = expand(charSet, numberOfWildCards)

	for i := 0; i < len(expanded); i++ {
		var word = saltBefore + expanded[i] + saltAfter
		var currentHash = hasher.Hash([]string{word})[0]

		if hasher.Compare(currentHash, hasher.DecodeInput(expectedDigest)) {
			return expanded[i]
		}
	}

	return ""
}

func expand(charset []string, expansionFactor int) []string {
	if expansionFactor == 0 {
		return []string{}
	}

	var result []string

	for i := 0; i < len(charset); i++ {
		var array = expand(charset, expansionFactor-1)
		if len(array) > 0 {
			for j := 0; j < len(array); j++ {
				result = append(result, charset[i]+array[j])
			}
		} else {
			result = append(result, charset[i])
		}
	}

	return result
}

