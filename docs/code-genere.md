# Index du code Go généré

Tous les fichiers restent dans le package `pennylane` à la racine du module.
Les noms de types, constantes, fonctions, méthodes et interfaces sont conservés.
Chaque domaine regroupe ses modèles, paramètres, constructeurs de requêtes,
appels HTTP et décodeurs de réponses. Les changelogs suivent leur ressource.

| Fichier | Domaine |
|---|---|
| [api_common.gen.go](../api_common.gen.go) | Transport HTTP, configuration et interfaces complètes du client |
| [api_company.gen.go](../api_company.gen.go) | Profil, fonctionnalités de société et immatriculations PA |
| [api_customers.gen.go](../api_customers.gen.go) | Clients entreprise/particulier, contacts et adresses électroniques |
| [api_customer_invoices.gen.go](../api_customer_invoices.gen.go) | Factures clients, imports, échéances, paiements et pièces jointes |
| [api_supplier_invoices.gen.go](../api_supplier_invoices.gen.go) | Factures fournisseurs, imports et rapprochements |
| [api_suppliers.gen.go](../api_suppliers.gen.go) | Fournisseurs |
| [api_accounting.gen.go](../api_accounting.gen.go) | Écritures, comptes, journaux, exercices, balance et exports |
| [api_banking.gen.go](../api_banking.gen.go) | Comptes bancaires, établissements et transactions |
| [api_mandates.gen.go](../api_mandates.gen.go) | Mandats SEPA, GoCardless et compte pro |
| [api_quotes.gen.go](../api_quotes.gen.go) | Devis |
| [api_commercial_documents.gen.go](../api_commercial_documents.gen.go) | Documents commerciaux |
| [api_subscriptions.gen.go](../api_subscriptions.gen.go) | Abonnements de facturation et modèles d’e-mails |
| [api_products.gen.go](../api_products.gen.go) | Produits |
| [api_categories.gen.go](../api_categories.gen.go) | Catégories et groupes |
| [api_purchase_requests.gen.go](../api_purchase_requests.gen.go) | Demandes d’achat |
| [api_attachments.gen.go](../api_attachments.gen.go) | Téléversement de fichiers |
| [api_webhooks.gen.go](../api_webhooks.gen.go) | Abonnements webhook |

Pour explorer le SDK, commencer par le domaine concerné, puis rechercher le
symbole précis avec `rg -n 'GetCustomers' api_customers.gen.go`. Exclure
`*.gen.go` des recherches générales sur le code écrit à la main.

Le découpage facilite les lectures ciblées ; il ne déduplique pas les schémas
et ne réduit pas le nombre total de déclarations. Les interfaces regroupant
toutes les opérations restent dans le fichier commun pour préserver l’API.

Ne pas modifier ces fichiers à la main. `go generate ./...` les reconstruit via
[internal/sdkgen](../internal/sdkgen/main.go). Voir la
[procédure de génération](generation.md).
