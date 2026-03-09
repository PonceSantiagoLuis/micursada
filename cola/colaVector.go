package cola

const INITIAL_INNER_VECTOR_SIZE int = 10
const REDIMENSION_FACTOR int = 2

type ColaVector[T any] struct {
	datos   []T
	primero int
	ultimo  int

	cantidad int
}

// Desencolar implements [Cola].
func (c *ColaVector[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	indexDesencolado := c.primero
	c.primero = c.avanzarIndex(c.primero)
	c.cantidad--
	return c.datos[indexDesencolado]

}

// Encolar implements [Cola].
func (c *ColaVector[T]) Encolar(dato T) {

	if c.cantidad == len(c.datos) { //si esta llená la cola redimensionas el arreglo inicial
		var datosAux []T = make([]T, len(c.datos)*REDIMENSION_FACTOR)

		for i, actual := 0, c.primero; i < c.cantidad; i++ {
			datosAux[i] = c.datos[actual]
			actual = c.avanzarIndex(actual)

		}
		c.datos = datosAux
		c.primero = 0
		c.ultimo = c.cantidad
	}
	c.datos[c.ultimo] = dato
	c.ultimo = c.avanzarIndex(c.ultimo)
	c.cantidad++
}

// EstaVacia implements [Cola].
func (c *ColaVector[T]) EstaVacia() bool {
	return c.cantidad == 0
}

// VerPrimero implements [Cola].
func (c *ColaVector[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	return c.datos[c.primero]
}

func CrearColaVector[T any]() Cola[T] {
	var cola *ColaVector[T] = new(ColaVector[T])
	cola.datos = make([]T, INITIAL_INNER_VECTOR_SIZE)
	return cola
}

func (c *ColaVector[T]) avanzarIndex(index int) int {
	return (index + 1) % len(c.datos)

}
