//  Author: Gabriel Faria
//  https://github.com/DevTechG

// Exemplo simples dos comandos de entrada e saída em Golang.
// var -> declaração de variável
// fmt -> biblioteca de funções de entrada e saída
// Scan, Scanln e Scanf -> comando de entrada
// Print, Println e Printf -> comandos de saída

// Simple example of input and output commands in Golang.
// var -> variable declaration
// fmt -> library of input and output functions
// Scan, Scanln and Scanf -> input commands
// Print, Println and Printf -> output commands

package main

import "fmt"

func main() {
	var destination, departure string                                   //local de destino, local de saída
	var startingKm, endingKm, distanceTraveled int                      //quilômetro de partida, quilômetro de chegada
	var fuelPrice, averageConsumption, litersFuel, amountFueled float32 // preço combustível, consumo médio
	fmt.Print("Informe a cidade do qual está partindo:\n")
	fmt.Scan(&departure)
	fmt.Println("Agora informe o quilômetro inicial:")
	fmt.Scanf("%d", &startingKm)
	fmt.Print("Informe cidade de destino: ")
	fmt.Scanln(&destination)
	fmt.Print("Informe o quilômetro final:\n")
	fmt.Scan(&endingKm)
	fmt.Println("Por último, o preço médio do combustível e o consumo médio do seu veículo:")
    fmt.Scanln(&fuelPrice, &averageConsumption)
	distanceTraveled = endingKm - startingKm
	litersFuel = float32(distanceTraveled) / averageConsumption
	amountFueled = fuelPrice * litersFuel
	fmt.Printf("A cidade de partida é %s no quilômetro %d\nA cidade de destino é %s no quilômetro %d\nConsiderando que o veículo utilizado possui um consumo médio de %.2f e que foram percorridos %d km foram gastos %f litros de combustível no trajeto\nLevando em conta o preço do combustível de %.2f\nA viagem custou %.2f somente de combustível\n", departure, startingKm, destination, endingKm, averageConsumption, distanceTraveled, litersFuel, fuelPrice, amountFueled)
}
