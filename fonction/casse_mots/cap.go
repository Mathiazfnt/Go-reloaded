package fonction

import ("strconv"
		"strings"
	)

func tocap(mot string) string {
	var tempo rune
	var result string
	fait := true

	for i := 0; i < len(mot); i++ {
		if mot[i] >= 'a' && mot[i] <= 'z' && fait {
			fait = false
			tempo = rune(mot[i])
			tempo -= 32
			result += string(tempo)
		} else {
			result += string(rune(mot[i]))
		}
	}

	return result
}


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

			for i := 0; i < nombre; i++ {
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

