package main

import (
	"fmt"
	"training/app"
)

func main() {
	array := []int{1, 2, 3, 4, 5}
	resultado := app.Calculate(array)
	fmt.Println(resultado)
}
