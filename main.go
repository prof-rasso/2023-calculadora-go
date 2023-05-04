package main

import "fmt"

// Ler primeiro valor
// Ler operação
// Ler segundo valor
// Realizar o cálculo
// Imprimir o resultado do cálculo v0.1.0
// Garantir que os valores sejam informados e que sejam númericos v0.1.1
// Garantir que uma operação válida seja informada v0.1.2
// Perguntar se deseja continuar calculando ou parar de calcular v0.2.0
// Registrar os resultados
// Perguntar se deseja imprimir os ultimos resultados ao fechar o programa. v0.3.0
// Transformação em arquivos. v0.4.0
// Transformação para 'Orientação a Objetos' v1.0.0

func main() {
	fmt.Println("Bem-vindos a magnifica calculadora feita em Golang!!!!!!!!!")

	fmt.Println("Informe um valor!")
	var umValor float64
	fmt.Scan(&umValor)

	fmt.Println("Informe uma operação")
	fmt.Println("1 - Somar")
	fmt.Println("2 - Subtrair")
	fmt.Println("3 - Multiplicar")
	fmt.Println("4 - Dividr")
	var operacao int
	fmt.Scan(&operacao)

	fmt.Println("Informe outro valor!")
	var outroValor float64
	fmt.Scan(&outroValor)

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
