package fonction

func Ponct_spl(matrice [][]string) [][]string {
	ponct := []string{".", ",", "!", "?", ":", ";"}

	for mot := range matrice {
		if mot == 0 {  // evite de regarder le mot précédent quand on est sur le premier mot.
			continue
		}

		for cherche := range ponct {
			if matrice[mot][0] == ponct[cherche] {
				matrice[mot-1] = append(matrice[mot-1], matrice[mot][0])
				matrice[mot] = matrice[mot][1:]
			}
		}
	}

	return matrice
}