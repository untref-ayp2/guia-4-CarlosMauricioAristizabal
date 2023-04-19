package ejercicios

import (
	"guia4/set"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func Letras(texto string) *set.Set[string] {
	//panic("Not implemented")}
	conjunto := set.NewSet[string]()
	for _, letra := range texto {
		if !stringUtils.IsBlank(string(letra)) {
			conjunto.Add(string(letra))
		}
	}
	return conjunto
}
