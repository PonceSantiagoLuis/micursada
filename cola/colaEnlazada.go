package cola

type Node[T any] struct {
	data T
	next *Node[T]
}

type NodeList[T any] struct {
	first *Node[T]
	last  *Node[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(NodeList[T])
}
func CrearNodo[T any](dato T) *Node[T] {
	ptrNode := new(Node[T])
	ptrNode.data = dato
	return ptrNode
}

// Desencolar implements [Cola].
func (n *NodeList[T]) Desencolar() T {
	if n.EstaVacia() {
		panic("La cola esta vacia")
	}

	ptrNodo := n.first
	n.first = n.first.next

	if n.first == nil {
		n.last = nil
	}

	return ptrNodo.data
}

// Encolar implements [Cola].
func (n *NodeList[T]) Encolar(dato T) {
	if n.first == nil {
		n.first = CrearNodo(dato)
		n.last = n.first
	} else {
		n.last.next = CrearNodo(dato)
		n.last = n.last.next
	}
}

// EstaVacia implements [Cola].
func (n *NodeList[T]) EstaVacia() bool {
	return n.first == nil
}

// VerPrimero implements [Cola].
func (n *NodeList[T]) VerPrimero() T {
	if n.EstaVacia() {
		panic("La cola esta vacia")
	}
	return n.first.data
}
