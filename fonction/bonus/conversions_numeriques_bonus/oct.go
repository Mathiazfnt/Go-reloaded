package fonction

import "strconv"

func Oct(tab []string) []string { // Nombre octal → décimal
	for mot := 0; mot < len(tab); mot++ {
		if tab[mot] == "(oct)" {
			// Supprime (oct)
			tab = append(tab[:mot], tab[mot+1:]...)

			// Vérifie qu'il existe un mot avant
			if mot > 0 {
				mot--

				// octal -> string(décimal)
				nombre, err := strconv.ParseInt(tab[mot], 8, 64)
				if err == nil {
					// décimal -> string
					tab[mot] = strconv.FormatInt(nombre, 10)
				}
			}
		}
	}

	return tab
}
