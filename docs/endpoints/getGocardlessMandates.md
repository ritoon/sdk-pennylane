# List gocardless mandates

`GET /api/external/v2/gocardless_mandates`

List gocardless mandates

> ℹ️
> This endpoint requires one of the following scopes: `customer_mandates:all`, `customer_mandates:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getgocardlessmandates) · [Index du SDK](../operations.md)

## Appel Go

```go
GetGocardlessMandatesWithResponse(ctx context.Context, params *GetGocardlessMandatesParams, reqEditors ...RequestEditorFn) (*GetGocardlessMandatesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `customer_id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `external_reference`: `eq`<br>example `"[{\"field\": \"id\", \"operator\": \"eq\", \"value\": \"42\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes.<br>Sort field may be prefixed with `-` for descending order.<br>Example: `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields: `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

The list of gocardless mandates

`application/json` → `réponse.JSON200` (type `*GetGocardlessMandatesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].created_at` | string (date-time) | obligatoire | Creation date of the Gocardless mandate<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.items[].customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].customer.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.items[].external_customer_account` | string | obligatoire | Identifier for the Gocardless customer account.<br>example `"test_account_id"` |
| `$.items[].external_customer_label` | string ou null | obligatoire | Label for the Gocardless customer account.<br>example `"test_label"` |
| `$.items[].external_reference` | string ou null | obligatoire | Unique external identifier for the Gocardless mandate.<br>example `"MD000VDD0HN1RP"` |
| `$.items[].id` | integer (int64) | obligatoire | example `42` |
| `$.items[].status` | string | obligatoire | enum `["pending_customer_approval","pending_submission","submitted","active","failed","cancelled","expired","consumed","replaced","blocked","bank_disconnected"]` |
| `$.items[].updated_at` | string (date-time) | obligatoire | Last update date of the Gocardless mandate<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "created_at": "2023-08-07T14:23:12.000Z",
      "customer": {
        "id": 42,
        "url": "https://app.pennylane.com/api/external/v2/customers/42"
      },
      "external_customer_account": "test_account_id",
      "external_customer_label": "test_label",
      "external_reference": "MD000VDD0HN1RP",
      "id": 42,
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
