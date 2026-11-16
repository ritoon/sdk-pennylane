# Update a category

`PUT /api/external/v2/categories/{id}`

This endpoint updates a category.

> ℹ️
> This endpoint requires the following scope: `categories:all`

[Source Pennylane](https://pennylane.readme.io/reference/updatecategory) · [Index du SDK](../operations.md)

## Appel Go

```go
UpdateCategoryWithResponse(ctx context.Context, id string, params *UpdateCategoryParams, body UpdateCategoryJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateCategoryResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | string | obligatoire | The unique identifier of the category |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.analytical_code` | string ou null | facultatif | Cannot be changed when the `kind` of the category group is `structured_analytical`.<br><br>example `"CODE123"` |
| `$.direction` | string | facultatif | Only applicable for _treasury_ categories<br>enum `["cash_in","cash_out"]` |
| `$.label` | string | facultatif | example `"Alimentaire"` |
| `$.ledger_account_id` | integer (int64) | facultatif | The ID of the ledger account to wire this category to. Only valid for categories belonging to a category group with self-service accounting enabled; rejected otherwise. Bank, pending, and, when applicable, cash accounts cannot be used as transaction counterparts. Archived ledger accounts cannot be used. Omit the field to leave the current ledger account unchanged. Self-service accounting categories must always have a ledger account, so it cannot be cleared.<br>example `124` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "analytical_code": "CODE123",
  "label": "Alimentaire",
  "ledger_account_id": 124
}
```

</details>

## Sortie

### HTTP 200

Returns the updated category

`application/json` → `réponse.JSON200` (type `*GetCategoryStatus200`).

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
