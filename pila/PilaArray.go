package pila

const INIT_SIZE int = 100
const FACTOR_REDIMENSION int = 2

/*
	0 1 2 3 4 5

|A|B|C| | | |  size = 3

| | | | | | |
*/
type PilaArray[T any] struct {
	innerVector []T
	size        int
}

func CrearPila[T any]() Pila[T] {
	pila := new(PilaArray[T])
	pila.innerVector = make([]T, INIT_SIZE)
	return pila
}

// Apilar implements [Pila].
func (p *PilaArray[T]) Apilar(dato T) {

	if p.size == cap(p.innerVector) { //crecimiento del inner vector
		p.redimensionar()
	}

	p.innerVector[p.size] = dato // caso ideal, el size da el index del proximo elemento a apilar
	p.size++
}

// Desapilar implements [Pila].
func (p *PilaArray[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila está vacia")
	}

	//redimension por eficiencia de memoria
	if p.size < cap(p.innerVector)/4 && cap(p.innerVector) > INIT_SIZE {
		p.redimensionar()
	}

	p.size--
	return p.innerVector[p.size]
}

func (p *PilaArray[T]) redimensionar() {
	aux := make([]T, p.size*FACTOR_REDIMENSION)
	copy(aux, p.innerVector)
	p.innerVector = aux
}

// EstaVacia implements [Pila].
func (p *PilaArray[T]) EstaVacia() bool {
	return p.size == 0
}

// VerTope implements [Pila].
func (p *PilaArray[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	return p.innerVector[p.size-1]
}
