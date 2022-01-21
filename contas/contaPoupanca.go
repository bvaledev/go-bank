package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta string
	Operacao                   int
	saldo                      float64
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente."
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito >= 10
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Valor depositado com sucesso", c.saldo
	} else {
		return "Não foi possível depositar", c.saldo
	}
}

