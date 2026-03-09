package lista

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]

	len int
}

func CrearLista[T any]() Lista[T] {
	return new(LinkedList[T])
}

// EstaVacia implements [Lista].
func (l *LinkedList[T]) EstaVacia() bool {
	return l.len == 0
}

// InsertarPrimero implements [Lista].
func (l *LinkedList[T]) InsertarPrimero(dato T) {

	nodo := crearNodo(dato, nil)
	if l.EstaVacia() {
		l.head = nodo
		l.tail = nodo
	} else {
		nodo.next = l.head
		l.head = nodo
	}
	l.len++
}

// InsertarUltimo implements [Lista].
func (l *LinkedList[T]) InsertarUltimo(dato T) {
	nodo := crearNodo(dato, nil)
	if l.EstaVacia() {
		l.head = nodo
		l.tail = nodo
	} else {
		l.tail.next = nodo
		l.tail = nodo
	}
	l.len++

}

// VerPrimero implements [Lista].
func (l *LinkedList[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.head.data
}

// VerUltimo implements [Lista].
func (l *LinkedList[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.tail.data
}

func crearNodo[T any](dato T, next *Node[T]) *Node[T] {
	nodo := new(Node[T])
	nodo.data = dato
	nodo.next = next

	return nodo
}

func (l *LinkedList[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("Esta vacia")
	}
	dato := l.head.data
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	}
	l.len--
	return dato
}

func (l *LinkedList[T]) Largo() int {
	return l.len
}
