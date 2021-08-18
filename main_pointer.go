package main

import (
	"fmt"
)

//alterando uma variavel implicitamente
func main() {
	v := 2
	incrementFromAPointer(&v) //voce esta passando o endereço de memoria da variavel v
	fmt.Println(v)            //ao printar esse mesmo valor, ele estara alterado por aquela função que o acessou
}

//essa função recebe um endereço de memoria, e acessa seu valor
func incrementFromAPointer(p *int) int {
	*p++      //soma mais um valor nesse endereço
	return *p //retorna esse valor
}
