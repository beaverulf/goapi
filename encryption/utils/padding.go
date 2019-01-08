package utils

//Pkcs7Padding pads the incoming src block with PKCS7 padding of a specific block size and returns new slice.
func Pkcs7Padding(src []byte, blockSize int) []byte {
	padding := byte(0x04)
	srcLen := len(src)

	paddingLen := 0
	if srcLen < blockSize {
		paddingLen = blockSize - srcLen
	} else if (srcLen % blockSize) != 0 {
		paddingLen = blockSize - (srcLen % blockSize)
	}

	var paddedDst = src[:]
	for i := 0; i < paddingLen; i++ {
		paddedDst = append(paddedDst, padding)
	}

	//fmt.Printf("String: %s \t Values: %o\n", string(paddedDst), paddedDst)
	return paddedDst
}
