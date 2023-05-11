package main

import (
	"calculadora-go/calculo"
	"calculadora-go/cli/terminal"
	"calculadora-go/resultado"
	"fmt"
)

// Transformação para 'Orientação a Objetos' v1.0.0

func main() {
	fmt.Println("Bem-vindos a magnifica calculadora feita em Golang!!!!!!!!!")
	continuar := true
	for continuar {
		umValor := terminal.LerFloat("Informe um valor!")
		operacao := terminal.LerOperacao()
		outroValor := terminal.LerFloat("Informe outro valor!")
		resultadoFloat, sinal := calculo.Calcular(operacao, umValor, outroValor)
		resultadoStr := resultado.Registrar(umValor, sinal, outroValor, resultadoFloat)
		fmt.Println(resultadoStr)
		continuar = terminal.LerSimNao("Deseja continuar calculando? 1-Sim 2-Não")
	}
	imprimir := terminal.LerSimNao("Deseja imprimir os resultados anteriores? 1-Sim 2-Não")
	if imprimir {
		resultado.Imprimir()
	}
	fmt.Println("Bye bye!")
}
