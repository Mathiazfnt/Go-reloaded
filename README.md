# Go-reloaded

## Présentation

Go-reloaded est un programme en ligne de commande qui lit un fichier texte,
applique des règles de conversion et de formatage, puis écrit le résultat dans
un fichier de sortie. Le fichier d'entrée reste inchangé.

Le texte est traité ligne par ligne. Les espaces multiples d'une ligne sont
normalisés en espaces simples dans le résultat.

## Prérequis

- Go 1.27.1 ou une version compatible avec celle déclarée dans `go.mod`.
- Git uniquement pour cloner le dépôt.

Le projet n'utilise que la bibliothèque standard de Go et ne nécessite aucune
configuration externe.

## Installation

```powershell
git clone https://github.com/Mathiazfnt/Go-reloaded.git
cd Go-reloaded
```

## Utilisation

Depuis la racine du dépôt, indique le fichier d'entrée puis le fichier de
sortie :

```powershell
go run . sample.txt result.txt
```

Le fichier de sortie est créé s'il n'existe pas et remplacé s'il existe déjà.
Pour compiler puis lancer le programme :

```powershell
go build -o go-reloaded.exe .
.\go-reloaded.exe sample.txt result.txt
```

Le programme attend exactement deux chemins en argument. Il affiche une
instruction d'utilisation si le nombre d'arguments est incorrect.

## Règles prises en charge

Les commandes de casse, de titre, d'inversion et les conversions numériques
agissent sur le ou les mots placés avant leur marqueur. Le nombre est facultatif
et indique combien de mots précédents traiter.

| Règle | Exemple d'entrée | Résultat |
| --- | --- | --- |
| Hexadécimal vers décimal | `2a (hex)` | `42` |
| Binaire vers décimal | `1010 (bin)` | `10` |
| Octal vers décimal | `17 (oct)` | `15` |
| Décimal vers binaire | `10 (to bin)` | `1010` |
| Décimal vers hexadécimal | `255 (to hex)` | `ff` |
| Tout en majuscules | `bonjour (up)` | `BONJOUR` |
| Tout en minuscules | `BONJOUR (low)` | `bonjour` |
| Première lettre en majuscule | `bonjour (cap)` | `Bonjour` |
| Casse de titre | `the lord of the rings (title, 5)` | `the Lord of the Rings` |
| Inverser un mot | `bonjour (rev)` | `ruojnob` |
| Inverser deux mots | `bonjour salut (rev, 2)` | `ruojnob tulas` |
| Ajuster `a` / `an` | `a apple and an banana` | `an apple and a banana` |
| Rapprocher la ponctuation | `bonjour ,monde !` | `bonjour, monde!` |
| Points et `!?` | `attends .. vraiment !?` | `attends.. vraiment!?` |
| Apostrophes et guillemets | `' bonjour '` | `'bonjour'` |

Les formes avec un nombre comprennent par exemple `(up, 2)`, `(low, 3)`,
`(cap, 2)`, `(title, 5)`, `(rev, 2)`, `(to bin, 2)` et `(to hex, 2)`.
La conversion octale `(oct)` traite un seul nombre précédent.

Les règles sont appliquées dans l'ordre défini dans `main.go`. Les conversions
numériques sont indépendantes : elles ne réalisent pas d'opérations
arithmétiques entre les valeurs.

## Limites connues

- L'ajustement de `a` et `an` examine seulement le premier caractère du mot
  suivant et utilise une liste limitée aux voyelles ASCII et à `h`.
- La règle des points de suspension colle actuellement deux points (`..`),
  même lorsque l'entrée en contient trois.
- Les nombres invalides pour certaines conversions peuvent rester inchangés ;
  une syntaxe ou un nombre invalide dans `(rev, nombre)` est signalé comme une
  erreur avec le numéro de ligne.
- Les règles de conversion romaine sont présentes dans le dossier `roman/`,
  mais ne sont pas activées dans `main.go`.

## Exemple fourni

`sample.txt` contient un exemple par règle active et `result.txt` contient le
résultat généré par cette commande :

```powershell
go run . sample.txt result.txt
```

## Tests

Exécuter les tests et vérifier la compilation de tous les packages avec :

```powershell
go test ./...
```

## Architecture

```text
.
├── main.go
├── sample.txt
├── result.txt
└──; fonction/
    ├── tableau.go
    ├── tableau_test.go
    ├── casse_mots/
    │   ├── cap.go
    │   ├── cap_test.go
    │   ├── low.go
    │   ├── low_test.go
    │   ├── upp.go
    │   ├── upp_test.go
    │   ├── depassement.go
    │   └── depassement_test.go
    ├── conversions_numeriques__a_an/
    │   ├── a_an.go
    │   ├── a_an_test.go
    │   ├── bin.go
    │   ├── bin_test.go
    │   └── hex.go
    ├── ponctuation/
    │   ├── apostrophes.go
    │   ├── apostrophe_test.go
    │   ├── ponctuation_spl.go
    │   ├── ponctuation_spl_test.go
    │   ├── exclama_inter.go
    │   └── exclama_inter_test.go
    └── bonus/
        ├── conversions_numeriques_bonus/
        │   ├── oct.go
        │   ├── oct_test.go
        │   ├── to_bin.go
        │   ├── to_bin_test.go
        │   ├── to_hex.go
        │   └── to_hex_test.go
        ├── nouveaux_modificateurs/
        │   ├── rev.go
        │   ├── rev_test.go
        │   ├── title.go
        │   └── title_test.go
        └── roman/
            ├── roman.go
            ├── to_roman.go
            └── valide.go
```

- `main.go` lit le fichier d'entrée, applique les règles et écrit le résultat.
- `fonction/tableau.go` regroupe les mots d'une ligne et conserve ensemble les
  commandes entre parenthèses.
- `fonction/casse_mots/` contient les règles de casse et le calcul du nombre de
  mots à traiter.
- `fonction/conversions_numeriques__a_an/` contient les conversions
  hexadécimale et binaire vers décimal et l'ajustement de `a` et `an`.
- `fonction/ponctuation/` regroupe le traitement de la ponctuation, des
  apostrophes et des guillemets.
- `fonction/bonus/` contient les conversions supplémentaires et les
  modificateurs `rev` et `title`. Le sous-dossier `roman/` n'est pas encore
  raccordé au traitement principal.
