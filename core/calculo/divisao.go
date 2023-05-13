package calculo

import "calculadora-go/core/resultado"

type Divisao struct {
	umValor    float64
	outroValor float64
}

func NovaDivisao(umValor, outroValor float64) *Divisao {
	return &Divisao{umValor, outroValor}
}

func (d *Divisao) Calcular() resultado.Resultado {
	return resultado.Novo("/", d.umValor, d.outroValor, d.umValor/d.outroValor)
}
