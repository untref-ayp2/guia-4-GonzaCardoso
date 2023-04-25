package ejercicios

import (
	"guia4/set"
)

func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {
	n := set.NewSet[T]()
	//s := set.NewSet[T]()
	for _, v := range conjuntos {
		n = v

	}
	/*for _, k := range conjuntos {
		s = set.Intersection(n, k)
		break
	}
	*/
	return n
}

func algo[T comparable](con *set.Set[T]) *set.Set[T] {
	m := set.NewSet[T]()
	return m
}
