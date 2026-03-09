package lista

type Iterador[T any] interface {
	HasNext() bool
	Next()
	Actual() T

	//Inserta un nuevo elemento antes  del iterador, mueve el iterador en el elemento insertado.
	Insertar(T)
	//Borrar el elemento donde esta el iterador, el iterador avanza al siguiente.
	Borrar() T
}
