package ejercicios

import (
	"guia4/set"

	strUtil "github.com/agrison/go-commons-lang/stringUtils"
)

// By AgusLacomi Punto 1
func Letras(s string) *set.Set[string] {

	conjuntoLetras := set.NewSet[string]()

	for _, letra := range s {

		if !strUtil.IsBlank(string(letra)) {
			conjuntoLetras.Add(string(letra))
		}
	}

	return conjuntoLetras

}
