package main

import (
	"fmt"

	"github.com/GiriChandana/GoLangLearning/cryptit/decrypt"
	"github.com/GiriChandana/GoLangLearning/cryptit/encrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("KodeCloud"))
	encryptedStr := encrypt.Nimbus("KodeCloud")
	fmt.Println("Decrypted string: ", decrypt.Nimbus(encryptedStr))
}
