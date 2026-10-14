package ponctuation

import "strings"

func Trois_point_Excla_inter(tableau []string) []string {
	for mot := 0; mot < len(tableau); mot++ {
		if ((tableau[mot] == ".." || strings.HasPrefix(tableau[mot], "..")) ||
			(tableau[mot] == "!?" || strings.HasPrefix(tableau[mot], "!?"))) && mot > 0 {

			if strings.HasPrefix(tableau[mot], "..") {
				tableau[mot-1] = tableau[mot-1] + ".."
			} else {
				tableau[mot-1] = tableau[mot-1] + "!?"
			}

			tableau = append(tableau[:mot], tableau[mot+1:]...)
			mot--
		}
	}

	return tableau
}
