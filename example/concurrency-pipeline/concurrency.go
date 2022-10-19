package main

import (
	"fmt"
	"sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Insert Word 1: Inside my goroutine\n")
	fmt.Println("Insert Word 2: Inside my waitgroup\n")
	waitgroup.Done()
}

func main() {
	fmt.Println("First : Hello Guys\n")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	waitgroup.Wait()

	fmt.Println("Last : Finished Execution")
}
