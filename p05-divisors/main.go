package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the desired number:")
	fmt.Scan(&n)

	fmt.Printf("The divisors of %d are:\n", n)

	for i := 1; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
