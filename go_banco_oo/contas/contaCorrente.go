package contas

import (
	"strconv"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	if valorDoSaque < 0 {
		return "Valor do saque inválido!"
	}

	if valorDoSaque <= c.Saldo {
		c.Saldo -= valorDoSaque
		return "Valor sacado com sucesso. \n" + "Valor restante: " + strconv.FormatFloat(c.Saldo, 'f', 2, 64)
	}

	return "Saldo insuficiente " + strconv.FormatFloat(c.Saldo, 'f', 2, 64)
}

func (c *ContaCorrente) Depositar(deposito float64) string {
	if deposito < 0 {
		return "Valor do depoisito inválido!"
	}

	c.Saldo += deposito
	return "Valor depositado com sucesso. \n" + "Valor em conta: " + strconv.FormatFloat(c.Saldo, 'f', 2, 64)
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (mensagem string, valorRestanteConta float64, ValorTotalContaDep float64) {

	if valorDaTransferencia <= 0 {
		mensagem = "Valor inválido para transferência"
		return
	}

	if c.Saldo < valorDaTransferencia {
		mensagem = "Saldo insuficiente"
		return
	}

	c.Saldo -= valorDaTransferencia
	contaDestino.Depositar(valorDaTransferencia)

	mensagem = "Transferência realizada com sucesso"
	valorRestanteConta = c.Saldo
	ValorTotalContaDep = contaDestino.Saldo
	return

}
