
package conversions_numeriques__a_an

import "strings"

func An(tab []string) []string {
	liste := "aeiouhAEIOUH"

	for mot := 0; mot < len(tab)-1; mot++ {
		if len(tab[mot+1]) == 0 {
			continue
		}

		if tab[mot] == "a" || tab[mot] == "A" {
			if strings.Contains(liste, string(tab[mot+1][0])) {
				tab[mot] += "n"
			}
		} else if tab[mot] == "an" || tab[mot] == "An" {
			if !strings.Contains(liste, string(tab[mot+1][0])) {
				tab[mot] = string(tab[mot][0])
			}
		}
	}

	return tab
}
