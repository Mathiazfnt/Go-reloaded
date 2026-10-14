package fonction

import "strings"

func Creertab(phrase string) []string {
	mots := strings.Fields(phrase)
	var matrice []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		for strings.HasPrefix(mot, "(") && !strings.HasSuffix(mot, ")") && i+1 < len(mots) {
			i++
			mot += " " + mots[i]
		}

		matrice = append(matrice, mot)
	}

	return matrice
}
