package ejercicios

import (
	"guia4/set"
)

func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {
	s3 := set.Intersection(s1, s2)
	s1 = set.Union(s1, s2)
	s1 = set.Difference(s1, s3)
	return s1
}
