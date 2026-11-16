# Link Analytical Categories to a Ledger Entry line

`PUT /api/external/v2/ledger_entry_lines/{ledger_entry_line_id}/categories`

This endpoint replaces already existing categories on the Ledger Entry line with new values.
If an empty array of categories_ids is provided, it will remove all categories from the Ledger Entry line.


> ℹ️
> This endpoint requires the following scope: `ledger_entries:all`

[Source Pennylane](https://pennylane.readme.io/reference/putledgerentrylinescategories) · [Index du SDK](../operations.md)

## Appel Go

```go
PutLedgerEntryLinesCategoriesWithResponse(ctx context.Context, ledgerEntryLineId float32, params *PutLedgerEntryLinesCategoriesParams, body PutLedgerEntryLinesCategoriesJSONRequestBody, reqEditors ...RequestEditorFn) (*PutLedgerEntryLinesCategoriesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `ledger_entry_line_id` | `ledgerEntryLineId` | path | number | obligatoire | Existing Ledger Entry line (id) |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | array | racine |  |
| `$[]` | object | par élément | additionalProperties `false` |
| `$[].id` | integer (int64) | obligatoire | example `394` |
| `$[].weight` | string | obligatoire | A number between `0` and `1`, including `1`. It must have a maximum of 7 decimals. `0.85` means 85%.<br>example `"0.6575"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
[
  {
    "id": 394,
    "weight": "0.6575"
  }
]
```

</details>

## Sortie

### HTTP 200

Returns the Ledger Entry line with attached Analytical Categories

`application/json` → `réponse.JSON200` (type `*PutLedgerEntryLinesCategoriesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.ledger_entry_line` | object | obligatoire | additionalProperties `false` |
| `$.ledger_entry_line.categories` | array | obligatoire |  |
| `$.ledger_entry_line.categories[]` | object | par élément | additionalProperties `false` |
| `$.ledger_entry_line.categories[].analytical_code` | string ou null | obligatoire | example `"CODE123"` |
| `$.ledger_entry_line.categories[].category_group` | object | obligatoire | additionalProperties `false` |
| `$.ledger_entry_line.categories[].category_group.id` | integer (int64) | obligatoire | example `229` |
| `$.ledger_entry_line.categories[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.ledger_entry_line.categories[].id` | integer (int64) | obligatoire | example `421` |
| `$.ledger_entry_line.categories[].label` | string | obligatoire | example `"HR - Salaries"` |
| `$.ledger_entry_line.categories[].ledger_account` | object ou null | obligatoire | The ledger account the category is wired to, or `null` when none is set.<br>additionalProperties `false` |
| `$.ledger_entry_line.categories[].ledger_account.id` | integer (int64) | obligatoire | example `124` |
| `$.ledger_entry_line.categories[].ledger_account.url` | string | obligatoire | The URL to the ledger account<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/124"` |
| `$.ledger_entry_line.categories[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.ledger_entry_line.categories[].weight` | string | obligatoire | example `"0.25"` |
| `$.ledger_entry_line.id` | integer (int64) | obligatoire | example `2` |
| `$.ledger_entry_line.label` | string | obligatoire | example `"Employee - remuneration and contributions"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "ledger_entry_line": {
    "categories": [
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
    ],
    "id": 2,
    "label": "Employee - remuneration and contributions"
  }
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
