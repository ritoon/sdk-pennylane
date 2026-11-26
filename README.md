# SDK Pennylane pour Go

Ce module expose les **186 opérations de la Company API v2** répertoriées sur le portail Pennylane le 16 novembre 2026. Les requêtes, les réponses par statut HTTP et les méthodes du client sont générées en Go à partir des définitions OpenAPI de chaque fiche. Les trois fiches d’événements webhook décrivent des notifications entrantes. Les API Firm et Firm Group disposent de portails distincts. [Index officiel](https://pennylane.readme.io/llms.txt) · [Familles d’API](https://pennylane.readme.io/docs/what-apis-are-available)

## Installation

```sh
go get github.com/ritoon/sdk-pennylane@v1.0.1
```

Le [guide du client](docs/utilisation.md) montre comment vérifier la connexion et parcourir les journaux. La [référence des endpoints](docs/operations.md) détaille les paramètres Go et HTTP, les corps de requête et les payloads retournés par statut. Exemples de fiches : [lister les journaux](docs/endpoints/getJournals.md) et [créer un compte bancaire](docs/endpoints/postBankAccount.md).

Le [petit index du code généré](docs/code-genere.md) permet de retrouver les fichiers Go par domaine.

## Tester la connexion

L’[exemple Go](exemple/main.go) appelle `Ping` pour vérifier que Pennylane reconnaît le token et la société. Cet appel est en lecture seule.

Depuis la racine du dépôt, renseignez `PENNYLANE_TOKEN` dans votre environnement. Avec un token Firm, exportez aussi `PENNYLANE_COMPANY_ID` avec l’identifiant numérique de la société ciblée. Lancez ensuite la commande suivante.

```sh
go run ./exemple
```

En cas de succès, le programme affiche `Connexion à Pennylane réussie.`. En cas d’échec, il affiche l’erreur et se termine avec un code non nul.

## Exemples métier

Le [guide des exemples exécutables](exemple/README.md) montre comment lire la balance comptable, envoyer une facture par e-mail, créer un client entreprise et enregistrer un mandat SEPA signé pour ce client.

## Tests et intégration continue

Chaque PR fusionnée est suivie d’une nouvelle version sémantique. Les règles
d’incrément et de publication sont décrites dans [CONTRIBUTING.md](CONTRIBUTING.md).

Les tests utilisent des transports HTTP simulés : aucun token ni accès à Pennylane
n’est nécessaire. Depuis la racine du dépôt :

```sh
go test -race -shuffle=on -count=1 -timeout=5m -covermode=atomic -coverpkg=./... -coverprofile=coverage.out ./...
python3 scripts/check-coverage.py coverage.out
go tool cover -html=coverage.out -o coverage.html
```

La [CI GitHub Actions](.github/workflows/ci.yml) s’exécute sur les PR et les pushes
sur `main`, avec Go 1.24 et la version stable. Elle contrôle le formatage, les
dépendances, `go vet`, la compilation, les tests avec détection des accès
concurrents et la reproductibilité du SDK et de la documentation générés. Les
rapports de couverture sont disponibles dans les artefacts pendant 14 jours.

Les seuils portent sur les instructions exécutées et sont définis dans
[`scripts/check-coverage.py`](scripts/check-coverage.py) :

| Périmètre | Minimum |
|---|---:|
| Projet complet | 57 % |
| Client généré | 57 % |
| SDK écrit à la main | 94 % |
| Exemples | 85 % |
| Générateur de documentation | 93 % |
| Préparation de la spécification | 91 % |
| Découpage du client généré | 80 % |

Le calcul fusionne les blocs dupliqués par `-coverpkg=./...` pour inclure les
appels au SDK depuis les tests des exemples sans les compter plusieurs fois.
Un seuil non atteint fait échouer la CI. Pour bloquer aussi la fusion d’une PR,
les checks `Go 1.24.x` et `Go stable` doivent être requis dans les règles de
protection de `main`.

Les tests de contrat exécutent les 186 méthodes typées et vérifient les routes,
verbes HTTP, en-têtes de société, paramètres obligatoires, types de corps,
uploads et décodage des statuts documentés. Leurs payloads minimaux sont
synthétiques : ils ne valident pas les règles métier du service distant.
Des scénarios ciblés couvrent notamment les erreurs de transport, le JSON
malformé, la fermeture des réponses, les champs multipart optionnels, les
entrées invalides et les curseurs de pagination manquants ou répétés.
