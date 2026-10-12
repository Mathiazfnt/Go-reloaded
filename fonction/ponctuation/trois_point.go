package fonction

import "strings"

func Trois_point(tableau []string) []string {
	for mot := 0; mot < len(tableau); mot++ {
		if  (tableau[mot] == ".." && mot > 0 ) ||
			(strings.HasPrefix(tableau[mot], "..") && mot > 0) {

			tableau[mot-1] = tableau[mot-1] + ".."
			tableau = append(tableau[:mot], tableau[mot+1:]...)
			mot--
		}
	}

	return tableau
}
