package lista

type Lista[T any] interface {
	VerPrimero() T
	VerUltimo() T
	InsertarPrimero(T)
	InsertarUltimo(T)
	EstaVacia() bool
	BorrarPrimero() T
	Largo() int
	Iterador() Iterador[T]
}
