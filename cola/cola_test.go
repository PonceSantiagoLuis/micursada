package cola_test

import (
	"fmt"
	"tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con la creacion de la cola")
	cola := cola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestFIFO(t *testing.T) {
	t.Log("Hacemos pruebas para ver si se mantiene la invariante de la cola")
	cola := cola.CrearColaEnlazada[string]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	cola.Encolar("Hola")
	require.EqualValues(t, "Hola", cola.VerPrimero())
	cola.Encolar("que")
	cola.Encolar("tal!")
	require.EqualValues(t, "Hola", cola.Desencolar())
	require.EqualValues(t, "que", cola.Desencolar())
	require.EqualValues(t, "tal!", cola.Desencolar())
	TestColaVacia(t)
}

func TestVolumen(t *testing.T) {
	cola := cola.CrearColaEnlazada[int]()
	for i := 0; i < 1000; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 1000; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
	TestColaVacia(t)
}

const CAPACIDADINICIAL int = 5

func partir[T any](org cola.Cola[T], k int) []cola.Cola[T] {
	if org.EstaVacia() {
		return make([]cola.Cola[T], 0)
	}
	capacidadActual := CAPACIDADINICIAL
	colasFinales := make([]cola.Cola[T], CAPACIDADINICIAL)
	var colasUsadas int
	yaTermine := false
	for colasUsadas = 0; !yaTermine; colasUsadas++ {
		if colasUsadas == capacidadActual {
			capacidadActual *= 2
			arr := make([]cola.Cola[T], capacidadActual)
			copy(arr, colasFinales)
			colasFinales = arr
		}
		colaAuxiliar := cola.CrearColaVector[T]()
		for j := 0; j < k; j++ {
			if org.EstaVacia() {
				yaTermine = true
				break
			}
			colaAuxiliar.Encolar(org.Desencolar())
		}
		colasFinales[colasUsadas] = colaAuxiliar
	}
	return colasFinales[:colasUsadas]
}

const DIAOK int = 20

func todoOkElDia(n int) bool {
	if n <= DIAOK {
		return true
	} else {
		return false
	}
}

func buscarDiaFalla(diasTotales int) int {
	return _buscarDia(diasTotales, 0, diasTotales-1)
}

func _buscarDia(dia, inicio, fin int) int {
	if fin-inicio <= 0 {
		return dia
	}
	medio := (inicio + fin) / 2
	if todoOkElDia(medio) {
		return _buscarDia(medio, medio+1, fin)
	} else {
		return _buscarDia(medio, inicio, medio)
	}

}

// 1er Parcial
func Test1erParcialito(t *testing.T) {
	org := cola.CrearColaEnlazada[int]()
	elementosEncolar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(elementosEncolar); i++ {
		org.Encolar(elementosEncolar[i])
	}
	k := 2
	colasFinales := partir[int](org, k)
	for i := 0; i < len(colasFinales); i++ {
		colaActual := colasFinales[i]
		fmt.Printf("Cola %d:", i)
		for !colaActual.EstaVacia() {
			fmt.Printf("%d, ", colaActual.Desencolar())
		}
		fmt.Printf("\n")
	}
	require.True(t, true, true)
	diaFinal := buscarDiaFalla(60)
	require.Equal(t, 21, diaFinal)
}

/*------EJ 12--------
Implementar una función func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) , que elimine los elementos encolados para los cuales la función filtro devuelve false.
Aquellos elementos que no son eliminados deben permanecer en el mismo orden en el que estaban antes de invocar a la función.
No es necesario destruir los elementos que sí fueron eliminados.
Se pueden utilizar las estructuras auxiliares que se consideren necesarias y no está permitido acceder a la estructura interna de la cola (es una función).
¿Cuál es el orden del algoritmo implementado?
*/

/*

func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) {
	if cola.EstaVacia() {
		return
	}
	colaAuxiliar := CrearColaEnlazada[K]()
	for !cola.EstaVacia() {
		if !filtro(cola.VerPrimero()) {
			cola.Desencolar()
			continue
		}
		colaAuxiliar.Encolar(cola.Desencolar())
	}
	cola = colaAuxiliar
}

*/

//test dados por chatgpt

func TestColaNuevaEstaVacia(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	if !c.EstaVacia() {
		t.Fatalf("Una cola recién creada debería estar vacía")
	}
}

func TestEncolarYVerPrimero(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	c.Encolar(1)
	c.Encolar(2)

	if c.EstaVacia() {
		t.Fatalf("No debería estar vacía después de encolar")
	}

	if primero := c.VerPrimero(); primero != 1 {
		t.Fatalf("Se esperaba 1 como primero, se obtuvo %d", primero)
	}
}

func TestFIFOcgpt(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	c.Encolar(10)
	c.Encolar(20)
	c.Encolar(30)

	if c.Desencolar() != 10 {
		t.Fatalf("No respeta FIFO")
	}
	if c.Desencolar() != 20 {
		t.Fatalf("No respeta FIFO")
	}
	if c.Desencolar() != 30 {
		t.Fatalf("No respeta FIFO")
	}

	if !c.EstaVacia() {
		t.Fatalf("Debería estar vacía al final")
	}
}

func TestDesencolarEnVaciaPaniquea(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("Se esperaba panic al desencolar en cola vacía")
		}

		if r != "La cola esta vacia" {
			t.Fatalf("Mensaje de panic incorrecto. Se obtuvo: %v", r)
		}
	}()

	c.Desencolar()
}

