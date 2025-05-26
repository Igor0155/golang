package main

import (
	"fmt"
	// "reflect"
)


func main() {
	// var name string = "Igor"
	// var numFloat float32 = 22.2
	// var idade int = 22
	// fmt.Printf("\no tipo da váriavel Nome é %s", reflect.TypeOf(name))
	// fmt.Printf("\no tipo da váriavel numFloat é %s", reflect.TypeOf(numFloat))
	// fmt.Printf("\no tipo da váriavel idade é %s", reflect.TypeOf(idade))

	name := "Igor G."
	version := 1.0
	// idade := 22
	
	fmt.Printf("Olá %s, \nEste programa está na versão %0.1f\n", name, version)
	fmt.Println("1 - Iniciar o Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O valor da variável comando é:", comando)
}
