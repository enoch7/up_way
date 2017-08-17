package main

import "fmt"
import "time"

func worker(done chan<- bool,num int) {
	fmt.Println("worker",num,"is working....")
	time.Sleep(time.Second*2)
	fmt.Println("worker",num," done!")
	done<- true
}

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

	c3 := make(chan bool)

	go worker(c3,1)
	go worker(c3,2)

	<-c3
	<-c3


}