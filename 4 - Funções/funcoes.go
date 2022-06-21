package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8{
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função F")
	}

	f()

	var x = func(txt string) {
		fmt.Println(txt)
	}

	x("teste")

	var y = func(txt string) string {	
		return txt
	}

	results := y("laranja")

	fmt.Println(results)

	//resultsMath, _ := calculosMatematicos(20, 15)
	resultsMath, teste := calculosMatematicos(20, 15)
	fmt.Println(resultsMath, teste)
}