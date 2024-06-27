package main

import (
	"fmt"
	"time"
)

func odd() {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 10)
		if i&1 == 1 {
			fmt.Println(i)
		}
	}
}

func even() {
	for i := 2; i <= 100; i++ {
		if i&1 == 0 {
			time.Sleep(time.Millisecond * 12)
			fmt.Println(i)
		}
	}
}

func main() {
	//ch := make (chan int)
	fmt.Println("Start")
	go odd()
	go even()
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 20)
	}
}
