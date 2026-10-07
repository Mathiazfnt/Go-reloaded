package fonction

func CreerMatrice(phrase string) [][]string {
	var matrice [][]string
	var mot []string
	runes := []rune(phrase)

	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			if len(mot) > 0 {
				matrice = append(matrice, mot)
				mot = nil
			}

			parenthese, fin := Parenthese(runes, i)
			matrice = append(matrice, parenthese)
			i = fin
		} else if runes[i] == ' ' {
			if len(mot) > 0 {
				matrice = append(matrice, mot)
				mot = nil
			}
		} else {
			mot = append(mot, string(runes[i]))
		}
	}

	if len(mot) > 0 {
		matrice = append(matrice, mot)
	}

	return matrice
}

func Parenthese(phrase []rune, debut int) ([]string, int) {
	var tab []string

	for i := debut; i < len(phrase); i++ {
		if phrase[i] != ' ' && phrase[i] != ',' {
			tab = append(tab, string(phrase[i]))
		}

		if phrase[i] == ')' {
			return tab, i
		}
	}

	return tab, len(phrase)
}

func RefaireTexte_mat(matrice [][]string) string {
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

func RefaireTexte_tab(tab []string) string {
	texte := ""

	for _, lettre := range tab {
		texte += lettre
	}

	return texte
}

func TransformerMot(texte string) []string {
	var mot []string

	for _, lettre := range texte {
		mot = append(mot, string(lettre))
	}

	return mot
}