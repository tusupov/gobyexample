package main

import "fmt"

func main() {

	message, sum := make(chan string), make(chan int);

	go func() {
		all := 0
		for i := 1; i < 200; i++ {
			fmt.Println(i)
			all += i
		}
		sum <- all
	}()
	go func() { message <- "ping" }()

	all := <- sum
	fmt.Println(all)

	msg := <- message
	fmt.Println(msg)

}
