package hashs

import (
	"errors"
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"strings"
)

func HasherCreator(hashType string) (func() Hasher, error) {
	var hasherMap = buildHasherMap()

	if val, present := hasherMap[hashType]; present {
		var creator = func() Hasher {
			return val()
		}
    	return creator, nil
	}

	return nil, errors.New(hashType + " is not a valid hash type, must be one of " + strings.Join(AllHasherTypes(), ", "))
}

func AllHasherTypes() []string {
	var values []string

	for k := range buildHasherMap() {
		values = append(values, k)
	}

	return values
}

func IsValidhash(hash conf.HashConf) bool {
	if hasherCreator, e := HasherCreator(hash.HashType); e == nil {
		return hasherCreator().IsValid(hash.Value)
	} else {
		return true
	}
}

func buildHasherMap() map[string] func() Hasher{
	var hasherMap = make(map[string] func() Hasher)
	hasherMap["sha256"] = func() Hasher {return NewHasherSha256()}
	hasherMap["md5"] = func() Hasher {return NewHasherMd5()}
	hasherMap["sha1"] = func() Hasher {return NewHasherSha1()}
	hasherMap["sha512"] = func() Hasher {return NewHasherSha512()}
	return hasherMap
}
