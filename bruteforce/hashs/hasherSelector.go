package hashs

import (
	"errors"
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"sort"
	"strings"
)

func HasherCreator(hashType string) (func() Hasher, error) {
	var hasherMap = buildHasherMap()

	if description, present := hasherMap[hashType]; present {
		var creator = func() Hasher {
			return description.builder()
		}
		return creator, nil
	}

	return nil, errors.New(hashType + " is not a valid hash type, must be one of " + strings.Join(AllHasherTypes(), ", "))
}

func HasherBenchmarkDescription(hashType string) string {
	var hasherMap = buildHasherMap()
	if description, present := hasherMap[hashType]; present {
		return description.benchDescription
	}

	return hashType
}

func AllHasherTypes() []string {
	var values []string

	for k := range buildHasherMap() {
		values = append(values, k)
	}

	sort.Strings(values)

	return values
}

func IsValidhash(hash conf.HashConf) bool {
	if hasherCreator, e := HasherCreator(hash.HashType); e == nil {
		return hasherCreator().IsValid(hash.Value)
	} else {
		return true
	}
}

func buildHasherMap() map[string]hasherInfos {
	var hasherMap = make(map[string]hasherInfos)
	hasherMap["sha256"] = NewHasherInfos("SHA256", func() Hasher { return NewHasherSha256() })
	hasherMap["md5"] = NewHasherInfos("MD5", func() Hasher { return NewHasherMd5() })
	hasherMap["sha1"] = NewHasherInfos("SHA1", func() Hasher { return NewHasherSha1() })
	hasherMap["sha512"] = NewHasherInfos("SHA512", func() Hasher { return NewHasherSha512() })
	hasherMap["bcrypt"] = NewHasherInfos("bcrypt (cost 10)", func() Hasher { return NewHasherBcrypt() })
	return hasherMap
}
