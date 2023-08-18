package main

import "fmt"

func main() {
	sum := 0
	switch {
	case sum > 0:
		fmt.Println("Positivo ")
	case sum < 0:
		fmt.Println("Negativo ")
	default:
		fmt.Println("0 ")
	}
}
