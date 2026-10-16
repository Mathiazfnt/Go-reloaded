package nouveaux_modificateurs

import "testing"

func TestRev(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{
			"Inversion d'un mot",
			[]string{"bonjour", "(rev)"},
			[]string{"ruojnob"},
		},
		{
			"Inversion de deux mots",
			[]string{"bonjour", "salut", "(rev, 2)"},
			[]string{"ruojnob", "tulas"},
		},
		{
			"Conservation des accents",
			[]string{"été", "(rev)"},
			[]string{"été"},
		},
		{
			"Mot avec plusieurs accents",
			[]string{"éléphant", "(rev)"},
			[]string{"tnahpélé"},
		},
		{
			"Mot avec une majuscule",
			[]string{"Bonjour", "(rev)"},
			[]string{"ruojnoB"},
		},
		{
			"Sans marqueur",
			[]string{"bonjour", "salut"},
			[]string{"bonjour", "salut"},
		},
		{
			"Nombre supérieur aux mots disponibles",
			[]string{"bonjour", "(rev, 5)"},
			[]string{"ruojnob"},
		},
	}

	for _, test := range tests {
		resultat, err := Rev(test.entree)

		if err != nil {
			t.Errorf("%s : erreur inattendue : %v", test.nom, err)
			continue
		}

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

func TestRevErreur(t *testing.T) {
	tests := []struct {
		nom    string
		entree []string
	}{
		{
			"Syntaxe incorrecte",
			[]string{"bonjour", "(rev,abc)"},
		},
		{
			"Nombre non numérique",
			[]string{"bonjour", "(rev, x)"},
		},
		{
			"Marqueur sans parenthèse fermante",
			[]string{"bonjour", "(rev, 2"},
		},
	}

	for _, test := range tests {
		_, err := Rev(test.entree)

		if err == nil {
			t.Errorf("%s : une erreur était attendue", test.nom)
		}
	}
}