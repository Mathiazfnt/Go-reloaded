package nouveaux_modificateurs

import (
	fc "Go-reloaded/fonction/casse_mots"
	"strconv"
	"strings"
)

func MajPremiere(mot string) string {
	lettres := []rune(mot)
	if len(lettres) == 0 {
		return mot
	}
	return strings.ToUpper(string(lettres[0])) + strings.ToLower(string(lettres[1:]))
}

func Title(tab []string) []string {
	motsOutils := map[string]bool{
		"a": true, "an": true, "the": true, "of": true, "and": true,
		"or": true, "in": true, "on": true, "at": true, "to": true,
	}

	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(title") {

			nombre := 1

			if tab[mot] != "(title)" {
				texte := tab[mot][8 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			limite := fc.Depassement(tab, mot+1, nombre)
			for i := 0; i < limite; i++ {
				if mot < 0 {
					break
				}

				if motsOutils[strings.ToLower(tab[mot])] {
					tab[mot] = strings.ToLower(tab[mot])
				} else {
					tab[mot] = MajPremiere(tab[mot])
				}
				mot--
			}

			mot += nombre
		}
	}

	return tab
}
