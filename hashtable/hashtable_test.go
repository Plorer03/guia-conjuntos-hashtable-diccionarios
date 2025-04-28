package hashtable

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// hashFunc calcula el índice del bucket para una clave dada. Se utiliza la técnica de Mulitiplicación Polinómica.
func hashFn(key string) uint {
	const a float64 = 11.0
	l := len(key)
	var hash uint = 0
	for i, c := range key {
		hash += uint(c) * uint(math.Pow(a, float64(l-i-1)))
	}
	return hash
}

func TestNewHashTable(t *testing.T) {
	ht := NewHashTable[string, int](0, 0, hashFn)
	require.NotNil(t, ht)
	assert.Equal(t, uint(17), ht.capacity)
	assert.Equal(t, uint(0), ht.size)
	assert.Equal(t, float32(0.75), ht.loadFactor)
	assert.Equal(t, uint(12), ht.threshold)
}

func TestHashTable(t *testing.T) {
	ht := NewHashTable[string, int](0, 0, hashFn)
	require.NotNil(t, ht)

	// Verificar que la tabla está vacía
	assert.Equal(t, uint(0), ht.Size())
	assert.Equal(t, uint(17), ht.capacity)

	// Agregar elementos
	ht.Put("key1", 1)
	ht.Put("key2", 2)
	ht.Put("key3", 3)

	// Actulizar un valor existente
	ht.Put("key2", 4)

	// Verificar que los elementos se han agregado correctamente
	assert.Equal(t, uint(3), ht.size)
	assert.Equal(t, uint(17), ht.capacity)

	// Obtener valores
	val1, found1 := ht.Get("key1")
	assert.True(t, found1)
	assert.Equal(t, 1, val1)

	val2, found2 := ht.Get("key2")
	assert.True(t, found2)
	assert.Equal(t, 4, val2)

	val3, found3 := ht.Get("key3")
	assert.True(t, found3)
	assert.Equal(t, 3, val3)

	// Probar la eliminación de un elemento
	ht.Remove("key2")
	val2Removed, found2Removed := ht.Get("key2")
	assert.False(t, found2Removed)
	assert.Equal(t, 0, val2Removed)

	// Verificar el tamaño después de eliminar un elemento
	assert.Equal(t, uint(2), ht.Size())
}
