# Create a General Ledger export

`POST /api/external/v2/exports/general_ledgers`

This endpoint allows you to create a General Ledger export. The generated export file is an xlsx file.

> ℹ️
> This endpoint requires the following scope: `exports:gl`

[Source Pennylane](https://pennylane.readme.io/reference/exportgeneralledger) · [Index du SDK](../operations.md)

## Appel Go

```go
ExportGeneralLedgerWithResponse(ctx context.Context, params *ExportGeneralLedgerParams, body ExportGeneralLedgerJSONRequestBody, reqEditors ...RequestEditorFn) (*ExportGeneralLedgerResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.period_end` | string (date) | obligatoire | End date of the period to export<br>example `"2023-08-30"` |
| `$.period_start` | string (date) | obligatoire | Start date of the period to export<br>example `"2023-08-30"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "period_end": "2023-08-30",
  "period_start": "2023-08-30"
}
```

</details>

## Sortie

### HTTP 201

Returns the generated export status

`application/json` → `réponse.JSON201` (type `*ExportAnalyticalGeneralLedgerStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.id` | integer (int64) | obligatoire | ID of the export<br>example `124` |
| `$.status` | string | obligatoire | The state of the export<br>enum `["pending","ready","error"]` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "created_at": "2023-08-30T10:08:08.146343Z",
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
