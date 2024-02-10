package main

import (
	"fmt"
	"strings"
)

func truncatesentence(s string, k int) string {
	//spilt the strings into words
	words := strings.Fields(s)
	fmt.Println("*****words", words)
	//select the first k words from the list of words
	truncatedWords := words[:k]
	//join the words
	fmt.Println("*****truncatewords", truncatedWords)
	truncatesentence := strings.Join(truncatedWords, " ")
	return truncatesentence
}

func main() {
	s := "Hello how are you priyanka"
	k := 4
	fmt.Println(truncatesentence(s, k))
}
