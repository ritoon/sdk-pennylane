# List customers (company and individual)

`GET /api/external/v2/customers`

This endpoint returns a list of both company and individual customers

> ℹ️
> This endpoint requires one of the following scopes: `customers:all`, `customers:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcustomers) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCustomersWithResponse(ctx context.Context, params *GetCustomersParams, reqEditors ...RequestEditorFn) (*GetCustomersResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `customer_type`: `eq`, `not_eq`<br>- `ledger_account_id`: `eq`, `not_eq`, `in`, `not_in`<br>- `name`: `start_with`, `eq`<br>- `external_reference`: `start_with`, `eq`, `not_eq`, `in`, `not_in`<br>- `reg_no`: `eq`, `not_eq`, `in`, `not_in`<br>- `emails`: `in`<br>example `"[{\"field\": \"ledger_account_id\", \"operator\": \"eq\", \"value\": \"123\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a list of both company and individual customers

`application/json` → `réponse.JSON200` (type `*GetCustomersStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | oneOf | par élément |  |
| `$.items[] (oneOf 1)` | object | selon variante | title `"Company customer"`<br>additionalProperties `false` |
| `$.items[] (oneOf 1).billing_address` | object | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 1).billing_address.address` | string | obligatoire |  |
| `$.items[] (oneOf 1).billing_address.city` | string | obligatoire |  |
| `$.items[] (oneOf 1).billing_address.country_alpha2` | string | obligatoire |  |
| `$.items[] (oneOf 1).billing_address.postal_code` | string | obligatoire |  |
| `$.items[] (oneOf 1).billing_iban` | string ou null | obligatoire | example `"FR1420041010050500013M02606"` |
| `$.items[] (oneOf 1).billing_language` | string | obligatoire | The language in which the customer will receive invoices. Default is `fr_FR`<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.items[] (oneOf 1).contacts` | object | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 1).contacts.url` | string | obligatoire | URL to get the contacts of the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42/contacts"` |
| `$.items[] (oneOf 1).created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[] (oneOf 1).customer_type` | string | obligatoire | enum `["company"]` |
| `$.items[] (oneOf 1).delivery_address` | object | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 1).delivery_address.address` | string | obligatoire |  |
| `$.items[] (oneOf 1).delivery_address.city` | string | obligatoire |  |
| `$.items[] (oneOf 1).delivery_address.country_alpha2` | string | obligatoire |  |
| `$.items[] (oneOf 1).delivery_address.postal_code` | string | obligatoire |  |
| `$.items[] (oneOf 1).electronic_billing_address` | string ou null | obligatoire | The electronic billing address of the company customer used for e-invoicing routing. Based on the customer's SIREN or SIRET.<br>example `"123456789_00012"` |
| `$.items[] (oneOf 1).emails` | array | obligatoire | example `["hello@example.org"]` |
| `$.items[] (oneOf 1).emails[]` | string | par élément |  |
| `$.items[] (oneOf 1).external_reference` | string | obligatoire | The unique external reference assigned to this customer, assigned on creation either by you or Pennylane.<br>example `"0e67fc3c-c632-4feb-ad34-e18ed5fbf66a"` |
| `$.items[] (oneOf 1).id` | integer (int64) | obligatoire | example `42` |
| `$.items[] (oneOf 1).ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 1).ledger_account.id` | integer (int64) | obligatoire |  |
| `$.items[] (oneOf 1).mandates` | object | obligatoire | GoCardless mandates associated with this customer<br>additionalProperties `false` |
| `$.items[] (oneOf 1).mandates.url` | string | obligatoire | URL to list GoCardless mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/gocardless_mandates?filter=%5B%7B%5C\"field%5C\":%5C\"customer_id%5C\",%5C\"operator%5C\":%5C\"eq%5C\",%5C\"value%5C\":%5C\"42%5C\"%7D%5D"` |
| `$.items[] (oneOf 1).name` | string | obligatoire | example `"My company"` |
| `$.items[] (oneOf 1).notes` | string ou null | obligatoire | example `"Some notes"` |
| `$.items[] (oneOf 1).payment_conditions` | string | obligatoire | Note that the `custom` option is only used on Pennylane's web app to avoid pre-filling the deadline when creating an invoice. On the API it has no effect and you will still have to provide a deadline when creating an invoice.<br>enum `["upon_receipt","custom","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$.items[] (oneOf 1).phone` | string | obligatoire | example `"+33612345678"` |
| `$.items[] (oneOf 1).pro_account_mandates` | object | obligatoire | Pro Account mandates associated with this customer<br>additionalProperties `false` |
| `$.items[] (oneOf 1).pro_account_mandates.url` | string | obligatoire | URL to list Pro Account mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=%5B%7B%22field%22%3A%22customer_id%22%2C%22operator%22%3A%22eq%22%2C%22value%22%3A%2242%22%7D%5D"` |
| `$.items[] (oneOf 1).public_administration` | boolean | obligatoire | Whether the company customer is a French public administration. Based on the public administration check when available, otherwise on the customer's SIREN.<br>example `false` |
| `$.items[] (oneOf 1).recipient` | string | obligatoire | The name of the person to whom the invoice is addressed<br>example `"John Doe"` |
| `$.items[] (oneOf 1).reference` | string ou null | obligatoire | example `"REF-1234"` |
| `$.items[] (oneOf 1).reg_no` | string | obligatoire | example `"123456789"` |
| `$.items[] (oneOf 1).updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[] (oneOf 1).vat_number` | string | obligatoire | example `"FR12345678901"` |
| `$.items[] (oneOf 2)` | object | selon variante | title `"Individual customer"`<br>additionalProperties `false` |
| `$.items[] (oneOf 2).billing_address` | object | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 2).billing_address.address` | string | obligatoire |  |
| `$.items[] (oneOf 2).billing_address.city` | string | obligatoire |  |
| `$.items[] (oneOf 2).billing_address.country_alpha2` | string | obligatoire |  |
| `$.items[] (oneOf 2).billing_address.postal_code` | string | obligatoire |  |
| `$.items[] (oneOf 2).billing_iban` | string ou null | obligatoire | example `"FR1420041010050500013M02606"` |
| `$.items[] (oneOf 2).billing_language` | string | obligatoire | The language in which the customer will receive invoices. Default is `fr_FR`<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.items[] (oneOf 2).contacts` | object | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 2).contacts.url` | string | obligatoire | URL to get the contacts of the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42/contacts"` |
| `$.items[] (oneOf 2).created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[] (oneOf 2).customer_type` | string | obligatoire | enum `["individual"]` |
| `$.items[] (oneOf 2).delivery_address` | object | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 2).delivery_address.address` | string | obligatoire |  |
| `$.items[] (oneOf 2).delivery_address.city` | string | obligatoire |  |
| `$.items[] (oneOf 2).delivery_address.country_alpha2` | string | obligatoire |  |
| `$.items[] (oneOf 2).delivery_address.postal_code` | string | obligatoire |  |
| `$.items[] (oneOf 2).emails` | array | obligatoire | example `["hello@example.org"]` |
| `$.items[] (oneOf 2).emails[]` | string | par élément |  |
| `$.items[] (oneOf 2).external_reference` | string | obligatoire | The unique external reference assigned to this customer, assigned on creation either by you or Pennylane. (Same attribute as `source_id` in the API v1)<br>example `"0e67fc3c-c632-4feb-ad34-e18ed5fbf66a"` |
| `$.items[] (oneOf 2).first_name` | string | obligatoire | example `"John"` |
| `$.items[] (oneOf 2).id` | integer (int64) | obligatoire | example `42` |
| `$.items[] (oneOf 2).last_name` | string | obligatoire | example `"Doe"` |
| `$.items[] (oneOf 2).ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[] (oneOf 2).ledger_account.id` | integer (int64) | obligatoire |  |
| `$.items[] (oneOf 2).mandates` | object | obligatoire | GoCardless mandates associated with this customer<br>additionalProperties `false` |
| `$.items[] (oneOf 2).mandates.url` | string | obligatoire | URL to list GoCardless mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/gocardless_mandates?filter=%5B%7B%5C\"field%5C\":%5C\"customer_id%5C\",%5C\"operator%5C\":%5C\"eq%5C\",%5C\"value%5C\":%5C\"42%5C\"%7D%5D"` |
| `$.items[] (oneOf 2).name` | string | obligatoire | example `"John Doe"` |
| `$.items[] (oneOf 2).notes` | string ou null | obligatoire | example `"Some notes"` |
| `$.items[] (oneOf 2).payment_conditions` | string | obligatoire | Note that the `custom` option is only used on Pennylane's web app to avoid pre-filling the deadline when creating an invoice. On the API it has no effect and you will still have to provide a deadline when creating an invoice.<br>enum `["upon_receipt","custom","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$.items[] (oneOf 2).phone` | string | obligatoire | example `"+33612345678"` |
| `$.items[] (oneOf 2).pro_account_mandates` | object | obligatoire | Pro Account mandates associated with this customer<br>additionalProperties `false` |
| `$.items[] (oneOf 2).pro_account_mandates.url` | string | obligatoire | URL to list Pro Account mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=%5B%7B%22field%22%3A%22customer_id%22%2C%22operator%22%3A%22eq%22%2C%22value%22%3A%2242%22%7D%5D"` |
| `$.items[] (oneOf 2).recipient` | string | obligatoire | example `"John Doe"` |
| `$.items[] (oneOf 2).reference` | string ou null | obligatoire | example `"REF-1234"` |
| `$.items[] (oneOf 2).updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
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
