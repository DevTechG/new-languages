package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Número aleatório:", rand.Intn(100), ".")
	fmt.Printf("Número aleatório: %d.\n", rand.Intn(100))
	fmt.Println("Número aleatório:", rand.Intn(100), ".")
}
