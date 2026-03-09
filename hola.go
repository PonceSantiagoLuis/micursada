package main

import (
	"fmt"
)

/*
func main() {

	fmt.Println("Prueba merge generico")
	var tablaA []int = []int{1, 4, 7, 8, 9}
	var tablaB []int = []int{3, 5, 7, 10, 11}
	fmt.Println(tablaA)
	fmt.Println(tablaB)
	fmt.Println(ordenamiento.Merge[int](tablaA, tablaB, func(i1, i2 int) int {
		return i1 - i2
	})) //1,3,4,5,7,7,8,9,10,11

	fmt.Println("Prueba mergesort")
	var source []int = []int{0, -10, 2, 4, 11, 1, 8, 63, 7, 21, -5, 8}

	fmt.Println(ordenamiento.MergeSort(source, func(i1, i2 int) int {
		return i1 - i2
	}))

	fmt.Println("Test counting sort simple")

	corazon_ocho := cartas.CrearCarta(8, cartas.PALO_CORAZON)
	diamante_2 := cartas.CrearCarta(2, cartas.PALO_DIAMANTE)
	pica_ocho := cartas.CrearCarta(8, cartas.PALO_PICA)
	trebol_7 := cartas.CrearCarta(7, cartas.PALO_TREBOL)
	corazon_3 := cartas.CrearCarta(3, cartas.PALO_CORAZON)
	trebol_8 := cartas.CrearCarta(8, cartas.PALO_TREBOL)
	corazon_1 := cartas.CrearCarta(1, cartas.PALO_CORAZON)
	pica_siete := cartas.CrearCarta(7, cartas.PALO_PICA)
	trebon_siete := cartas.CrearCarta(7, cartas.PALO_TREBOL)
	diamante_3 := cartas.CrearCarta(3, cartas.PALO_DIAMANTE)
	diamante_1 := cartas.CrearCarta(1, cartas.PALO_DIAMANTE)

	mazo := []*cartas.Carta{corazon_ocho, diamante_2, pica_ocho,
		trebol_7, corazon_3, corazon_1, diamante_1, diamante_3,
		trebon_siete, trebol_8, pica_siete}

	cartas.DibujarCartas(mazo)
	fmt.Println("counting")
	cartas.DibujarCartas(ordenamiento.CountingSort(mazo))

	fmt.Println("Prueba de iterador de lista")

	listaNumeros := lista.CrearLista[string]()
	listaNumeros.InsertarPrimero("hola")
	listaNumeros.InsertarUltimo("mundo")
	listaNumeros.InsertarUltimo("Feliz")

	for it := listaNumeros.Iterador(); it.HasNext(); it.Next() {
		fmt.Println(it.Actual())
	}
}
*/

/*
// prueba de mapa implementado el hash
func main() {
	fmt.Println("Prueba de hash")

	cumpleanos := mapa.CrearMapa[string]()

	cumpleanos.Set("Barbi", "09-04-1994")
	cumpleanos.Set("santi", "29-09-1996")
	cumpleanos.Set("pepo", "21-01-1999")
	cumpleanos.Set("lichi", "28-04-2007")
	cumpleanos.Set("mami", "28-01-1975")
	cumpleanos.Set("pesspo", "21-01-1999")
	cumpleanos.Set("lisschi", "28-04-2007")
	cumpleanos.Set("mamssi", "28-01-1975")
	cumpleanos.Set("ssss", "21-01-1999")
	cumpleanos.Set("licssshi", "28-04-2007")

	cumpleanos.Set("pes333spo", "21-01-1999")

	cumpleanos.Mostrar()
	fmt.Println(cumpleanos.Get("mami"))

	var a *map[string]string = new(map[string]string)
	a

}
*/

func main() {
	fmt.Println("Prueba abb ")

}

func compararRunas(a, b rune) int {
	return int(a) - int(b)
}

func imprimirRuna(c rune) {
	fmt.Printf("%c ", c)
}
