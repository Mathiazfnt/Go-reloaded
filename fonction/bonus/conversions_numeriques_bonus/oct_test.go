package conversions_numeriques_bonus

import "testing"

func TestOct(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"Conversion octale simple",
			[]string{"17", "(oct)"},
			[]string{"15"},
		},
		{
			"Conversion de zéro",
			[]string{"0", "(oct)"},
			[]string{"0"},
		},
		{
			"Conversion de dix en octal",
			[]string{"10", "(oct)"},
			[]string{"8"},
		},
		{
			"Deux nombres en octal",
			[]string{"10", "17", "(oct, 2)"},
			[]string{"8", "15"},
		},
		{
			"Sans marqueur",
			[]string{"10", "17"},
			[]string{"10", "17"},
		},
		{
			"Nombre supérieur aux mots disponibles",
			[]string{"17", "(oct, 5)"},
			[]string{"15"},
		},
		{
			"Marqueur sans nombre avant",
			[]string{"(oct)"},
			[]string{},
		},
		{
			"Plusieurs marqueurs",
			[]string{"10", "(oct)", "17", "(oct)"},
			[]string{"8", "15"},
		},
		{
			"Texte avant le marqueur",
			[]string{"bonjour", "(oct)"},
			[]string{"bonjour"},
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Oct(test.entree)

			if len(resultat) != len(test.attendu) {
				t.Errorf("%s : longueur obtenue = %d, longueur attendue = %d",
					test.nom, len(resultat), len(test.attendu))
				return
			}

			for i := 0; i < len(resultat); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("%s : resultat[%d] = %q, attendu %q",
						test.nom, i, resultat[i], test.attendu[i])
				}
			}
		})
	}
}