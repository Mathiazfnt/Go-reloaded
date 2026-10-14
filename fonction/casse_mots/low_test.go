package casse_mots

import "testing"

func TestLow(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"Un mot en minuscule",
			[]string{"BONJOUR", "(low)"},
			[]string{"bonjour"},
		},
		{
			"Mot déjà en minuscule",
			[]string{"bonjour", "(low)"},
			[]string{"bonjour"},
		},
		{
			"Mot avec majuscules et minuscules",
			[]string{"BoNjOuR", "(low)"},
			[]string{"bonjour"},
		},
		{
			"Deux mots en minuscule",
			[]string{"BONJOUR", "SALUT", "(low, 2)"},
			[]string{"bonjour", "salut"},
		},
		{
			"Sans marqueur",
			[]string{"BONJOUR", "SALUT"},
			[]string{"BONJOUR", "SALUT"},
		},
		{
			"Nombre supérieur aux mots disponibles",
			[]string{"BONJOUR", "(low, 5)"},
			[]string{"bonjour"},
		},
	}

	for _, test := range tests {
		resultat := Low(test.entree)

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
