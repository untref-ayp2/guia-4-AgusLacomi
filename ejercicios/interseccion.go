package ejercicios

import (
	"guia4/set"
)

// By AgusLacomi punto 4
func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {

	conjuntoAretornar := conjuntos[0]

	for i := 1; i < len(conjuntos); i++ {
		conjuntoAretornar = set.Intersection(conjuntos[i], conjuntoAretornar)
	}

	return conjuntoAretornar
}
