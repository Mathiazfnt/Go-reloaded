package roman

import (
	"strconv"
)

// 4. Vérifie que le chiffre romain est correctement écrit,
// puis renvoie l'entier correspondant.

func Roman(tab []string) []string { // Chiffre romain → décimal
	for mot := 0; mot < len(tab); mot++ {
		if tab[mot] == "(roman)" {
			// Supprime (roman)
			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			// Vérifie qu'il existe un mot avant
			if mot >= 0 {
				// Vérifie que le chiffre romain est valide
				if Valide(tab[mot]) && tab[mot] == to_roman(roman(tab[mot])){
					nombre := roman(tab[mot])

					// décimal -> string
					tab[mot] = strconv.Itoa(nombre)
				}
			}
		}
	}

	return tab
}

// 2. Convertit un chiffre romain en entier.
func roman(texte string) int {
	if !Valide(texte) {
		return 0
	}

	valeurs := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	resultat := 0
	lettres := []rune(texte)

	for i := 0; i < len(lettres); i++ {
		if i+1 < len(lettres) && valeurs[lettres[i]] < valeurs[lettres[i+1]] {
			resultat -= valeurs[lettres[i]]
		} else {
			resultat += valeurs[lettres[i]]
		}
	}

	return resultat
}
