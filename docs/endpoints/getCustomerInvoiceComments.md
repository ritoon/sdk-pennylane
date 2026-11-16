# List comments of a customer invoice

`GET /api/external/v2/customer_invoices/{customer_invoice_id}/comments`

List the comments of a customer invoice, newest first. These are the
comments shown in the **Comments** section of the invoice in Pennylane.

Mentions of users are rendered as `@First name Last name`.
Listing comments does not mark them as read.


> ℹ️
> This endpoint requires one of the following scopes: `customer_invoices:all`, `customer_invoices:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcustomerinvoicecomments) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCustomerInvoiceCommentsWithResponse(ctx context.Context, customerInvoiceId int64, params *GetCustomerInvoiceCommentsParams, reqEditors ...RequestEditorFn) (*GetCustomerInvoiceCommentsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_invoice_id` | `customerInvoiceId` | path | integer (int64) | obligatoire | The unique identifier of the customer invoice.<br>example `42` |
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order. By default uses created_at in descending order.<br>Example : `created_at` will sort by ascending order, `-created_at` will sort by descending order.<br>Keep the same value when you request the next pages with the cursor.<br>Available fields :<br>- `created_at`<br>default `"-created_at"`<br>example `"created_at"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

The list of comments of the customer invoice

`application/json` → `réponse.JSON200` (type `*GetCustomerInvoiceCommentsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].content` | string | obligatoire | Plain text of the comment.<br>example `"Reminder sent to the customer on October 14."` |
| `$.items[].created_at` | string (date-time) | obligatoire | The time the comment was added<br>example `"2026-10-14T16:00:00.000000Z"` |
| `$.items[].id` | integer (int64) | obligatoire | Comment identifier<br>example `9001` |
| `$.items[].updated_at` | string (date-time) | obligatoire | The last time the comment has been updated, for example when its author edits it in Pennylane.<br>example `"2026-10-14T16:00:00.000000Z"` |
| `$.items[].user` | object ou null | obligatoire | Author of the comment. `null` when the comment was not written by a user.<br>additionalProperties `false` |
| `$.items[].user.first_name` | string | obligatoire | First name of the author<br>example `"Jane"` |
| `$.items[].user.last_name` | string | obligatoire | Last name of the author<br>example `"Doe"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "content": "Reminder sent to the customer on October 14.",
      "created_at": "2026-10-14T16:00:00.000000Z",
      "id": 9001,
      "updated_at": "2026-10-14T16:00:00.000000Z",
      "user": {
        "first_name": "Jane",
        "last_name": "Doe"
      }
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
