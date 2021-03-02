package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func forceDeadlock3() {
	fmt.Println("vim-go")
	c := make(chan bool)
	fmt.Println("vim-go2")
	<-c
	fmt.Println("vim-go3")
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
// go run ./cmd/milmove-tasks test-deadlock
func testDeadlock3(cmd *cobra.Command, args []string) error {

	fmt.Println("Hello, playground")
	forceDeadlock3()
	fmt.Println("done")

	return nil
}
