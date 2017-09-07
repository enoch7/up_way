package main

import (
	"fmt"
	"os"
	"net"
)

func main() {
	name := os.Args[1]
	addr := net.ParseIP(name)
	fmt.Println(addr.String())
	os.Exit(0)
	
}