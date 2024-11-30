package main

// 37

import "fmt"

func calculator(r, m int) int {
	if m <= 6 {
		n := (m-1)*31 + r
		return n
	} else if m < 11 {
		n := 186 + (m-6)*30 + r
		return n
	} else {
		n := 336 + r
		return n
	}
}

func main() {
	var r, m int
	fmt.Println("enter the day (r) of month (m):")
	fmt.Scanf("%d %d", &r, &m)

	fmt.Printf("we are on the %dth of the year", calculator(r, m))

}
