package fonction

func An(tab []string) []string {
	tab=a_an(tab)
	tab=amaj_an(tab)
	tab=an_a(tab)
	tab=an_maj_a(tab)

	return tab
}

func est_dedans(tab []string, recherche string) bool {
	for elt := range tab{
		if recherche == tab[elt]{
			return true
		}
	}
	return false
}

func a_an(tab []string) []string {
	liste := []string{"a", "e", "i", "o", "u", "h"}

	for mot := range tab {
		if mot < len(tab)-1 && tab[mot] == "a" {
			if est_dedans(liste, tab[mot+1]) {
				tab[mot] = "an"
			}
		}
	}
	return tab
}

func amaj_an(tab []string) []string {
	liste := []string{"a", "e", "i", "o", "u", "h"}

	for mot := range tab {
		if mot < len(tab)-1 && tab[mot] == "A" {
			if est_dedans(liste, tab[mot+1]) {
				tab[mot] = "An"
			}
		}
	}
	return tab
}

func an_a(tab []string) []string {
	liste := []string{"a", "e", "i", "o", "u", "h"}

	for mot := range tab {
		if mot < len(tab)-1 && tab[mot] == "an" {
			if !est_dedans(liste, tab[mot+1]) {
				tab[mot] = "a"
			}
		}
	}
	return tab
}

func an_maj_a(tab []string) []string {
	liste := []string{"a", "e", "i", "o", "u", "h"}

	for mot := range tab {
		if mot < len(tab)-1 && tab[mot] == "An" {
			if !est_dedans(liste, tab[mot+1]) {
				tab[mot] = "A"
			}
		}
	}
	return tab
}
