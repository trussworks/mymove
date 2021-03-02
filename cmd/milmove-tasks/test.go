package main

import "fmt"

func main11() {
	fmt.Println("vim-go")
	c := make(chan bool)
	<-c
}
