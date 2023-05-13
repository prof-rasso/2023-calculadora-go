package terminal

import (
	"bufio"
	"calculadora-go/core/operacao"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func LerFloat(label string) float64 {
	for {
		fmt.Println(label)
		if scanner.Scan() {
			text := scanner.Text()
			umValor, err := strconv.ParseFloat(text, 64)
			if err == nil {
				return umValor
			}
		}
		fmt.Println("O valor informado é inválido")
	}
}

func LerOperacao() (operacao.Operacao, error) {
	for {
		fmt.Println("Escolha uma operação.")
		fmt.Println("1 - Somar")
		fmt.Println("2 - Subtrair")
		fmt.Println("3 - Multiplicar")
		fmt.Println("4 - Dividir")
		if scanner.Scan() {
			text := scanner.Text()
			op, err := strconv.Atoi(text)
			if err == nil && op >= 1 && op <= 4 {
				return operacao.NovaOperacao(op)
			}
		}
		fmt.Println("O valor informado é inválido")
	}
}

func LerSimNao(label string) bool {
	for {
		fmt.Println(label)
		if scanner.Scan() {
			text := scanner.Text()
			textInt, err := strconv.Atoi(text)
			if err == nil && (textInt == 1 || textInt == 2) {
				return textInt == 1
			}
		}
		fmt.Println("O valor informado é inválido")
	}
}
