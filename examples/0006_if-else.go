package main

import "fmt"

func main() {

	var n = 8

	if n % 2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}

	if n % 4 == 0 {
		fmt.Println(n, "is divisible by 4")
	}

	if num := 9; n < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digit")
	}

}
