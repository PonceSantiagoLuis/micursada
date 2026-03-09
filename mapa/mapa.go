package mapa

type Map[K comparable, V any] interface {
	Get(K) V
	Set(K, V)
	HasKey(K) bool
	Delete(K) V
}
