// +build opencl

package gpu

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/maths"
	"gitlab.com/ngirot/blackcl"
	"strconv"
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

	output, _ := buildUintBuffer(device, make([]uint32, hashSizeInByte*len(datas)))
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
		for j := 0; j < hashSizeInByte/4; j++ {
			endianness.PutUint32(digests[i][j*4:(j+1)*4], newData[i*8+j])
		}
	}

	return digests
}

func genericProcessWithGpu(device *blackcl.Device, kernel *blackcl.Kernel, endianness binary.ByteOrder,
	charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {

	var de, _ = hex.DecodeString(expectedDigest)
	var asUint32 []uint32
	for i := 0; i < len(de)/4; i++ {
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

func detectEndianness(device *blackcl.Device, expectedHash string) binary.ByteOrder {
	kernelTest := device.Kernel(genericKernelCryptName)

	var bigEndianResult = genericHashWithGpu(device, kernelTest, binary.BigEndian, []string{"test"}, len(expectedHash)/2)[0]

	var endianness binary.ByteOrder = binary.LittleEndian
	if hex.EncodeToString(bigEndianResult) == expectedHash {
		endianness = binary.BigEndian
	}

	return endianness
}

func buildByteBuffer(d *blackcl.Device, data []byte) (*blackcl.Bytes, error) {
	v, err := d.NewBytes(len(data))
	if err != nil {
		panic("could not allocate buffer")
	}

	err = <-v.Copy(data)
	if err != nil {
		panic("could not copy data to buffer")
	}
	return v, err
}

func buildUintBuffer(d *blackcl.Device, data []uint32) (*blackcl.Uint32, error) {
	v, err := d.NewUint32(len(data))
	if err != nil {
		panic("could not allocate buffer")
	}

	err = <-v.Copy(data)
	if err != nil {
		panic("could not copy data to buffer")
	}
	return v, err
}

func convert(s string) []byte {
	return []byte(s)
}

const genericKernelCryptName = "crypt_kernel"
const genericKernelCryptAndWorderName = "crypt_and_worder_kernel"

func buildGenericKernel(size int) string {

	const parametrized = `

	uint custom_pow(uint a, uint b) {
	  int i;
	  int result = 1;
	  for(i=0;i<b;i++) {
		result *=a;
	  }
	  return result;
	}
	
	int comp_arrays(uint *a1, __global uint *a2, int size) {
	  int i;
	  for (i=0;i<size;i++) {
		if ( a1[i] != a2[i] ) return 0;
	  }
	  return 1;
	}
	
	__kernel void crypt_and_worder_kernel(
										  __global char *buffer,
										  __global char *salt_before,
										  __global char *salt_after,
										  __global char *char_set,
										  __global uint *sizes,
										  __global char *matching_wild_card,
										  __global uint *digest_expected
										  ) {
	
	  int index=get_global_id(0);
	  int i;
	
	  uint number_of_wildcards = sizes[0];
	  uint size_salt_before = sizes[1];
	  uint size_salt_after = sizes[2];
	  uint size_char_set = sizes[3];
	
	  int word_size = size_salt_before + size_salt_after + number_of_wildcards;
	  global char *my_buffer = &buffer[index * word_size];
	
	  for (i=0;i < word_size;i++) {
		if (i < size_salt_before) {
		  my_buffer[i] = salt_before[i];
		} else if (i < size_salt_before + number_of_wildcards) {
		  int pos = number_of_wildcards - (i - size_salt_before) - 1;
		  int base = size_char_set;
		  int div =  index / custom_pow(base, pos);
		  int current =  div % base;
		  my_buffer[i] = char_set[current];
		} else {
		  my_buffer[i] = salt_after[i - size_salt_before - number_of_wildcards];
		}
	  }
	
	  uint current_digest[___SIZE___];
	  hash(my_buffer, current_digest, word_size);
	  
	  if (comp_arrays(current_digest, digest_expected, ___SIZE___) == 1) {
		global char *r = &matching_wild_card[0];
		int j;
		for (j=0;j<number_of_wildcards;j++) {
		  matching_wild_card[j] = my_buffer[size_salt_before+j];
		}
	  }
	}
	
	
	__kernel void crypt_kernel(__global uint *data_info,__global char *plain_keyMulti,  __global uint *digestMulti){
	
	  int index=get_global_id(0);
	
	  global char *plain_key=&plain_keyMulti[data_info[index]];
	  global uint *digest=&digestMulti[index*8];
	  uint ulen = data_info[index+1] - data_info[index];
	
	  uint current_digest[___SIZE___];
	  hash(plain_key, current_digest, ulen);
	
	  int i;
	  for(i=0;i< ___SIZE___ ;i++) {
		digest[i] = current_digest[i];
	  }
	`

	return strings.ReplaceAll(parametrized, "___SIZE___", strconv.Itoa(size))
}
