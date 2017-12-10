package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "string"
	messages <- "ping"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
