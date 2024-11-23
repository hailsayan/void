package main

import "fmt"

//24

func main() {
	even()
}

func even() {
	for i := 1000; i <= 2000; i += 2 {
		fmt.Println(i)
		if i == 2000 {
			break
			fmt.Println("done")
		}

	}
}
