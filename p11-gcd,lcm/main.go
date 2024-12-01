package main

import "fmt"

// 39
// BMM -> GCD
// KMM -> LCM

func gcd(a, b int) int {
	// while loop
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n, m int
	fmt.Println("enter two positive integers. n then m:")
	fmt.Scanf("%d %d", &n, &m)
}
