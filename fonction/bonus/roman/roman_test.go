package roman

import "testing"

// Test de la fonction roman
func TestRoman(t *testing.T) {
	if roman("III") != 3 {
		t.Error("III devrait donner 3")
	}

	if roman("IV") != 4 {
		t.Error("IV devrait donner 4")
	}

	if roman("XIV") != 14 {
		t.Error("XIV devrait donner 14")
	}

	if roman("MCMXCIV") != 1994 {
		t.Error("MCMXCIV devrait donner 1994")
	}

	if roman("IIII") != 4 {
		t.Error("roman convertit les symboles, même si l'écriture est incorrecte")
	}
	
}

// Test de la fonction Roman
func TestRomanTab(t *testing.T) {
	tab := []string{"XIV", "(roman)"}
	resultat := Roman(tab)

	if len(resultat) != 1 || resultat[0] != "14" {
		t.Error("XIV (roman) devrait donner [14]")
	}

	tab = []string{"IIII", "(roman)"}
	resultat = Roman(tab)

	if len(resultat) != 1 || resultat[0] != "IIII" {
		t.Error("IIII (roman) devrait rester inchangé")
	}

	tab = []string{"IC", "(roman)"}
	resultat = Roman(tab)

	if len(resultat) != 1 || resultat[0] != "IC" {
		t.Error("IC (roman) devrait rester inchangé")
	}

	tab = []string{"ABC", "(roman)"}
	resultat = Roman(tab)

	if len(resultat) != 1 || resultat[0] != "ABC" {
		t.Error("ABC (roman) devrait rester inchangé")
	}

	tab = []string{"(roman)"}
	resultat = Roman(tab)

	if len(resultat) != 0 {
		t.Error("(roman) sans chiffre avant devrait être supprimé")
	}

	tab = []string{"X", "(roman)", "V", "(roman)"}
	resultat = Roman(tab)

	if len(resultat) != 2 || resultat[0] != "10" || resultat[1] != "5" {
		t.Error("Les deux chiffres romains devraient être convertis")
	}
}
