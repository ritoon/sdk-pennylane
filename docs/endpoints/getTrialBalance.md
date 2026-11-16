# Get the trial balance

`GET /api/external/v2/trial_balance`

This endpoint returns the trial balance of the current company
for the given period.

> ℹ️
> This endpoint requires the following scope: `trial_balance:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/gettrialbalance) · [Index du SDK](../operations.md)

## Appel Go

```go
GetTrialBalanceWithResponse(ctx context.Context, params *GetTrialBalanceParams, reqEditors ...RequestEditorFn) (*GetTrialBalanceResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `period_start` | `params.PeriodStart` | query | string (date) | obligatoire | The start of the period you want the trial balance for.<br>example `"2022-01-01"` |
| `period_end` | `params.PeriodEnd` | query | string (date) | obligatoire | The end of the period you want the trial balance for.<br>example `"2022-12-31"` |
| `is_auxiliary` | `params.IsAuxiliary` | query | boolean | facultatif | Whether to include auxiliary accounts or not.<br>example `true` |
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"MjAyNS0wMS0wOVQwODoyNDozOC44MTI0NTha"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 1000.<br>minimum `1`<br>maximum `1000` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a list of balances grouped by ledger accounts.

`application/json` → `réponse.JSON200` (type `*GetTrialBalanceStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].credits` | string | obligatoire | Ledger account credits<br>example `"2343.05"` |
| `$.items[].debits` | string | obligatoire | Ledger account debits<br>example `"45353.95"` |
| `$.items[].formatted_number` | string | obligatoire | Ledger account number with padded<br>example `"51200000"` |
| `$.items[].label` | string | obligatoire | Ledger account label<br>example `"Capital souscrit - appelé, versé"` |
| `$.items[].number` | string | obligatoire | Ledger account number<br>example `"512"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results. |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "credits": "2343.05",
      "debits": "45353.95",
      "formatted_number": "51200000",
      "label": "Capital souscrit - appelé, versé",
      "number": "512"
    }
  ]
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

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.
