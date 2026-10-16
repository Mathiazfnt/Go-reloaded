package conversions_numeriques_bonus

import "testing"

func TestToHex(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"Conversion classique",
			[]string{"255", "(to hex)"},
			[]string{"ff"},
		},
		{
			"Conversion vers zéro",
			[]string{"0", "(to hex)"},
			[]string{"0"},
		},
		{
			"Conversion d'un nombre positif",
			[]string{"30", "(to hex)"},
			[]string{"1e"},
		},
		{
			"Conversion d'un nombre négatif",
			[]string{"-10", "(to hex)"},
			[]string{"-a"},
		},
		{
			"Deux nombres à convertir",
			[]string{"255", "16", "(to hex, 2)"},
			[]string{"ff", "10"},
		},
		{
			"Sans marqueur",
			[]string{"255", "16"},
			[]string{"255", "16"},
		},
		{
			"Nombre supérieur aux mots disponibles",
			[]string{"255", "(to hex, 5)"},
			[]string{"ff"},
		},
		{
			"Plusieurs conversions",
			[]string{"255", "(to hex)", "16", "(to hex)"},
			[]string{"ff", "10"},
		},
	}

	for _, test := range tests {
		resultat := To_hex(test.entree)

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