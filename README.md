# Go-reloaded

## Présentation

Go-reloaded est un programme en ligne de commande qui lit un fichier texte,
applique des règles de formatage et écrit le résultat dans le fichier de sortie
indiqué en argument. Le fichier d'entrée n'est pas modifié.

## Fonctionnalités

- Convertir une valeur hexadécimale suivie de `(hex)` en décimal.
- Convertir une valeur binaire suivie de `(bin)` en décimal.
- Modifier la casse des mots avec `(up)`, `(low)` et `(cap)`. Un nombre
  permet d'appliquer la règle à plusieurs mots précédents, par exemple
  `(up, 2)`.
- Tenter d'ajuster `a` et `an` selon le token qui suit.
- Rapprocher certains signes de ponctuation séparés par des espaces.
- Traiter les points de suspension et les apostrophes entourant un mot.

Les règles sont appliquées dans l'ordre défini dans `main.go`. L'ajustement de
`a` et `an` ne reconnaît pas systématiquement la première lettre d'un mot ;
le traitement des points de suspension est lui aussi limité.

## Technologies

- Go (version déclarée dans `go.mod` : Go 1.27.1).
- Bibliothèque standard Go pour la lecture des fichiers, le traitement du
  texte et les conversions.

## Prérequis

- Installer une version de Go compatible avec le fichier `go.mod`.
- Git est nécessaire uniquement pour cloner le dépôt.

Aucune base de données, variable d'environnement ou configuration externe
n'est nécessaire.

## Installation

```powershell
git clone https://github.com/Mathiazfnt/Go-reloaded.git
cd Go-reloaded
```

## Lancement

Depuis la racine du dépôt, fournir le chemin du fichier d'entrée puis celui du
fichier de sortie :

```powershell
go run . sample.txt result.txt
```

Le fichier de sortie est créé s'il n'existe pas et son contenu est remplacé à
chaque exécution. Les deux chemins peuvent être adaptés.

Il est aussi possible de compiler le programme puis de l'exécuter :

```powershell
go build -o go-reloaded.exe .
.\go-reloaded.exe sample.txt result.txt
```

Le programme attend exactement deux arguments. Sans le chemin d'entrée et le
chemin de sortie, il affiche la commande d'utilisation.

## Exemple d'utilisation

Le fichier `sample.txt` du dépôt sert d'exemple d'entrée et couvre plusieurs
règles :

- `(cap)` et `(cap, 6)` pour capitaliser un ou plusieurs mots précédents ;
- `(up)` et `(low, 3)` pour modifier leur casse ;
- `(hex)` et `(bin)` pour convertir des nombres ;
- des signes de ponctuation séparés par des espaces et des points de
  suspension.

Lancer l'exemple avec :

```powershell
go run . sample.txt result.txt
```

Le programme écrit le texte transformé dans `result.txt`. Les conversions
sont individuelles : `42 (hex)` devient `66` et `10 (bin)` devient `2` ;
le programme ne fait pas l'addition suggérée par la phrase d'exemple. De
même, l'ajustement de `a` et `an` ne détecte pas encore de façon fiable la
première lettre du mot suivant, et le traitement de `...` peut le réduire à
`..`. Ces exemples permettent donc aussi d'observer les limites actuelles du
traitement.

## Tests

Lancer les tests et la vérification de compilation depuis la racine :

```powershell
go test ./...
```

Les packages du projet ne contiennent pas encore de tests automatisés ;
cette commande vérifie principalement leur compilation.

## Architecture

```text
.
├── main.go
├── sample.txt
├── result.txt
├── fonction/
│   ├── tableau.go
│   ├── casse_mots/
│   │   ├── cap.go
│   │   ├── low.go
│   │   └── upp.go
│   ├── conversions_numeriques__a_an/
│   │   ├── a_an.go
│   │   ├── bin.go
│   │   └── hex.go
│   └── ponctuation/
│       ├── apostrophes.go
│       ├── exclama_inter.go
│       ├── ponctuation_spl.go
│       └── trois_point.go
```

- `main.go` lit le fichier d'entrée passé en premier argument, traite le texte
  ligne par ligne et écrit dans le fichier de sortie passé en second argument.
- `fonction/tableau.go` regroupe les mots d'une ligne en tenant compte des
  directives écrites entre parenthèses.
- `fonction/casse_mots/` contient les règles de casse.
- `fonction/conversions_numeriques__a_an/` contient les conversions numériques
  et l'ajustement de `a` et `an`.
- `fonction/ponctuation/` contient les traitements de ponctuation.
- `sample.txt` est un exemple de fichier d'entrée ; `result.txt` est le fichier
  de sortie utilisé dans les exemples de lancement.
