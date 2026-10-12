package fonction

func An(tab []string) []string {
	liste := []string{"a", "e", "i", "o", "u", "h"}
	liste2 := []string{"A", "E", "I", "O", "U", "H"}

	for mot := 0; mot < len(tab)-1; mot++ {
		if tab[mot] == "a" {
			if est_dedans(liste, tab[mot+1]) || est_dedans(liste2, tab[mot+1]) {
				tab[mot] = "an"
			}
		} else if tab[mot] == "A" {
			if est_dedans(liste, tab[mot+1]) || est_dedans(liste2, tab[mot+1]) {
				tab[mot] = "An"
			}
		} else if tab[mot] == "an" {
			if !est_dedans(liste, tab[mot+1]) && !est_dedans(liste2, tab[mot+1]) {
				tab[mot] = "a"
			}
		} else if tab[mot] == "An" {
			if !est_dedans(liste, tab[mot+1]) && !est_dedans(liste2, tab[mot+1]) {
				tab[mot] = "A"
			}
		}
	}

	return tab
}

func est_dedans(tab []string, recherche string) bool {
	for elt := range tab {
		if recherche == tab[elt] {
			return true
		}
	}
	return false
}
