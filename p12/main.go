package main

import "fmt"

// 40

func main() {
	var n int
	fmt.Println("enter a number: ")
	fmt.Scanf("%d", &n)

	digitCount := 0
	sumOfDigits := 0

	for n > 0 {
		digit := n % 10
		sumOfDigits += digit
		digitCount++
		n /= 10
	}

	fmt.Printf("Number of digits: %d\n", digitCount)
	fmt.Printf("Sum of digits: %d\n", sumOfDigits)
}
