package ejercicios

import (
	"guia4/set"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func Letras(s string) *set.Set[string] {
	conjunto := set.NewSet[string]()
	for _, v := range s {
		if !stringUtils.IsBlank(string(v)) {
			conjunto.Add(string(v))
		}
	}
	return conjunto
}
