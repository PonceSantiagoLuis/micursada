package arboles

type ArbolBinarioDeBusqueda[K comparable, V any] struct {
	raiz       *Nodo[K, V]
	comparador func(K, K) int
	cantidad   int
}

func CrearABB[K comparable, V any](comparador func(K, K) int) *ArbolBinarioDeBusqueda[K, V] {
	abb := new(ArbolBinarioDeBusqueda[K, V])
	abb.comparador = comparador
	return abb
}

func (abb *ArbolBinarioDeBusqueda[K, V]) Insertar(key K, value V) {

	ptrToNode := abb.raiz.buscar(key, abb.comparador, &abb.raiz) //busco la posicion en donde esta o deberia estar el nodo

	if (*ptrToNode) != nil {
		(*ptrToNode).value = value
		return
	}
	(*ptrToNode) = crearNodo(key, value, nil, nil)
	abb.cantidad++
}

func (abb *ArbolBinarioDeBusqueda[K, V]) Pertenece(key K) bool {
	ptrToNode := abb.raiz.buscar(key, abb.comparador, &abb.raiz)
	return (*ptrToNode) != nil
}

func (abb *ArbolBinarioDeBusqueda[K, V]) Obtener(key K) V {
	ptrToNode := abb.raiz.buscar(key, abb.comparador, &abb.raiz)

	if (*ptrToNode) == nil {
		panic("El elemento no pertenece al arbol")
	}
	return (*ptrToNode).value
}

func (abb *ArbolBinarioDeBusqueda[K, V]) Borrar(key K) V {
	ptrToNode := abb.raiz.buscar(key, abb.comparador, &abb.raiz)
	if (*ptrToNode) == nil {
		panic("El elemento no pertenece al arbol")
	}

	value := (*ptrToNode).value
	abb.cantidad--
	//ahora distinguimos por casos, caso se quiere borrar una hoja
	if (*ptrToNode).hijoIzq == nil && (*ptrToNode).hijoDer == nil {
		*ptrToNode = nil
		return value
	}

	// si tiene un solo hijo, apunto su padre al hijo
	if (*ptrToNode).hijoDer == nil && (*ptrToNode).hijoIzq != nil {
		(*ptrToNode) = (*ptrToNode).hijoIzq
		return value
	}
	if (*ptrToNode).hijoDer != nil && (*ptrToNode).hijoIzq == nil {
		(*ptrToNode) = (*ptrToNode).hijoDer
		return value
	}

	//caso 2 hijo: el algo es toma el maximo de el subarbol izq, o el minimo del subarbol derecho
	//remplaza su valor en el nodo que queres borrar, y borra el max/min

	maximo := (*ptrToNode).hijoIzq.maximo(&(*ptrToNode).hijoIzq)
	maxValue := (*maximo).value
	maxKey := (*maximo).key
	(*ptrToNode).key = maxKey
	(*ptrToNode).value = maxValue

	//ahora tengo que borrar el maximo

	(*maximo) = (*maximo).hijoIzq
	return value

}
