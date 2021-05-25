// +build opencl

package gpu

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/maths"
	"gitlab.com/ngirot/blackcl"
	"strings"
)

func genericHashWithGpu(device *blackcl.Device, kernel *blackcl.Kernel, endianness binary.ByteOrder, datas []string, hashSizeInByte int) [][]byte {

	var sizeInput = make([]uint32, len(datas)+1)
	var encoded = make([][]byte, len(datas))
	var sumEncoded = 0
	sizeInput[0] = 0

	var previous = uint32(0)
	for i, s := range datas {
		/// Addresses
		var address = uint32(len(s)) + previous
		sizeInput[i+1] = address
		previous = address

		// Encoding
		var converted = convert(s)
		encoded[i] = converted
		sumEncoded += len(converted)
	}

	var textInput = make([]byte, sumEncoded)
	var position = 0
	for _, s := range encoded {
		var endPosition = position + len(s)
		copy(textInput[position:endPosition], s)
		position = endPosition
	}

	infos, _ := buildUintBuffer(device, sizeInput)
	defer infos.Release()

	key, _ := buildByteBuffer(device, textInput)
	defer key.Release()

	output, _ := buildUintBuffer(device, make([]uint32, 32*len(datas)))
	defer output.Release()

	err := <-kernel.Global(len(datas)).Local(1).Run(infos, key, output)
	if err != nil {
		panic("could not run kernel")
	}

	newData, err := output.Data()
	if err != nil {
		panic("could not get data from buffer")
	}

	var digests = make([][]byte, len(datas))
	for i := 0; i < len(datas); i++ {
		temp := make([]byte, hashSizeInByte)
		digests[i] = temp
		for j:=0;j<hashSizeInByte/4;j++ {
			endianness.PutUint32(digests[i][j*4:(j+1)*4], newData[i*8+j])
		}
	}

	return digests
}

func genericProcessWithGpu(device *blackcl.Device, kernel *blackcl.Kernel, endianness binary.ByteOrder,
	charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {

	var de, _ = hex.DecodeString(expectedDigest)
	var asUint32 []uint32
	for i:=0;i<len(de)/4;i++ {
		asUint32 = append(asUint32, endianness.Uint32(de[i*4:(i+1)*4]))
	}

	digestExpected, _ := buildUintBuffer(device, asUint32)
	defer digestExpected.Release()
	numberOfWordToTest := maths.PowInt(len(charSet), numberOfWildCards)
	wordSize := len(saltBefore) + len(saltAfter) + numberOfWildCards

	bBuffer, _ := buildByteBuffer(device, make([]byte, numberOfWordToTest*wordSize))
	defer bBuffer.Release()

	bSaltBefore, _ := buildByteBuffer(device, []byte(saltBefore+"X"))
	defer bSaltBefore.Release()

	mergedCharset := []byte(strings.Join(charSet, ""))
	sizes, _ := buildUintBuffer(device, []uint32{
		uint32(numberOfWildCards),
		uint32(len(saltBefore)),
		uint32(len(saltAfter)),
		uint32(len(mergedCharset)),
	})
	defer sizes.Release()

	bSaltAfter, _ := buildByteBuffer(device, []byte(saltAfter+"X"))
	defer bSaltAfter.Release()

	bMatchingWildcard, _ := buildByteBuffer(device, make([]byte, numberOfWildCards))
	defer bMatchingWildcard.Release()

	bCharSet, _ := buildByteBuffer(device, mergedCharset)
	defer bCharSet.Release()

	err := <-kernel.Global(numberOfWordToTest).Local(1).Run(
		bBuffer,
		bSaltBefore,
		bSaltAfter,
		bCharSet,
		sizes,
		bMatchingWildcard,
		digestExpected,
	)

	if err != nil {
		panic("could not run kernel")
	}

	result, err := bMatchingWildcard.Data()
	if err != nil {
		panic("could not get data from buffer")
	}

	for i := 0; i < numberOfWildCards; i++ {
		if result[i] != 0 {
			return string(result)
		}
	}

	return ""
}
