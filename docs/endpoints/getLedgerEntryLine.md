# Retrieve a Ledger entry line

`GET /api/external/v2/ledger_entry_lines/{id}`

Retrieve a ledger entry line

> ℹ️
> This endpoint requires one of the following scopes: `ledger_entries:readonly`, `ledger_entries:all`

[Source Pennylane](https://pennylane.readme.io/reference/getledgerentryline) · [Index du SDK](../operations.md)

## Appel Go

```go
GetLedgerEntryLineWithResponse(ctx context.Context, id int, params *GetLedgerEntryLineParams, reqEditors ...RequestEditorFn) (*GetLedgerEntryLineResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | The unique identifier of the ledger entry line<br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A ledger entry line

`application/json` → `réponse.JSON200` (type `*GetLedgerEntryLineStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.alternative_quantity` | string ou null | obligatoire | Quantity expressed in the alternative unit<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.alternative_unit` | string ou null | obligatoire | Alternative unit of the quantity (UN/ECE recommendation 20 code)<br>example `"LTR"` |
| `$.categories` | array | obligatoire |  |
| `$.categories[]` | object | par élément | additionalProperties `false` |
| `$.categories[].analytical_code` | string ou null | obligatoire | example `"CODE123"` |
| `$.categories[].category_group` | object | obligatoire | additionalProperties `false` |
| `$.categories[].category_group.id` | integer (int64) | obligatoire | example `229` |
| `$.categories[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.categories[].id` | integer (int64) | obligatoire | example `421` |
| `$.categories[].label` | string | obligatoire | example `"HR - Salaries"` |
| `$.categories[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.categories[].weight` | string | obligatoire | example `"0.25"` |
| `$.created_at` | string (date-time) | obligatoire | The time the entry line has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.credit` | string | obligatoire | Credit amount for the entry line<br>example `"0.00"` |
| `$.date` | string (date) | obligatoire | Date of the ledger entry line (ISO 8601)<br>example `"2023-08-30"` |
| `$.debit` | string | obligatoire | Debit amount for the entry line<br>example `"100.00"` |
| `$.id` | integer (int64) | obligatoire | ID of the entry line<br>example `42` |
| `$.journal` | object | obligatoire | additionalProperties `false` |
| `$.journal.id` | integer (int64) | obligatoire | Journal ID<br>example `123` |
| `$.journal.url` | string | obligatoire | The URL to the journal<br>example `"https://app.pennylane.com/api/external/v2/journals/42"` |
| `$.label` | string | obligatoire | Label that describes the entry line<br>example `"Employees - Jean Dupont / Salary May 2023"` |
| `$.ledger_account` | object | obligatoire | additionalProperties `false` |
| `$.ledger_account.id` | integer | obligatoire | Ledger account ID<br>example `987` |
| `$.ledger_account.number` | string | obligatoire | Ledger account number<br>example `"411000"` |
| `$.ledger_account.url` | string | obligatoire | The URL to the ledger account<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/42"` |
| `$.ledger_entry` | object | obligatoire | additionalProperties `false` |
| `$.ledger_entry.id` | integer (int64) | obligatoire | The ledger entry ID of the ledger entry line<br>example `123001` |
| `$.lettered_ledger_entry_lines` | object | obligatoire | Ledger entry lines that are lettered with this entry line.<br>additionalProperties `false` |
| `$.lettered_ledger_entry_lines.ids` | array | obligatoire | IDs of all ledger entry lines that share the same lettering, including the ID of this entry line itself. Will be empty if this entry line is not lettered.<br>example `[42,1271004]` |
| `$.lettered_ledger_entry_lines.ids[]` | integer | par élément |  |
| `$.lettered_ledger_entry_lines.url` | string | obligatoire | URL to fetch the lettered ledger entry lines.<br>example `"https://app.pennylane.com/api/external/v2/ledger_entry_lines/42/lettered_ledger_entry_lines"` |
| `$.quantity` | string ou null | obligatoire | Quantity of goods or services recorded on the entry line. `null` when no quantity is recorded.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.unit` | string ou null | obligatoire | Unit code as stored on the line (UN/ECE recommendation 20 code)<br>example `"KGM"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the entry line has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "alternative_quantity": "100.00",
  "alternative_unit": "LTR",
  "categories": [
    {
      "analytical_code": "CODE123",
      "category_group": {
        "id": 229
      },
      "created_at": "2023-08-30T10:08:08.146343Z",
      "id": 421,
      "label": "HR - Salaries",
      "updated_at": "2023-08-30T10:08:08.146343Z",
      "weight": "0.25"
    }
  ],
  "created_at": "2023-08-30T10:08:08.146343Z",
  "credit": "0.00",
  "date": "2023-08-30",
  "debit": "100.00",
  "id": 42,
  "journal": {
    "id": 123,
    "url": "https://app.pennylane.com/api/external/v2/journals/42"
  },
  "label": "Employees - Jean Dupont / Salary May 2023",
  "ledger_account": {
    "id": 987,
    "number": "411000",
    "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/42"
  },
  "ledger_entry": {
    "id": 123001
  },
  "lettered_ledger_entry_lines": {
    "ids": [
      42,
      1271004
    ],
    "url": "https://app.pennylane.com/api/external/v2/ledger_entry_lines/42/lettered_ledger_entry_lines"
  },
  "quantity": "100.00",
  "unit": "KGM",
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
