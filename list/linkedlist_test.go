package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINTERNALNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestINTERNALLinkedListPrependOnEmptyList(t *testing.T) {
	list := NewLinkedList[string]()

	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestINTERNALLinkedListPrependOnNonEmptyList(t *testing.T) {
	list := NewLinkedList[string]()
	list.Append("1")

	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestINTERNALLinkedListAppendOnEmptyList(t *testing.T) {
	list := NewLinkedList[string]()

	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

func TestINTERNALLinkedListAppendOnNonEmptyList(t *testing.T) {
	list := NewLinkedList[string]()
	list.Append("1")

	list.Append("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "2", list.Tail().Data())
}

func TestINTERNALLinkedListClear(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Clear()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestINTERNALLinkedListIsEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	assert.True(t, list.IsEmpty())

	list.Append(1)
	assert.False(t, list.IsEmpty())
}

func TestINTERNALLinkedListSize(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Equal(t, 0, list.Size())

	list.Append(1)
	assert.Equal(t, 1, list.Size())

	list.Append(2)
	assert.Equal(t, 2, list.Size())
}

func TestINTERNALLinkedListHead(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Nil(t, list.Head())

	list.Append(1)
	assert.Equal(t, 1, list.Head().Data())

	list.Append(2)
	assert.Equal(t, 1, list.Head().Data())
}

func TestINTERNALLinkedListTail(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Nil(t, list.Tail())

	list.Append(1)
	assert.Equal(t, 1, list.Tail().Data())

	list.Append(2)
	assert.Equal(t, 2, list.Tail().Data())
}

func TestINTERNALLinkedListFind(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	nodo := list.Find(2)
	assert.NotNil(t, nodo)
	assert.Equal(t, 2, nodo.Data())

	nodo = list.Find(4)
	assert.Nil(t, nodo)
}

func TestINTERNALLinkedListRemoveFirst(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.RemoveFirst()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.RemoveFirst()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestINTERNALLinkedListRemoveLast(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.RemoveLast()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 2, list.Tail().Data())

	list.RemoveLast()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())

	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestINTERNALLinkedListRemove(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Remove(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.Remove(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())

	list.Remove(3)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

func TestINTERNALLinkedListRemoveOnLastElement(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Remove(4)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
}

func TestINTERNALLinkedListRemoveNotExists(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Remove(4)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3, list.Tail().Data())
}

func TestINTERNALLinkedListRemoveFirstOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
}

func TestINTERNALLinkedListRemoveLastOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
}

func TestINTERNALLinkedListStringOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	assert.Equal(t, "LinkedList: []", list.String())
}

func TestINTERNALLinkedListString(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.Equal(t, "LinkedList: [1] → [2] → [3]", list.String())
}
