package main

import (
	"time"
	"fmt"
)

func main() {

	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 2")
	}

}
