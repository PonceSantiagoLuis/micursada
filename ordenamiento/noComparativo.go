package ordenamiento

import "tdas/cartas"

func CountingSort(mano []*cartas.Carta) []*cartas.Carta {
	ocurrencias := make([]int, 13)
	indices := make([]int, 13)

	final := make([]*cartas.Carta, len(mano))

	//contar ocurrencias

	for _, carta := range mano {
		ocurrencias[carta.Numero-1]++
	}

	//calculas indices
	indice := 0
	for i, cantidad := range ocurrencias {
		indices[i] = indice
		indice += cantidad

	}

	for _, carta := range mano {
		final[indices[carta.Numero-1]] = carta
		indices[carta.Numero-1]++
	}

	return final
}
