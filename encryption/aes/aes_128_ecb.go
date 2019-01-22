package aes

import (
	"crypto/aes"
	"errors"
	"fmt"
	"log"

	"github.com/beaverulf/goapi/encryption/utils"
	"github.com/beaverulf/goapi/global"
)

//EncryptAesEcb128 encrypts input with algorithm AES 128-bit in ECB mode.
func EncryptAesEcb128(src, key []byte) ([]byte, error) {
	if len(key)%8 != 0 {
		return nil, fmt.Errorf("key size must be a multiple of 8, currently: %d", len(key))
	} else if len(key) > 32 {
		return nil, fmt.Errorf("key size must be min: 16, max 32, currently: %d", len(key))
	} else if len(key) < 16 {
		return nil, fmt.Errorf("key size must be min: 16, max 32, currently: %d", len(key))
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	blockSize := cipher.BlockSize()
	paddedSrc := utils.Pkcs7Padding(src, blockSize)

	dst := make([]byte, len(paddedSrc))

	if global.DebugEncryption {
		log.Printf("Padded: %d, src: %d", len(paddedSrc), len(src))
	}

	for i := 0; i < len(paddedSrc); i += blockSize {
		cipher.Encrypt(dst[i:i+blockSize], paddedSrc[i:i+blockSize])
	}

	return dst, nil
}

//DecryptAesEcb128 decrypts input with algorithm AES 128-bit in ECB mode.
func DecryptAesEcb128(src, key []byte) ([]byte, error) {
	if len(key)%8 != 0 {
		return nil, fmt.Errorf("key size must be a multiple of 8, currently: %d", len(key))
	} else if len(key) > 32 {
		return nil, errors.New("key size must be min: 16, max 32")
	} else if len(key) < 16 {
		return nil, errors.New("key size must be min: 16, max 32")
	}

	dst := make([]byte, len(src))

	if len(src)%len(key) != 0 {
		if global.DebugEncryption {
			log.Printf("key len: %d src len: %d", len(key), len(src))
		}

		return nil, errors.New("src needs to be multiple of key length(blocksize)")
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	var blockSize = cipher.BlockSize()

	for i := 0; i < len(src); i += blockSize {
		cipher.Decrypt(dst[i:i+blockSize], src[i:i+blockSize])
	}

	return dst, nil
}
