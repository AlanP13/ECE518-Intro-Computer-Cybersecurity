// crypto/crypto.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func findPassword() {
	nonce := make([]byte, 12)
	data := "jwang34@iit.edu"

	ciphermac, _ := hex.DecodeString(
		"7d4eb640daf43844bd605bb1c2a66046fd33cba7d2ba35828d25056c953834d93b9a04c54fa86147f62a")

	fmt.Printf("finding password for nonce=%x, data=%s, ciphermac=%x...\n",
		nonce, data, ciphermac)

	for i := 0; i < 10000; i++ {
		password := fmt.Sprintf("%04d", i)
		/*
			Modify code below to check that if sha256(password) is a correct
			key to decrypt the ciphertext with MAC in ciphermac, which is
			obtained via AES-GCM with a 0 nouce and the data being jwang34@iit.edu
			Once found, show the correct password as well as the plaintext.
		*/
		sha := sha256.New()         //Creating a new SHA256
		sha.Write([]byte(password)) //SHA of the new password
		hash := sha.Sum(nil)        //SHA hash
		//fmt.Printf("SHA256(%s)=%x\n",
		//	password, hash) //Printing each 4-digit password and hash
		block, _ := aes.NewCipher(hash)                                              //From AES function in validate.go
		aesgcm, _ := cipher.NewGCM(block)                                            //From AES function in validate.go
		plaintext, error := aesgcm.Open(nil, nonce, []byte(ciphermac), []byte(data)) //From AES function in validate.go
		if error == nil {
			fmt.Printf("Correct password=%s\n", password)
			fmt.Printf("%s\n", string(plaintext))
			break
		}
	}
}

func SHA256time() {
	message := make([]byte, 16) //Creating a 16-bit message
	var t time.Duration = 0     //Variable t for the duration of computation

	for i := 0; i < 100000; i++ {
		start := time.Now()        //Start time
		sha := sha256.New()        //Creating a new SHA256
		sha.Write([]byte(message)) //Taking the SHA of the message
		t += time.Since(start)
	}
	fmt.Printf("SHA256 hash computation Time=%s\n", (t / 100000)) // Formatted string as "2.5ns"
}
func AESGSM256time() {
	password := fmt.Sprintf("%04d", 1234)
	plaintext := make([]byte, 1024*1024) //Creating 1M-byte message
	rand.Read(plaintext)                 //From AES function in validate.go
	sha := sha256.New()                  //Creating a new SHA256
	sha.Write([]byte(password))          //Taking the SHA of the password
	hash := sha.Sum(nil)                 //Creating a hash for the SHA
	c, _ := aes.NewCipher(hash)          // create a cipher
	out := make([]byte, 1024*1024)       // allocate space for ciphered data
	var e time.Duration = 0              //Variable e for the duration of encryption
	var d time.Duration = 0              //Variable t for the duration of decryption
	for i := 0; i < 10000; i++ {
		en_start := time.Now() //Start time
		c.Encrypt(out, []byte(plaintext))
		e += time.Since(en_start) //Add encryption time for the process
		pt := make([]byte, len(out))
		de_start := time.Now()
		c.Decrypt(pt, out)
		d += time.Since(de_start) //Add decryption time for the process
	}
	fmt.Printf("AESGCM256 encryption time=%s\n", (e / 100000)) // Formatted string as "2.5ns"
	fmt.Printf("AESGCM256 decrpytion time=%s\n", (d / 100000)) // Formatted string as "2.5ns"
}

func main() {
	validateSHA256()
	validateAESGCM()
	SHA256time()
	AESGSM256time()
	findPassword()
}
