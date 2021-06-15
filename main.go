package main

import (
	"fmt"
	"training/app"
)

func main() {
	array := []int{-4, 3, -9, 0, 4, 1}
	resultado := app.Calculate(array)
	for i := 0; i < len(resultado); i++ {
		fmt.Printf("%.6f \n", resultado[i])
	}

}
