package formas

import (
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (r Circulo) Area() float64 {
	return math.Pi * (r.raio * r.raio)
}
