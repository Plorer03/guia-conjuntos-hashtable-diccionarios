package dictionary

import (
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/hashtable"
)

type Dictionary[K comparable] struct {
	hash *hashtable.HashTable[K, struct{}]
}

func NewDictionary[K comparable](hashFunc func(key K) uint) *Dictionary[K] {
	return &Dictionary[K]{
		hash: hashtable.NewHashTable[K, struct{}](0, 0, hashFunc),
	}
}

func (d *Dictionary[K, V]) Set(key K, value V) {
	d.hash.Put(key, value)
}
