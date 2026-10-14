package casse_mots

import "testing"

func TestUpp(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"Un mot en majuscules",
			[]string{"bonjour", "(up)"},
			[]string{"BONJOUR"},
		},
		{
			"Mot déjà en majuscules",
			[]string{"BONJOUR", "(up)"},
			[]string{"BONJOUR"},
		},
		{
			"Mot avec majuscules et minuscules",
			[]string{"BoNjOuR", "(up)"},
			[]string{"BONJOUR"},
		},
		{
			"Deux mots en majuscules",
			[]string{"bonjour", "salut", "(up, 2)"},
			[]string{"BONJOUR", "SALUT"},
		},
		{
			"Sans marqueur",
			[]string{"bonjour", "salut"},
			[]string{"bonjour", "salut"},
		},
		{
			"Nombre supérieur aux mots disponibles",
			[]string{"bonjour", "(up, 5)"},
			[]string{"BONJOUR"},
		},
	}

	for _, test := range tests {
		resultat := Upp(test.entree)

		if len(resultat) != len(test.attendu) {
			t.Errorf("%s : longueur obtenue = %d, longueur attendue = %d",
				test.nom, len(resultat), len(test.attendu))
			continue
		}

		for i := 0; i < len(resultat); i++ {
			if resultat[i] != test.attendu[i] {
				t.Errorf("%s : resultat[%d] = %q, attendu %q",
					test.nom, i, resultat[i], test.attendu[i])
			}
		}
	}
}
