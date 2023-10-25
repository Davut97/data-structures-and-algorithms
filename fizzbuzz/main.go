package main

import "fmt"

func main() {
	Fizzbuzz(10)
}
func Fizzbuzz(num int) {
	for i := 0; i <= num; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

}
