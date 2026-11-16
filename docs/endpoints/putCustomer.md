# Update a customer (company or individual)

`PUT /api/external/v2/customers/{id}`

This endpoint updates a customer. Only the fields provided in the
request body are updated. The customer type can be changed by providing the
`customer_type` field.

Some fields are exclusive to one customer type. Sending a field that does
not belong to the current customer type returns a 422.
- Company-exclusive: `name`, `vat_number`, `reg_no`
- Individual-exclusive: `first_name`, `last_name`

When changing `customer_type`, some fields of the previous type are cleared automatically.
See each field's description for details.


> ℹ️
> This endpoint requires the following scope: `customers:all`

[Source Pennylane](https://pennylane.readme.io/reference/putcustomer) · [Index du SDK](../operations.md)

## Appel Go

```go
PutCustomerWithResponse(ctx context.Context, id int, params *PutCustomerParams, body PutCustomerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutCustomerResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | Customer identifier |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.billing_address` | object ou null | facultatif | additionalProperties `false` |
| `$.billing_address.address` | string | facultatif | example `"8 rue de la paix"` |
| `$.billing_address.city` | string | facultatif | example `"Paris"` |
| `$.billing_address.country_alpha2` | string | facultatif | example `"FR"` |
| `$.billing_address.postal_code` | string | facultatif | example `"75002"` |
| `$.billing_iban` | string ou null | facultatif | example `"FR1420041010050500013M02606"` |
| `$.billing_language` | string | facultatif | The language in which the customer will receive invoices. Default is `fr_FR`<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.customer_type` | string | facultatif | The type of the customer. Changing this value clears the fields of the previous type automatically.<br>enum `["company","individual"]` |
| `$.delivery_address` | object ou null | facultatif | additionalProperties `false` |
| `$.delivery_address.address` | string | facultatif | example `"8 rue de la paix"` |
| `$.delivery_address.city` | string | facultatif | example `"Paris"` |
| `$.delivery_address.country_alpha2` | string | facultatif | example `"FR"` |
| `$.delivery_address.postal_code` | string | facultatif | example `"75002"` |
| `$.electronic_billing_address` | string ou null | facultatif | The electronic billing address of the company customer used for e-invoicing routing. Based on the customer's SIREN or SIRET. Use a value returned by the GET /api/external/v2/customers/{id}/electronic_billing_addresses endpoint. Sending null or an empty string leaves the current address unchanged.<br>example `"123456789_00012"` |
| `$.emails` | array | facultatif | example `["hello@example.org"]` |
| `$.emails[]` | string | par élément |  |
| `$.external_reference` | string | facultatif | example `"0e67fc3c-c632-4feb-ad34-e18ed5fbf66a"` |
| `$.first_name` | string | facultatif | Individual customers only. Sending this field for a company customer returns a 422. Existing value cleared automatically when switching to `company`.<br>maxLength `1000`<br>example `"John"` |
| `$.last_name` | string | facultatif | Individual customers only. Sending this field for a company customer returns a 422. Existing value cleared automatically when switching to `company`.<br>maxLength `1000`<br>example `"Doe"` |
| `$.name` | string | facultatif | Company customers only. Sending this field for an individual customer returns a 422.<br>example `"My company"` |
| `$.notes` | string ou null | facultatif | example `"Some notes"` |
| `$.payment_conditions` | string | facultatif | Note that the `custom` option is only used on Pennylane's web app to avoid pre-filling the deadline when creating an invoice. On the API it has no effect and you will still have to provide a deadline when creating an invoice.<br>enum `["upon_receipt","custom","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$.phone` | string | facultatif | example `"+33612345678"` |
| `$.recipient` | string | facultatif | The name of the person to whom the invoice is addressed<br>example `"John Doe"` |
| `$.reference` | string ou null | facultatif | example `"REF-1234"` |
| `$.reg_no` | string | facultatif | Company customers only. Sending this field for an individual customer returns a 422. Existing value cleared automatically when switching to `individual`.<br>example `"123456789"` |
| `$.vat_number` | string | facultatif | Company customers only. Sending this field for an individual customer returns a 422. Existing value cleared automatically when switching to `individual`.<br>example `"FR12345678901"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "billing_address": {
    "address": "8 rue de la paix",
    "city": "Paris",
    "country_alpha2": "FR",
    "postal_code": "75002"
  },
  "billing_iban": "FR1420041010050500013M02606",
  "delivery_address": {
    "address": "8 rue de la paix",
    "city": "Paris",
    "country_alpha2": "FR",
    "postal_code": "75002"
  },
  "electronic_billing_address": "123456789_00012",
  "emails": [
    "hello@example.org"
  ],
  "external_reference": "0e67fc3c-c632-4feb-ad34-e18ed5fbf66a",
  "first_name": "John",
  "last_name": "Doe",
  "name": "My company",
  "notes": "Some notes",
  "phone": "+33612345678",
  "recipient": "John Doe",
  "reference": "REF-1234",
  "reg_no": "123456789",
  "vat_number": "FR12345678901"
}
```

</details>

## Sortie

### HTTP 200

Returns the updated customer

`application/json` → `réponse.JSON200` (type `*GetCustomerStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | oneOf | racine |  |
| `$ (oneOf 1)` | object | selon variante | title `"Company customer"`<br>additionalProperties `false` |
| `$ (oneOf 1).billing_address` | object | obligatoire | additionalProperties `false` |
| `$ (oneOf 1).billing_address.address` | string | obligatoire |  |
| `$ (oneOf 1).billing_address.city` | string | obligatoire |  |
| `$ (oneOf 1).billing_address.country_alpha2` | string | obligatoire |  |
| `$ (oneOf 1).billing_address.postal_code` | string | obligatoire |  |
| `$ (oneOf 1).billing_iban` | string ou null | obligatoire | example `"FR1420041010050500013M02606"` |
| `$ (oneOf 1).billing_language` | string | obligatoire | The language in which the customer will receive invoices. Default is `fr_FR`<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$ (oneOf 1).contacts` | object | obligatoire | additionalProperties `false` |
| `$ (oneOf 1).contacts.url` | string | obligatoire | URL to get the contacts of the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42/contacts"` |
| `$ (oneOf 1).created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$ (oneOf 1).customer_type` | string | obligatoire | enum `["company"]` |
| `$ (oneOf 1).delivery_address` | object | obligatoire | additionalProperties `false` |
| `$ (oneOf 1).delivery_address.address` | string | obligatoire |  |
| `$ (oneOf 1).delivery_address.city` | string | obligatoire |  |
| `$ (oneOf 1).delivery_address.country_alpha2` | string | obligatoire |  |
| `$ (oneOf 1).delivery_address.postal_code` | string | obligatoire |  |
| `$ (oneOf 1).electronic_billing_address` | string ou null | obligatoire | The electronic billing address of the company customer used for e-invoicing routing. Based on the customer's SIREN or SIRET.<br>example `"123456789_00012"` |
| `$ (oneOf 1).emails` | array | obligatoire | example `["hello@example.org"]` |
| `$ (oneOf 1).emails[]` | string | par élément |  |
| `$ (oneOf 1).external_reference` | string | obligatoire | The unique external reference assigned to this customer, assigned on creation either by you or Pennylane.<br>example `"0e67fc3c-c632-4feb-ad34-e18ed5fbf66a"` |
| `$ (oneOf 1).id` | integer (int64) | obligatoire | example `42` |
| `$ (oneOf 1).ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$ (oneOf 1).ledger_account.id` | integer (int64) | obligatoire |  |
| `$ (oneOf 1).mandates` | object | obligatoire | GoCardless mandates associated with this customer<br>additionalProperties `false` |
| `$ (oneOf 1).mandates.url` | string | obligatoire | URL to list GoCardless mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/gocardless_mandates?filter=%5B%7B%5C\"field%5C\":%5C\"customer_id%5C\",%5C\"operator%5C\":%5C\"eq%5C\",%5C\"value%5C\":%5C\"42%5C\"%7D%5D"` |
| `$ (oneOf 1).name` | string | obligatoire | example `"My company"` |
| `$ (oneOf 1).notes` | string ou null | obligatoire | example `"Some notes"` |
| `$ (oneOf 1).payment_conditions` | string | obligatoire | Note that the `custom` option is only used on Pennylane's web app to avoid pre-filling the deadline when creating an invoice. On the API it has no effect and you will still have to provide a deadline when creating an invoice.<br>enum `["upon_receipt","custom","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$ (oneOf 1).phone` | string | obligatoire | example `"+33612345678"` |
| `$ (oneOf 1).pro_account_mandates` | object | obligatoire | Pro Account mandates associated with this customer<br>additionalProperties `false` |
| `$ (oneOf 1).pro_account_mandates.url` | string | obligatoire | URL to list Pro Account mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=%5B%7B%22field%22%3A%22customer_id%22%2C%22operator%22%3A%22eq%22%2C%22value%22%3A%2242%22%7D%5D"` |
| `$ (oneOf 1).public_administration` | boolean | obligatoire | Whether the company customer is a French public administration. Based on the public administration check when available, otherwise on the customer's SIREN.<br>example `false` |
| `$ (oneOf 1).recipient` | string | obligatoire | The name of the person to whom the invoice is addressed<br>example `"John Doe"` |
| `$ (oneOf 1).reference` | string ou null | obligatoire | example `"REF-1234"` |
| `$ (oneOf 1).reg_no` | string | obligatoire | example `"123456789"` |
| `$ (oneOf 1).updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$ (oneOf 1).vat_number` | string | obligatoire | example `"FR12345678901"` |
| `$ (oneOf 2)` | object | selon variante | title `"Individual customer"`<br>additionalProperties `false` |
| `$ (oneOf 2).billing_address` | object | obligatoire | additionalProperties `false` |
| `$ (oneOf 2).billing_address.address` | string | obligatoire |  |
| `$ (oneOf 2).billing_address.city` | string | obligatoire |  |
| `$ (oneOf 2).billing_address.country_alpha2` | string | obligatoire |  |
| `$ (oneOf 2).billing_address.postal_code` | string | obligatoire |  |
| `$ (oneOf 2).billing_iban` | string ou null | obligatoire | example `"FR1420041010050500013M02606"` |
| `$ (oneOf 2).billing_language` | string | obligatoire | The language in which the customer will receive invoices. Default is `fr_FR`<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$ (oneOf 2).contacts` | object | obligatoire | additionalProperties `false` |
| `$ (oneOf 2).contacts.url` | string | obligatoire | URL to get the contacts of the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42/contacts"` |
| `$ (oneOf 2).created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$ (oneOf 2).customer_type` | string | obligatoire | enum `["individual"]` |
| `$ (oneOf 2).delivery_address` | object | obligatoire | additionalProperties `false` |
| `$ (oneOf 2).delivery_address.address` | string | obligatoire |  |
| `$ (oneOf 2).delivery_address.city` | string | obligatoire |  |
| `$ (oneOf 2).delivery_address.country_alpha2` | string | obligatoire |  |
| `$ (oneOf 2).delivery_address.postal_code` | string | obligatoire |  |
| `$ (oneOf 2).emails` | array | obligatoire | example `["hello@example.org"]` |
| `$ (oneOf 2).emails[]` | string | par élément |  |
| `$ (oneOf 2).external_reference` | string | obligatoire | The unique external reference assigned to this customer, assigned on creation either by you or Pennylane. (Same attribute as `source_id` in the API v1)<br>example `"0e67fc3c-c632-4feb-ad34-e18ed5fbf66a"` |
| `$ (oneOf 2).first_name` | string | obligatoire | example `"John"` |
| `$ (oneOf 2).id` | integer (int64) | obligatoire | example `42` |
| `$ (oneOf 2).last_name` | string | obligatoire | example `"Doe"` |
| `$ (oneOf 2).ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$ (oneOf 2).ledger_account.id` | integer (int64) | obligatoire |  |
| `$ (oneOf 2).mandates` | object | obligatoire | GoCardless mandates associated with this customer<br>additionalProperties `false` |
| `$ (oneOf 2).mandates.url` | string | obligatoire | URL to list GoCardless mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/gocardless_mandates?filter=%5B%7B%5C\"field%5C\":%5C\"customer_id%5C\",%5C\"operator%5C\":%5C\"eq%5C\",%5C\"value%5C\":%5C\"42%5C\"%7D%5D"` |
| `$ (oneOf 2).name` | string | obligatoire | example `"John Doe"` |
| `$ (oneOf 2).notes` | string ou null | obligatoire | example `"Some notes"` |
| `$ (oneOf 2).payment_conditions` | string | obligatoire | Note that the `custom` option is only used on Pennylane's web app to avoid pre-filling the deadline when creating an invoice. On the API it has no effect and you will still have to provide a deadline when creating an invoice.<br>enum `["upon_receipt","custom","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$ (oneOf 2).phone` | string | obligatoire | example `"+33612345678"` |
| `$ (oneOf 2).pro_account_mandates` | object | obligatoire | Pro Account mandates associated with this customer<br>additionalProperties `false` |
| `$ (oneOf 2).pro_account_mandates.url` | string | obligatoire | URL to list Pro Account mandates associated with this customer<br>example `"https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=%5B%7B%22field%22%3A%22customer_id%22%2C%22operator%22%3A%22eq%22%2C%22value%22%3A%2242%22%7D%5D"` |
| `$ (oneOf 2).recipient` | string | obligatoire | example `"John Doe"` |
| `$ (oneOf 2).reference` | string ou null | obligatoire | example `"REF-1234"` |
| `$ (oneOf 2).updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

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

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)

### HTTP 503

The service is temporarily unavailable

`application/json` → `réponse.JSON503` (type `*ImportCustomerInvoicesStatus503`).

[Champs du payload d’erreur](erreurs.md#importcustomerinvoicesstatus503)
