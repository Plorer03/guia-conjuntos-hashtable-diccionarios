package dictionary

import (
	"untref-ayp2/guia-conjuntos-hashes-diccionarios/hashtable"
)

type Dictionary[K comparable, V any] struct {
	hash *hashtable.HashTable[K, V]
}

func NewDictionary[K comparable, V any](hashFunc func(key K) uint) *Dictionary[K, V] {

}

func (d *Dictionary[K, V]) Set(key K, value V) {
}

func (d *Dictionary[K, V]) Get(key K) V {
}

func (d *Dictionary[K, V]) Size() int {
}

func (d *Dictionary[K, V]) Remove(key K) {
}

func (d *Dictionary[K, V]) Contains(key K) bool {
}

func (d *Dictionary[K, V]) Keys() []K {
}

func (d *Dictionary[K, V]) Values() []V {
}

func (d *Dictionary[K, V]) IsEmpty() bool {
}

func (d *Dictionary[K, V]) Clear() {
}
