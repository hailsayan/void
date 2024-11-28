package main

import "fmt"

func isEvenOrOdd(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	A := isEvenOrOdd(12)
	fmt.Println(A)
}
