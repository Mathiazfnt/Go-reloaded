package fonction

func Apost_Guill(tab []string) []string {
	tab = apostrophe(tab)
	tab = guillemet(tab)

	return tab
}

func apostrophe(tab []string) []string {
	compteur := 0

	for mot := 0; mot < len(tab); mot++ {
		if tab[mot] == "'" {
			compteur++

			if compteur%2 != 0 {
				if mot+1 < len(tab) {
					tab[mot+1] = "'" + tab[mot+1]
				} else {
					tab = append(tab, "'")
				}
			} else {
				if mot > 0 {
					tab[mot-1] += "'"
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

func guillemet(tab []string) []string {
	compteur := 0

	for mot := 0; mot < len(tab); mot++ {
		if tab[mot] == "\"" {
			compteur++

			if compteur%2 != 0 {
				if mot+1 < len(tab) {
					tab[mot+1] = "\"" + tab[mot+1]
				} else {
					tab = append(tab, "\"")
				}
			} else {
				if mot > 0 {
					tab[mot-1] += "\""
				}
			}

			tab = append(tab[:mot], tab[mot+1:]...)
			mot--
		}
	}

	if compteur%2 != 0 && len(tab) > 0 {
		tab[len(tab)-1] += "\""
	}

	return tab
}