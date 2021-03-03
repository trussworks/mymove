package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func forceDeadlock3() {
	defer fmt.Println("def in foreceDeadlock")
	c := make(chan bool)
	fmt.Println("i'm running")
	//panic("panic - before deadlock")
	<-c
	fmt.Println("i'm done")
}

// func main() {
// 	fmt.Println("Hello, playground")
// 	forceDeadlock()
// 	fmt.Println("done")
// }

// func initTestDeadlockFlags(flag *pflag.FlagSet) {
// 	// Logging Levels
// 	cli.InitLoggingFlags(flag)
//
// 	// Don't sort flags
// 	flag.SortFlags = false
// }
//
// func forceDeadlock() {
// 	fmt.Println("vim-go")
// 	c := make(chan bool)
// 	fmt.Println("vim-go2")
// 	<-c
// 	fmt.Println("vim-go3")
// }
//
// func main() {
// 	forceDeadlock()
// }
//
func random() {
	panic("oops... a panic")
}

// go run ./cmd/milmove-tasks test-deadlock
func testDeadlock3(cmd *cobra.Command, args []string) error {
	defer fmt.Println("first defer")
	fmt.Println("Hello, playground")
	forceDeadlock3()
	random()
	fmt.Println("really all done")

	return nil
}
