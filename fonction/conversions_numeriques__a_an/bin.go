package fonction

import "strconv"

func Bin(tab []string) []string { // Binaire → décimal
	for mot := 0; mot < len(tab); mot++ {
		if len(tab[mot]) >= 5 {
			if tab[mot][1] == 'b' &&
				tab[mot][2] == 'i' &&
				tab[mot][3] == 'n' {

				// Supprime (bin)
				tab = append(tab[:mot], tab[mot+1:]...)

				// On modifie le mot avant
				mot--

				// bin -> décimal
				nombre, _ := strconv.ParseInt(tab[mot], 2, 64)

				// Décimal
				tab[mot] = strconv.FormatInt(nombre, 10)
			}
		}
	}

	return tab
}
			