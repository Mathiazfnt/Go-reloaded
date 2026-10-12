package fonction

func Apostrophe(tab []string) []string {
	for mot := 1; mot < len(tab)-1; mot++ {
		if tab[mot] == "'" && mot+2 < len(tab) && tab[mot+2] == "'" {
			tab[mot] = "'" + tab[mot+1]
			tab = append(tab[:mot+1], tab[mot+2:]...)
			tab[mot] += "'"
			tab = append(tab[:mot+1], tab[mot+2:]...)
		}
	}
	return tab
}
