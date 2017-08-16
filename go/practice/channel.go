package main

import "fmt"

func main() {
	c1 := make(chan string)

	go func(msg string) { 
		c1 <- msg
	}("give you a message")

	mes := <- c1

	fmt.Println(mes)

	fmt.Println("---------")

	c2 := make(chan string,2)

	c2 <- "message 1" //will not block here
	c2 <- "message 2"
	close(c2)

	for m := range c2 {
		fmt.Println(m)
	}


}