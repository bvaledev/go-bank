package main

import (
	"banco/clientes"
	"banco/contas"
	"banco/pagamentos"
	"fmt"
)

func main() {
	cliente := clientes.Titular{
		Nome:      "Brendo",
		Cpf:       "02898287245",
		Profissao: "Programador",
	}

	contaBrendo := contas.ContaCorrente{
		Titular:       cliente,
		NumeroConta:   "1233-3",
		NumeroAgencia: "0001",
	}
	fmt.Println(contaBrendo.ObterSaldo())
	pagamentos.PagarBoleto(&contaBrendo, 120)
	fmt.Println(contaBrendo.ObterSaldo())

	contaPoupanca := contas.ContaPoupanca{
		Titular:       cliente,
		NumeroConta:   "1233-3",
		NumeroAgencia: "0001",
	}
	fmt.Println(contaPoupanca.ObterSaldo())
	pagamentos.PagarBoleto(&contaPoupanca, 120)
	fmt.Println(contaPoupanca.ObterSaldo())
}
