package main

// 37

import "fmt"

func main() {
	for i := 2; i <= 500; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				fmt.Printf("divisor of %d is %d\n", i, j)
			}
		}
	}
}
