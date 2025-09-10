package dictionary

import (
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/hashtable"
)

type Dictionary[K comparable, V any] struct {
	hash *hashtable.HashTable[K, V]
}

func NewDictionary[K comparable, V any](hashFunc func(key K) uint) *Dictionary[K, V] {
	return &Dictionary[K, V]{
		hash: hashtable.NewHashTable[K, V](0, 0, hashFunc),
	}
}

func (d *Dictionary[K, V]) Set(key K, value V) {
	d.hash.Put(key, value)
}
