package fonction

func Ponct_spl(tab []string) []string {
	ponct := []string{".", ",", "!", "?", ":", ";"}

	for mot := 1; mot < len(tab); mot++ {
		for cherche := range ponct {
			if len(tab[mot]) > 0 && tab[mot][0] == ponct[cherche][0] {
				tab[mot-1] += string(tab[mot][0])
				tab[mot] = tab[mot][1:]
			}
		}
	}

	return tab
}
