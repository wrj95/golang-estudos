package main

import (
	"fmt"
)

func main() {
	var var1 string = "variavel 1"
	fmt.Println(var1)
	var2 := "variavel 2"
	fmt.Println(var2)

	var (
		var3 string = "lalalal"
		var4 string
	)

	var4 = "variavel 4"
	fmt.Println(var3, var4)

	var5, var6 := "variavel 5", "variavel 6"

	fmt.Println(var5, var6)

	//constante
	const const1 string = "Contante 1"
	fmt.Println(const1)

	var5, var6 = var6, var5

	fmt.Println(var5, var6)
}