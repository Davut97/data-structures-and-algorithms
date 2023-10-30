package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	fmt.Println(capitalize("a short sentence"))
}
func capitalize(input string) string {

	words := strings.Fields(input) // Split the input string into words
	capitalizedWords := make([]string, len(words))

	for i, v := range words {
		capitalizedWords[i] = capitalizeWord(v)

	}

	return strings.Join(capitalizedWords, " ")
}
func capitalizeWord(word string) string {
	if len(word) == 0 {
		return ""
	}
	// Capitalize the first letter of the word
	return string(unicode.ToUpper(rune(word[0]))) + word[1:]
}
