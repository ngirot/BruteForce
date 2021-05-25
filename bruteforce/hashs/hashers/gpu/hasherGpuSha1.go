// +build opencl

package gpu

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"github.com/ngirot/BruteForce/bruteforce/maths"
	"gitlab.com/ngirot/blackcl"
	"strings"
)

type hasherGpuSha1 struct {
	device           *blackcl.Device
	kernelDictionary *blackcl.Kernel
	kernelAlphabet   *blackcl.Kernel
	endianness       binary.ByteOrder
}

func NewHasherGpuSha1() hashers.Hasher {
	gpus, err := blackcl.GetDevices(blackcl.DeviceTypeGPU)
	if err == nil {
		for _, device := range gpus {
			device.AddProgram(kernelSourceImport2)
			kernelTest := device.Kernel("sha1_crypt_kernel")

			var bigEndianResult = hashWithGpu2(device, kernelTest, binary.BigEndian, []string{"test"})[0]

			var endianness binary.ByteOrder = binary.LittleEndian
			if hex.EncodeToString(bigEndianResult) == "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" {
				endianness = binary.BigEndian
			}

			return &hasherGpuSha1{device, device.Kernel("sha1_crypt_kernel"), device.Kernel("sha1_crypt_and_worder"), endianness}
		}
	}

	return nil
}

func (h *hasherGpuSha1) Example() string {
	return hex.EncodeToString(h.Hash([]string{"1234567890"})[0])
}

func (h *hasherGpuSha1) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherGpuSha1) Hash(datas []string) [][]byte {
	return hashWithGpu2(h.device, h.kernelDictionary, h.endianness, datas)
}

func (h *hasherGpuSha1) IsValid(data string) bool {
	return hashers.GenericBase64Validator(h, data)
}

func (h *hasherGpuSha1) Compare(transformedData []byte, referenceData []byte) bool {
	return bytes.Equal(transformedData, referenceData)
}

func convert2(s string) []byte {
	return []byte(s)
}

