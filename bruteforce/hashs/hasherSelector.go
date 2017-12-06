package hashs

import (
	"errors"
	"strings"
)

func HasherCreator(hashType string) (func() Hasher, error) {

	var hasherMap = make(map[string] func() Hasher)
	hasherMap["sha256"] = func() Hasher {return NewHasherSha256()};
	hasherMap["md5"] = func() Hasher {return NewHasherMd5()};

	if val, present := hasherMap[hashType]; present {
		var creator = func() Hasher {
			return val()
		};
    	return creator, nil;
	}

	return nil, errors.New(hashType + " is not a valid hash type, must be one of " + listAllType(hasherMap));
}

func listAllType(hasherMap (map[string]func() Hasher)) string {
	var values = []string{};

	for k := range hasherMap {
		values = append(values, k)
	}

	return strings.Join(values, ", ");
}
