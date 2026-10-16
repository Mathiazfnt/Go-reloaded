package conversions_numeriques__a_an

import (
	d "Go-reloaded/fonction/casse_mots"
	"strconv"
	"strings"
)

func Bin(tab []string) []string { // Binaire → décimal
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(bin") {
			nombre := 1

			if tab[mot] != "(bin)" {
				texte := tab[mot][6 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			// Supprime (bin)
			tab = append(tab[:mot], tab[mot+1:]...)
			mot--

			// Applique la règle aux mots précédents
			limite := d.Depassement(tab, mot+1, nombre)
			for i := 0; i < limite; i++ {
				if mot < 0 {
					break
				}

				// Binaire -> string(décimal)
				resultat, err := strconv.ParseInt(tab[mot], 2, 64)
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
