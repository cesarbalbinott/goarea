package goarea

import (
	"math"
)

//Pi valor de pi
const Pi = 3.1416

//Circ calcula area da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
