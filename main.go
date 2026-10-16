package main

import (
	"fmt"
	"os"
	"strings"

	ft "Go-reloaded/fonction"

	fc "Go-reloaded/fonction/casse_mots"
	fn "Go-reloaded/fonction/conversions_numeriques__a_an"
	fp "Go-reloaded/fonction/ponctuation"

	fbcn "Go-reloaded/fonction/bonus/conversions_numeriques_bonus"
	fro "Go-reloaded/fonction/bonus/roman"
	fm "Go-reloaded/fonction/bonus/nouveaux_modificateurs"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Utilisation : go run . sample.txt result.txt")
		return
	}

	texte, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Il y a une erreur :", err)
		return
	}

	contenu := string(texte)
	lignes := strings.Split(contenu, "\n")

	for i := 0; i < len(lignes); i++ {
		tab := ft.Creertab(lignes[i])
		tab, err = regles(tab)
		if err != nil {
			fmt.Printf("Erreur ligne %d : %v\n", i+1, err)
			return
		}
		lignes[i] = strings.Join(tab, " ")
	}

	result := strings.Join(lignes, "\n")

	sortie := os.Args[2]
	tabbyte := []byte(result)

	err = os.WriteFile(sortie, tabbyte, os.FileMode(os.O_CREATE|os.O_WRONLY|os.O_TRUNC))
	if err != nil {
		fmt.Println("Il y a une erreur :", err)
		return
	}
}

func regles(tab []string) ([]string, error) {
	// on commence par les regle avec les  ()
	tab = fn.Hex(tab)
	tab = fn.Bin(tab)
	tab = fn.An(tab)
	tab = fc.Upp(tab)
	tab = fc.Low(tab)
	tab = fc.Cap(tab)
	// ensuite la ponctuation
	tab = fp.Ponct_spl(tab)
	tab = fp.Trois_point_Excla_inter(tab)
	tab = fp.Apost_Guill(tab)

	// bonus :
	tab = fbcn.Oct(tab)
	tab = fbcn.To_bin(tab)
	tab = fbcn.To_hex(tab)
	tab = fro.Roman(tab)
	tab = fro.To_roman(tab)
	var err error
	tab, err = fm.Rev(tab)
	if err != nil {
		return nil, err
	}
	tab = fm.Title(tab)

	return tab, nil
}
