package main

import (
	"fmt"
)

func main() {

	fmt.Println(chunk([]int{1, 2, 3, 4, 5, 6}, 1))

}
func chunk(slice []int, size int) [][]int {
	newSlice := make([][]int, 0)
	for i := 0; i < len(slice); i += size {
		if (i+size)-len(slice) > 0 {
			tempSlice := slice[i:]
			newSlice = append(newSlice, tempSlice)
		} else {
			tempSlice := slice[i : i+size]
			newSlice = append(newSlice, tempSlice)
		}

	}

	return newSlice
}
