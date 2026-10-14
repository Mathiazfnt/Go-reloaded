package casse_mots
import (
	"strconv"
	"strings"
)

func Cap(tab []string) []string {
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(cap") {

			nombre := 1

			if tab[mot] != "(cap)" {
				texte := tab[mot][6 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			for i := 0; i < Depassement(tab, mot, nombre); i++ {
				if mot < 0 {
					break
				}

				mot2 := strings.ToLower(tab[mot])

				if len(mot2) > 0 {
					tab[mot] = strings.ToUpper(mot2[:1]) + mot2[1:]
				}

				mot--
			}

			mot += nombre
		}
	}

	return tab
}
