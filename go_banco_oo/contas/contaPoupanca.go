package contas

import (
	"strconv"

	"github.com/Igor0155/golang/go_banco_oo/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao 	  int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	if valorDoSaque < 0 {
		return "Valor do saque inválido!"
	}

	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Valor sacado com sucesso. \n" + "Valor restante: " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
	}

	return "Saldo insuficiente " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
}

func (c *ContaPoupanca) Depositar(deposito float64) string {
	if deposito < 0 {
		return "Valor do depoisito inválido!"
	}

	c.saldo += deposito
	return "Valor depositado com sucesso. \n" + "Valor em conta: " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
}

 func (c *ContaPoupanca) ObterSaldo() float64{
	return c.saldo
 }