func TestVerPrimeroEnVaciaPaniquea(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("Se esperaba panic al ver primero en cola vacía")
		}

		if r != "La cola esta vacia" {
			t.Fatalf("Mensaje de panic incorrecto. Se obtuvo: %v", r)
		}
	}()

	c.VerPrimero()
}

func TestRedimensionamiento(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	const N = 10000

	// Encolar muchos elementos
	for i := 0; i < N; i++ {
		c.Encolar(i)
	}

	// Desencolar la mitad
	for i := 0; i < N/2; i++ {
		if val := c.Desencolar(); val != i {
			t.Fatalf("Error en redimensionamiento. Esperado %d, obtenido %d", i, val)
		}
	}

	// Encolar más para forzar wrap-around
	for i := N; i < N+5000; i++ {
		c.Encolar(i)
	}

	// Vaciar completamente
	expected := N / 2
	for !c.EstaVacia() {
		val := c.Desencolar()
		if val != expected {
			t.Fatalf("Error en comportamiento FIFO tras resize. Esperado %d, obtenido %d", expected, val)
		}
		expected++
	}
}

func TestRedimensionamientoHaciaAbajo(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	const N = 10000

	// Forzar crecimiento
	for i := 0; i < N; i++ {
		c.Encolar(i)
	}

	// Desencolar casi todo (deja pocos elementos)
	for i := 0; i < N-5; i++ {
		val := c.Desencolar()
		if val != i {
			t.Fatalf("Error antes del shrink. Esperado %d, obtenido %d", i, val)
		}
	}

	// Ahora la estructura debería haberse reducido internamente
	// si implementaste shrink cuando ocupa poco espacio.

	// Volvemos a encolar muchos elementos
	for i := N; i < N+5000; i++ {
		c.Encolar(i)
	}

	// Verificamos orden completo
	expected := N - 5
	for !c.EstaVacia() {
		val := c.Desencolar()
		if val != expected {
			t.Fatalf("Error después del shrink. Esperado %d, obtenido %d", expected, val)
		}
		expected++
	}
}

func TestResizeConArregloPartido(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	// Paso 1: llenar lo suficiente para forzar crecimiento inicial
	for i := 0; i < 8; i++ {
		c.Encolar(i)
	}

	// Paso 2: desencolar algunos para mover inicio
	for i := 0; i < 5; i++ {
		val := c.Desencolar()
		if val != i {
			t.Fatalf("Error inicial, esperado %d, obtenido %d", i, val)
		}
	}

	// Ahora internamente el arreglo está partido si usás circular

	// Paso 3: encolar más para forzar wrap-around
	for i := 8; i < 16; i++ {
		c.Encolar(i)
	}

	// En este punto probablemente inicio > fin

	// Paso 4: forzar resize estando partido
	for i := 16; i < 40; i++ {
		c.Encolar(i)
	}

	// Paso 5: verificar orden completo
	expected := 5
	for !c.EstaVacia() {
		val := c.Desencolar()
		if val != expected {
			t.Fatalf("Fallo tras resize con arreglo partido. Esperado %d, obtenido %d", expected, val)
		}
		expected++
	}
}

type Persona struct {
	Nombre string
	Edad   int
}

func TestColaEnlazadaFIFO(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	c.Encolar(1)
	c.Encolar(2)
	c.Encolar(3)

	if c.Desencolar() != 1 {
		t.Fatal("Debe salir 1 primero")
	}
	if c.Desencolar() != 2 {
		t.Fatal("Debe salir 2 segundo")
	}
	if c.Desencolar() != 3 {
		t.Fatal("Debe salir 3 tercero")
	}
}

func TestColaEnlazadaVaciaLuegoDeDesencolarTodo(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	c.Encolar(10)
	c.Desencolar()

	if !c.EstaVacia() {
		t.Fatal("La cola debería estar vacía")
	}

	c.Encolar(20)
	if c.VerPrimero() != 20 {
		t.Fatal("Debería poder reutilizarse luego de vaciarse")
	}
}

func TestColaEnlazadaAlternando(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	for i := 0; i < 1000; i++ {
		c.Encolar(i)
		val := c.Desencolar()
		if val != i {
			t.Fatalf("Esperado %d, obtenido %d", i, val)
		}
	}

	if !c.EstaVacia() {
		t.Fatal("Debería estar vacía")
	}
}

func TestColaEnlazadaMuchosElementos(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	const N = 50000

	for i := 0; i < N; i++ {
		c.Encolar(i)
	}

	for i := 0; i < N; i++ {
		val := c.Desencolar()
		if val != i {
			t.Fatalf("Error en %d: esperado %d, obtenido %d", i, i, val)
		}
	}

	if !c.EstaVacia() {
		t.Fatal("Debe quedar vacía")
	}
}

func TestColaEnlazadaPanics(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Debería haber panic al desencolar vacía")
		}
	}()

	c.Desencolar()
}

func TestColaEnlazadaStruct(t *testing.T) {
	c := cola.CrearColaEnlazada[Persona]()

	p := Persona{"Juan", 30}
	c.Encolar(p)

	salida := c.Desencolar()

	if salida.Nombre != "Juan" || salida.Edad != 30 {
		t.Fatal("Error con struct")
	}
}

func TestColaEnlazadaReutilizacionExtrema(t *testing.T) {
	c := cola.CrearColaEnlazada[int]()

	for k := 0; k < 100; k++ {
		for i := 0; i < 1000; i++ {
			c.Encolar(i)
		}
		for i := 0; i < 1000; i++ {
			if c.Desencolar() != i {
				t.Fatal("FIFO roto")
			}
		}
		if !c.EstaVacia() {
			t.Fatal("Debe vaciarse completamente")
		}
	}
}
