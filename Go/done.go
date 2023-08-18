package main

import (
	"fmt"
	"time"
)

func number2() {
	for i := 1; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

func number1(done chan string) {
	for i := 100; i < 110; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
	done <- "Terminei"
}

func main() {
	done := make(chan string)
	go number1(done)
	go number2()
	fmt.Println(<-done)
}
