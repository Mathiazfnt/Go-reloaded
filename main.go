package main

import (
	"fmt"
	"os"

	ft "Go-reloaded/fonction"
	fc "Go-reloaded/fonction/casse_mots"
	fn "Go-reloaded/fonction/conversions_numeriques__a_an"
	fp "Go-reloaded/fonction/ponctuation"
)

func main() { // Il lit le fichier, applique chaque règle de modification dans l'ordre.
	texte, _ := os.ReadFile(os.Args[1])
	contenu := string(texte)

	matrice := ft.CreerMatrice(contenu)
	matrice = regles(matrice)

	result := ft.RefaireTexte_mat(matrice)
	os.WriteFile("result.txt", []byte(result), 0644)

	fmt.Print("tout marche")
}


func regles(matrice [][]string) [][]string{
	matrice = fn.Hex(matrice)
	matrice = fn.Bin(matrice)
	matrice = fc.Upp(matrice)
	matrice = fc.Low(matrice)
	matrice = fc.Cap(matrice)
	matrice = fp.Ponct_spl(matrice)
	return matrice
}

