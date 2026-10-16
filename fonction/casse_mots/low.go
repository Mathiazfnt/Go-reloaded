package casse_mots

import (
	"strconv"
	"strings"
)

func Low(tab []string) []string {
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(low") {

			nombre := 1

			if tab[mot] != "(low)" {
				texte := tab[mot][6 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			// Supprime (low)
			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			limite := Depassement(tab, mot+1, nombre)
			for i := 0; i < limite; i++ {
				if mot < 0 {
					break
				}

				tab[mot] = strings.ToLower(tab[mot])
				mot--
			}

			mot += nombre
		}
	}

	return tab
}
