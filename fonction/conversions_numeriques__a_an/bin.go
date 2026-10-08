package fonction

import "strconv"

func Bin(matrice [][]string) [][]string { // Binaire → décimal
	for mot := 0; mot < len(matrice); mot++ {
		if len(matrice[mot]) >= 4 {
			if matrice[mot][1] == "b" &&
				matrice[mot][2] == "i" &&
				matrice[mot][3] == "n" {

				// Supprime la ligne bin
				matrice = append(matrice[:mot], matrice[mot+1:]...)
				
				// On va modifier le mot avant
				mot--

				// On récupère tout le mot bin
				bin := ""
				for _, chiffre := range matrice[mot] {
					bin += chiffre
				}

				// bin -> décimal
				nombre, _ := strconv.ParseInt(bin, 2, 64)

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