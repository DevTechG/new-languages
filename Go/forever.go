package main

import "fmt"

func main() {
	sum := 1
	//for sum < 1000 {
	for {
		sum += 1
		fmt.Printf("%v, ", sum)
	}
}
