package conversions_numeriques_bonus

import "testing"

func TestToBin(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"Conversion classique",
			[]string{"10", "(to bin)"},
			[]string{"1010"},
		},
		{
			"Conversion vers zéro",
			[]string{"0", "(to bin)"},
			[]string{"0"},
		},
		{
			"Conversion d'un nombre impair",
			[]string{"15", "(to bin)"},
			[]string{"1111"},
		},
		{
			"Conversion d'un nombre négatif",
			[]string{"-10", "(to bin)"},
			[]string{"-1010"},
		},
		{
			"Deux nombres à convertir",
			[]string{"10", "5", "(to bin, 2)"},
			[]string{"1010", "101"},
		},
		{
			"Trois nombres à convertir",
			[]string{"10", "5", "3", "(to bin, 3)"},
			[]string{"1010", "101", "11"},
		},
		{
			"Sans marqueur",
			[]string{"10", "5"},
			[]string{"10", "5"},
		},
		{
			"Nombre supérieur aux mots disponibles",
			[]string{"10", "(to bin, 5)"},
			[]string{"1010"},
		},
		{
			"Plusieurs conversions",
			[]string{"10", "(to bin)", "8", "(to bin)"},
			[]string{"1010", "1000"},
		},
	}

	for _, test := range tests {
		resultat := To_bin(test.entree)

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