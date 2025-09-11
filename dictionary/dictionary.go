package dictionary

import (
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/hashtable"
)

type Dictionary[K comparable, V any] struct {
	hash *hashtable.HashTable[K, V]
}

func NewDictionary[K comparable, V any](hashFunc func(key K) uint) *Dictionary[K, V] {
	return &Dictionary[K, V]{hash: hashtable.NewHashTable[K, V](0, 0, hashFunc)}
}

func (d *Dictionary[K, V]) Put(key K, value V) {
	d.hash.Put(key, value)
}

func (d *Dictionary[K, V]) Get(key K) V {
	value, _ := d.hash.Get(key)
	return value
}

func (d *Dictionary[K, V]) Size() int {
	return int(d.hash.Size())
}

func (d *Dictionary[K, V]) Remove(key K) {
	d.hash.Remove(key)
}

func (d *Dictionary[K, V]) Contains(key K) bool {
	_, exists := d.hash.Get(key)
	return exists
}

func (d *Dictionary[K, V]) Keys() []K {
	return d.hash.Keys()
}

func (d *Dictionary[K, V]) Values() []V {
	return d.hash.Values()
}

func (d *Dictionary[K, V]) IsEmpty() bool {
	return d.hash.IsEmpty()
}

func (d *Dictionary[K, V]) Clear() {
	d.hash.Clear()
}
