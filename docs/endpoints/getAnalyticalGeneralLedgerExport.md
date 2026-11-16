# Retrieve an Analytical General Ledger export

`GET /api/external/v2/exports/analytical_general_ledgers/{id}`

The endpoint returns a specific Analytical General Ledger export. The export file is an xlsx file, using the in-line analytical mode.

> ℹ️
> This endpoint requires the following scope: `exports:agl`

[Source Pennylane](https://pennylane.readme.io/reference/getanalyticalgeneralledgerexport) · [Index du SDK](../operations.md)

## Appel Go

```go
GetAnalyticalGeneralLedgerExportWithResponse(ctx context.Context, id int, params *GetAnalyticalGeneralLedgerExportParams, reqEditors ...RequestEditorFn) (*GetAnalyticalGeneralLedgerExportResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | Existing export identifier (id)<br>example `1` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns the export

`application/json` → `réponse.JSON200` (type `*GetAnalyticalGeneralLedgerExportStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.file_url` | string ou null | obligatoire | URL to download the export file. The URL will expire after 30 minutes.<br>example `"https://app.pennylane.com/my_export.xlsx"` |
| `$.id` | integer (int64) | obligatoire | ID of the export<br>example `124` |
| `$.status` | string | obligatoire | The state of the export<br>enum `["pending","ready","error"]` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "created_at": "2023-08-30T10:08:08.146343Z",
  "file_url": "https://app.pennylane.com/my_export.xlsx",
  "id": 124,
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
