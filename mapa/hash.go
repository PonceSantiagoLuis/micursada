package mapa

import (
	"fmt"
)

type Par[K comparable, V any] struct {
	value V
	key   K
}

type status_t int

// enums
const (
	STATUS_FREE status_t = iota
	STATUS_DELETED
	STATUS_OCUPPIED
)

const INNER_VECTOR_INIT_SIZE int = 100
const FNV_PRIME uint64 = 0x100000001B3
const FNV_OFFSET_BASIS uint64 = 0xCBF29CE484222325

const FACTOR_REDIM int = 2

type Slot[K comparable, V any] struct {
	par    *Par[K, V]
	status status_t
}

type Hash[K comparable, V any] struct {
	tabla           []Slot[K, V]
	hashingFunction func(key K, longTabla int) int
	cantidad        int
}

func crearPar[K comparable, V any](key K, value V) *Par[K, V] {
	par := new(Par[K, V])
	par.key = key
	par.value = value
	return par
}

func CrearMapa[K comparable, V any]() Map[K, V] {
	mapa := new(Hash[K, V])
	mapa.tabla = make([]Slot[K, V], INNER_VECTOR_INIT_SIZE)
	mapa.hashingFunction = fnv1Hash
	return mapa
}

func fnv1Hash[K comparable](key K, longTabla int) int {

	hash := FNV_OFFSET_BASIS

	for _, byte := range []byte(fmt.Sprintf("%v", key)) {
		hash *= FNV_PRIME
		hash = hash ^ uint64(byte)
	}

	return int(hash % uint64(longTabla))
}

func dispersion(pos int, size int) int {
	return (pos + 1) % size
}

// Set implements [Map].
func (h *Hash[K, V]) Set(key K, value V) {

	if float32(h.cantidad)/float32(cap(h.tabla)) > 0.7 { //redimensionar tabla
		tablaAntigua := h.tabla
		h.tabla = make([]Slot[K, V], cap(h.tabla)*FACTOR_REDIM)
		h.cantidad = 0
		for _, slot := range tablaAntigua {
			if slot.status == STATUS_OCUPPIED {
				h.Set(slot.par.key, slot.par.value)
			}

		}
	}

	pos := h.hashingFunction(key, len(h.tabla))
	for h.tabla[pos].status == STATUS_OCUPPIED && h.tabla[pos].par.key != key { //sale por libre, deleted u ocupado y misma key
		pos = dispersion(pos, len(h.tabla))
	}

	if h.tabla[pos].status == STATUS_OCUPPIED {
		h.tabla[pos].par.value = value
		return
	}

	h.tabla[pos].par = crearPar(key, value)
	h.tabla[pos].status = STATUS_OCUPPIED
	h.cantidad++

}

// te da la posicion donde esta o deberia estar el elemento
// puede ser:
// ocupada y misma key
// libre
func (h *Hash[K, V]) buscarPosicion(key K) int {
	pos := h.hashingFunction(key, len(h.tabla))

	for (h.tabla[pos].status == STATUS_OCUPPIED && h.tabla[pos].par.key != key) || h.tabla[pos].status == STATUS_DELETED {
		pos = dispersion(pos, len(h.tabla))
	}
	return pos
}

// Delete implements [Map].
func (h *Hash[K, V]) Delete(key K) V {
	pos := h.buscarPosicion(key)

	if !(h.tabla[pos].status == STATUS_OCUPPIED) {
		panic("No existe en el diccionario")
	}

	par := h.tabla[pos].par
	h.tabla[pos].par = nil
	h.tabla[pos].status = STATUS_DELETED
	h.cantidad--

	return par.value

}

// HasKey implements [Map].
func (h *Hash[K, V]) HasKey(key K) bool {
	pos := h.buscarPosicion(key)

	return h.tabla[pos].status == STATUS_OCUPPIED
}

func (h *Hash[K, V]) Get(key K) V {
	pos := h.buscarPosicion(key)

	if !(h.tabla[pos].status == STATUS_OCUPPIED) {
		panic("No existe en el diccionario")
	}

	return h.tabla[pos].par.value
}
