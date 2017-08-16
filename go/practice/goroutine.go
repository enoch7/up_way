package main

import "fmt"

func gogo(str string) {
	fmt.Println(str)
}


func main() {
	fmt.Println("go start")
	go gogo("goging")
	fmt.Println("dont know if it is running")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	fmt.Println("done")
}