
package roman

import "testing"

func TestValide(t *testing.T) {
	if !Valide("I") {
		t.Error("I devrait être valide")
	}

	if !Valide("IVXLCDM") {
		t.Error("IVXLCDM devrait être valide")
	}

	if !Valide("MCMXCIV") {
		t.Error("MCMXCIV devrait être valide")
	}

	if Valide("") {
		t.Error("Une chaîne vide ne devrait pas être valide")
	}

	if Valide("ABC") {
		t.Error("ABC ne devrait pas être valide")
	}

	if Valide("xiv") {
		t.Error("xiv ne devrait pas être valide car les minuscules ne sont pas autorisées")
	}

	if Valide("X V") {
		t.Error("X V ne devrait pas être valide car il contient un espace")
	}

	if Valide("IVZ") {
		t.Error("IVZ ne devrait pas être valide")
	}

	if Valide("123") {
		t.Error("123 ne devrait pas être valide")
	}
}
