package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex
	val int
}

func main() {
	// Waitgroup
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()

	// Mutex
	runtime.GOMAXPROCS(2)

	var wg1 sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg1.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}

			wg1.Done()
		}()
	}

	wg1.Wait()
	fmt.Println(meter.Value())
}

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) Value() (x int) {
	return c.val
}
