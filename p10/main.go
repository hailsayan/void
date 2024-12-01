package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter a positive even integer: ")
	fmt.Scanf("%d", &n)

	if n%2 != 0 || n <= 0 {
		fmt.Println("your number must be a positive even integer")
		return
	}

	var s float64
	sign := 1.0
	for i := 1; i < n; i++ {
		s += sign * float64(i) / float64(i+1)
		sign *= -1
	}
	fmt.Printf("the sum of the series is : %.6\n", s)
}
