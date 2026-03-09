package cola

type Cola[T any] interface {
	Encolar(T)
	Desencolar() T
	VerPrimero() T
	EstaVacia() bool
}
