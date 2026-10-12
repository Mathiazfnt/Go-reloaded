package main

import (
	"fmt"
	"os"
	"strings"

	ft "Go-reloaded/fonction"
	fc "Go-reloaded/fonction/casse_mots"
	fn "Go-reloaded/fonction/conversions_numeriques__a_an"
	fp "Go-reloaded/fonction/ponctuation"
)

func main() { // Il lit le fichier, applique chaque règle de modification dans l'ordre.
	texte, _ := os.ReadFile(os.Args[1])
	contenu := string(texte)

	tab := ft.CreerMatrice(contenu)
	tab = regles(tab)

	result := strings.Join(tab, " ")
	os.WriteFile("result.txt", []byte(result), 0644)

	fmt.Print("tout marche")
}


func regles(tab []string) []string{
	tab = fn.Hex(tab)
	tab = fn.Bin(tab)
	tab = fn.An(tab)
	tab = fc.Upp(tab)
	tab = fc.Low(tab)
	tab = fc.Cap(tab)
	tab = fp.Ponct_spl(tab)
	tab = fp.Trois_point(tab)
	
	return tab
}

