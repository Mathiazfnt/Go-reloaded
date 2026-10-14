package casse_mots

import (
	"strconv"
	"strings"
)

func Upp(tab []string) []string {
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(up") {

			nombre := 1

			if tab[mot] != "(up)" {
				texte := tab[mot][5 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			for i := 0; i < Depassement(tab, mot, nombre); i++ {
				if mot < 0 {
					break
				}

				tab[mot] = strings.ToUpper(tab[mot])
				mot--
			}

			mot += nombre
		}
	}

	return tab
}
