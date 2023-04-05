package ejercicios

import "guia4/set"

func EliminarRepetidos[T comparable](arreglo []T) []T {

	conjunto := set.NewSet[T]()
	for _, valor := range arreglo {
		conjunto.Add(valor)
	}
	return conjunto.Values()
}
func EliminarRepetidos2[T comparable](arreglo []T) []T {
	conjunto := set.NewSet(arreglo...)
	return conjunto.Values()
}
