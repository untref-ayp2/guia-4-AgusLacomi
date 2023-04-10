package tests

import (
	"guia4/ejercicios"
	"guia4/set"
	"reflect"
	"testing"
)

func TestInterseccion(t *testing.T) {
	set1 := set.NewSet("a", "b", "c", "d")
	set2 := set.NewSet("b", "c", "e", "d")
	set3 := set.NewSet("b", "d", "f")

	resultado := ejercicios.Interseccion(set1, set2, set3)
	expected := set.NewSet("d", "b")

	if !reflect.DeepEqual(resultado, expected) {
		t.Errorf("La intersección debería ser %v, pero se obtuvo %v", expected, resultado)
	}
}
