package main

import (
	"fmt"
)

func fizzbuzzWithConstraint(N int) {
	for i := 1; i <= N; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// Input dari pengguna
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)
	fizzbuzzWithConstraint(N)
}
