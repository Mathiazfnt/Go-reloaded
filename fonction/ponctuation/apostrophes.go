package ponctuation

func Apost_Guill(tab []string) []string {
	tab = regrouperSigne(tab, "'")
	return regrouperSigne(tab, "\"")
}

func regrouperSigne(tab []string, signe string) []string {
	total := 0
	for _, mot := range tab {
		if mot == signe {
			total++
		}
	}

	orphelin := total%2 != 0
	resultat := make([]string, 0, len(tab))
	occurrences := 0
	ouvertureEnAttente := false

	for i := 0; i < len(tab); i++ {
		if tab[i] != signe {
			mot := tab[i]
			if ouvertureEnAttente {
				mot = signe + mot
				ouvertureEnAttente = false
			}
			resultat = append(resultat, mot)
			continue
		}

		occurrences++
		if orphelin && occurrences == total {
			if signe == "'" && len(resultat) > 0 && i+1 < len(tab) && estElision(resultat[len(resultat)-1]) {
				resultat[len(resultat)-1] += "'" + tab[i+1]
				i++
				orphelin = false
				continue
			}
			continue
		}

		if occurrences%2 != 0 {
			ouvertureEnAttente = true
		} else if len(resultat) > 0 {
			resultat[len(resultat)-1] += signe
		} else {
			resultat = append(resultat, signe)
		}
	}

	if ouvertureEnAttente {
		resultat = append(resultat, signe)
	}
	if orphelin {
		if len(resultat) > 0 {
			resultat[len(resultat)-1] += signe
		} else {
			resultat = append(resultat, signe)
		}
	}

	return resultat
}

func estElision(mot string) bool {
	switch mot {
	case "l", "d", "j", "m", "t", "s", "n", "c", "qu":
		return true
	default:
		return false
	}
}
