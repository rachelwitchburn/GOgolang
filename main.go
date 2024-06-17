package main

import (
	"fmt"
)

func produtoDeNumeros(numero1 int, numero2 int) int {
	var produto int = numero1 * numero2
	return produto
}

func main() {

	//pedir 1° variavel
	//pedir 2° variavel
	//guardar a multip delas em uma outra variavel
	//printar correspondentemente no exemplo

	var numero1 int
	var numero2 int
	fmt.Scan(&numero1)
	fmt.Scan(&numero2)

	fmt.Println("PROD =", (produtoDeNumeros(numero1, numero2)))

}
