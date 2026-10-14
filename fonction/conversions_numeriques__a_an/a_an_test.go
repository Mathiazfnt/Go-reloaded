
package conversions_numeriques__a_an

import "testing"

func TestAn(t *testing.T) {
	tests := []struct {
		nom     string
		entree  []string
		attendu []string
	}{
		{"a devant voyelle", []string{"a", "apple"}, []string{"an", "apple"}},
		{"A devant voyelle", []string{"A", "Apple"}, []string{"An", "Apple"}},
		{"a devant consonne", []string{"a", "banana"}, []string{"a", "banana"}},
		{"an devant consonne", []string{"an", "banana"}, []string{"a", "banana"}},
		{"An devant consonne", []string{"An", "Banana"}, []string{"A", "Banana"}},
		{"an devant voyelle", []string{"an", "apple"}, []string{"an", "apple"}},
		{"un seul mot", []string{"apple"}, []string{"apple"}},
		{"mot sans article", []string{"hello", "world"}, []string{"hello", "world"}},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := An(append([]string{}, test.entree...))

			if len(resultat) != len(test.attendu) {
				t.Fatalf("An(%v) = %v, attendu %v",
					test.entree, resultat, test.attendu)
			}

			for i := 0; i < len(test.attendu); i++ {
				if resultat[i] != test.attendu[i] {
					t.Errorf("An(%v) = %v, attendu %v",
						test.entree, resultat, test.attendu)
					break
				}
			}
		})
	}
}
