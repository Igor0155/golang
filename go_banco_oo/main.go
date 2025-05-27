package main

import (
	"fmt"

	c "github.com/Igor0155/golang/go_banco_oo/contas"
)

func main() {

	contaIgu := c.ContaCorrente{Titular: "Igor", Saldo: 500}

	contaJoel := c.ContaCorrente{Titular: "Joel", Saldo: 300}

	fmt.Println(contaIgu)
	fmt.Println(contaJoel)

	mensagem, saldo1, saldo2 := contaIgu.Transferir(-600, &contaJoel)

	fmt.Println(mensagem, saldo1, saldo2)

		fmt.Println(contaIgu)
	fmt.Println(contaJoel)

	// contaIgu := ContaCorrente{}
	// contaIgu.titular = "Igor"
	// contaIgu.saldo = 500

	// // mensagem :=  contaIgu.Sacar(-800)
	// mensagem := contaIgu.Depositar(100)

	// fmt.Println(mensagem)	
}
