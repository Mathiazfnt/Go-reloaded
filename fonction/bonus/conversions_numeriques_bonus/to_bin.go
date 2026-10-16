package conversions_numeriques_bonus

import (
	"strconv"
	"strings"

	f "Go-reloaded/fonction/casse_mots"
)

func To_bin(tab []string) []string { // Décimal → binaire
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(to bin") {
			nombre := 1

			if tab[mot] != "(to bin)" {
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

				// Décimal -> string(binaire)
				resultat, err := strconv.ParseInt(tab[mot], 10, 64)
				if err == nil {
					tab[mot] = strconv.FormatInt(resultat, 2)
				}

				mot--
			}

			mot += nombre
		}
	}

	return tab
}
