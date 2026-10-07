package fonction

import "strconv"

func IsUpper(sentence string) bool {
	result := false
	var sentence2 = []byte(sentence)

	for i := 0; i < len(sentence); i++ {
		result = false

		if sentence2[i] >= 65 && sentence2[i] <= 90 {
			result = true
		}

		if result == false {
			return false
		}
	}

	return result
}

func toupper(sentence string) string {
	var tempo rune
	var result string

	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 97 && sentence[i] <= 122 {
			tempo = rune(sentence[i])
			tempo -= 32
			result += string(tempo)
		} else {
			result += string(rune(sentence[i]))
		}
	}

	return result
}

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
				texte := RefaireTexte_tab(matrice[mot])
				texte = toupper(texte)

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