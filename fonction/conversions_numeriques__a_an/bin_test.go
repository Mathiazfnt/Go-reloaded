package conversions_numeriques__a_an

import "testing"

func TestBin(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{"binaire simple", []string{"101", "(bin)"}, []string{"5"}},
		{"binaire zero", []string{"0", "(bin)"}, []string{"0"}},
		{"zero devant", []string{"0010", "(bin)"}, []string{"2"}},
		{"binaire invalide", []string{"102", "(bin)"}, []string{"102"}},
		{"sans marqueur", []string{"101", "hello"}, []string{"101", "hello"}},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := Bin(append([]string{}, test.entree...))

			if len(resultat) != len(test.attendu) {
				t.Fatalf("Bin(%v) = %v, attendu %v",
					test.entree, resultat, test.attendu)
			}

			for i := 0; i < len(test.attendu); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("Bin(%v) = %v, attendu %v",
						test.entree, resultat, test.attendu)
					break
				}
			}
		})
	}
}
