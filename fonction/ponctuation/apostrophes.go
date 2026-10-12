package fonction

import "strings"

func Apostrophes(tableau []string) []string {
	for mot := 0; mot < len(tableau); mot++ {
		if strings.HasPrefix(tableau[mot], "'") || strings.HasSuffix(tableau[mot], "\""){

		}
	}

	return tableau
}
