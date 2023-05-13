package calculo

import "calculadora-go/core/resultado"

type Multiplicacao struct {
	umValor    float64
	outroValor float64
}

func NovaMultiplicacao(umValor, outroValor float64) *Multiplicacao {
	return &Multiplicacao{umValor, outroValor}
}

func (d *Multiplicacao) Calcular() resultado.Resultado {
	return resultado.Novo("*", d.umValor, d.outroValor, d.umValor*d.outroValor)
}
