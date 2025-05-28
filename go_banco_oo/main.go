package main

import (
	"fmt"

	t "github.com/Igor0155/golang/go_banco_oo/clientes"
	c "github.com/Igor0155/golang/go_banco_oo/contas"
)

func main() {
	clienteIgu := t.Titular{ Nome: "Igor", Cpf: "999.999.999.99", Profissao: "Dev",}
	contaCorrenteIgu := c.ContaCorrente{Titular: clienteIgu, NumeroAgencia: 1513813, NumeroConta: 153113838}
	contaPoupancaIgu := c.ContaCorrente{Titular: clienteIgu, NumeroAgencia: 1513813, NumeroConta: 153113838}
	contaCorrenteIgu.Depositar(500)
	contaPoupancaIgu.Depositar(300)

	PagarBoleto(&contaCorrenteIgu, 60)

	fmt.Println("Saldo Conta Corrente:", contaCorrenteIgu.ObterSaldo())


	

	// contaJoel := c.ContaCorrente{Titular: "Joel", Saldo: 300}

	// fmt.Println(contaIgu)
	// fmt.Println(contaJoel)

	// mensagem, saldo1, saldo2 := contaIgu.Transferir(-600, &contaJoel)

	// fmt.Println(mensagem, saldo1, saldo2)

	// fmt.Println(contaIgu)
	// fmt.Println(contaJoel)

	// contaIgu := ContaCorrente{}
	// contaIgu.titular = "Igor"
	// contaIgu.saldo = 500

	// // mensagem :=  contaIgu.Sacar(-800)
	// mensagem := contaIgu.Depositar(100)

	// fmt.Println(mensagem)	
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {

	conta.Sacar(valorDoBoleto)
	
}

type verificarConta interface {
	Sacar(valorDoSaque float64) string
}
