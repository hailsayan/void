package main

// 36

import "fmt"

func factorialAndItsSum(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	count := 1
	sum := 0
	for i := 2; i <= n; i++ {
		count *= i
		sum += count
	}
	return sum
}

func main() {
	var n int
	fmt.Println("enter a number: ")
	fmt.Scan(&n)

	if factorialAndItsSum(n) < 0 {
		fmt.Println("factorial is not defined for negative numbers")
	} else {
		fmt.Printf("the sum of factorial is: %d\n", factorialAndItsSum(n))
	}

}
