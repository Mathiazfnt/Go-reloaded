package nouveaux_modificateurs

import (
	"strconv"
	"strings"
)

var titleMotsOutils = map[string]bool{
	"a": true, "an": true, "the": true, "of": true, "and": true,
	"or": true, "in": true, "on": true, "at": true, "to": true,
}

func MajPremiere(mot string) string {
	lettres := []rune(mot)
	if len(lettres) == 0 {
		return mot
	}
	return strings.ToUpper(string(lettres[0])) + strings.ToLower(string(lettres[1:]))
}

func Title(tab []string) []string {
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(title") {
			nombre := 1
			if tab[mot] != "(title)" {
				if strings.HasSuffix(tab[mot], ")") && len(tab[mot]) > 8 {
					texte := strings.TrimSpace(tab[mot][8 : len(tab[mot])-1])
					if valeur, err := strconv.Atoi(texte); err == nil {
						nombre = valeur
					}
				}
			}

			if nombre < 0 {
				nombre = -nombre
			}
			debut := mot - nombre
			if debut < 0 {
				debut = 0
			}

			tab = append(tab[:mot], tab[mot+1:]...)

			for i := debut; i < mot; i++ {
				if i == debut || !titleMotsOutils[strings.ToLower(tab[i])] {
					tab[i] = MajPremiere(tab[i])
				} else {
					tab[i] = strings.ToLower(tab[i])
				}
			}

			mot--
		}
	}

	return tab
}
