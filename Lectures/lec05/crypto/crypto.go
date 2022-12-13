// crypto/crypto.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func cbc() {
	fmt.Println("AES-CBD demo")
	key, _ := hex.DecodeString("000102030405060708090A0B0C0D0E0F")
	fmt.Printf("key=*%v*\n", hex.EncodeToString(key))

	plaintxt := "0123456789ABCDEF0123456789ABCDEF"
	pbuf := []byte(plaintxt)
	iv := make([]byte, 16)
	rand.Read(iv)
	fmt.Printf("iv=%v\n", hex.EncodeToString(iv))
	fmt.Printf("plaintext=%v\n", hex.EncodeToString(pbuf))

	aes, _ := aes.NewCipher(key)

	cbcEnc := cipher.NewCBCEncrypter(aes, iv)
	ciphertxt := make([]byte, len(pbuf))
	cbcEnc.CryptBlocks(ciphertxt, pbuf)
	fmt.Printf("ciphertext=%v\n", hex.EncodeToString(ciphertxt))

	cbcDec := cipher.NewCBCDecrypter(aes, iv)
	pbuf2 := make([]byte, len(ciphertxt))
	cbcDec.CryptBlocks(pbuf2, ciphertxt)
	decrypted := string(pbuf2)
	fmt.Printf("decrypted plaintext=%v\n", decrypted)
}

func ctr() {
	fmt.Println("AES-CTR demo")
	key, _ := hex.DecodeString("000102030405060708090A0B0C0D0E0F")
	fmt.Printf("key=*%v*\n", hex.EncodeToString(key))

	plaintxt := "0123456789ABCDEF0123456789"
	pbuf := []byte(plaintxt)
	iv := make([]byte, 16)
	rand.Read(iv)
	fmt.Printf("iv=%v\n", hex.EncodeToString(iv))
	fmt.Printf("plaintext=%v\n", hex.EncodeToString(pbuf))

	aes, _ := aes.NewCipher(key)

	ctrStream := cipher.NewCTR(aes, iv)
	ciphertxt := make([]byte, len(pbuf))
	ctrStream.XORKeyStream(ciphertxt, pbuf)
	fmt.Printf("ciphertext=%v\n", hex.EncodeToString(ciphertxt))

	ctrStream2 := cipher.NewCTR(aes, iv)
	pbuf2 := make([]byte, len(ciphertxt))
	ctrStream2.XORKeyStream(pbuf2, ciphertxt)
	decrypted := string(pbuf2)
	fmt.Printf("decrypted plaintext=%v\n", decrypted)
}

func main() {
	cbc()
	ctr()
}
