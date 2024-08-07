package main

import (
	"fmt"
	"sync"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		wg.Done()
	}
}

// thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	//thread 2
	go task("A", &waitGroup)
	//thread 3
	go task("B", &waitGroup)
	//thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
