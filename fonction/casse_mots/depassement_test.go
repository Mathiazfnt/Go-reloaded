package casse_mots

import "testing"

////////////////////////////////////////////////////////////////////////// a refaire 

func TestDepassement(t *testing.T) {
	tests := []struct {
		nom    string
		mot    int
		nombre int
		attendu int
	}{
		{"Nombre normal", 5, 2, 2},
		{"Nombre égal à mot", 5, 5, 5},
		{"Nombre supérieur à mot", 3, 8, 3},
		{"Nombre négatif", 5, -2, 2},
		{"Nombre négatif supérieur à mot", 3, -8, 3},
		{"Nombre égal à zéro", 5, 0, 0},
		{"Mot égal à zéro", 0, 3, 0},
		{"Les deux à zéro", 0, 0, 0},
		{"Mot négatif", -2, 5, -2},
		{"Nombre négatif et mot négatif", -3, -5, -3},
	}

	for _, test := range tests {
		resultat := Depassement([]string{}, test.mot, test.nombre)

		if resultat != test.attendu {
			t.Errorf("%s : Depassement(..., %d, %d) = %d, attendu %d",
				test.nom, test.mot, test.nombre, resultat, test.attendu)
		}
	}
}
