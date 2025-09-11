package set

type IntSet struct {
	elements map[int]bool
}

func NewIntSet(elements ...int) *IntSet { //O(n)
	s := &IntSet{elements: make(map[int]bool)}
	for _, v := range elements {
		s.Add(v)
	}
	return s
}

func (s *IntSet) Add(element int) { //O(1)
	s.elements[element] = true
}

func (s *IntSet) Remove(element int) { //O(1)
	delete(s.elements, element)
}

func (s *IntSet) Contains(element int) bool { //O(1)
	return s.elements[element]
}

func (s *IntSet) Size() int { //O(1)
	return len(s.elements)
}

func (s *IntSet) Values() []int { //O(n)
	values := make([]int, 0, s.Size())
	for k := range s.elements {
		values = append(values, k)
	}
	return values
}

// Dado un conjunto A y un conjunto B, la unión de los conjuntos A y B será otro
// conjunto que estará formado por todos los elementos de A, con todos los
// elementos de B sin repetir ningún elemento.
func (s *IntSet) Union(other *IntSet) *IntSet {
	newC := NewIntSet()
	for k := range s.elements {
		newC.Add(k)
	}
	for k := range other.elements {
		newC.Add(k)
	}
	return newC
}

// Dado un conjunto A y un conjunto B, la intersección de los conjuntos A y B
// será otro conjunto que estará formado por los elementos de A y los elementos
// de B que sean comunes, los elementos no comunes entre A y B, serán excluidos.
func (s *IntSet) Intersection(other *IntSet) *IntSet {
	newC := NewIntSet()
	for k := range s.elements {
		if other.Contains(k) {
			newC.Add(k)
		}
	}
	return newC
}

// Dado un conjunto A y un conjunto B, la diferencia de los conjuntos A y B será
// otro conjunto que estará formado por los elementos de A que no estan
// presentes en B.
func (s *IntSet) Difference(other *IntSet) *IntSet {
	newC := NewIntSet()
	for k := range s.elements {
		if !other.Contains(k) {
			newC.Add(k)
		}
	}
	return newC
}

// Dado un conjunto A y un conjunto B, la diferencia simétrica de los conjuntos
// A y B será otro conjunto que estará formado por todos los elementos no
// comunes a los conjuntos A y B.
func (s *IntSet) SymmetricDifference(other *IntSet) *IntSet {
	difAB := s.Difference(other)
	difBA := other.Difference(s)
	return difAB.Union(difBA)
	//return s.Difference(other).Union(other.Difference(s))

}

// Un conjunto A es igual a un conjunto B si ambos conjuntos tienen los mismos
// elementos.
func (s *IntSet) Equal(other *IntSet) bool {
	for k := range s.elements {
		if !other.Contains(k) {
			return false
		}
	}
	for k := range other.elements {
		if !s.Contains(k) {
			return false
		}
	}
	return true
}

// El conjunto `other` es subconjunto de `s` si todos los elementos de `other`
// están incluidos en `s`.
func (s *IntSet) Subset(other *IntSet) bool {
	/*
		for k := range other.elements {
			if !s.Contains(k) {
				return false
			}
		}
		return true
	*/

	return other.Equal(s.Intersection(other))
}

// El conjunto `other` es superconjunto de `s` si `other` contiene todos los
// elementos de `s`.
func (s *IntSet) Superset(other *IntSet) bool {
	for k := range s.elements {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}
