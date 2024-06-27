package Synchronization

import (
	"fmt"
	"sync"
)

func print_odd(wg *sync.WaitGroup, oddChan, evenChan chan struct{}, maxInt int) {
	defer wg.Done()
	for i := 1; i <= maxInt; i += 2 {
		<-oddChan
		fmt.Println(i)

		evenChan <- struct{}{}
	}
}

func print_even(wg *sync.WaitGroup, oddChan, evenChan chan struct{}, maxInt int) {
	defer wg.Done()
	for i := 2; i <= maxInt; i += 2 {
		<-evenChan
		fmt.Println(i)
		if i < maxInt {
			oddChan <- struct{}{}
		}
	}
	close(oddChan)

}

func main() {

	var wg sync.WaitGroup
	fmt.Println("This code is and example for the synchronization between multiple goroutines to ensure a  number sequence as in output")
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})
	fmt.Println("Please add the last digit untill which the sequence is desired")
	maxInt := 100

	wg.Add(2)
	go print_odd(&wg, oddChan, evenChan, maxInt)
	go print_even(&wg, oddChan, evenChan, maxInt)

	oddChan <- struct{}{}

	wg.Wait()

}
