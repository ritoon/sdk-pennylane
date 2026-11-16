# Vérifier et actualiser le SDK

Go 1.24 ou ultérieur. Le contrat consolidé est dans `internal/spec/company.json`. Chaque opération contient le lien vers sa fiche officielle. Depuis la racine du dépôt, le générateur Markdown reconstruit l’index et les fiches dans `docs/endpoints/` à partir du contrat et des méthodes Go compilées.

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
