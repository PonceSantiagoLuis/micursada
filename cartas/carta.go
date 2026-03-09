package cartas

import "fmt"

type palo_t int
type numero_t int

const (
	PALO_CORAZON palo_t = iota
	PALO_TREBOL
	PALO_DIAMANTE
	PALO_PICA
)

var simbolos []rune = []rune{
	'♥', '♣', '♦', '♠'}

const (
	NUMERO_1 numero_t = iota + 1
	NUMERO_2
	NUMERO_3
	NUMERO_4
	NUMERO_5
	NUMERO_6
	NUMERO_7
	NUMERO_8
	NUMERO_9
	NUMERO_10
	NUMERO_11
	NUMERO_12
)

type Carta struct {
	Palo   palo_t
	Numero numero_t
}

func CrearCarta(numero numero_t, palo palo_t) *Carta {
	carta := new(Carta)
	carta.Numero = numero
	carta.Palo = palo
	return carta
}

func (carta Carta) Dibujar() {
	fmt.Printf(" _______ \n")
	fmt.Printf("|%d	|\n", carta.Numero)
	fmt.Printf("|	|\n")
	fmt.Printf("|   %c   |\n", simbolos[carta.Palo])
	fmt.Printf("|	|\n")
	fmt.Printf("|_______|\n")

}

func (carta Carta) DibujarLineas() []string {
	return []string{
		" _______ ",
		fmt.Sprintf("|%-2d     |", carta.Numero),
		"|       |",
		fmt.Sprintf("|   %c   |", simbolos[carta.Palo]),
		"|       |",
		"|_______|",
	}
}

func DibujarCartas(cartas []*Carta) {
	if len(cartas) == 0 {
		return
	}

	// Obtener las líneas de cada carta
	lineasCartas := make([][]string, len(cartas))
	for i, carta := range cartas {
		lineasCartas[i] = carta.DibujarLineas()
	}

	// Todas tienen la misma cantidad de líneas
	for linea := 0; linea < len(lineasCartas[0]); linea++ {
		for carta := 0; carta < len(lineasCartas); carta++ {
			fmt.Print(lineasCartas[carta][linea] + " ")
		}
		fmt.Println()
	}
}
