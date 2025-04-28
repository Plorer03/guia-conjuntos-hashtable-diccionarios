package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSliceSet(t *testing.T) {
	set := NewSliceSet[int]()

	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())
}

func TestSliceSetAdd(t *testing.T) {
	set := NewSliceSet[int]()

	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestSliceSetAddMultiple(t *testing.T) {
	set := NewSliceSet[int]()

	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Size())
}

func TestSliceSetAddExistenteNoRepite(t *testing.T) {
	set := NewSliceSet[int]()

	set.Add(1)
	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestSliceSetContains(t *testing.T) {
	set := NewSliceSet[int]()
	set.Add(1)

	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestSliceSetRemove(t *testing.T) {
	set := NewSliceSet[int]()
	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())
}

func TestSliceSetRemoveNonExistent(t *testing.T) {
	set := NewSliceSet[int]()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Remove(2)
	assert.Equal(t, 1, set.Size())
}

func TestSliceSetSize(t *testing.T) {
	set := NewSliceSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.Equal(t, 2, set.Size())
}

func TestSliceSetValuesOnAnEmptySet(t *testing.T) {
	set := NewSliceSet[int]()
	values := set.Values()

	assert.Empty(t, values)
}

func TestSliceSetValuesOnANonEmptySet(t *testing.T) {
	set := NewSliceSet(1, 2)
	values := set.Values()

	assert.Len(t, values, 2)
	assert.ElementsMatch(t, []int{1, 2}, values)
}

func TestSliceSetStringEnSetVacio(t *testing.T) {
	set := NewSliceSet[int]()
	assert.Equal(t, "Set: {}", set.String())
}

func TestSliceSetStringEnSetNoVacio(t *testing.T) {
	set := NewSliceSet(1, 2)
	possibleRepresentations := []string{
		"Set: {1, 2}",
		"Set: {2, 1}",
	}

	assert.Contains(t, possibleRepresentations, set.String())
}
