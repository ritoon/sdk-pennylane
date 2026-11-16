# List Pro Account payment mandates

`GET /api/external/v2/pro_account/mandates`

This endpoint allows you to retrieve all payment mandates associated
with your company's pro account.

Requirements:
- Company must have a Pro Account (returns 404 if not)
- Company must have an enabled merchant profile (returns 403 if not)


> ℹ️
> This endpoint requires one of the following scopes: `customer_mandates:readonly`, `customer_mandates:all`

[Source Pennylane](https://pennylane.readme.io/reference/getproaccountmandates) · [Index du SDK](../operations.md)

## Appel Go

```go
GetProAccountMandatesWithResponse(ctx context.Context, params *GetProAccountMandatesParams, reqEditors ...RequestEditorFn) (*GetProAccountMandatesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and operators:<br>- `status`: `eq`, `not_eq`, `in`, `not_in`<br>- `customer_id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"status\", \"operator\": \"eq\", \"value\": \"enabled\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes.<br>Sort field may be prefixed with `-` for descending order.<br>Example: `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields: `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of payment mandates

`application/json` → `réponse.JSON200` (type `*GetProAccountMandatesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | Payment mandate for Pro Account<br>additionalProperties `false` |
| `$.items[].active_billing_subscription` | boolean | obligatoire | Whether the customer has at least one active billing subscription<br>example `false` |
| `$.items[].created_at` | string (date-time) | obligatoire | Date and time when the mandate was created<br>example `"2024-01-15T10:30:00Z"` |
| `$.items[].customer` | object ou null | obligatoire | Associated customer information<br>additionalProperties `false` |
| `$.items[].customer.id` | integer (int64) | obligatoire | Customer ID<br>example `47334785` |
| `$.items[].customer.url` | string | obligatoire | URL to the customer resource<br>example `"https://api.example.com/api/external/v2/customers/47334785"` |
| `$.items[].early_execution_date_permitted` | boolean | obligatoire | Whether early execution date is permitted for this mandate<br>example `false` |
| `$.items[].pdf_url` | string ou null | obligatoire | URL to the PDF file of the mandate (only present for mandate type, null for mandate_request)<br>example `"https://example.com/mandate.pdf"` |
| `$.items[].signed_at` | string (date) ou null | obligatoire | Date when the mandate was signed<br>example `"2024-01-15"` |
| `$.items[].status` | string | obligatoire | Current status of the mandate<br>example `"enabled"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "active_billing_subscription": false,
      "created_at": "2024-01-15T10:30:00Z",
      "customer": {
        "id": 47334785,
        "url": "https://api.example.com/api/external/v2/customers/47334785"
      },
      "early_execution_date_permitted": false,
      "pdf_url": "https://example.com/mandate.pdf",
      "signed_at": "2024-01-15",
      "status": "enabled"
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
