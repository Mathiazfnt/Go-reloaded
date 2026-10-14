package conversions_numeriques__a_an

import (
	d "Go-reloaded/fonction/casse_mots"
	"strconv"
	"strings"
)

func Hex(tab []string) []string { // Hexadécimal → décimal
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(hex") {
			nombre := 1

			if tab[mot] != "(hex)" {
				texte := tab[mot][6 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			// Supprime (hex)
			tab = append(tab[:mot], tab[mot+1:]...)
			mot--

			// Applique la règle aux mots précédents
			for i := 0; i < d.Depassement(tab, mot, nombre); i++ {
				if mot < 0 {
					break
				}

				// Hexadécimal -> string(décimal)
				tab[mot] = strings.ToUpper(tab[mot])
				resultat, err := strconv.ParseInt(tab[mot], 16, 64)
				if err == nil {
					tab[mot] = strconv.FormatInt(resultat, 10)
				}

				mot--
			}

			mot += nombre
		}
	}

	return tab
}
