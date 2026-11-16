# List SEPA mandates

`GET /api/external/v2/sepa_mandates`

This endpoint allows you to retrieve all SEPA mandates associated with your company

> ℹ️
> This endpoint requires one of the following scopes: `customer_mandates:all`, `customer_mandates:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getsepamandates) · [Index du SDK](../operations.md)

## Appel Go

```go
GetSepaMandatesWithResponse(ctx context.Context, params *GetSepaMandatesParams, reqEditors ...RequestEditorFn) (*GetSepaMandatesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `customer_id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"id\", \"operator\": \"eq\", \"value\": \"42\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes.<br>Sort field may be prefixed with `-` for descending order.<br>Example: `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields: `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of SEPA mandates

`application/json` → `réponse.JSON200` (type `*GetSepaMandatesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].bank` | string ou null | obligatoire | Name of the customer's bank<br>example `"BNP Paribas"` |
| `$.items[].bic` | string | obligatoire | Bank Identifier Code (BIC) of the customer's bank<br>example `"BNPAFRPP"` |
| `$.items[].created_at` | string (date-time) | obligatoire | Creation date of the SEPA mandate<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.items[].customer` | object | obligatoire | additionalProperties `false` |
| `$.items[].customer.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.items[].iban` | string | obligatoire | International Bank Account Number (IBAN) of the customer<br>example `"FR3612739000506556687647Z31"` |
| `$.items[].id` | integer (int64) | obligatoire | ID of the created SEPA mandate<br>example `12345` |
| `$.items[].identifier` | string | obligatoire | Unique identifier for the mandate<br>example `"K-02-2023-12345"` |
| `$.items[].sequence_type` | string | obligatoire | SEPA mandate sequence type that defines the payment process.<br>- `FRST`: First payment in a series of recurring payments<br>- `OOFF`: One-off payment<br>- `RCUR`: Recurring payment that is not the first payment<br>- `FNAL`: Final payment in a series of recurring payments<br><br>enum `["FRST","OOFF","RCUR","FNAL"]`<br>default `"RCUR"`<br>example `"RCUR"` |
| `$.items[].signed_at` | string (date) | obligatoire | Date when the mandate was signed<br>example `"2025-04-24"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | Last update date of the SEPA mandate<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "bank": "BNP Paribas",
      "bic": "BNPAFRPP",
      "created_at": "2023-08-07T14:23:12.000Z",
      "customer": {
        "id": 42,
        "url": "https://app.pennylane.com/api/external/v2/customers/42"
      },
      "iban": "FR3612739000506556687647Z31",
      "id": 12345,
      "identifier": "K-02-2023-12345",
      "sequence_type": "RCUR",
      "signed_at": "2025-04-24",
      "updated_at": "2023-08-07T14:23:12.000Z"
    }
  ],
  "next_cursor": "dXBkYXRlZF9hdDoxNjc0MTIzNDU2"
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
