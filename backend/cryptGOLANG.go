package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/

func createHash(key string) string {

	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func checkPassword(pass string) (bool, string) {
	// Si je crypte 2 fois le même mot de passe, j'obtiens 2 chaines différentes
	// L'important est qu'en décryptant ces 2 chaines, je retombe sur le même résultat
	//first := encrypt([]byte(pass), "ThisIsMyPasswordToUseEachTimeToEncryptOrDeccrypt")
	//firstd := decrypt(first, "ThisIsMyPasswordToUseEachTimeToEncryptOrDeccrypt")
	second := []byte{65, 107, 66, 123, 101, 173, 226, 225, 248, 112, 251, 5, 242, 105, 100, 142, 40, 41, 34, 179, 135, 0, 65, 200, 27, 228, 231, 92, 249, 55, 127, 27, 43, 204, 111, 80, 190, 163, 241, 45, 99, 171}
	secondd := decrypt(second, "ThisIsMyPasswordToUseEachTimeToEncryptOrDeccrypt")

	//	if !bytes.Equal(firstd, secondd) {
	if pass != string(secondd) {
		return false, "Mot de passe incorrect"
	}
	return true, "Mot de passe vérifié"
}
