# Get ledger entry category change events

`GET /api/external/v2/changelogs/ledger_entries_categories`

Returns the list of changes based on the provided `start_date`.
If no `start_date` is provided it returns the oldest set of recorded changes.
Changes for the last 4 weeks are retained. The items will be returned using
`processed_at` in ASC order (oldest first).

A ledger entry category change indicates that a category was added to or removed from a
ledger entry (e.g. a customer invoice, a supplier invoice). Use `ledger_entry.id`
and `ledger_entry.type` to identify the underlying resource to refetch.


> ℹ️
> This endpoint requires one of the following scopes: `ledger_entries:readonly`, `ledger_entries:all`

[Source Pennylane](https://pennylane.readme.io/reference/getledgerentriescategorychanges) · [Index du SDK](../operations.md)

## Appel Go

```go
GetLedgerEntriesCategoryChangesWithResponse(ctx context.Context, params *GetLedgerEntriesCategoryChangesParams, reqEditors ...RequestEditorFn) (*GetLedgerEntriesCategoryChangesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results. The cursor is an opaque string returned in the previous response's metadata. Leave empty for the first request.<br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 1000.<br>minimum `1`<br>maximum `1000` |
| `start_date` | `params.StartDate` | query | string (date-time) | facultatif | Filter the changes based on the event date. The date should follow RFC3339 format. If no date is provided, the oldest changes will be returned. Changes for the last 4 weeks are retained, thus providing a `start_date` older than that will result in a 422 response. Providing both `start_date` and `cursor` parameters will result in a 400 response.<br>example `"2025-06-25T11:54:18.589480Z"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of ledger entry category changes

`application/json` → `réponse.JSON200` (type `*GetLedgerEntriesCategoryChangesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].created_at` | string (date-time) | obligatoire | Timestamp when the record was initially created |
| `$.items[].id` | integer (int64) | obligatoire | Unique identifier of the ledger entry category record |
| `$.items[].ledger_entry` | object | obligatoire | The ledger entry the category was applied to<br>additionalProperties `false` |
| `$.items[].ledger_entry.id` | integer (int64) | obligatoire | Identifier of the ledger entry the category was applied to |
| `$.items[].ledger_entry.type` | string | obligatoire | Type of the ledger entry the category was applied to |
| `$.items[].operation` | string | obligatoire | enum `["insert","update","delete"]` |
| `$.items[].processed_at` | string (date-time) | obligatoire | Timestamp when the event arrived in the change log pipeline |
| `$.items[].updated_at` | string (date-time) | obligatoire | Timestamp when the record was updated in the database (can vary due to data restoration) |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
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

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)

### HTTP 500

Internal server error

Aucun schéma de corps documenté.
