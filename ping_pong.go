package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
			ch <- "pong"
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
			ch <- "ping"
		}
	}()

	wg.Wait()
	close(ch)
}
