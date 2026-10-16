package conversions_numeriques_bonus

import (
	"strconv"
	"strings"

	f "Go-reloaded/fonction/casse_mots"
)

func To_hex(tab []string) []string { // Nombre décimal → hexadécimal, en majuscules
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(to hex") {
			nombre := 1

			if tab[mot] != "(to hex)" {
				texte := tab[mot][9 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			// Supprime (to bin)
			tab = append(tab[:mot], tab[mot+1:]...)
			mot--

			// Applique la règle aux mots précédents
			limite := f.Depassement(tab, mot+1, nombre)
			for i := 0; i < limite; i++ {
				if mot < 0 {
					break
				}

				// Nombre décimal → hexadécimal, en majuscules
				resultat, err := strconv.ParseInt(tab[mot], 10, 64)
				if err == nil {
					tab[mot] = strconv.FormatInt(resultat, 16)
				}

				mot--
			}

			mot += nombre
		}
	}

	return tab
}
