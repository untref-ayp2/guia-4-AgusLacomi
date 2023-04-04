package ejercicios

import (
	"guia4/set"
)
// By AgusLacomi Punto 3
func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {

	s3 := set.NewSet[T]()

	for _, elemento := range s1.Values() {
		if !s2.Contains(elemento) {
			s3.Add(elemento)
		}
	}

	return s3
}
