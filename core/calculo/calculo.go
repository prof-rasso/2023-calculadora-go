package calculo

import "calculadora-go/core/resultado"

type Calculo interface {
	Calcular() resultado.Resultado
}
