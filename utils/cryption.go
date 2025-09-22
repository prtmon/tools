package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
)

// 默认字符集（字母+数字）
const defaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// RandomString 生成随机字符串
func RandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be positive number")
	}
	// 处理字符集
	selectedCharset := defaultCharset
	charsetLen := big.NewInt(int64(len(selectedCharset)))
	// 预分配结果缓冲区
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		// 生成加密安全的随机数（0 ~ charsetLen-1）
		n, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = selectedCharset[n.Int64()]
	}
	return string(result), nil
}

// MD5 md5字符串
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}

// AesCbcEncrypt AES-CBC加密
func AesCbcEncrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充明文
	plaintext = PKCS7Padding(plaintext, aes.BlockSize)

	// 生成随机IV
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// AesCbcDecrypt AES-CBC解密
func AesCbcDecrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 去除填充
	return PKCS7UnPadding(ciphertext)
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7UnPadding PKCS7去除填充
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid data")
	}
	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("invalid padding")
	}
	return data[:length-padding], nil
}
