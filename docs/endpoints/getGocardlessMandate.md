# Get a Gocardless mandate

`GET /api/external/v2/gocardless_mandates/{id}`

This endpoint allows you to retrieve a specific Gocardless mandate by ID.

> ℹ️
> This endpoint requires one of the following scopes: `customer_mandates:all`, `customer_mandates:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getgocardlessmandate) · [Index du SDK](../operations.md)

## Appel Go

```go
GetGocardlessMandateWithResponse(ctx context.Context, id int, params *GetGocardlessMandateParams, reqEditors ...RequestEditorFn) (*GetGocardlessMandateResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | ID of the Gocardless mandate to retrieve |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A Gocardless mandate

`application/json` → `réponse.JSON200` (type `*GetGocardlessMandateStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.created_at` | string (date-time) | obligatoire | Creation date of the Gocardless mandate<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.customer.id` | integer (int64) | obligatoire | example `42` |
| `$.customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.external_customer_account` | string | obligatoire | Identifier for the Gocardless customer account.<br>example `"test_account_id"` |
| `$.external_customer_label` | string ou null | obligatoire | Label for the Gocardless customer account.<br>example `"test_label"` |
| `$.external_reference` | string ou null | obligatoire | Unique external identifier for the Gocardless mandate.<br>example `"MD000VDD0HN1RP"` |
| `$.id` | integer (int64) | obligatoire | example `42` |
| `$.status` | string | obligatoire | enum `["pending_customer_approval","pending_submission","submitted","active","failed","cancelled","expired","consumed","replaced","blocked","bank_disconnected"]` |
| `$.updated_at` | string (date-time) | obligatoire | Last update date of the Gocardless mandate<br>example `"2023-08-07T14:23:12.000Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "created_at": "2023-08-07T14:23:12.000Z",
  "customer": {
    "id": 42,
    "url": "https://app.pennylane.com/api/external/v2/customers/42"
  },
  "external_customer_account": "test_account_id",
  "external_customer_label": "test_label",
  "external_reference": "MD000VDD0HN1RP",
  "id": 42,
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
