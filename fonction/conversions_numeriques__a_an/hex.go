package fonction

import "strconv"

func Hex(tab []string) []string { // Hexadécimal → décimal
	for mot := 0; mot < len(tab); mot++ {
		if len(tab[mot]) >= 5 {
			if tab[mot][1] == 'h' &&
				tab[mot][2] == 'e' &&
				tab[mot][3] == 'x' {

				// Supprime (hex)
				tab = append(tab[:mot], tab[mot+1:]...)

				// On modifie le mot avant
				mot--

				// Hexadécimal -> décimal
				nombre, _ := strconv.ParseInt(tab[mot], 16, 64)

				// Décimal
				tab[mot] = strconv.FormatInt(nombre, 10)
			}
		}
	}

	return tab
}