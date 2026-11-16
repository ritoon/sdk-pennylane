# Create a category

`POST /api/external/v2/categories`

Create a category

> ℹ️
> This endpoint requires the following scope: `categories:all`

[Source Pennylane](https://pennylane.readme.io/reference/postcategories) · [Index du SDK](../operations.md)

## Appel Go

```go
PostCategoriesWithResponse(ctx context.Context, params *PostCategoriesParams, body PostCategoriesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCategoriesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.analytical_code` | string ou null | facultatif | Categories cannot be created in a category group whose `kind` is `structured_analytical`<br>through this endpoint, even when an `analytical_code` is provided.<br><br>example `"CODE123"` |
| `$.category_group_id` | integer (int64) | obligatoire | example `23` |
| `$.direction` | anyOf | facultatif | Only applicable for _treasury_ categories. If not provided, defaults to "cash_out". |
| `$.direction (anyOf 1)` | string | selon variante | Only applicable for _treasury_ categories<br>enum `["cash_in","cash_out"]` |
| `$.direction (anyOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.label` | string | obligatoire | example `"Loyers"` |
| `$.ledger_account_id` | integer (int64) ou null | facultatif | Id of the ledger account to map to this category.<br>Only valid for categories belonging to a category group with self-service accounting enabled; rejected otherwise.<br>Bank, pending, and, when applicable, cash accounts cannot be used as transaction counterparts.<br><br>example `124` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "analytical_code": "CODE123",
  "category_group_id": 23,
  "label": "Loyers",
  "ledger_account_id": 124
}
```

</details>

## Sortie

### HTTP 201

Returns the created category

`application/json` → `réponse.JSON201` (type `*GetCategoryStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.analytical_code` | string ou null | obligatoire | Always set for categories of a category group whose `kind` is `structured_analytical`.<br><br>example `"CODE123"` |
| `$.category_group` | object | obligatoire | additionalProperties `false` |
| `$.category_group.id` | integer (int64) | obligatoire | example `348` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.direction` | anyOf | obligatoire |  |
| `$.direction (anyOf 1)` | string | selon variante | Only applicable for _treasury_ categories<br>enum `["cash_in","cash_out"]` |
| `$.direction (anyOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.id` | integer (int64) | obligatoire | example `42` |
| `$.label` | string | obligatoire | example `"Alimentaire"` |
| `$.ledger_account` | object ou null | obligatoire | The ledger account the category is wired to, or `null` when none is set.<br>additionalProperties `false` |
| `$.ledger_account.id` | integer (int64) | obligatoire | example `124` |
| `$.ledger_account.url` | string | obligatoire | The URL to the ledger account<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/124"` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "analytical_code": "CODE123",
  "category_group": {
    "id": 348
  },
  "created_at": "2023-08-30T10:08:08.146343Z",
  "id": 42,
  "label": "Alimentaire",
  "ledger_account": {
    "id": 124,
    "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/124"
  },
  "updated_at": "2023-08-30T10:08:08.146343Z"
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
