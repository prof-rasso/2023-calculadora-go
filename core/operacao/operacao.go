package operacao

import (
	"calculadora-go/core/calculo"
	"errors"
)

type Operacao int

const (
	Soma Operacao = iota + 1
	Subtrai
	Multiplica
	Divide
)

var operacoes = map[int]Operacao{
	1: Soma,
	2: Subtrai,
	3: Multiplica,
	4: Divide,
}

func NovaOperacao(valor int) (Operacao, error) {
	if op, opEncontrada := operacoes[valor]; opEncontrada {
		return op, nil
	}
	return -1, errors.New("Operação inválida!")
}

func (op Operacao) CriarCalculoEntre(umValor, outroValor float64) (calculo.Calculo, error) {
	switch op {
	case Soma:
		return calculo.NovaSoma(umValor, outroValor), nil
	case Subtrai:
		return calculo.NovaSubtracao(umValor, outroValor), nil
	case Multiplica:
		return calculo.NovaMultiplicacao(umValor, outroValor), nil
	case Divide:
		return calculo.NovaDivisao(umValor, outroValor), nil
	default:
		return nil, errors.New("Não foi possível criar o cálculo!")
	}
}
