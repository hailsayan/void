package main

// 42

import "fmt"

func main() {
	var n int
	fmt.Println("enter a number: ")
	fmt.Scanf("%d", &n)

	maxDigit := 0

	for n > 0 {
		digit := n % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		n /= 10
	}
	fmt.Printf("the largest digit is: %d\n", maxDigit)
}
