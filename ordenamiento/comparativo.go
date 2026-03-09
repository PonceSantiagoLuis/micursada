package ordenamiento

func Merge[T any](tablaA, tablaB []T, cmp func(T, T) int) []T {
	var tablaC []T = make([]T, len(tablaA)+len(tablaB))
	var a, b, c = 0, 0, 0
	for ; a < len(tablaA) && b < len(tablaB); c++ {
		if cmp(tablaA[a], tablaB[b]) > 0 { //a > b
			tablaC[c] = tablaB[b]
			b++
		} else {
			tablaC[c] = tablaA[a]
			a++
		}
	}

	if a == len(tablaA) { //consumiste la tabla a -> copias tabla b
		for b < len(tablaB) {
			tablaC[c] = tablaB[b]
			b++
			c++
		}
	} else {
		for a < len(tablaA) {
			tablaC[c] = tablaA[a]
			a++
			c++
		}
	}
	return tablaC
}

// a b   a-b  si cmp(a,b)> 0 a> b

func MergeSort[T any](source []T, cmp func(T, T) int) []T {

	if len(source) == 1 {
		return source

	}

	izq := MergeSort(source[:len(source)/2], cmp)
	der := MergeSort(source[len(source)/2:], cmp)
	return Merge(izq, der, cmp)

}
