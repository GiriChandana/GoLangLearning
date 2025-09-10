package main

import "fmt"

func maps() {

	codes := map[string]string{"en": "english", "fr": "french"}
	fmt.Println(codes)

	maparr := make(map[string]int)
	fmt.Println(maparr)

	for key, value := range codes {
		fmt.Println(key, "=>", value)
	}
	delete(codes, "en")
	fmt.Println("after deletion of element with Key 'en':")
	fmt.Println(codes)
}
