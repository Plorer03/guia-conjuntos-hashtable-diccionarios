package hashtable

// OpenHashTable es una tabla hash cerrada que utiliza un arreglo para almacenar
// elementos. La tabla solo soporta string como claves y cualquier tipo como
// valores. En cada posición del arreglo se almacena un par clave-valor.
type OpenHashTable[K comparable, V any] struct {
	// arreglo de entradas de la tabla hash.
	buckets [][]*hashTableEntry[K, V]
	// size es el número de elementos en la tabla.
	size uint
	// capacity es la capacidad de la tabla.
	capacity uint
	// loadFactor es el factor de carga de la tabla.
	loadFactor float32
	// threshold es el umbral de carga para redimensionar la tabla.
	threshold uint
	// hashFunc calcula el índice del bucket para una clave dada.
	hashFunc func(key K) uint
}
