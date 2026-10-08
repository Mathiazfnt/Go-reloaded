package fonction

import ("strconv"
		ft "Go-reloaded/fonction"
		"strings"
)

func Upp(matrice [][]string) [][]string {
	for mot := 0; mot < len(matrice); mot++ {
		if len(matrice[mot]) >= 3 &&
			matrice[mot][1] == "u" &&
			matrice[mot][2] == "p" {

			// Si c'est (up), on ajoute 1
			if len(matrice[mot]) == 4 {
				matrice[mot] = []string{"(", "u", "p", "1", ")"}
			}

			nombre, _ := strconv.Atoi(matrice[mot][3])

			// Supprime la ligne (up)
			matrice = append(matrice[:mot], matrice[mot+1:]...)

			// On commence par le mot juste avant
			mot--

			for i := 0; i < nombre; i++ {
				texte := ft.RefaireTexte_tab(matrice[mot])
				texte = strings.ToUpper(texte)

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