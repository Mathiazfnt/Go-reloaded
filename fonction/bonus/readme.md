# Bonus Go Reforged

Le dossier `bonus/` regroupe une sélection de fonctionnalités inspirées de
certaines parties du sujet **Go Reforged**, qui prolonge le projet
**Go-reloaded**. Il ne constitue pas une implémentation complète de
Go Reforged : seules quelques règles supplémentaires ont été reprises et
intégrées au programme existant.

## Fonctionnalités ajoutées

- Conversions numériques : octal vers décimal `(oct)`, décimal vers binaire
  `(to bin)` et décimal vers hexadécimal `(to hex)`.
- Conversions entre nombres décimaux et chiffres romains `(roman)` et
  `(to_roman)`.
- Modificateur `(rev)` pour inverser les lettres d'un ou plusieurs mots
  précédents, en conservant les caractères accentués.
- Modificateur `(title)` pour appliquer une casse de titre à un ou plusieurs
  mots précédents.

Les commandes suivies d'un nombre permettent de traiter plusieurs mots lorsque
cette règle le prend en charge, par exemple `(rev, 2)` ou `(title, 4)`.

## Ce qui n'est pas inclus

Ce bonus reprend seulement certaines fonctionnalités du sujet Go Reforged. Il
n'implémente pas son architecture complète ni les autres exigences comme
l'analyse lexicale dédiée, les blocs, l'échappement, les diagnostics détaillés,
le traitement en flux de gros fichiers, les options de ligne de commande, le
traitement concurrent, le fuzzing ou les benchmarks.

## Organisation

- `conversions_numeriques_bonus/` contient les conversions numériques
  supplémentaires.
- `roman/` contient les conversions et la validation des chiffres romains.
- `nouveaux_modificateurs/` contient les règles `(rev)` et `(title)`.

```text
bonus/
├── conversions_numeriques_bonus/
│   ├── oct.go
│   ├── to_bin.go
│   └── to_hex.go
├── nouveaux_modificateurs/
│   ├── rev.go
│   └── title.go
├── roman/
│   ├── roman.go
│   ├── to_roman.go
│   └── valide.go
└── readme.md
```

Les fonctions bonus sont appelées par le programme principal à la racine du
dépôt. Depuis cette racine, lancer les tests avec :

```powershell
go test ./...
```