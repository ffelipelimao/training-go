package main

import (
	"fmt"
	"strings"
)

func main() {
	printStaircase(3)
}

func printStaircase(n int32) {
	//para converter o tipo, só passar o nome como função
	for i := int32(1); i <= n; i++ {
		fmt.Printf("%s%s\n",
			strings.Repeat(" ", int(n-i)),
			strings.Repeat("#", int(i)))
	}
}
