package casse_mots

import "testing"

func TestCap(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"majuscule simple",
			[]string{"bonjour", "(cap)"},
			[]string{"Bonjour"},
		},
		{
			"mot deja en majuscule",
			[]string{"Bonjour", "(cap)"},
			[]string{"Bonjour"},
		},
		{
			"mot en majuscules",
			[]string{"BONJOUR", "(cap)"},
			[]string{"Bonjour"},
		},
		{
			"deux mots",
			[]string{"bonjour", "tout", "(cap, 2)"},
			[]string{"Bonjour", "Tout"},
		},
		{
			"aucun marqueur",
			[]string{"bonjour", "monde"},
			[]string{"bonjour", "monde"},
		},
		{
			"nombre superieur aux mots disponibles",
			[]string{"bonjour", "(cap, 5)"},
			[]string{"Bonjour"},
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Cap(append([]string{}, test.entree...))

			if len(resultat) != len(test.attendu) {
				t.Fatalf("Cap(%v) = %v, attendu %v",
					test.entree, resultat, test.attendu)
			}

			for i := 0; i < len(test.attendu); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("Cap(%v) = %v, attendu %v",
						test.entree, resultat, test.attendu)
					break
				}
			}
		})
	}
}
