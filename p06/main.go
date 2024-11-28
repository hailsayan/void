package main

import "fmt"

func isPerfectNumber(n int) bool {
	// initialize the sum of divisors
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if isPerfectNumber(n) {
		fmt.Printf("%d is a perfect number.\n", n)
	} else {
		fmt.Printf("%d is not a perfect number.\n", n)
	}
}
