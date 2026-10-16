package nouveaux_modificateurs

import "testing"

func TestTitle(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"Un mot en minuscule",
			[]string{"bonjour", "(title)"},
			[]string{"Bonjour"},
		},
		{
			"Mot déjà en majuscule",
			[]string{"Bonjour", "(title)"},
			[]string{"Bonjour"},
		},
		{
			"Mot entièrement en majuscules",
			[]string{"BONJOUR", "(title)"},
			[]string{"Bonjour"},
		},
		{
			"Première lettre accentuée",
			[]string{"éCOLE", "(title)"},
			[]string{"École"},
		},
		{
			"Deux mots",
			[]string{"bonjour", "monde", "(title, 2)"},
			[]string{"Bonjour", "Monde"},
		},
		{
			"Mot outil a en première position",
			[]string{"a", "(title)"},
			[]string{"A"},
		},
		{
			"Mot outil the en première position",
			[]string{"the", "lord", "of", "the", "rings", "(title, 5)"},
			[]string{"The", "Lord", "of", "the", "Rings"},
		},
		{
			"Mot outil of en première position",
			[]string{"of", "mice", "and", "men", "(title, 4)"},
			[]string{"Of", "Mice", "and", "Men"},
		},
		{
			"Mot outil au milieu",
			[]string{"the", "lord", "of", "the", "rings", "(title, 3)"},
			[]string{"the", "lord", "Of", "the", "Rings"},
		},
		{
			"Sans marqueur",
			[]string{"bonjour", "monde"},
			[]string{"bonjour", "monde"},
		},
		{
			"Nombre supérieur aux mots disponibles",
			[]string{"bonjour", "(title, 5)"},
			[]string{"Bonjour"},
		},
		{
			"Nombre égal à zéro",
			[]string{"bonjour", "(title, 0)"},
			[]string{"bonjour"},
		},
		{
			"Plusieurs marqueurs",
			[]string{"bonjour", "(title)", "monde", "(title)"},
			[]string{"Bonjour", "Monde"},
		},
		{
			"Marqueur avec un seul mot outil",
			[]string{"the", "(title, 1)"},
			[]string{"The"},
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Title(test.entree)

			if len(resultat) != len(test.attendu) {
				t.Fatalf("longueur obtenue = %d, attendue = %d",
					len(resultat), len(test.attendu))
			}

			for i := 0; i < len(resultat); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("resultat[%d] = %q, attendu %q",
						i, resultat[i], test.attendu[i])
				}
			}
		})
	}
}
