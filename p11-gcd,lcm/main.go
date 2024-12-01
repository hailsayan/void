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

	GCD := gcd(n, m)
	LCM := n * m / gcd(n, m)

	fmt.Printf("GCD: %d\n", GCD)
	fmt.Printf("LCM: %d\n", LCM)
	fmt.Println("tadaaaaa")

}
