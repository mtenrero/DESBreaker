package main

import "crypto/des"

import "github.com/andreburgaud/crypt2go/ecb"

// DesDecryption decrypts the given cipherText using the provided key
func DesDecryption(key, cipherText []byte) []byte {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil
	}

	blockMode := ecb.NewECBDecrypter(block)
	originalData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(originalData, cipherText)
	return originalData
}
