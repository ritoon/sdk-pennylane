# List contacts of a customer

`GET /api/external/v2/customers/{customer_id}/contacts`

List contacts of a customer

> ℹ️
> This endpoint requires one of the following scopes: `customers:all`, `customers:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcustomercontacts) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCustomerContactsWithResponse(ctx context.Context, customerId int, params *GetCustomerContactsParams, reqEditors ...RequestEditorFn) (*GetCustomerContactsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_id` | `customerId` | path | integer | obligatoire | <br>example `42` |
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

The list of contacts of a customer

`application/json` → `réponse.JSON200` (type `*GetCustomerContactsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].email` | string | obligatoire | example `"jane.doe@example.com"` |
| `$.items[].first_name` | string | obligatoire | example `"Jane"` |
| `$.items[].id` | integer (int64) | obligatoire | example `421` |
| `$.items[].last_name` | string | obligatoire | example `"Doe"` |
| `$.items[].mobile_number` | string | obligatoire |  |
| `$.items[].role` | string | obligatoire | example `"accountant"` |
| `$.items[].telephone_number` | string | obligatoire |  |
| `$.items[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "created_at": "2023-08-30T10:08:08.146343Z",
      "email": "jane.doe@example.com",
      "first_name": "Jane",
      "id": 421,
      "last_name": "Doe",
      "role": "accountant",
      "updated_at": "2023-08-30T10:08:08.146343Z"
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
