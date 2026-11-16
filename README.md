# SDK Pennylane pour Go

Ce module expose les **186 opérations de la Company API v2** répertoriées sur le portail Pennylane le 16 novembre 2026. Les requêtes, les réponses par statut HTTP et les méthodes du client sont générées en Go à partir des définitions OpenAPI de chaque fiche. Les trois fiches d’événements webhook décrivent des notifications entrantes. Les API Firm et Firm Group disposent de portails distincts. [Index officiel](https://pennylane.readme.io/llms.txt) · [Familles d’API](https://pennylane.readme.io/docs/what-apis-are-available)

## Installation

```sh
go get github.com/ritoon/sdk-pennylane@v0.1.0
```

Le [guide du client](docs/utilisation.md) montre comment vérifier la connexion et parcourir les journaux. La [documentation du SDK](docs/README.md) recense les méthodes et les références Pennylane.

## Tester la connexion

L’[exemple Go](exemple/main.go) appelle `Ping` pour vérifier que Pennylane reconnaît le token et la société. Cet appel est en lecture seule.

Depuis la racine du dépôt, renseignez `PENNYLANE_TOKEN` dans votre environnement. Avec un token Firm, exportez aussi `PENNYLANE_COMPANY_ID` avec l’identifiant numérique de la société ciblée. Lancez ensuite la commande suivante.

```sh
go run ./exemple
```

En cas de succès, le programme affiche `Connexion à Pennylane réussie.`. En cas d’échec, il affiche l’erreur et se termine avec un code non nul.
