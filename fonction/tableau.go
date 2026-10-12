package fonction

import "strings"

func Creertab(phrase string) []string {
	var matrice []string
	mots := strings.Fields(phrase)

	for i := 0; i < len(mots); i++ {
		if strings.HasPrefix(mots[i], "(") && !strings.HasSuffix(mots[i], ")") {
			mot := mots[i]

			for i+1 < len(mots) {
				i++
				mot += " " + mots[i]

				if strings.HasSuffix(mots[i], ")") {
					break
				}
			}

			matrice = append(matrice, mot)
		} else {
			matrice = append(matrice, mots[i])
		}
	}

	return matrice
}

func RefaireTexte_tab(tab []string) string {
	return strings.Join(tab, " ")
}