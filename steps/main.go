package main

import "fmt"

func main() {
	steps(10)
}

func steps(n int) {
	printSprints(n, 0, 0)
}
func printSprints(n, row, col int) {
	if row == n {
		return
	}
	if col <= row {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}

	if col == n-1 {
		fmt.Println()
		printSprints(n, row+1, 0)
		return
	}
	printSprints(n, row, col+1)

}
