// viginere_cypher project main.go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	keywordLength := utf8.RuneCountInString(keyword)
	repeatedString := ""
	for idx, _ := range cipherText {
		keywordIdx := idx % keywordLength
		keywordChar := keyword[keywordIdx]
		offset := keywordChar - 'A'
		cipherChar := cipherText[idx]
		decodedChar := (cipherChar-'A'-offset)%byte(26) + 'A'
		fmt.Printf("%c", decodedChar)
	}

	fmt.Printf(repeatedString)
}
