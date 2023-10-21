package utils

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
)

const Sign = "c62b84f6c92f0aa33c9834ea6132dde8"

// AESEncodeNormal 通用加密
func AESEncodeNormal(m map[string]string, baiduTokenSecretKeySign string) string {
	sourceOfByte, _ := json.Marshal(m)
	encryptCode := AESEncrypt(sourceOfByte, []byte(baiduTokenSecretKeySign))
	encryptCode1 := base64.StdEncoding.EncodeToString(encryptCode)
	encryptCode2 := url.QueryEscape(encryptCode1)
	return encryptCode2
}

// AESEncrypt ecb加密
func AESEncrypt(src, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))

	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}

	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

const keyLen = 16

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, keyLen)
	copy(genKey, key)

	for i := 16; i < len(key); {
		for j := 0; j < keyLen && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}

	return genKey
}

func AESDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	cipher, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return nil
	}

	blockSize := cipher.BlockSize()
	if len(encrypted)%blockSize != 0 {
		return nil
	}
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+blockSize, be+blockSize {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	if trim < 0 || trim >= len(decrypted) {
		return nil
	}
	return decrypted[:trim]
}

// AESDecodeNormal 通用解密
func AESDecodeNormal(token string, baiduTokenSecretKeySign string) (m map[string]string, err error) {
	encryptCode1, _ := url.QueryUnescape(token)
	encryptCode2, _ := base64.StdEncoding.DecodeString(encryptCode1)

	decryptCode := AESDecrypt(encryptCode2, []byte(baiduTokenSecretKeySign))

	err = json.Unmarshal(decryptCode, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
