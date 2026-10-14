package ponctuation

import "testing"

func TestPonct_spl(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"virgule",
			[]string{"bonjour", ",", "monde"},
			[]string{"bonjour,", "", "monde"},
		},
		{
			"point",
			[]string{"bonjour", ".", "monde"},
			[]string{"bonjour.", "", "monde"},
		},
		{
			"point d'exclamation",
			[]string{"bonjour", "!", "monde"},
			[]string{"bonjour!", "", "monde"},
		},
		{
			"point d'interrogation",
			[]string{"bonjour", "?", "monde"},
			[]string{"bonjour?", "", "monde"},
		},
		{
			"deux-points",
			[]string{"attention", ":", "danger"},
			[]string{"attention:", "", "danger"},
		},
		{
			"point-virgule",
			[]string{"bonjour", ";", "monde"},
			[]string{"bonjour;", "", "monde"},
		},
		{
			"ponctuation collee au texte",
			[]string{"bonjour", ",monde"},
			[]string{"bonjour,", "monde"},
		},
		{
			"aucune ponctuation",
			[]string{"bonjour", "monde"},
			[]string{"bonjour", "monde"},
		},
		{
			"plusieurs ponctuations",
			[]string{"bonjour", ",", "monde", "!"},
			[]string{"bonjour,", "", "monde!", ""},
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Ponct_spl(append([]string{}, test.entree...))

			if len(resultat) != len(test.attendu) {
				t.Fatalf("Ponct_spl(%v) = %v, attendu %v",
					test.entree, resultat, test.attendu)
			}

			for i := 0; i < len(test.attendu); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("Ponct_spl(%v) = %v, attendu %v",
						test.entree, resultat, test.attendu)
					break
				}
			}
		})
	}
}
