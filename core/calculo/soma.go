package calculo

import "calculadora-go/core/resultado"

type Soma struct {
	umValor    float64
	outroValor float64
}

func NovaSoma(umValor, outroValor float64) *Soma {
	return &Soma{umValor, outroValor}
}

func (d *Soma) Calcular() resultado.Resultado {
	return resultado.Novo("+", d.umValor, d.outroValor, d.umValor+d.outroValor)
}