func buildByteBuffer2(d *blackcl.Device, data []byte) (*blackcl.Bytes, error) {
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

func buildUintBuffer2(d *blackcl.Device, data []uint32) (*blackcl.Uint32, error) {
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

func (h *hasherGpuSha1) ProcessWithWildcard(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	return processWithGpu2(h.device, h.kernelAlphabet, h.endianness,
		charSet, saltBefore, saltAfter, numberOfWildCards, expectedDigest)
}

func processWithGpu2(device *blackcl.Device, kernel *blackcl.Kernel, endianness binary.ByteOrder,
	charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {

	var de, _ = hex.DecodeString(expectedDigest)
	digestExpected, _ := buildUintBuffer2(device, []uint32{
		endianness.Uint32(de[0:4]),
		endianness.Uint32(de[4:8]),
		endianness.Uint32(de[8:12]),
		endianness.Uint32(de[12:16]),
		endianness.Uint32(de[16:20]),
	})
	defer digestExpected.Release()
	numberOfWordToTest := maths.PowInt(len(charSet), numberOfWildCards)
	wordSize := len(saltBefore) + len(saltAfter) + numberOfWildCards

	bBuffer, _ := buildByteBuffer2(device, make([]byte, numberOfWordToTest*wordSize))
	defer bBuffer.Release()

	bSaltBefore, _ := buildByteBuffer2(device, []byte(saltBefore+"X"))
	defer bSaltBefore.Release()

	mergedCharset := []byte(strings.Join(charSet, ""))
	sizes, _ := buildUintBuffer2(device, []uint32{
		uint32(numberOfWildCards),
		uint32(len(saltBefore)),
		uint32(len(saltAfter)),
		uint32(len(mergedCharset)),
	})
	defer sizes.Release()

	bSaltAfter, _ := buildByteBuffer2(device, []byte(saltAfter+"X"))
	defer bSaltAfter.Release()

	bMatchingWildcard, _ := buildByteBuffer2(device, make([]byte, numberOfWildCards))
	defer bMatchingWildcard.Release()

	bCharSet, _ := buildByteBuffer2(device, mergedCharset)
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

func hashWithGpu2(device *blackcl.Device, kernel *blackcl.Kernel, endianness binary.ByteOrder, datas []string) [][]byte {

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
		var converted = convert2(s)
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

	infos, _ := buildUintBuffer2(device, sizeInput)
	defer infos.Release()

	key, _ := buildByteBuffer2(device, textInput)
	defer key.Release()

	output, _ := buildUintBuffer2(device, make([]uint32, 20*len(datas)))
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
		temp := make([]byte, 20)
		digests[i] = temp

		endianness.PutUint32(digests[i][0:4], newData[i*8])
		endianness.PutUint32(digests[i][4:8], newData[i*8+1])
		endianness.PutUint32(digests[i][8:12], newData[i*8+2])
		endianness.PutUint32(digests[i][12:16], newData[i*8+3])
		endianness.PutUint32(digests[i][16:20], newData[i*8+4])
	}

	return digests
}

// https://github.com/Fruneng/opencl_sha_al_im
const kernelSourceImport2 = `
#ifdef cl_khr_byte_addressable_store
#pragma OPENCL EXTENSION cl_khr_byte_addressable_store : disable
#endif

#ifdef cl_nv_pragma_unroll
#define NVIDIA
#pragma OPENCL EXTENSION cl_nv_pragma_unroll : enable
#endif

#ifdef NVIDIA
inline uint SWAP32(uint x)
{
	x = rotate(x, 16U);
	return ((x & 0x00FF00FF) << 8) + ((x >> 8) & 0x00FF00FF);
}
#else
#define SWAP32(a)	(as_uint(as_uchar4(a).wzyx))
#endif

#define K0  0x5A827999
#define K1  0x6ED9EBA1
#define K2  0x8F1BBCDC
#define K3  0xCA62C1D6

#define H1 0x67452301
#define H2 0xEFCDAB89
#define H3 0x98BADCFE
#define H4 0x10325476
#define H5 0xC3D2E1F0

#ifndef uint32_t
#define uint32_t unsigned int
#endif

uint32_t SHA1CircularShift(int bits, uint32_t word)
{
	return ((word << bits) & 0xFFFFFFFF) | (word) >> (32 - (bits));
}

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

void hash(global char *plain_key, uint *digest, uint ulen) {
int t, gid, msg_pad;
    int stop, mmod;
    uint i, item, total;
    uint W[80], temp, A,B,C,D,E;
	int current_pad;

	msg_pad=0;

	total = ulen%64>=56?2:1 + ulen/64;

	//printf("ulen: %u total:%u\n", ulen, total);

    digest[0] = 0x67452301;
	digest[1] = 0xEFCDAB89;
	digest[2] = 0x98BADCFE;
	digest[3] = 0x10325476;
	digest[4] = 0xC3D2E1F0;
	for(item=0; item<total; item++)
	{

		A = digest[0];
		B = digest[1];
		C = digest[2];
		D = digest[3];
		E = digest[4];

	#pragma unroll
		for (t = 0; t < 80; t++){
		W[t] = 0x00000000;
		}
		msg_pad=item*64;
		if(ulen > msg_pad)
		{
			current_pad = (ulen-msg_pad)>64?64:(ulen-msg_pad);
		}
		else
		{
			current_pad =-1;		
		}

		//printf("current_pad: %d\n",current_pad);
		if(current_pad>0)
		{
			i=current_pad;

			stop =  i/4;
			//printf("i:%d, stop: %d msg_pad:%d\n",i,stop, msg_pad);
			for (t = 0 ; t < stop ; t++){
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 1]) << 16;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 2]) << 8;
				W[t] |= (uchar)  plain_key[msg_pad + t * 4 + 3];
				//printf("W[%u]: %u\n",t,W[t]);
			}
			mmod = i % 4;
			if ( mmod == 3){
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 1]) << 16;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 2]) << 8;
				W[t] |=  ((uchar) 0x80) ;
			} else if (mmod == 2) {
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 1]) << 16;
				W[t] |=  0x8000 ;
			} else if (mmod == 1) {
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |=  0x800000 ;
			} else /*if (mmod == 0)*/ {
				W[t] =  0x80000000 ;
			}
			
			if (current_pad<56)
			{
				W[15] =  ulen*8 ;
				//printf("w[15] :%u\n", W[15]);
			}
		}
		else if(current_pad <0)
		{
			if( ulen%64==0)
				W[0]=0x80000000;
			W[15]=ulen*8;
			//printf("w[15] :%u\n", W[15]);
		}

		

		for (t = 16; t < 80; t++)
		{
			W[t] = SHA1CircularShift(1, W[t - 3] ^ W[t - 8] ^ W[t - 14] ^ W[t - 16]);
		}

		for (t = 0; t < 20; t++)
		{
			temp = SHA1CircularShift(5, A) +
				((B & C) | ((~B) & D)) + E + W[t] + K0;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		for (t = 20; t < 40; t++)
		{
			temp = SHA1CircularShift(5, A) + (B ^ C ^ D) + E + W[t] + K1;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		for (t = 40; t < 60; t++)
		{
			temp = SHA1CircularShift(5, A) +
				((B & C) | (B & D) | (C & D)) + E + W[t] + K2;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		for (t = 60; t < 80; t++)
		{
			temp = SHA1CircularShift(5, A) + (B ^ C ^ D) + E + W[t] + K3;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		digest[0] = (digest[0] + A) & 0xFFFFFFFF;
		digest[1] = (digest[1] + B) & 0xFFFFFFFF;
		digest[2] = (digest[2] + C) & 0xFFFFFFFF;
		digest[3] = (digest[3] + D) & 0xFFFFFFFF;
		digest[4] = (digest[4] + E) & 0xFFFFFFFF;

		//for(i=0;i<80;i++)
			//printf("W[%u]: %u\n", i,W[i] );

		//printf("%u\n",  digest[0]);
		//printf("%u\n",  digest[1]);
		//printf("%u\n",  digest[2]);
		//printf("%u\n",  digest[3]);
		//printf("%u\n",  digest[4]);
	}
}

__kernel void sha1_crypt_and_worder(
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

  uint current_digest[20];
  hash(my_buffer, current_digest, word_size);
  
  if (comp_arrays(current_digest, digest_expected, 20) == 1) {
    global char *r = &matching_wild_card[0];
    int j;
    for (j=0;j<number_of_wildcards;j++) {
      matching_wild_card[j] = my_buffer[size_salt_before+j];
    }
  }
}


__kernel void sha1_crypt_kernel(__global uint *data_info,__global char *plain_keyMulti,  __global uint *digestMulti){

  int index=get_global_id(0);

  global char *plain_key=&plain_keyMulti[data_info[index]];
  global uint *digest=&digestMulti[index*8];
  uint ulen = data_info[index+1] - data_info[index];

  uint current_digest[20];
  hash(plain_key, current_digest, ulen);

  int i;
  for(i=0;i<20;i++) {
    digest[i] = current_digest[i];
  }

}
`
