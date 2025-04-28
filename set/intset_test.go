package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIntSet(t *testing.T) {
	set := NewIntSet()

	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())
}

func TestIntSetAdd(t *testing.T) {
	set := NewIntSet()

	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestIntSetAddMultiple(t *testing.T) {
	set := NewIntSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.Equal(t, 3, set.Size())
}

func TestIntSetAddExistenteNoRepite(t *testing.T) {
	set := NewIntSet()

	set.Add(1)
	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestIntSetContains(t *testing.T) {
	set := NewIntSet()
	set.Add(1)

	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestIntSetRemove(t *testing.T) {
	set := NewIntSet(1, 2)

	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())
}

func TestIntSetRemoveNonExistent(t *testing.T) {
	set := NewIntSet()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Remove(2)
	assert.Equal(t, 1, set.Size())
}

func TestIntSetSize(t *testing.T) {
	set := NewIntSet()
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.Equal(t, 2, set.Size())
}

func TestIntSetValuesOnAnEmptySet(t *testing.T) {
	set := NewIntSet()
	values := set.Values()

	assert.Empty(t, values)
}

func TestIntSetValuesOnANonEmptySet(t *testing.T) {
	set := NewIntSet()
	set.Add(1)
	set.Add(2)
	values := set.Values()

	assert.Len(t, values, 2)
	assert.ElementsMatch(t, []int{1, 2}, values)
}

func TestUnion(t *testing.T) {
	set1 := NewIntSet(1, 2)
	set2 := NewIntSet(1, 3)

	result1 := set1.Union(set2)

	require.NotNil(t, result1)
	assert.Equal(t, 3, result1.Size())
	for _, v := range []int{3, 1, 2} {
		assert.True(t, result1.Contains(v))
	}

	set3 := NewIntSet()
	set4 := NewIntSet(1, 3)

	result2 := set3.Union(set4)

	require.NotNil(t, result2)
	assert.Equal(t, 2, result2.Size())
	for _, v := range []int{3, 1} {
		assert.True(t, result2.Contains(v))
	}
}

func TestIntersection(t *testing.T) {
	set1 := NewIntSet(1, 2, 5)
	set2 := NewIntSet(1, 3, 5)

	result1 := set1.Intersection(set2)

	require.NotNil(t, result1)
	assert.Equal(t, 2, result1.Size())
	for _, v := range []int{1, 5} {
		assert.True(t, result1.Contains(v))
	}

	set3 := NewIntSet()
	set4 := NewIntSet(1, 3, 5)

	result2 := set3.Intersection(set4)

	require.NotNil(t, result2)
	assert.Equal(t, 0, result2.Size())
}

func TestDifference(t *testing.T) {
	set1 := NewIntSet(1, 2, 4, 7)
	set2 := NewIntSet(1, 3, 7)

	result1 := set1.Difference(set2)

	require.NotNil(t, result1)
	assert.Equal(t, 2, result1.Size())
	for _, v := range []int{2, 4} {
		assert.True(t, result1.Contains(v))
	}

	set3 := NewIntSet(1, 2, 4, 7)
	set4 := NewIntSet()

	result2 := set3.Difference(set4)

	require.NotNil(t, result2)
	assert.Equal(t, 4, result2.Size())
	for _, v := range []int{1, 2, 4, 7} {
		assert.True(t, result2.Contains(v))
	}
}

func TestSymmetricDifference(t *testing.T) {
	set1 := NewIntSet(1, 2, 3, 4, 7, 11)
	set2 := NewIntSet(3, 4, 11, 12)

	result1 := set1.SymmetricDifference(set2)

	require.NotNil(t, result1)
	assert.Equal(t, 4, result1.Size())
	for _, v := range []int{1, 2, 7, 12} {
		assert.True(t, result1.Contains(v))
	}

	set3 := NewIntSet(1, 2, 3, 4, 7, 11)
	set4 := NewIntSet()

	result2 := set3.SymmetricDifference(set4)

	require.NotNil(t, result2)
	assert.Equal(t, 6, result2.Size())
	for _, v := range []int{1, 2, 3, 4, 7, 11} {
		assert.True(t, result2.Contains(v))
	}
}

func TestEqual(t *testing.T) {
	set1 := NewIntSet(1, 2)
	set2 := NewIntSet(2, 1)
	assert.True(t, set1.Equal(set2))

	set3 := NewIntSet(1, 2)
	set4 := NewIntSet(2, 3)
	assert.False(t, set3.Equal(set4))
}

func TestSubset(t *testing.T) {
	set1 := NewIntSet(1, 2)
	set2 := NewIntSet(1)
	assert.True(t, set1.Subset(set2))

	set3 := NewIntSet(1, 2)
	set4 := NewIntSet(5)
	assert.False(t, set3.Subset(set4))
}

func TestSuperset(t *testing.T) {
	set1 := NewIntSet(1)
	set2 := NewIntSet(1, 2)
	assert.True(t, set1.Superset(set2))

	set3 := NewIntSet(5)
	set4 := NewIntSet(1, 2)
	assert.False(t, set3.Superset(set4))
}
