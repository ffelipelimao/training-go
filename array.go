package main

import (
	"fmt"
)

func main() {
	array := []int{-4, 3, -9, 0, 4, 1}
	for i := range array[0:] {
		fmt.Println(array[i] + 1)
	}

}
