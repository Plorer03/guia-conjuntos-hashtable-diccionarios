package ejercicios

import (
	"math"
	"strings"
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/dictionary"
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/list"
)

func hashFn(key string) uint {
	const a float64 = 11.0
	l := len(key)
	var hash uint = 0
	for i, c := range key {
		hash += uint(c) * uint(math.Pow(a, float64(l-i-1)))
	}
	return hash
}

func Traducir(texto string, dict dictionary.Dictionary[string, string]) string {
	palabras := strings.Split(texto, " ")
	var traduccion []string
	for _, palabra := range palabras {
		if dict.Contains(palabra) {
			traduccion = append(traduccion, dict.Get(palabra))
		} else {
			traduccion = append(traduccion, "error")
		}
	}
	return strings.Join(traduccion, " ")
}

func Frecuencia(texto string) *dictionary.Dictionary[string, int] {
	dict := dictionary.NewDictionary[string, int](hashFn)
	for _, letra := range texto {
		if string(letra) != " " {
			cant := dict.Get(string(letra))
			if cant == 0 {
				cant = 1
			} else {
				cant++
			}
			dict.Put(string(letra), cant)
		}
	}
	return dict
}

func Interseccion(s1 []string, s2 []string) *list.LinkedList[string] {
	//Implementar
	return nil
}

func InformacionSolicitada(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string] {
	//Implementar
	return nil
}
