# List ledger entry lines of a Ledger Entry

`GET /api/external/v2/ledger_entries/{ledger_entry_id}/ledger_entry_lines`

List ledger entry lines of a Ledger Entry


> ℹ️
> This endpoint requires one of the following scopes: `ledger_entries:readonly`, `ledger_entries:all`

[Source Pennylane](https://pennylane.readme.io/reference/getledgerentriesledgerentrylines) · [Index du SDK](../operations.md)

## Appel Go

```go
GetLedgerEntriesLedgerEntryLinesWithResponse(ctx context.Context, ledgerEntryId float32, params *GetLedgerEntriesLedgerEntryLinesParams, reqEditors ...RequestEditorFn) (*GetLedgerEntriesLedgerEntryLinesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `ledger_entry_id` | `ledgerEntryId` | path | number | obligatoire | Existing Ledger Entry (id) |
| `cursor` | `params.Cursor` | query | string | facultatif | Use this to fetch the next set of results. The cursor is an opaque string returned in the previous response's metadata. Leave empty for the first request.<br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields : `ledger_account_id`<br>Available operators : `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"ledger_account_id\", \"operator\": \"eq\", \"value\": \"42\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields :<br>- `id`<br>example `"id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns Ledger Entry lines of requested Ledger Entry


`application/json` → `réponse.JSON200` (type `*GetLedgerEntriesLedgerEntryLinesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean ou null | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire | Array of entry lines of requested Ledger Entry |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].alternative_quantity` | string ou null | obligatoire | Quantity expressed in the alternative unit<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.items[].alternative_unit` | string ou null | obligatoire | Alternative unit of the quantity (UN/ECE recommendation 20 code)<br>example `"LTR"` |
| `$.items[].credit` | string | obligatoire | Credit amount for the entry line<br>example `"0.00"` |
| `$.items[].debit` | string | obligatoire | Debit amount for the entry line<br>example `"100.00"` |
| `$.items[].id` | integer (int64) | obligatoire | ID of the entry line<br>example `42` |
| `$.items[].label` | string | obligatoire | Label that describes the entry line<br>example `"Employees - Jean Dupont / Salary May 2023"` |
| `$.items[].ledger_account` | object | obligatoire | additionalProperties `false` |
| `$.items[].ledger_account.id` | integer (int64) | obligatoire | Ledger account ID<br>example `987` |
| `$.items[].ledger_account.number` | string | obligatoire | Ledger account number<br>example `"411000"` |
| `$.items[].ledger_account.url` | string | obligatoire | The URL to the ledger account<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/42"` |
| `$.items[].ledger_account_id` | integer (int64) | obligatoire | This attribute is deprecated. Please use `id` from `ledger_account` instead.<br>deprecated `true`<br>example `987` |
| `$.items[].quantity` | string ou null | obligatoire | Quantity of goods or services recorded on the entry line. `null` when no quantity is recorded.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.items[].unit` | string ou null | obligatoire | Unit code as stored on the line (UN/ECE recommendation 20 code)<br>example `"KGM"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "alternative_quantity": "100.00",
      "alternative_unit": "LTR",
      "credit": "0.00",
      "debit": "100.00",
      "id": 42,
      "label": "Employees - Jean Dupont / Salary May 2023",
      "ledger_account": {
        "id": 987,
        "number": "411000",
        "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/42"
      },
      "ledger_account_id": 987,
      "quantity": "100.00",
      "unit": "KGM"
    }
  ],
  "next_cursor": "MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"
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
