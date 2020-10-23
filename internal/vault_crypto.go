package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func decrypt(aeskey []byte, ciphervault []byte) ([]byte, error) {

	block, err := aes.NewCipher(aeskey[:32])
	if err != nil {
		panic(err)
	}

	if len(ciphervault) < aes.BlockSize {
		return []byte{}, nil
	}

	iv := ciphervault[:aes.BlockSize]

	ciphervault = ciphervault[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphervault, ciphervault)

	return ciphervault, err
}

func encrypt(aeskey []byte, vault []byte) ([]byte, error) {

	block, err := aes.NewCipher(aeskey[:32])
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(vault))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], vault)

	return ciphertext, err

}
