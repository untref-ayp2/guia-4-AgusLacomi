package ejercicios

import (
	"guia4/set"
)

// By AgusLacomi Punto 3
func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {

	return set.Union(set.Difference(s1, s2), set.Difference(s2, s1))
}
