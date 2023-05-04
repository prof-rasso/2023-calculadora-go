package main

import (
	"fmt"
	"os"
	"strconv"
)
import "bufio"

// Garantir que os valores sejam informados e que sejam númericos v0.1.1
// Garantir que uma operação válida seja informada v0.1.2
// Perguntar se deseja continuar calculando ou parar de calcular v0.2.0
// Registrar os resultados
// Perguntar se deseja imprimir os ultimos resultados ao fechar o programa. v0.3.0
// Transformação em arquivos. v0.4.0
// Transformação para 'Orientação a Objetos' v1.0.0

func main() {
	fmt.Println("Bem-vindos a magnifica calculadora feita em Golang!!!!!!!!!")
	reader := bufio.NewReader(os.Stdin)

	var invalido = true
	var umValor float64
	for invalido {
		fmt.Println("Informe um valor!")
		lido, _ := reader.ReadString('\n')
		lido = lido[:len(lido)-1] // remove o 'enter' no final
		if valor, err := strconv.ParseFloat(lido, 64); err == nil {
			umValor = valor
			invalido = false
		} else {
			fmt.Println("O valor informado é inválido!")
		}
	}

	fmt.Println("Informe uma operação")
	fmt.Println("1 - Somar")
	fmt.Println("2 - Subtrair")
	fmt.Println("3 - Multiplicar")
	fmt.Println("4 - Dividr")
	var operacao int
	fmt.Scan(&operacao)

	invalido = true
	var outroValor float64
	for invalido {
		fmt.Println("Informe outro valor!")
		lido, _ := reader.ReadString('\n')
		lido = lido[:len(lido)-1] // remove o 'enter' no final
		if valor, err := strconv.ParseFloat(lido, 64); err == nil {
			outroValor = valor
			invalido = false
		} else {
			fmt.Println("O valor informado é inválido!")
		}
	}

	var sinal string
	var resultado float64
	switch operacao {
	case 1:
		sinal = "+"
		resultado = umValor + outroValor
	case 2:
		sinal = "-"
		resultado = umValor - outroValor
	case 3:
		sinal = "*"
		resultado = umValor * outroValor
	case 4:
		sinal = "/"
		resultado = umValor / outroValor
	}

	fmt.Printf("O resultado de %.2f %s %.2f é igual a %.2f", umValor, sinal, outroValor, resultado)

}
