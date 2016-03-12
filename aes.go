package libs

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(plantText, key []byte) ([]byte, error) {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, key)

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func Decrypt(ciphertext, key []byte) ([]byte, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
