package main

import "fmt"

func payramid(n int) {
	midpoint := (2*n - 1) / 2

	for row := 0; row < n; row++ {
		level := " "
		for col := 0; col < 2*n-1; col++ {
			if midpoint-row <= col && midpoint+row >= col {
				level += "#"
			} else {
				level += " "
			}

		}
		fmt.Println(level)
	}

}

func main() {
	payramid(5)
}
