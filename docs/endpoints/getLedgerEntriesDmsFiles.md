# List DMS files of a Ledger Entry

`GET /api/external/v2/ledger_entries/{ledger_entry_id}/dms_files`

List DMS files attached to a Ledger Entry


> ℹ️
> This endpoint requires one of the following scopes: `ledger_entries:readonly`, `ledger_entries:all`

[Source Pennylane](https://pennylane.readme.io/reference/getledgerentriesdmsfiles) · [Index du SDK](../operations.md)

## Appel Go

```go
GetLedgerEntriesDmsFilesWithResponse(ctx context.Context, ledgerEntryId float32, params *GetLedgerEntriesDmsFilesParams, reqEditors ...RequestEditorFn) (*GetLedgerEntriesDmsFilesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `ledger_entry_id` | `ledgerEntryId` | path | number | obligatoire | Existing Ledger Entry (id) |
| `cursor` | `params.Cursor` | query | string | facultatif | Use this to fetch the next set of results. The cursor is an opaque string returned in the previous response's metadata. Leave empty for the first request.<br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns DMS files of requested Ledger Entry


`application/json` → `réponse.JSON200` (type `*GetLedgerEntriesDmsFilesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire | Array of DMS files of requested Ledger Entry |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].created_at` | string (date-time) | obligatoire | Date and time when the DMS file was created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].id` | integer (int64) | obligatoire | ID of the DMS file<br>example `42` |
| `$.items[].name` | string | obligatoire | Name of the DMS file<br>example `"Invoice_2024.pdf"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | Date and time when the DMS file was last updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].url` | string ou null | obligatoire | URL to download the DMS file.<br>This URL expires after 30 minutes.<br><br>example `"https://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/Invoice.pdf"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "created_at": "2023-08-30T10:08:08.146343Z",
      "id": 42,
      "name": "Invoice_2024.pdf",
      "updated_at": "2023-08-30T10:08:08.146343Z",
      "url": "https://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/Invoice.pdf"
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
