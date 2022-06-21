package main

import (
	"fmt"
)

func main() {
	//Operadores Aritméticos

	/*
	+
	-
	/
	*
	%
	*/

	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 2 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//Atribuição
	var var1 string = "teste"
	var2 := "teste2"
	fmt.Println(var1, var2)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Lógicos
	fmt.Println("---------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//Unários
	num := 10
	num++
	fmt.Println(num)

	num2 := 10
	num2 += 15
	fmt.Println(num2)

	num3 := 15
	num3--
	fmt.Println(num3)

	num3 -= 10
	fmt.Println(num3)

	num *= 3
	fmt.Println(num)
	num /= 3
	fmt.Println(num)
	num %= 1
	fmt.Println(num)

	//Ternário
	//texto := num3 > 5 ? "Maior que 5" : "Menor que 5"
	//Nativamente não existe operador ternário em Golang, portanto para fazer um, pode ser necessário usar o if/else
	var texto string
	if num3 > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}