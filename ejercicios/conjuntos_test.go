package ejercicios

import (
	"testing"
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/set"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEliminarRepetidos(t *testing.T) {
	array := []int{1, 2, 7, 2, 1, 3}

	result := EliminarRepetidos(array)

	require.NotNil(t, result)
	assert.Equal(t, 4, len(result))
	for _, v := range []int{1, 2, 7, 3} {
		assert.Contains(t, result, v)
	}
}

func TestInterseccionMultiple(t *testing.T) {
	set1 := set.NewListSet(1, 2, 3)
	set2 := set.NewListSet(2, 3)
	set3 := set.NewListSet(2)

	result1 := InterseccionMultiple(set1, set2, set3)

	require.NotNil(t, result1)
	assert.Equal(t, 1, result1.Size())
	for _, v := range []int{2} {
		assert.True(t, result1.Contains(v))
	}

	set4 := set.NewListSet(7, 11)
	set5 := set.NewListSet[int]()
	set6 := set.NewListSet[int]()

	result2 := InterseccionMultiple(set4, set5, set6)

	require.NotNil(t, result2)
	assert.Equal(t, 0, result2.Size())

	set7 := set.NewListSet(7, 11)

	result3 := InterseccionMultiple(set7)

	require.NotNil(t, result3)
	assert.Equal(t, 0, result3.Size())
}
