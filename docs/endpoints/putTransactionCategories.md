# Categorize a bank transaction

`PUT /api/external/v2/transactions/{transaction_id}/categories`

Update the categories of a transaction. You can pass categories that don't belong to the same category group. The sum of categories of a same group must equal `1`. In the following example, the two first categories belong to the same category group A, the sum of the weights is `1`. The third category belongs to a category group B, its weight is `1`.
```
[
  { "id": 59, "weight": "0.5" }, // category group A
  { "id": 33, "weight": "0.5" }, // category group A
  { "id": 65, "weight": "1" } // category group B
]
```


> ℹ️
> This endpoint requires the following scope: `transactions:all`

[Source Pennylane](https://pennylane.readme.io/reference/puttransactioncategories) · [Index du SDK](../operations.md)

## Appel Go

```go
PutTransactionCategoriesWithResponse(ctx context.Context, transactionId int, params *PutTransactionCategoriesParams, body PutTransactionCategoriesJSONRequestBody, reqEditors ...RequestEditorFn) (*PutTransactionCategoriesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `transaction_id` | `transactionId` | path | integer | obligatoire | <br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | array | racine | additionalProperties `false` |
| `$[]` | object | par élément | additionalProperties `false` |
| `$[].id` | integer (int64) | obligatoire | example `426` |
| `$[].weight` | string | obligatoire | A number between `0` and `1`, including `1`. It must have a maximum of 7 decimals. `0.85` means 85%.<br>example `"0.6575"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
[
  {
    "id": 426,
    "weight": "0.6575"
  }
]
```

</details>

## Sortie

### HTTP 200

The list of categories of the bank transaction

`application/json` → `réponse.JSON200` (type `*`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | array | racine |  |
| `$[]` | object | par élément | additionalProperties `false` |
| `$[].analytical_code` | string ou null | obligatoire | example `"CODE123"` |
| `$[].category_group` | object | obligatoire | additionalProperties `false` |
| `$[].category_group.id` | integer (int64) | obligatoire | example `229` |
| `$[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$[].id` | integer (int64) | obligatoire | example `421` |
| `$[].label` | string | obligatoire | example `"HR - Salaries"` |
| `$[].ledger_account` | object ou null | obligatoire | The ledger account the category is wired to, or `null` when none is set.<br>additionalProperties `false` |
| `$[].ledger_account.id` | integer (int64) | obligatoire | example `124` |
| `$[].ledger_account.url` | string | obligatoire | The URL to the ledger account<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/124"` |
| `$[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$[].weight` | string | obligatoire | example `"0.25"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
[
  {
    "analytical_code": "CODE123",
    "category_group": {
      "id": 229
    },
    "created_at": "2023-08-30T10:08:08.146343Z",
    "id": 421,
    "label": "HR - Salaries",
    "ledger_account": {
      "id": 124,
      "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/124"
    },
    "updated_at": "2023-08-30T10:08:08.146343Z",
    "weight": "0.25"
  }
]
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
