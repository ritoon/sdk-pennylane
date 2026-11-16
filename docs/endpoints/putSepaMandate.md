# Update a SEPA mandate

`PUT /api/external/v2/sepa_mandates/{id}`

This endpoint allows you to update an existing SEPA mandate

> ℹ️
> This endpoint requires the following scope: `customer_mandates:all`

[Source Pennylane](https://pennylane.readme.io/reference/putsepamandate) · [Index du SDK](../operations.md)

## Appel Go

```go
PutSepaMandateWithResponse(ctx context.Context, id int, params *PutSepaMandateParams, body PutSepaMandateJSONRequestBody, reqEditors ...RequestEditorFn) (*PutSepaMandateResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | ID of the SEPA mandate to update |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.bank` | string ou null | facultatif | Name of the customer's bank<br>example `"BNP Paribas"` |
| `$.bic` | string | facultatif | Bank Identifier Code (BIC) of the customer's bank<br>example `"BNPAFRPP"` |
| `$.customer_id` | integer (int64) | facultatif | ID of the customer for which the mandate is created<br>example `47334785` |
| `$.iban` | string | facultatif | International Bank Account Number (IBAN) of the customer<br>example `"FR3612739000506556687647Z31"` |
| `$.identifier` | string | facultatif | Unique identifier for the mandate<br>example `"K-02-2025-12345"` |
| `$.sequence_type` | string | facultatif | SEPA mandate sequence type that defines the payment process.<br>- `FRST`: First payment in a series of recurring payments<br>- `OOFF`: One-off payment<br>- `RCUR`: Recurring payment that is not the first payment<br>- `FNAL`: Final payment in a series of recurring payments<br><br>enum `["FRST","OOFF","RCUR","FNAL"]`<br>default `"RCUR"`<br>example `"RCUR"` |
| `$.signed_at` | string (date) | facultatif | Date when the mandate was signed<br>example `"2025-04-24"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "bank": "BNP Paribas",
  "bic": "BNPAFRPP",
  "customer_id": 47334785,
  "iban": "FR3612739000506556687647Z31",
  "identifier": "K-02-2025-12345",
  "sequence_type": "RCUR",
  "signed_at": "2025-04-24"
}
```

</details>

## Sortie

### HTTP 200

Returns the updated SEPA mandate

`application/json` → `réponse.JSON200` (type `*PostSepaMandatesStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.bank` | string ou null | obligatoire | Name of the customer's bank<br>example `"BNP Paribas"` |
| `$.bic` | string | obligatoire | Bank Identifier Code (BIC) of the customer's bank<br>example `"BNPAFRPP"` |
| `$.created_at` | string (date-time) | obligatoire | Creation date of the SEPA mandate<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.customer` | object | obligatoire | additionalProperties `false` |
| `$.customer.id` | integer (int64) | obligatoire | example `42` |
| `$.customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.iban` | string | obligatoire | International Bank Account Number (IBAN) of the customer<br>example `"FR3612739000506556687647Z31"` |
| `$.id` | integer (int64) | obligatoire | ID of the created SEPA mandate<br>example `12345` |
| `$.identifier` | string | obligatoire | Unique identifier for the mandate<br>example `"K-02-2023-12345"` |
| `$.sequence_type` | string | obligatoire | SEPA mandate sequence type that defines the payment process.<br>- `FRST`: First payment in a series of recurring payments<br>- `OOFF`: One-off payment<br>- `RCUR`: Recurring payment that is not the first payment<br>- `FNAL`: Final payment in a series of recurring payments<br><br>enum `["FRST","OOFF","RCUR","FNAL"]`<br>default `"RCUR"`<br>example `"RCUR"` |
| `$.signed_at` | string (date) | obligatoire | Date when the mandate was signed<br>example `"2025-04-24"` |
| `$.updated_at` | string (date-time) | obligatoire | Last update date of the SEPA mandate<br>example `"2023-08-07T14:23:12.000Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "bank": "BNP Paribas",
  "bic": "BNPAFRPP",
  "created_at": "2023-08-07T14:23:12.000Z",
  "customer": {
    "id": 42,
    "url": "https://app.pennylane.com/api/external/v2/customers/42"
  },
  "iban": "FR3612739000506556687647Z31",
  "id": 12345,
  "identifier": "K-02-2023-12345",
  "sequence_type": "RCUR",
  "signed_at": "2025-04-24",
  "updated_at": "2023-08-07T14:23:12.000Z"
}
```

</details>

### HTTP 400

Bad request

`application/json` → `réponse.JSON400` (type `*GetJournalsStatus400`).

[Champs du payload d’erreur](erreurs.md#getjournalsstatus400)

### HTTP 401

Access token is missing or invalid

`application/json` → `réponse.JSON401` (type `*GetJournalsStatus401`).

[Champs du payload d’erreur](erreurs.md#getjournalsstatus401)

### HTTP 403

Access to this resource forbidden

`application/json` → `réponse.JSON403` (type `*GetJournalsStatus403`).

[Champs du payload d’erreur](erreurs.md#getjournalsstatus403)

### HTTP 404

The resource was not found

`application/json` → `réponse.JSON404` (type `*GetJournalsStatus404`).

[Champs du payload d’erreur](erreurs.md#getjournalsstatus404)

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
