package fonction

import "strconv"

func Hex(matrice [][]string) [][]string { // Hexadécimal → décimal
	for mot := 0; mot < len(matrice); mot++ {
		if len(matrice[mot]) >= 4 {
			if  matrice[mot][1] == "h" &&
				matrice[mot][2] == "e" &&
				matrice[mot][3] == "x" {

				// Supprime la ligne hex
				matrice = append(matrice[:mot], matrice[mot+1:]...)

				// On va modifier le mot avant
				mot--

				// On récupère tout le mot hexadécimal
				hex := ""
				for _, chiffre := range matrice[mot] {
					hex += chiffre
				}

				// Hexadécimal -> décimal
				nombre, _ := strconv.ParseInt(hex, 16, 64)

				// Décimal -> texte
				decimal := strconv.FormatInt(nombre, 10)

				// On remet chaque chiffre dans le tableau
				nouveau := []string{}
				for _, chiffre := range decimal {
					nouveau = append(nouveau, string(chiffre))
				}

				// On remplace l'ancien mot
				matrice[mot] = nouveau
			}
		}
	}

	return matrice
}