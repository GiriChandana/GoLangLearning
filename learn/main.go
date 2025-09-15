package main

import (
	"fmt"

	"github.com/GiriChandana/cryptit/decrypt"
	"github.com/GiriChandana/cryptit/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	fmt.Println(encrypt.Nimbus("KodeCloud"))
	encryptedStr := encrypt.Nimbus("KodeCloud")
	fmt.Println("Decrypted string: ", decrypt.Nimbus(encryptedStr))
}
