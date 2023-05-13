package resultado

import "fmt"

var resultados = []Resultado{}

type Resultado struct {
	resultado string
}

func (r Resultado) String() string {
	return r.resultado
}

func Novo(sinal string, umValor, outroValor, resultado float64) Resultado {
	valor := fmt.Sprintf("O resultado de %0.2f %s %0.2f é igual á %0.2f", umValor, sinal, outroValor, resultado)
	return Resultado{valor}
}

func Registrar(resultado Resultado) {
	resultados = append(resultados, resultado)
}

func Imprimir() {
	for _, item := range resultados {
		fmt.Println(item)
	}
}
