package main

import (
	"fmt"
	"time"
)

func say1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func number2() {
	for i := 1; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

func number1() {
	for i := 100; i < 110; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	go number1()
	go number2()
	//go say1("world")
	//say1("hello")
	time.Sleep(10 * time.Second)
}
