package main

import "fmt"

func main() {
	fmt.Println(string(MaxChar("oweso")))
	fmt.Printf("The most used character in the string is: %c\n", MaxChar("oweso"))

}

func MaxChar(s string) rune {
	charCountMap := make(map[rune]int)

	mostCharFreq := 0
	var mostUsedChar rune
	for _, char := range s {
		charCountMap[char]++
	}
	for char, freq := range charCountMap {
		if freq > mostCharFreq {
			mostCharFreq = freq
			mostUsedChar = char
		}
	}

	return mostUsedChar

}
