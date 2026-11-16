# Exemples métier Pennylane

Lancer les commandes depuis la racine du dépôt avec `PENNYLANE_TOKEN` dans l’environnement. Avec un token Firm, configurer aussi `PENNYLANE_COMPANY_ID`. La commande `go run ./exemple` seule vérifie la connexion. Les commandes de création et d’envoi agissent sur la société configurée.

## Lire la balance comptable

```sh
go run ./exemple balance 2026-01-01 2026-03-31
```

La commande lit toutes les pages de la période et écrit une ligne JSON par compte, avec son numéro, son libellé, ses débits et ses crédits. Les montants conservent leur représentation décimale en chaîne. Scope requis : `trial_balance:readonly`. [Code Go](balance.go) · [Contrat](../docs/endpoints/getTrialBalance.md)

## Envoyer une facture existante par e-mail

```sh
go run ./exemple envoyer-facture 12345 facturation@example.org
```

Remplacer les deux arguments par l’ID de la facture et l’adresse du destinataire. Utiliser une facture finalisée ou importée dont le PDF est disponible. Un statut 204 confirme la prise en charge de l’envoi. Si le PDF est en préparation, Pennylane peut répondre 409. Pour une facture éligible à la facturation électronique, l’envoi par e-mail empêche la transmission ultérieure par plateforme agréée. Scope requis : `customer_invoices:all`. [Code Go](facture.go) · [Contrat e-mail](../docs/endpoints/sendByEmailCustomerInvoice.md) · [Transmission par plateforme agréée](../docs/endpoints/sendToPaCustomerInvoice.md)

## Créer un client entreprise

Adapter le [modèle de client](donnees/client.json) dans une copie locale. La commande retourne l’identifiant Pennylane du client créé sur la sortie standard. Le conserver pour le mandat. Scope requis : `customers:all`. [Code Go](client.go) · [Contrat](../docs/endpoints/postCustomer.md)

```sh
cp exemple/donnees/client.json /tmp/client-pennylane.json
# Modifier le nom et l’adresse de facturation dans la copie.
CLIENT_ID=$(go run ./exemple creer-client < /tmp/client-pennylane.json)
```

## Enregistrer un mandat SEPA signé pour ce client

Renseigner dans un fichier local privé la référence unique, l’IBAN, le BIC et la date réelle de signature du mandat. Le [modèle JSON](donnees/mandat.json) donne les champs attendus. La commande affecte `CLIENT_ID` au champ `customer_id` et affiche l’ID du mandat créé. On peut aussi passer l’ID d’un client existant. Scope requis : `customer_mandates:all`. [Code Go](mandat.go) · [Contrat](../docs/endpoints/postSepaMandates.md)

```sh
go run ./exemple creer-mandat "$CLIENT_ID" < /chemin/prive/mandat.json
```
