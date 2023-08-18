package main

import "fmt"

func main() {
	sum := 1
	if sum > 0 {
		fmt.Println("Positivo ")
	}
	sum -= 10
	if sum < 0 {
		fmt.Println("Negativo ")
	}
}
