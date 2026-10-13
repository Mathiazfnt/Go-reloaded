package fonction


	valeurs := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symboles := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}


	roman := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
func valide(texte string) bool {


	caracteres := []rune(texte)
	somme := 0

	if texte == "" {
		return false
	}

	for i := 0; i < len(caracteres); i++ {
		valeur, existe := roman[caracteres[i]]
		if !existe {
			return false
		}

		if i+1 < len(caracteres) && valeur < roman[caracteres[i+1]] {
			somme -= valeur
		} else {
			somme += valeur
		}
	}


	reconstruit := ""
	reste := somme

	for i := 0; i < len(valeurs); i++ {
		for reste >= valeurs[i] {
			reconstruit += symboles[i]
			reste -= valeurs[i]
		}
	}

	return somme >= 1 && somme <= 3999 && reconstruit == texte
}