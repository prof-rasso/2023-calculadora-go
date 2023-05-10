package main

import (
	"fmt"
	"os"
	"strconv"
)
import "bufio"

// Registrar os resultados
// Perguntar se deseja imprimir os ultimos resultados ao fechar o programa. v0.3.0
// Transformação em arquivos. v0.4.0
// Transformação para 'Orientação a Objetos' v1.0.0

func main() {
	fmt.Println("Bem-vindos a magnifica calculadora feita em Golang!!!!!!!!!")
	scanner := bufio.NewScanner(os.Stdin)

	continuar := true

	for continuar {
		var umValor float64
		var err error
		for {
			fmt.Println("Informe um valor!")
			if scanner.Scan() {
				text := scanner.Text()
				umValor, err = strconv.ParseFloat(text, 64)
				if err == nil {
					break
				}
			}
			fmt.Println("O valor informado é inválido")
		}

		var operacao int
		for {
			fmt.Println("Informe uma operação")
			fmt.Println("1 - Somar")
			fmt.Println("2 - Subtrair")
			fmt.Println("3 - Multiplicar")
			fmt.Println("4 - Dividr")
			if scanner.Scan() {
				text := scanner.Text()
				operacao, err = strconv.Atoi(text)
				if err == nil && operacao >= 1 && operacao <= 4 {
					break
				}
			}
			fmt.Println("O valor informado é inválido")
		}

		var outroValor float64
		for {
			fmt.Println("Informe outro valor!")
			if scanner.Scan() {
				text := scanner.Text()
				outroValor, err = strconv.ParseFloat(text, 64)
				if err == nil {
					break
				}
			}
			fmt.Println("O valor informado é inválido")
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

		fmt.Printf("O resultado de %.2f %s %.2f é igual a %.2f \n", umValor, sinal, outroValor, resultado)

		for {
			fmt.Println("Deseja continuar calculando? 1-Sim 2-Não")
			if scanner.Scan() {

				text := scanner.Text()
				var textInt int
				textInt, err = strconv.Atoi(text)
				if err == nil && (textInt == 1 || textInt == 2) {
					continuar = textInt == 1
					break
				}
			}
			fmt.Println("O valor informado é inválido")
		}

	}
	fmt.Println("Bye bye!")
}
