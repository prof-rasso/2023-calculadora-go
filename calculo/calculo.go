package calculo

func Calcular(operacao int, umValor float64, outroValor float64) (resultado float64, sinal string) {
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
	return
}
