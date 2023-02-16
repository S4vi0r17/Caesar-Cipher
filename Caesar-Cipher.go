package main

import "fmt"

func main() {
	fmt.Print("Enter the word to encrypt: ")
	var word string
	fmt.Scanln(&word)

	fmt.Print("Enter the displacement: ")
	var displacement int
	fmt.Scanln(&displacement)

	wordBytes := []byte(word)

	for i := 0; i < len(wordBytes); i++ {
		ascii := int(wordBytes[i])

		cipheredAscii := (ascii + displacement - 65) % 26 + 65

		wordBytes[i] = byte(cipheredAscii)
	}

	fmt.Println("Encrypted word: " + string(wordBytes))
}

