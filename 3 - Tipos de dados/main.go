package main

import (
	"fmt"
	"errors"
)

func main(){
	//INT

	//int8, int16, int32 int64
	var num int8 = 127

	fmt.Println(num)

	//if you set a var int, golang will verify if PC is 32 or 64 bits and create one

	//unsigned int - uint
	var num2 uint32 = 15
	fmt.Println(num2)

	//alias
	//INT32 = RUNE
	var num3 rune = 323232
	fmt.Println(num3)

	//uint8 = byte
	var num4 byte = 18
	fmt.Println(num4)

	//FLOAT

	//float32, float64
	var numReal1 float32 = 123.45
	fmt.Println(numReal1)

	var numReal2 float64 = 123456.78
	fmt.Println(numReal2)

	//STRING

	var string1 string = "Texto"
	fmt.Println(string1)

	string2 := "Text2"
	fmt.Println(string2)

	//CHAR
	var char1 rune = 'B'
	fmt.Println(char1)

	char2 := 'w'
	fmt.Println(char2)

	//BOOLEAN

	var bool1 bool = true
	fmt.Println(bool1)
	bool2 := false
	fmt.Println(bool2)

	var bool3 bool
	fmt.Println(bool3)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro: XXXXXXXXXX")
	fmt.Println(erro2)
}