# Retrieve an individual customer

`GET /api/external/v2/individual_customers/{id}`

This endpoint returns an individual customer.

> ℹ️
> This endpoint requires one of the following scopes: `customers:all`, `customers:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getindividualcustomer) · [Index du SDK](../operations.md)

## Appel Go

```go
GetIndividualCustomerWithResponse(ctx context.Context, id string, params *GetIndividualCustomerParams, reqEditors ...RequestEditorFn) (*GetIndividualCustomerResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | string | obligatoire | Individual customer identifier |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns an individual customer

`application/json` → `réponse.JSON200` (type `*GetIndividualCustomerStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.billing_address` | object | obligatoire | additionalProperties `false` |
| `$.billing_address.address` | string | obligatoire |  |
| `$.billing_address.city` | string | obligatoire |  |
| `$.billing_address.country_alpha2` | string | obligatoire |  |
| `$.billing_address.postal_code` | string | obligatoire |  |
| `$.billing_iban` | string ou null | obligatoire | example `"FR1420041010050500013M02606"` |
| `$.billing_language` | string | obligatoire | The language in which the customer will receive invoices. Default is `fr_FR`<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.contacts` | object | obligatoire | additionalProperties `false` |
| `$.contacts.url` | string | obligatoire | URL to get the contacts of the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42/contacts"` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.delivery_address` | object | obligatoire | additionalProperties `false` |
| `$.delivery_address.address` | string | obligatoire |  |
| `$.delivery_address.city` | string | obligatoire |  |
| `$.delivery_address.country_alpha2` | string | obligatoire |  |
| `$.delivery_address.postal_code` | string | obligatoire |  |
| `$.emails` | array | obligatoire | example `["hello@example.org"]` |
| `$.emails[]` | string | par élément |  |
| `$.external_reference` | string | obligatoire | The unique external reference assigned to this customer, assigned on creation either by you or Pennylane. (Same attribute as `source_id` in the API v1)<br>example `"0e67fc3c-c632-4feb-ad34-e18ed5fbf66a"` |
| `$.first_name` | string | obligatoire | example `"John"` |
| `$.id` | integer (int64) | obligatoire | example `42` |
| `$.last_name` | string | obligatoire | example `"Doe"` |
| `$.ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$.ledger_account.id` | integer (int64) | obligatoire |  |
| `$.mandates` | object | obligatoire | GoCardless mandates associated with this customer<br>additionalProperties `false` |
| `$.mandates.url` | string | obligatoire | URL to list GoCardless mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/gocardless_mandates?filter=%5B%7B%5C\"field%5C\":%5C\"customer_id%5C\",%5C\"operator%5C\":%5C\"eq%5C\",%5C\"value%5C\":%5C\"42%5C\"%7D%5D"` |
| `$.name` | string | obligatoire | example `"John Doe"` |
| `$.notes` | string ou null | obligatoire | example `"Some notes"` |
| `$.payment_conditions` | string | obligatoire | Note that the `custom` option is only used on Pennylane's web app to avoid pre-filling the deadline when creating an invoice. On the API it has no effect and you will still have to provide a deadline when creating an invoice.<br>enum `["upon_receipt","custom","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$.phone` | string | obligatoire | example `"+33612345678"` |
| `$.pro_account_mandates` | object | obligatoire | Pro Account mandates associated with this customer<br>additionalProperties `false` |
| `$.pro_account_mandates.url` | string | obligatoire | URL to list Pro Account mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=%5B%7B%22field%22%3A%22customer_id%22%2C%22operator%22%3A%22eq%22%2C%22value%22%3A%2242%22%7D%5D"` |
| `$.recipient` | string | obligatoire | example `"John Doe"` |
| `$.reference` | string ou null | obligatoire | example `"REF-1234"` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "billing_iban": "FR1420041010050500013M02606",
  "contacts": {
    "url": "https://app.pennylane.com/api/external/v2/customers/42/contacts"
  },
  "created_at": "2023-08-30T10:08:08.146343Z",
  "emails": [
    "hello@example.org"
  ],
  "external_reference": "0e67fc3c-c632-4feb-ad34-e18ed5fbf66a",
  "first_name": "John",
  "id": 42,
  "last_name": "Doe",
  "mandates": {
    "url": "https://app.pennylane.com/api/external/v2/gocardless_mandates?filter=%5B%7B%5C\"field%5C\":%5C\"customer_id%5C\",%5C\"operator%5C\":%5C\"eq%5C\",%5C\"value%5C\":%5C\"42%5C\"%7D%5D"
  },
  "name": "John Doe",
  "notes": "Some notes",
  "phone": "+33612345678",
  "pro_account_mandates": {
    "url": "https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=%5B%7B%22field%22%3A%22customer_id%22%2C%22operator%22%3A%22eq%22%2C%22value%22%3A%2242%22%7D%5D"
  },
  "recipient": "John Doe",
  "reference": "REF-1234",
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
