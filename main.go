package main

import (
	"calculadora-go/cli/terminal"
	"calculadora-go/core/resultado"
	"fmt"
)

func main() {
	continuar := true
	for continuar {
		umValor := terminal.LerFloat("Informe um valor!")

		operacao, err := terminal.LerOperacao()
		if err != nil {
			fmt.Println(err)
			break
		}

		outroValor := terminal.LerFloat("Informe um valor!")

		calculo, err := operacao.CriarCalculoEntre(umValor, outroValor)
		if err != nil {
			fmt.Println(err)
			break
		}

		resultadoCalculo := calculo.Calcular()
		resultado.Registrar(resultadoCalculo)
		fmt.Println(resultadoCalculo)

		continuar = terminal.LerSimNao("Deseja continuar calculando? (1-Sim 2-Não)")
	}

	imprimir := terminal.LerSimNao("Deseja imprimir os resultados cálculados? (1-Sim 2-Não)")
	if imprimir {
		resultado.Imprimir()
	}

	fmt.Println("Bye bye!")
}
