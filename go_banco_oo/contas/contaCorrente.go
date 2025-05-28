package contas

import (
	"strconv"

	"github.com/Igor0155/golang/go_banco_oo/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	if valorDoSaque < 0 {
		return "Valor do saque inválido!"
	}

	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Valor sacado com sucesso. \n" + "Valor restante: " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
	}

	return "Saldo insuficiente " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
}

func (c *ContaCorrente) Depositar(deposito float64) string {
	if deposito < 0 {
		return "Valor do depoisito inválido!"
	}

	c.saldo += deposito
	return "Valor depositado com sucesso. \n" + "Valor em conta: " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (mensagem string, valorRestanteConta float64, ValorTotalContaDep float64) {

	if valorDaTransferencia <= 0 {
		mensagem = "Valor inválido para transferência"
		return
	}

	if c.saldo < valorDaTransferencia {
		mensagem = "Saldo insuficiente"
		return
	}

	c.saldo -= valorDaTransferencia
	contaDestino.Depositar(valorDaTransferencia)

	mensagem = "Transferência realizada com sucesso"
	valorRestanteConta = c.saldo
	ValorTotalContaDep = contaDestino.saldo
	return
}

 func (c *ContaCorrente) ObterSaldo() float64{
	return c.saldo
 }

