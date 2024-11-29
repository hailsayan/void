package main

//34

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("enter a number: ")
	fmt.Scan(&n)

	if isPrime(n) == true {
		fmt.Printf("%d is a prime number\n", n)
	} else {
		fmt.Printf("%d is not a prime number\n", n)
	}

}
