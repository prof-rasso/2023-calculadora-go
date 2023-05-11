package resultado

import "fmt"

var resultados = []string{}

func Registrar(umValor float64, sinal string, outroValor float64, resultado float64) string {
	resultadoStr := fmt.Sprintf("O resultado de %.2f %s %.2f é igual a %.2f", umValor, sinal, outroValor, resultado)

	resultados = append(resultados, resultadoStr)

	return resultadoStr
}

func Imprimir() {
	for i := len(resultados) - 1; i > 0; i-- {
		fmt.Println(resultados[i])
	}
}
