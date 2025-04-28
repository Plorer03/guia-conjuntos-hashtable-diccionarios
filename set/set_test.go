package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	set := NewSet[int]()

	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())
}

func TestSetAdd(t *testing.T) {
	set := NewSet[int]()

	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestSetAddMultiple(t *testing.T) {
	set := NewSet[int]()

	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Size())
}

func TestSetAddExistenteNoRepite(t *testing.T) {
	set := NewSet[int]()

	set.Add(1)
	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestSetContains(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)

	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestSetRemove(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())
}

func TestSetRemoveNonExistent(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Remove(2)
	assert.Equal(t, 1, set.Size())
}

func TestSetSize(t *testing.T) {
	set := NewSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.Equal(t, 2, set.Size())
}

func TestSetValuesOnAnEmptySet(t *testing.T) {
	set := NewSet[int]()
	values := set.Values()

	assert.Empty(t, values)
}

func TestSetValuesOnANonEmptySet(t *testing.T) {
	set := NewSet(1, 2)
	values := set.Values()

	assert.Len(t, values, 2)
	assert.ElementsMatch(t, []int{1, 2}, values)
}

func TestSetStringEnSetVacio(t *testing.T) {
	set := NewSet[int]()
	assert.Equal(t, "Set: {}", set.String())
}

func TestSetStringEnSetNoVacio(t *testing.T) {
	set := NewSet(1, 2)
	possibleRepresentations := []string{
		"Set: {1, 2}",
		"Set: {2, 1}",
	}

	assert.Contains(t, possibleRepresentations, set.String())
}
