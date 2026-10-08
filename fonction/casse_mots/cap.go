package fonction

import ("strconv"
		ft "Go-reloaded/fonction"
)
func tocap(mot string) string {
	var tempo rune
	var result string
	fait := true

	for i := 0; i < len(mot); i++ {
		if mot[i] >= 'a' && mot[i] <= 'z' && fait {
			fait = false
			tempo = rune(mot[i])
			tempo -= 32
			result += string(tempo)
		} else {
			result += string(rune(mot[i]))
		}
	}

	return result
}


func Cap(matrice [][]string) [][]string {
	for mot := 0; mot < len(matrice); mot++ {
		if len(matrice[mot]) >= 4 &&
			matrice[mot][1] == "c" &&
			matrice[mot][2] == "a" &&
			matrice[mot][3] == "p" {

			// Si c'est (cap), on ajoute 1
			if len(matrice[mot]) == 5 {
				matrice[mot] = []string{"(", "c", "a", "p", "1", ")"}
			}

			nombre, _ := strconv.Atoi(matrice[mot][4])

			// Supprime la ligne (cap)
			matrice = append(matrice[:mot], matrice[mot+1:]...)

			// On commence par le mot juste avant
			mot--

			for i := 0; i < nombre; i++ {
				if mot < 0 {
					break
				}

				texte := ft.RefaireTexte_tab(matrice[mot])
				texte = tocap(texte)

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
