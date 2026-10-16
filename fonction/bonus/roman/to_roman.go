package roman

import "strconv"

// 3. Convertit un entier en chiffre romain.
func to_roman(nombre int) string {
	if nombre < 1 || nombre > 3999 {
		return ""
	}

	valeurs := []int{
		1000, 900, 500, 400, 100, 90,
		50, 40, 10, 9, 5, 4, 1,
	}

	symboles := []string{
		"M", "CM", "D", "CD", "C", "XC",
		"L", "XL", "X", "IX", "V", "IV", "I",
	}

	resultat := ""

	for i := 0; i < len(valeurs); i++ {
		for nombre >= valeurs[i] {
			resultat += symboles[i]
			nombre -= valeurs[i]
		}
	}

	return resultat
}


func To_roman(tab []string) []string { // Décimal → chiffre romain
	for mot := 0; mot < len(tab); mot++ {
		if tab[mot] == "(to_roman)" {
			// Supprime (to_roman)
			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			// Vérifie qu'il existe un mot avant
			if mot >= 0 {
				// Convertit le décimal en chiffre romain
				nombre, err := strconv.Atoi(tab[mot])
				if err == nil {
					resultat := to_roman(nombre)

					if resultat != "" {
						tab[mot] = resultat
					}
				}
			}
		}
	}

	return tab
}
