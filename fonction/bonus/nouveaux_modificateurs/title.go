package nouveaux_modificateurs

import( fc "Go-reloaded/fonction/casse_mots"
			"strings"
			"strconv"
			
)

func MajPremiere(mot string) string {
	if mot == "" {
		return mot
	}
	return strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
}

func Title(tab []string) []string{
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(title") {

			nombre := 1

			if tab[mot] != "(title)" {
				texte := tab[mot][8 : len(tab[mot])-1]
				nombre, _ = strconv.Atoi(texte)
			}

			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			for i := 0; i < fc.Depassement(tab, mot, nombre); i++ {
				if mot < 0 {
					break
				}

				tab[mot] = MajPremiere(tab[mot])
				mot--
			}

			mot += nombre
		}
	}

	return tab
}