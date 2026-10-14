package ponctuation

import "testing"

func TestTrois_point_Excla_inter(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"deux points",
			[]string{"Je", "pense", "..", "Tu", "viens"},
			[]string{"Je", "pense..", "Tu", "viens"},
		},
		{
			"interrogation exclamation",
			[]string{"Quoi", "!?"},
			[]string{"Quoi!?"},
		},
		{
			"aucun signe a modifier",
			[]string{"Bonjour", "tout", "le", "monde", "!"},
			[]string{"Bonjour", "tout", "le", "monde", "!"},
		},
		{
			"deux points au debut",
			[]string{"..", "Bonjour"},
			[]string{"..", "Bonjour"},
		},
		{
			"plusieurs signes",
			[]string{"Salut", "..", "Quoi", "!?"},
			[]string{"Salut..", "Quoi!?"},
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Trois_point_Excla_inter(
				append([]string{}, test.entree...),
			)

			if len(resultat) != len(test.attendu) {
				t.Fatalf("Trois_point_Excla_inter(%v) = %v, attendu %v",
					test.entree, resultat, test.attendu)
			}

			for i := 0; i < len(test.attendu); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("Trois_point_Excla_inter(%v) = %v, attendu %v",
						test.entree, resultat, test.attendu)
					break
				}
			}
		})
	}
}
