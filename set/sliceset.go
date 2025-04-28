// Package set proporciona una implementación de un conjunto genérico basado en un mapa.
// Expone la estructura Set y sus métodos para manipular un conjunto.
package set

import (
	"fmt"
	"strings"
)

// SliceSet implementa un conjunto sobre una lista enlazada simple.
type SliceSet[T comparable] struct {
	elements []T
}

// NewSet crea un nuevo conjunto vacío y lo inicializa con los elementos especificados.
//
// Uso:
//
//	s1 := set.NewSet[int]() // Crea un nuevo conjunto vacío.
//	s2 := set.NewSet[int](1, 2, 3) // Crea un nuevo conjunto con los elementos 1, 2 y 3.
//
// Parámetros:
//   - `elements`: los elementos con los que inicializar el conjunto.
func NewSliceSet[T comparable](elements ...T) *SliceSet[T] {
	s := &SliceSet[T]{elements: []T{}}
	s.Add(elements...)
	return s
}

// Contains verifica si el conjunto contiene el elemento especificado.
//
// Uso:
//
//	if s.Contains(10) {
//		fmt.Println("El conjunto contiene el elemento 10.")
//	}
//
// Parámetros:
//   - `element`: el elemento a buscar en el conjunto.
//
// Retorna:
//   - `true` si el conjunto contiene el elemento; `false` en caso contrario.
func (s *SliceSet[T]) Contains(element T) bool {
	for _, v := range s.elements {
		if v == element {
			return true
		}
	}
	return false
}

// Add agrega los elementos especificados al conjunto.
//
// Uso:
//
//	s.Add(10, 20, 30) // Agrega los elementos 10, 20 y 30 al conjunto.
//
// Parámetros:
//   - `elements`: los elementos a agregar al conjunto.
func (s *SliceSet[T]) Add(elements ...T) {
	for _, element := range elements {
		if !s.Contains(element) {
			s.elements = append(s.elements, element)
		}
	}
}

// Remove elimina el elemento especificado del conjunto.
//
// Uso:
//
//	s.Remove(10) // Elimina el elemento 10 del conjunto.
//
// Parámetros:
//   - `element`: el elemento a eliminar del conjunto.
func (s *SliceSet[T]) Remove(element T) {
	for i, v := range s.elements {
		if v == element {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			return
		}
	}
}

// Size devuelve la cantidad de elementos en el conjunto.
//
// Uso:
//
//	size := s.Size() // Obtiene la cantidad de elementos en el conjunto.
//
// Retorna:
//   - la cantidad de elementos en el conjunto.
func (s *SliceSet[T]) Size() int {
	return len(s.elements)
}

// Values devuelve los elementos del conjunto.
//
// Uso:
//
//	values := s.Values() // Obtiene los elementos del conjunto como un slice.
//
// Retorna:
//   - los elementos del conjunto como un slice.
func (s *SliceSet[T]) Values() []T {
	// Se crea una copia de todos los elementos para evitar modificaciones externas.
	values := make([]T, 0, len(s.elements))
	values = append(values, s.elements...)
	return values
}

// String devuelve una representación en cadena del conjunto.
//
// Uso:
//
//	fmt.Println(s) // Muestra el conjunto como una cadena.
//
// Retorna:
//   - una representación en cadena del conjunto.
func (s *SliceSet[T]) String() string {
	str := "Set: {"
	strValues := make([]string, 0, len(s.elements))
	for _, v := range s.elements {
		strValues = append(strValues, fmt.Sprint(v))
	}
	str += strings.Join(strValues, ", ")
	str += "}"
	return str
}
