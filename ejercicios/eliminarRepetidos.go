package ejercicios

import "guia4/set"

//By AgusLacomi Punto 4
func EliminarRepetidos[T comparable](arreglo []T) []T {

	conjunto := set.NewSet[T]()

	for _, elemento := range arreglo {
		conjunto.Add(elemento)
	}

	return conjunto.Values()
}
