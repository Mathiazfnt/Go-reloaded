package nouveaux_modificateurs

import (
	"fmt"
	"strconv"
	"strings"

	fc "Go-reloaded/fonction/casse_mots"
)

// inverse les lettres du mot précédent.
// Les caractères accentués doivent être conservés intacts.

func Rev(tab []string) ([]string, error) {
	for mot := 0; mot < len(tab); mot++ {
		if strings.HasPrefix(tab[mot], "(rev") {

			nombre := 1

			if tab[mot] != "(rev)" {
				commande := tab[mot]
				if !strings.HasPrefix(commande, "(rev,") || !strings.HasSuffix(commande, ")") {
					return nil, fmt.Errorf("syntaxe invalide pour %q", commande)
				}

				texte := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(commande, "(rev,"), ")"))
				var err error
				nombre, err = strconv.Atoi(texte)
				if err != nil {
					return nil, fmt.Errorf("nombre invalide dans %q: %w", commande, err)
				}
			}

			tab = append(tab[:mot], tab[mot+1:]...)

			mot--

			limite := fc.Depassement(tab, mot+1, nombre)
			for i := 0; i < limite; i++ {
				if mot < 0 {
					break
				}

				lettres := []rune(tab[mot])
				resultat := ""
				for lettre := 0; lettre < len(lettres); lettre++ {
					resultat += string(lettres[len(lettres)-1-lettre])
				}
				tab[mot] = resultat

				mot--
			}

			mot += nombre
		}
	}

	return tab, nil
}
