package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	index := strings.IndexByte(word, ch)
	fmt.Println("this is index", index)

	if index == -1 {
		return word
	}

	runes := []rune(word)
	fmt.Println("this is runes", runes)
	for i, j := 0, index; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	word := "abcdef"
	ch := 'j'
	fmt.Println(reversePrefix(word, byte(ch)))
}
