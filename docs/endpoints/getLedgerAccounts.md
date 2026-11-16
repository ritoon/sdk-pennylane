# List Ledger Accounts

`GET /api/external/v2/ledger_accounts`

List Ledger Accounts


> ℹ️
> This endpoint requires one of the following scopes: `ledger_accounts:readonly`, `ledger_accounts:all`

[Source Pennylane](https://pennylane.readme.io/reference/getledgeraccounts) · [Index du SDK](../operations.md)

## Appel Go

```go
GetLedgerAccountsWithResponse(ctx context.Context, params *GetLedgerAccountsParams, reqEditors ...RequestEditorFn) (*GetLedgerAccountsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 1000.<br>minimum `1`<br>maximum `1000` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `number`: `start_with`, `eq`, `in`<br>- `enabled`: `eq`<br>example `"[{\"field\": \"number\", \"operator\": \"start_with\", \"value\": \"411\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields :<br>- `id`<br>example `"id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a list of ledger accounts


`application/json` → `réponse.JSON200` (type `*GetLedgerAccountsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean ou null | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].country_alpha2` | string | obligatoire | Ledger Account's country code (alpha2)<br>example `"FR"` |
| `$.items[].created_at` | string (date-time) | obligatoire | The time the ledger account has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].enabled` | boolean | obligatoire | example `true` |
| `$.items[].id` | integer (int64) | obligatoire | example `124` |
| `$.items[].label` | string | obligatoire | example `"Secondary Account"` |
| `$.items[].letterable` | boolean | obligatoire | example `false` |
| `$.items[].number` | string | obligatoire | example `"512"` |
| `$.items[].type` | string | obligatoire | example `"custom"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | The last time the ledger account has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].vat_rate` | string | obligatoire | Ledger Account's VAT rate in percentage<br>example `"any"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "country_alpha2": "FR",
      "created_at": "2023-08-30T10:08:08.146343Z",
      "enabled": true,
      "id": 124,
      "label": "Secondary Account",
      "letterable": false,
      "number": "512",
      "type": "custom",
      "updated_at": "2023-08-30T10:08:08.146343Z",
      "vat_rate": "any"
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
