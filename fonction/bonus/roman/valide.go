package roman

import "strings"

// 1. Vérifie si chaque caractère est un symbole romain autorisé.
func Valide(texte string) bool {
	if texte == "" {
		return false
	}

	for _, lettre := range texte {
		if !strings.ContainsRune("IVXLCDM", lettre) {
			return false
		}
	}

	return true
}
