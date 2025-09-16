// Package bitmap proporciona una implementación de un mapa de bits sobre un entero de 32 bits.
// Expone la estructura BitMap y sus métodos para manipular un mapa de bits.
package bitmap

import "errors"

const (
	// BitmapSize define el tamaño del mapa de bits.
	BitmapSize uint8 = 32
)

// BitMap implementa un mapa de bits sobre un entero de 32 bits.
type BitMap struct {
	bits uint32
}

// NewBitMap crea un mapa de bits y lo inicializa con todos los bits apagados.
//
// Uso:
//
//	bm := bitmap.NewBitMap() // Crea un nuevo mapa de bits.
func NewBitMap() *BitMap {
	return &BitMap{bits: 0b0} //00000000_00000000_00000000_000000000
}

// On enciende el bit de la posición indicada.
//
// Uso:
//
//	bm.On(5) // Enciende el bit en la posición 5.
//
// Parámetros:
//   - `pos`: la posición del bit a encender.
//
// Retorna:
//   - `error` si la posición es inválida.
func (bm *BitMap) On(pos uint8) error {
	//00000000_00000000_00000000_000000000
	//00000000_00000000_00000000_000000001
	if isOutOfRange(pos) {
		return errors.New("posición no válida")
	}
	bm.bits = bm.bits | 0b1<<pos
	return nil
}

// Off apaga el bit de la posición indicada.
//
// Uso:
//
//	bm.Off(5) // Apaga el bit en la posición 5.
//
// Parámetros:
//   - `pos`: la posición del bit a apagar.
//
// Retorna:
//   - `error` si la posición es inválida.
func (bm *BitMap) Off(pos uint8) error {
	//000000000
	//000000001
	//000010000
	//000010000
	if isOutOfRange(pos) {
		return errors.New("posición no válida")
	}
	//bm.bits = (bm.bits ^ 0b1<<pos) & 0b1 << pos
	bm.bits &= ^(0b1 << pos)
	return nil
}

// IsOn obtiene el estado del bit en la posición indicada.
//
// Uso:
//
//	on, _ := bm.IsOn(5) // Verifica si el bit en la posición 5 está encendido.
//
// Parámetros:
//   - `pos`: la posición del bit a verificar.
//
// Retorna:
//   - `true` si el bit está encendido; `false` si está apagado.
func (bm *BitMap) IsOn(pos uint8) (bool, error) {
	//101
	//011
	//001
	if isOutOfRange(pos) {
		return false, errors.New("posición no válida")
	}
	return bm.bits&(0b1<<pos) != 0b0, nil
}

// IsOff obtiene el estado del bit en la posición indicada.
//
// Uso:
//
//	off, _ := bm.IsOff(5) // Verifica si el bit en la posición 5 está apagado.
//
// Parámetros:
//   - `pos`: la posición del bit a verificar.
//
// Retorna:
//   - `true` si el bit está apagado; `false` si está encendido.
func (bm *BitMap) IsOff(pos uint8) (bool, error) {

	//101
	//011
	//000
	if isOutOfRange(pos) {
		return false, errors.New("posición no válida")
	}
	return bm.bits&(0b1<<pos) == 0b0, nil
}

// GetMap obtiene la representación interna del mapa de bits
//
// Uso:
//
//	bits := bm.GetMap() // Obtiene la representación interna del mapa de bits.
//
// Retorna:
//   - La representación interna del mapa de bits como un uint32.
func (bm *BitMap) GetMap() uint32 {
	//00000000_00000000_00000000_000000000
	return bm.bits
}

// CountOn devuelve la cantidad de bits encendidos en el mapa de bits.
//
// Uso:
//
//	count := bm.CountOn() // Obtiene la cantidad de bits encendidos.
//
// Retorna:
//   - La cantidad de bits encendidos como un entero.
func (bm *BitMap) CountOn() int {

	contador := 0
	for i := 0; i < 32; i++ {
		isOn, _ := bm.IsOn(uint8(i))
		if isOn {
			contador++
		}
	}
	return contador
}

// String devuelve una representación en cadena del mapa de bits.
//
// Uso:
//
//	fmt.Println(bm.String()) // Imprime la representación en binario del mapa.
//
// Retorna:
//   - Una cadena de 32 caracteres con los bits del mapa, desde el más significativo al menos significativo.
func (bm *BitMap) String() string {

	cadena := ""
	for i := 0; i < 32; i++ {
		isOn, _ := bm.IsOn(uint8(i))
		if isOn {
			cadena += "1"
		} else {
			cadena += "0"
		}
	}
	return cadena
}

// isOutOfRange verifica si la posición está fuera de rango.
//
// Uso:
//
//	if isOutOfRange(32) {
//		fmt.Println("Posición no válida")
//	}
//
// Retorna:
//   - `true` si la posición está fuera de rango; `false` en caso contrario.
func isOutOfRange(pos uint8) bool {
	return pos >= BitmapSize
}
