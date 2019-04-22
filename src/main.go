package main

import (
	"fmt"
	"4_collections/words"
)

func main() {
	text := "Hello World!"
	fmt.Printf("%v", words.CountFrequency(text))
}
