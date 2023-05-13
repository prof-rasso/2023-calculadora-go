package calculo

import "calculadora-go/core/resultado"

type Subtracao struct {
	umValor    float64
	outroValor float64
}

func NovaSubtracao(umValor, outroValor float64) *Subtracao {
	return &Subtracao{umValor, outroValor}
}

func (d *Subtracao) Calcular() resultado.Resultado {
	return resultado.Novo("-", d.umValor, d.outroValor, d.umValor-d.outroValor)
}
