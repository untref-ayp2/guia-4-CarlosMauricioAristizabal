package ejercicios

import (
	"guia4/set"
)

func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {
	//var interseccion []T
	interseccion := conjuntos[0]
	for i := range conjuntos {
		interseccion = set.Intersection(conjuntos[i], interseccion)
	}
	return interseccion
}
