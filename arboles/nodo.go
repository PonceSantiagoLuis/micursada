package arboles

type Nodo[K comparable, V any] struct {
	key   K
	value V

	hijoIzq *Nodo[K, V]
	hijoDer *Nodo[K, V]
}

func crearNodo[K comparable, V any](key K, value V, izq, der *Nodo[K, V]) *Nodo[K, V] {
	nodo := new(Nodo[K, V])
	nodo.key = key
	nodo.value = value
	nodo.hijoIzq = izq
	nodo.hijoDer = der
	return nodo
}

// el metodo buscar recibe la key a buscar, la funcion de comparacion y la dir del ptr a
// nodo que llama el metodo

func (n *Nodo[K, V]) buscar(key K, comparator func(K, K) int, self **Nodo[K, V]) **Nodo[K, V] {

	if n == nil || comparator(n.key, key) == 0 {
		return self
	}

	if comparator(n.key, key) > 0 {
		return n.hijoIzq.buscar(key, comparator, &n.hijoIzq)
	}
	return n.hijoDer.buscar(key, comparator, &n.hijoDer)

}

func (n *Nodo[K, V]) maximo(self **Nodo[K, V]) **Nodo[K, V] {

	if n == nil {
		return nil
	}
	if n.hijoDer == nil {
		return self
	}
	return n.hijoDer.maximo(&n.hijoDer)
}

func (nodo *Nodo[K, V]) altura() int {
	if nodo == nil {
		return 0
	}

	hizq := nodo.hijoIzq.altura()
	hder := nodo.hijoDer.altura()
	return maximo(hizq, hder)
}

func maximo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
