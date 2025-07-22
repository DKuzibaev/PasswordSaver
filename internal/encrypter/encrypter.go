package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY environment variable is not set")
	}
	return &Encrypter{Key: key}
}

// Метод для шифрования строки
func (e *Encrypter) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))

	if err != nil {
		panic("Failed to create cipher: " + err.Error())
	}

	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic("Failed to create GCM: " + err.Error())
	}

	nonce := make([]byte, aesGSM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic("Failed to read nonce: " + err.Error())
	}

	return aesGSM.Seal(nonce, nonce, plainStr, nil)
}

// Метод для дешифрования строки
// Возвращает []byte, так как строка может быть невалидной UTF-8
func (e *Encrypter) Decrypt(encryptedStr []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedStr) < nonceSize {
		return nil, fmt.Errorf("encrypted data too short")
	}

	nonce := encryptedStr[:nonceSize]
	ciphertext := encryptedStr[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func (e *Encrypter) NonceSize() int {
	block, _ := aes.NewCipher([]byte(e.Key))
	aesGCM, _ := cipher.NewGCM(block)
	return aesGCM.NonceSize()
}
