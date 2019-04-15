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
		// fmt.Printf("KeywordChar=%c\n", keywordChar)
		repeatedString = repeatedString + string(keyword[keywordIdx])
		offset := keywordChar - 'A'
		// fmt.Printf("Offset=%v\n", offset)
		cipherChar := cipherText[idx]
		decodedChar := (26 + cipherChar - 'A' - offset) % byte(26) + 'A'
		// fmt.Printf("CipherChar: %c, CipherCharVal: %v\n", cipherChar, cipherChar)
		// fmt.Printf("DecodedChar: %c, DecodedCharVal: %v\n", decodedChar, decodedChar)
		fmt.Printf("%c", decodedChar)
	}

	fmt.Printf("\n%v", repeatedString)
}
