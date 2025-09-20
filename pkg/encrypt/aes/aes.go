package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type Aes struct {
	key []byte
}

func NewAes(key []byte) *Aes {
	return &Aes{
		key: key,
	}
}

// padding 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// unPadding 去掉填充数据
func unPadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// Encrypt 加密
func (encrypt Aes) Encrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(encrypt.key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, encrypt.key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// Decrypt 解密
func (encrypt Aes) Decrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(encrypt.key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, encrypt.key)
	blockMode.CryptBlocks(src, src)
	src = unPadding(src)
	return src, nil
}
