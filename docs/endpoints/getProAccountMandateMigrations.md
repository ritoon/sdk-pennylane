# List mandate migration candidates

`GET /api/external/v2/pro_account/mandate_migrations`

This endpoint allows you to retrieve all mandate migration candidates
for your company. These are mandates that can be migrated to a Pro Account.

Requirements:
- Company must have a Pro Account (returns 404 if not)
- Company must have an enabled merchant profile (returns 403 if not)


> ℹ️
> This endpoint requires one of the following scopes: `customer_mandates:readonly`, `customer_mandates:all`

[Source Pennylane](https://pennylane.readme.io/reference/getproaccountmandatemigrations) · [Index du SDK](../operations.md)

## Appel Go

```go
GetProAccountMandateMigrationsWithResponse(ctx context.Context, params *GetProAccountMandateMigrationsParams, reqEditors ...RequestEditorFn) (*GetProAccountMandateMigrationsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `customer_id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `status`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"status\", \"operator\": \"eq\", \"value\": \"available\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes.<br>Sort field may be prefixed with `-` for descending order.<br>Example: `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields: `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

The list of mandate migration candidates

`application/json` → `réponse.JSON200` (type `*GetProAccountMandateMigrationsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].created_at` | string (date-time) | obligatoire | Creation date of the mandate migration candidate<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.items[].customer` | object ou null | obligatoire | The customer associated with this mandate<br>additionalProperties `false` |
| `$.items[].customer.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].customer.pro_account_mandate` | object | obligatoire | additionalProperties `false` |
| `$.items[].customer.pro_account_mandate.url` | string | obligatoire | Link to the pro account mandate for this customer.<br>example `"https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=..."` |
| `$.items[].customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.items[].direct_debit_method` | string | obligatoire | The direct debit method used by this mandate.<br>- `GoCardless`: GoCardless mandate<br>- `XML`: SEPA XML mandate<br>- `ProAccount`: Pro Account mandate<br><br>enum `["GoCardless","XML","ProAccount"]`<br>example `"GoCardless"` |
| `$.items[].error_message` | string ou null | obligatoire | Error message if the migration failed<br>example `null` |
| `$.items[].id` | integer (int64) | obligatoire | Numeric ID extracted from the display_id<br>example `42` |
| `$.items[].mandate` | object | obligatoire | A polymorphic mandate object can either be a SepaMandate or a Gocardless mandate.<br>additionalProperties `false` |
| `$.items[].mandate.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].mandate.type` | string | obligatoire | enum `["SepaMandate","Mandate"]`<br>example `"Mandate"` |
| `$.items[].migrated_at` | string (date-time) ou null | obligatoire | Date when the mandate was successfully migrated<br>example `null` |
| `$.items[].migration_started_at` | string (date-time) ou null | obligatoire | Date when the migration was started<br>example `null` |
| `$.items[].signed_at` | string (date-time) ou null | obligatoire | Date when the mandate was signed<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.items[].status` | string | obligatoire | Status of the mandate migration candidate.<br>- `available`: The mandate is available for migration<br>- `completed`: The mandate has been successfully migrated<br><br>enum `["available","completed"]`<br>example `"available"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | Last update date of the mandate migration candidate<br>example `"2023-08-07T14:23:12.000Z"` |
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
        "pro_account_mandate": {
          "url": "https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=..."
        },
        "url": "https://app.pennylane.com/api/external/v2/customers/42"
      },
      "direct_debit_method": "GoCardless",
      "error_message": null,
      "id": 42,
      "mandate": {
        "id": 42,
        "type": "Mandate"
      },
      "migrated_at": null,
      "migration_started_at": null,
      "signed_at": "2023-08-07T14:23:12.000Z",
      "status": "available",
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
