package terminal

import (
	"bufio"
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

func LerOperacao() int {
	for {
		fmt.Println("Informe uma operação")
		fmt.Println("1 - Somar")
		fmt.Println("2 - Subtrair")
		fmt.Println("3 - Multiplicar")
		fmt.Println("4 - Dividr")
		if scanner.Scan() {
			text := scanner.Text()
			operacao, err := strconv.Atoi(text)
			if err == nil && operacao >= 1 && operacao <= 4 {
				return operacao
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
