package ponctuation

import "testing"

func TestApost_Guill(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"apostrophes simples",
			[]string{"mot", "'", "autre", "'"},
			[]string{"mot", "'autre'"},
		},
		{
			"guillemets doubles",
			[]string{"mot", "\"", "autre", "\""},
			[]string{"mot", "\"autre\""},
		},
		{
			"apostrophe au milieu",
			[]string{"l", "'", "homme"},
			[]string{"l'homme"},
		},
		{
			"aucun signe",
			[]string{"bonjour", "tout", "le", "monde"},
			[]string{"bonjour", "tout", "le", "monde"},
		},
		{
			"plusieurs groupes",
			[]string{"'", "bonjour", "'", "et", "'", "salut", "'"},
			[]string{"'bonjour'", "et", "'salut'"},
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Apost_Guill(append([]string{}, test.entree...))

			if len(resultat) != len(test.attendu) {
				t.Fatalf("Apost_Guill(%v) = %v, attendu %v",
					test.entree, resultat, test.attendu)
			}

			for i := 0; i < len(test.attendu); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("Apost_Guill(%v) = %v, attendu %v",
						test.entree, resultat, test.attendu)
					break
				}
			}
		})
	}
}
