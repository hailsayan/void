package main

import "fmt"

// 35

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	count := 1
	for i := 2; i <= n; i++ {
		count *= i
	}
	return count
}

func main() {
	var n int
	fmt.Println("enter a number: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("factorial is not defined for negative numbers")
	} else {
		fmt.Printf("the factorial of %d is %d\n", n, factorial(n))
	}
}
