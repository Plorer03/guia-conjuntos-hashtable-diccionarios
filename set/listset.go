// Package set proporciona una implementación de un conjunto genérico basado en una lista dinámica.
// Expone la estructura ListSet y sus métodos para manipular un conjunto.
package set

import "untref-ayp2/guia-conjuntos-hashes-diccionarios/list"

// ListSet implementa un conjunto sobre una lista enlazada simple.
type ListSet[T comparable] struct {
	elements *list.LinkedList[T]
}

// NewListSet crea un nuevo conjunto vacío y lo inicializa con los elementos especificados.
//
// Uso:
//
//	s1 := set.NewListSet[int]() // Crea un nuevo conjunto vacío.
//	s2 := set.NewListSet[int](1, 2, 3) // Crea un nuevo conjunto con los elementos 1, 2 y 3.
//
// Parámetros:
//   - `elements`: los elementos con los que inicializar el conjunto.
func NewListSet[T comparable](elements ...T) *ListSet[T] {
	// Implementar
	return nil
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
func (s *ListSet[T]) Contains(element T) bool {
	// Implementar
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
func (s *ListSet[T]) Add(elements ...T) {
	// Implementar
}

// Remove elimina el elemento especificado del conjunto.
//
// Uso:
//
//	s.Remove(10) // Elimina el elemento 10 del conjunto.
//
// Parámetros:
//   - `element`: el elemento a eliminar del conjunto.
func (s *ListSet[T]) Remove(element T) {
	// Implementar
}

// Size devuelve la cantidad de elementos en el conjunto.
//
// Uso:
//
//	size := s.Size() // Obtiene la cantidad de elementos en el conjunto.
//
// Retorna:
//   - la cantidad de elementos en el conjunto.
func (s *ListSet[T]) Size() int {
	// Implementar
	return -1
}

// Values devuelve los elementos del conjunto.
//
// Uso:
//
//	values := s.Values() // Obtiene los elementos del conjunto como un slice.
//
// Retorna:
//   - los elementos del conjunto como un slice.
func (s *ListSet[T]) Values() []T {
	// Implementar
	return nil
}

// String devuelve una representación en cadena del conjunto.
//
// Uso:
//
//	fmt.Println(s) // Muestra el conjunto como una cadena.
//
// Retorna:
//   - una representación en cadena del conjunto.
func (s *ListSet[T]) String() string {
	// Implementar
	return ""
}
