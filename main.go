package main

import (
	"fmt"
	"os"

	ft "Go-reloaded/fonction"
)

func main() { // Il lit le fichier, applique chaque règle de modification dans l'ordre.
	fichier := os.Args[1]
	texte, _ := os.ReadFile(fichier)
	contenu := string(texte)

	matrice := ft.CrearinMatrice(contenu)
	fmt.Print(matrice)
	matrice = ft.Hex(matrice)
	matrice = ft.Bin(matrice)
	result := ft.RefaireTexte(matrice)
	os.WriteFile("result.txt", []byte(result), 0644)
}

