package roman

import "testing"

func TestTo_roman(t *testing.T) {
	// Test de la fonction to_roman
	if to_roman(1) != "I" {
		t.Error("1 devrait donner I")
	}

	if to_roman(4) != "IV" {
		t.Error("4 devrait donner IV")
	}

	if to_roman(14) != "XIV" {
		t.Error("14 devrait donner XIV")
	}

	if to_roman(1994) != "MCMXCIV" {
		t.Error("1994 devrait donner MCMXCIV")
	}

	if to_roman(3999) != "MMMCMXCIX" {
		t.Error("3999 devrait donner MMMCMXCIX")
	}

	if to_roman(0) != "" {
		t.Error("0 devrait donner une chaîne vide")
	}

	if to_roman(4000) != "" {
		t.Error("4000 devrait donner une chaîne vide")
	}

	// Test de la fonction To_roman
	tab := []string{"14", "(to_roman)"}
	resultat := To_roman(tab)

	if len(resultat) != 1 || resultat[0] != "XIV" {
		t.Error("14 (to_roman) devrait donner [XIV]")
	}

	tab = []string{"1994", "(to_roman)"}
	resultat = To_roman(tab)

	if len(resultat) != 1 || resultat[0] != "MCMXCIV" {
		t.Error("1994 (to_roman) devrait donner [MCMXCIV]")
	}

	tab = []string{"0", "(to_roman)"}
	resultat = To_roman(tab)

	if len(resultat) != 1 || resultat[0] != "0" {
		t.Error("0 devrait rester inchangé")
	}

	tab = []string{"abc", "(to_roman)"}
	resultat = To_roman(tab)

	if len(resultat) != 1 || resultat[0] != "abc" {
		t.Error("abc devrait rester inchangé")
	}

	tab = []string{"(to_roman)"}
	resultat = To_roman(tab)

	if len(resultat) != 0 {
		t.Error("(to_roman) sans nombre avant devrait être supprimé")
	}
}
