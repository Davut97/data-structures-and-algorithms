package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(anagrams("rail safety", "fairy tales"))
}

// func anagrams(stringA, stringB string) bool {
// 	m1 := make(map[byte]int)
// 	m2 := make(map[byte]int)
// 	for i := 0; i < len(stringA); i++ {
// 		m1[stringA[i]] += 1
// 	}
// 	for i := 0; i < len(stringB); i++ {
// 		m2[stringB[i]] += 1
// 	}

//		return reflect.DeepEqual(m1, m2)
//	}
func anagrams(stringA, stringB string) bool {

	return sortString(stringA) == sortString(stringB)

}
func sortString(input string) string {
	strChars := strings.Split(input, "")
	sort.Strings(strChars)

	return strings.Join(strChars, "")
}
 