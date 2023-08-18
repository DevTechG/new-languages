package main

import (
	"fmt"
)

var (
	isBool       bool       = false
	isString     string     = "False"
	isInt        int        = -1000000
	isInt8       int8       = 127   //mínimo -128 e máximo 127
	isUint       uint       = 10000 //apenas positivo
	isByte       byte       = 255   //apenas positivo, máximo 255
	isRune       rune       = -1000000
	isFloat32    float32    = 9999999999
	isFloat64    float64    = 9999999999
	isComplex64  complex64  = complex(5, 2)
	isComplex128 complex128 = complex(5, 2)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", isBool, isBool)
	fmt.Printf("Type: %T Value: %v\n", isString, isString)
	fmt.Printf("Type: %T Value: %v\n", isInt, isInt)
	fmt.Printf("Type: %T Value: %v\n", isInt8, isInt8)
	fmt.Printf("Type: %T Value: %v\n", isByte, isByte)
	fmt.Printf("Type: %T Value: %v\n", isRune, isRune)
	fmt.Printf("Type: %T Value: %v\n", isFloat32, isFloat32)
	fmt.Printf("Type: %T Value: %v\n", isFloat64, isFloat64)
	fmt.Printf("Type: %T Value: %v\n", isComplex64, isComplex64)
	fmt.Printf("Type: %T Value: %v\n", isComplex128, isComplex128)
}
