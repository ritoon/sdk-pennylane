# Vérifier et actualiser le SDK

Go 1.24 ou ultérieur. Le contrat consolidé est dans `internal/spec/company.json`. Chaque opération contient le lien vers sa fiche officielle. Depuis la racine du dépôt, le générateur Markdown reconstruit l’index et les fiches dans `docs/endpoints/` à partir du contrat et des méthodes Go compilées.

Les corps JSON contenant des unions `oneOf` ou `anyOf` sont placés dans `components.schemas` par l’assembleur, sous leur nom Go existant (`…JSONBody`). Avec oapi-codegen v2.5.1, cette représentation produit les méthodes `From…`, `As…`, `Merge…`, `MarshalJSON` et `UnmarshalJSON` nécessaires ; les schémas de requête inline n’obtiennent pas ces méthodes. Les références ne changent pas le contrat HTTP et le générateur Markdown les résout pour conserver les tableaux de champs et les exemples.

`payload_regression_test.go` vérifie le transport des champs concernés, y compris les valeurs nulles documentées. `union_regression_test.go` vérifie les variantes des unions de requête et leurs méthodes publiques. Ces tests utilisent un transport simulé et ne créent aucune donnée dans Pennylane.

```sh
go generate ./...
go run ./internal/docgen
go test ./...
```

Pour actualiser le contrat, télécharger l’index officiel et les fiches Markdown du répertoire `reference/` dans un dossier local, en conservant leurs noms. Assembler les fiches avant de régénérer le client.

```sh
go run ./internal/spec /chemin/llms.txt /chemin/fiches > company.next.json &&
    mv company.next.json internal/spec/company.json
go generate ./...
go run ./internal/docgen
go test ./...
```

## Télécharger les pages officielles en Markdown

Depuis la racine du dépôt, ces commandes constituent une copie locale de toutes les pages Markdown de l’index (guides, référence, conditions et changelog). Les fichiers gardent leur chemin sur le portail. Le téléchargement nécessite `curl` et `rg`.

```sh
mkdir -p docs/pennylane
curl -fsSL https://pennylane.readme.io/llms.txt -o docs/pennylane/index.md
rg -o 'https://pennylane\.readme\.io/(docs|reference|page|changelog)/[A-Za-z0-9_.-]+\.md' docs/pennylane/index.md |
  sort -u |
  while IFS= read -r url; do
    doc_path=${url#https://pennylane.readme.io/}
    mkdir -p "docs/pennylane/${doc_path%/*}" || exit 1
    curl --retry 3 -fsSL "$url" -o "docs/pennylane/$doc_path" || exit 1
  done
```

Les deux entrées de la commande d’assemblage sont alors `docs/pennylane/index.md` et `docs/pennylane/reference/`.

## Découpage reproductible par domaine

`go generate ./...` lance `internal/sdkgen`. Ce programme exécute la version
épinglée d’oapi-codegen dans un dossier temporaire, puis analyse les déclarations
Go avec `go/parser`. Il conserve leur source et leurs commentaires, répartit
les déclarations selon les opérations du contrat et ne retient que les imports
utilisés dans chaque fichier. Il ne réécrit ni les signatures ni les corps.

La table `domains` associe les ressources HTTP aux domaines. Une nouvelle
ressource inconnue fait échouer la génération : ajouter son association avant
de régénérer. Les noms sont associés au préfixe d’opération le plus long ; les
méthodes de modèles suivent leur type receveur. Les déclarations communes et
les interfaces complètes restent dans `api_common.gen.go`.

Les fichiers `api_*.gen.go` sont réservés au générateur. Après génération
réussie en mémoire, il remplace les sorties et supprime les anciens fichiers
de cette famille ainsi que l’ancien `api.gen.go`. Le contrat JSON reste unique.
L’[index du code généré](code-genere.md) décrit les domaines ; le mettre à jour
si la table des domaines change.

Le test `TestGenerationPreservesDeclarations` compare toutes les déclarations
(y compris les corps des méthodes) avec la sortie monolithique d’oapi-codegen,
vérifie les fichiers versionnés, le nettoyage des sorties obsolètes et la
stabilité d’une seconde génération. La CI contrôle aussi la documentation et
les nouveaux fichiers générés non versionnés.
