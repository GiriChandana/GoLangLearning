package main

import (
	"fmt"

	"github.com/GiriChandana/cryptit/decrypt"
	"github.com/GiriChandana/cryptit/encrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("KodeCloud"))
	encryptedStr := encrypt.Nimbus("KodeCloud")
	fmt.Println("Decrypted string: ", decrypt.Nimbus(encryptedStr))
}
