package casse_mots

func Depassement(tab []string, mot int, nombre int) int {
	if nombre < 0 {
		nombre = -nombre
	}

	if nombre > mot {
		nombre = mot
	}

	return nombre
}
