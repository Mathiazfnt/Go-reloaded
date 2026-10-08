package fonction

import ("strconv"
		ft "Go-reloaded/fonction"
		"strings"
)
func Low(matrice [][]string) [][]string {
	for mot := 0; mot < len(matrice); mot++ {
		if len(matrice[mot]) >= 4 &&
			matrice[mot][1] == "l" &&
			matrice[mot][2] == "o" &&
			matrice[mot][3] == "w" {

			// Si c'est (low), on ajoute 1
			if len(matrice[mot]) == 5 {
				matrice[mot] = []string{"(", "l", "o", "w", "1", ")"}
			}

			nombre, _ := strconv.Atoi(matrice[mot][4])

			// Supprime la ligne (low)
			matrice = append(matrice[:mot], matrice[mot+1:]...)

			// On commence par le mot juste avant
			mot--

			for i := 0; i < nombre; i++ {
				if mot < 0 {
					break
				}

				texte := ft.RefaireTexte_tab(matrice[mot])
				texte = strings.ToLower(texte)

				matrice[mot] = []string{}

				for _, lettre := range texte {
					matrice[mot] = append(matrice[mot], string(lettre))
				}

				mot--
			}

			mot += nombre
		}
	}

	return matrice
}
