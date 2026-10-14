package roman

import "strconv"

func Roman(tab []string) []string { // Chiffre romain valide → décimal
	for mot := 0; mot < len(tab); mot++ {
		if tab[mot] == "(roman)" {
			if mot > 0 && valide(tab[mot-1]) {
				tab[mot-1] = roman(tab[mot-1])

				// Supprime (roman)
				tab = append(tab[:mot], tab[mot+1:]...)
				mot--
			}
		}
	}

	return tab
}

func roman(str string) string { // Chiffre romain valide → décimal
	if !valide(str) {
		return str
	}

	roman := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50,
		"C": 100, "D": 500, "M": 1000,
	}

	somme := 0

	for elt := 0; elt < len(str); elt++ {
		if elt+1 < len(str) && roman[string(str[elt])] < roman[string(str[elt+1])] {
			somme -= roman[string(str[elt])]
		} else {
			somme += roman[string(str[elt])]
		}
	}

	return strconv.Itoa(somme)
}
