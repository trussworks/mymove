package main

import (
	"fmt"
)

func forceDeadlock2() {
	fmt.Println("vim-go")
	c := make(chan bool)
	fmt.Println("vim-go2")
	<-c
	fmt.Println("vim-go3")
}

func main2() {
	fmt.Println("Hello, playground")
	forceDeadlock()
	fmt.Println("done")
}
