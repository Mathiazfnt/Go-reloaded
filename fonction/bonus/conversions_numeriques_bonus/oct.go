package conversions_numeriques_bonus

import (
	"strconv"
	"strings"
)

func Oct(tab []string) []string {
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(oct") {
			nombre := 1

			if tab[mot] != "(oct)" {
				texte := tab[mot][6 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			limite := nombre
			if limite > mot+1 {
				limite = mot + 1
			}

			for i := 0; i < limite; i++ {
				if mot < 0 {
					break
				}

				valeur, err := strconv.ParseInt(tab[mot], 8, 64)
				if err == nil {
					tab[mot] = strconv.FormatInt(valeur, 10)
				}

				mot--
			}

			mot += nombre
		}
	}

	return tab
}