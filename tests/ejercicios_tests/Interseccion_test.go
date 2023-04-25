package tests

import (
	"fmt"
	"guia4/ejercicios"
	"guia4/set"
	"testing"
)

func TestInteseccionEntreDos(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := ejercicios.Interseccion(s1, s2)
	fmt.Println(s3)
	if s3.Size() != 2 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 2)
	}

}
func TestInteseccionEntreTres(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s4 := set.NewSet(6, 9, 2, 44)
	s3 := ejercicios.Interseccion(s1, s2, s4)
	fmt.Print(s3)
	if s3.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 4)
	}

}
