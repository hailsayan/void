package main

import "fmt"

//26

func main() {
	triangleChecker()
}

func triangleChecker() {
	var a, b, c int
	fmt.Print("Enter three values (a, b, c): ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)

	if a+b >= c && a+c >= b && b+c >= a {
		fmt.Println("Triangle exists")
	} else {
		fmt.Println("Triangle does not exist")
	}
}
