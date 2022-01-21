package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta string
	saldo                      float64
}

func (c *ContaCorrente) Transferir(valorDaTranferencia float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valorDaTranferencia > 0 && valorDaTranferencia <= c.saldo

	if podeTransferir {
		c.saldo -= valorDaTranferencia
		contaDestino.saldo += valorDaTranferencia
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente."
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito >= 10
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Valor depositado com sucesso", c.saldo
	} else {
		return "Não foi possível depositar", c.saldo
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
