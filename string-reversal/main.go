package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseString("alo"))
}

func reverseString(s string) string {
	bytes := []byte(s)
	length := len(bytes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

// alo
