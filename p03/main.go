package main

import "fmt"

//31

func compareValues(a, b int) {
	switch {
	case a < b:
		fmt.Printf("%d is less than %d\n", a, b)

	case a > b:
		fmt.Printf("%d is greater than %d\n", a, b)

	case a == b:
		fmt.Printf("%d is equal to %d\n", a, b)
	}
}

func main() {

	compareValues(45, 83)

}
