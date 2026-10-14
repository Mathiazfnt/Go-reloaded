package ponctuation

func Apost_Guill(tab []string) []string {
	compteur := 0

	for mot := 0; mot < len(tab); mot++ {
		if tab[mot] == "'" || tab[mot] == "\"" {
			signe := tab[mot]
			compteur++

			if compteur%2 != 0 {
				if mot+1 < len(tab) {
					tab[mot+1] = signe + tab[mot+1]
				} else {
					tab = append(tab, signe)
				}
			} else {
				if mot > 0 {
					tab[mot-1] += signe
				}
			}

			tab = append(tab[:mot], tab[mot+1:]...)
			mot--
		}
	}

	if compteur%2 != 0 && len(tab) > 0 {
		tab[len(tab)-1] += "'"
	}

	return tab
}
