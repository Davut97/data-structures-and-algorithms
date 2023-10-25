package main

import "fmt"

func main() {
	fmt.Println(ReverseInteger(-123))
}

func ReverseInteger(original int) int {

	num := original
	reversed := 0

	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num = num / 10
	}

	return reversed
}
