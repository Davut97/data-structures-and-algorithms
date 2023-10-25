package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "alo"
	reversed := isPalindrome(s)
	fmt.Println(reversed)

}

func isPalindrome(s string) bool {
	cleanedInput := strings.ReplaceAll(strings.ToLower(s), " ", "")

	length := len(cleanedInput)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		if cleanedInput[i] != cleanedInput[j] {
			return false
		}
	}
	return true
}
