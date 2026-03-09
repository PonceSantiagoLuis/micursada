package lista

type IteradorLista[T any] struct {
	next *Node[T]
	prev *Node[T]
}

func (it *IteradorLista[T]) HasNext() bool {
	return it.next != nil
}

func (it *IteradorLista[T]) Next() {
	it.prev = it.next
	if it.HasNext() {
		it.next = it.next.next
	}
}

func (it *IteradorLista[T]) Actual() T {
	return it.next.data
}

// Inserta un nuevo elemento antes  del iterador, mueve el iterador en el elemento insertado.
func (it *IteradorLista[T]) Insertar(dato T) {
	nodo := crearNodo(dato, it.next)
	if it.HasNext() { //estas parado en un nodo
		if it.prev == nil { //estas en el primero nodo
			it.prev = nodo
		} else {
			it.prev.next = nodo
		}
	}

}

// Borrar el elemento donde esta el iterador, el iterador avanza al siguiente.
func (it *IteradorLista[T]) Borrar() T {
	if it.HasNext() {
		actual := it.next
		if !it.HasNext() {
			panic("Iterador consumido")
		}

		if it.prev != nil {
			it.prev.next = it.next.next
			it.next = it.next.next
		} else { // un solo elemento
			it.next
		}
		return actual.data
	}

}
func (l *LinkedList[T]) Iterador() Iterador[T] {
	it := new(IteradorLista[T])
	it.next = l.head
	it.prev = nil
	return it
}
