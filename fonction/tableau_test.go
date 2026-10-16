package fonction

import "testing"

func TestCreertab(t *testing.T) {
	tests := []struct {
		nom     string
		entree  string
		attendu []string
	}{
		{
			"phrase simple",
			"Bonjour tout le monde",
			[]string{"Bonjour", "tout", "le", "monde"},
		},
		{
			"marqueur simple",
			"1E (hex)",
			[]string{"1E", "(hex)"},
		},
		{
			"marqueur avec nombre",
			"Hello (up, 2) world",
			[]string{"Hello", "(up, 2)", "world"},
		},
		{
			"marqueur separe sur plusieurs mots",
			"Hello (up, 2) world",
			[]string{"Hello", "(up, 2)", "world"},
		},
		{
			"parenthese sans fermeture",
			"Hello (up world",
			[]string{"Hello", "(up world"},
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Creertab(test.entree)

			if len(resultat) != len(test.attendu) {
				t.Fatalf("Creertab(%q) = %v, attendu %v",
					test.entree, resultat, test.attendu)
			}

			for i := 0; i < len(test.attendu); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("Creertab(%q) = %v, attendu %v",
						test.entree, resultat, test.attendu)
					break
				}
			}
		})
	}
}
