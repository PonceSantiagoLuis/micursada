package pila

type Pila[T any] interface {
	EstaVacia() bool
	VerTope() T
	Apilar(T)
	Desapilar() T
}
