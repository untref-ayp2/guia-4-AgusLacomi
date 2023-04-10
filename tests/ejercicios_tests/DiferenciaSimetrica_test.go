package tests

import (
	"guia4/ejercicios"
	"guia4/set"
	"testing"
)

func TestDiferenciaSimetrica(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := ejercicios.DiferenciaSimetrica(s1, s2)
	if s3.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 4)
	}
}

func TestNilpotnecia(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(1, 2, 3, 4)
	s3 := ejercicios.DiferenciaSimetrica(s1, s2)
	if s3.Size() != 0 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 4)
	}
}

func TestAsociativa(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := set.NewSet(4, 5, 6, 7)

	resultA := ejercicios.DiferenciaSimetrica(s1, s2)
	resultA = ejercicios.DiferenciaSimetrica(resultA, s3)

	resultB := ejercicios.DiferenciaSimetrica(s2, s3)
	resultB = ejercicios.DiferenciaSimetrica(s1, resultB)

	if !set.Equal(resultA, resultB) {
		t.Errorf("(A Δ B) Δ C = %s y A Δ (B Δ C) = %s", resultA.String(), resultB.String())
	}
}

func TestConmutativa(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)

	resultA := ejercicios.DiferenciaSimetrica(s1, s2)
	resultB := ejercicios.DiferenciaSimetrica(s2, s1)

	if !set.Equal(resultA, resultB) {
		t.Errorf("(A Δ B) = %s y (B Δ A) = %s", resultA.String(), resultB.String())
	}

}

func TestNulo(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)

	resultado := ejercicios.DiferenciaSimetrica(s1, set.NewSet[int]())

	if !set.Equal(s1, resultado) {
		t.Errorf("Error")
	}

}
