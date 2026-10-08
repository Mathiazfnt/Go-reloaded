package fonction

import "strings"

func CreerMatrice(phrase string) [][]string {
	var matrice [][]string
	var mot string
	runes := []rune(phrase)

	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			if mot != "" {
				matrice = append(matrice, TransformerMot(mot))
				mot = ""
			}

			parenthese, fin := Parenthese(runes, i)
			matrice = append(matrice, parenthese)
			i = fin

		} else if runes[i] == ' ' {
			if mot != "" {
				matrice = append(matrice, TransformerMot(mot))
				mot = ""
			}

		} else {
			mot += string(runes[i])
		}
	}

	if mot != "" {
		matrice = append(matrice, TransformerMot(mot))
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

func TransformerMot(texte string) []string {
	var mot []string

	for _, lettre := range texte {
		mot = append(mot, string(lettre))
	}

	return mot
}

func RefaireTexte_mat(matrice [][]string) string {
	var mots []string

	for _, mot := range matrice {
		mots = append(mots, RefaireTexte_tab(mot))
	}

	return strings.Join(mots, " ")
}

func RefaireTexte_tab(tab []string) string {
	return strings.Join(tab, "")
}
