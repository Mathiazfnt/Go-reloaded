package fonction

func CrearinMatrice(sentence string) [][]string {
	var matrice [][]string
	var mot []string
	var tempo string

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == 32 || sentence[i] == 9 || sentence[i] == 10 {
			if tempo != "" {
				for j := 0; j < len(tempo); j++ {
					mot = append(mot, string(tempo[j]))
				}

				matrice = append(matrice, mot)
				mot = nil
				tempo = ""
			}
			continue
		}

		tempo += string(sentence[i])
	}

	if tempo != "" {
		for j := 0; j < len(tempo); j++ {
			mot = append(mot, string(tempo[j]))
		}

		matrice = append(matrice, mot)
	}

	return matrice
}


func RefaireTexte(matrice [][]string) string {
	texte := ""

	for i, mot := range matrice {
		for _, lettre := range mot {
			texte += lettre
		}

		if i < len(matrice)-1 {
			texte += " "
		}
	}

	return texte
}

