package ejercicios

import (
	"testing"
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/dictionary"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTraducir(t *testing.T) {
	dic := dictionary.NewDictionary[string, string](hashFn)
	dic.Put("Dungeons", "Calabozos")
	dic.Put("Dragons", "Dragones")

	salida := Traducir("Dungeons", *dic)
	assert.Equal(t, "Calabozos", salida)

	salida = Traducir("Dwarf", *dic)
	assert.Equal(t, "error", salida)

	salida = Traducir("Dungeons & Dragons", *dic)
	assert.Equal(t, "Calabozos error Dragones", salida)
}

func TestFrecuencia(t *testing.T) {
	dict := Frecuencia("ahora")
	require.NotNil(t, dict)
	assert.Equal(t, 4, dict.Size())
	val := dict.Get("a")

	assert.Equal(t, 2, val)

	val = dict.Get("h")

	assert.Equal(t, 1, val)

	val = dict.Get("o")

	assert.Equal(t, 1, val)

	val = dict.Get("r")

	assert.Equal(t, 1, val)

	dict = Frecuencia("hoy es lunes")
	assert.Equal(t, 8, dict.Size())
	val = dict.Get("h")

	assert.Equal(t, 1, val)

	val = dict.Get("o")

	assert.Equal(t, 1, val)

	val = dict.Get("y")

	assert.Equal(t, 1, val)

	val = dict.Get("e")

	assert.Equal(t, 2, val)

	val = dict.Get("s")

	assert.Equal(t, 2, val)

	val = dict.Get("l")

	assert.Equal(t, 1, val)

	val = dict.Get("u")

	assert.Equal(t, 1, val)

	val = dict.Get("n")

	assert.Equal(t, 1, val)

}

func TestInterseccion(t *testing.T) {
	s1 := []string{"A", "B", "C"}
	s2 := []string{"A", "D", "E"}
	list := Interseccion(s1, s2)
	require.NotNil(t, list)
	assert.Equal(t, 1, list.Size())
	assert.NotNil(t, list.Find("A"))

	s1 = []string{"A", "B"}
	s2 = []string{"C", "D"}
	list = Interseccion(s1, s2)
	assert.Equal(t, 0, list.Size())

	s1 = []string{"A"}
	s2 = make([]string, 0)
	list = Interseccion(s1, s2)
	assert.Equal(t, 0, list.Size())

	s1 = make([]string, 0)
	s2 = []string{"A", "B"}
	list = Interseccion(s1, s2)
	assert.Equal(t, 0, list.Size())
}
func TestInformacionSolicitada(t *testing.T) {
	entrada := dictionary.NewDictionary[string, []string](hashFn)
	sl1 := []string{"Ana", "Pedro"}
	sl2 := []string{"Ana"}
	entrada.Put("Mie 10", sl1)
	entrada.Put("Vie 12", sl2)
	salida := InformacionSolicitada(*entrada)
	require.NotNil(t, salida)
	valPedro := salida.Get("Pedro")
	assert.ElementsMatch(t, []string{"Mie 10"}, valPedro)
	valAna := salida.Get("Ana")
	assert.ElementsMatch(t, []string{"Mie 10", "Vie 12"}, valAna)
}
