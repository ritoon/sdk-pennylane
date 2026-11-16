# Retrieve a commercial document

`GET /api/external/v2/commercial_documents/{id}`

This endpoint retrieves a commercial document.

> ℹ️
> This endpoint requires one of the following scopes: `commercial_documents:all`, `commercial_documents:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcommercialdocument) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCommercialDocumentWithResponse(ctx context.Context, id int, params *GetCommercialDocumentParams, reqEditors ...RequestEditorFn) (*GetCommercialDocumentResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | <br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A Commercial Document

`application/json` → `réponse.JSON200` (type `*UpdateCommercialDocumentStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.amount` | string | obligatoire | Document amount (total value of the document in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"230.32"` |
| `$.appendices` | object | obligatoire | additionalProperties `false` |
| `$.appendices.url` | string | obligatoire | URL to get the appendices of the document.<br>example `"https://app.pennylane.com/api/external/v2/commercial_documents/42/appendices"` |
| `$.archived_at` | string (date-time) ou null | obligatoire | The time the document has been archived<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.commercial_document_template` | object ou null | obligatoire | additionalProperties `false` |
| `$.commercial_document_template.id` | integer (int64) | obligatoire |  |
| `$.created_at` | string (date-time) | obligatoire | The time the document has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.currency_amount` | string | obligatoire | Document currency amount (total value of the document in the currency of the document)<br>example `"230.32"` |
| `$.currency_amount_before_tax` | string | obligatoire | Document currency amount before tax (total value before tax of the document in the currency of the document)<br>example `"196.32"` |
| `$.currency_tax` | string | obligatoire | Document taxable amount (in document currency)<br>example `"34.0"` |
| `$.customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.customer.id` | integer (int64) | obligatoire | example `42` |
| `$.customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.date` | string (date) ou null | obligatoire | Document issue date (ISO 8601)<br>example `"2020-01-01"` |
| `$.deadline` | string (date) ou null | obligatoire | Document deadline (ISO 8601)<br>example `"2020-12-31"` |
| `$.discount` | object | obligatoire | additionalProperties `false` |
| `$.discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.discount.value` | string ou null | obligatoire | Discount value on the total amount before tax of the line<br>example `"25"` |
| `$.document_number` | string | obligatoire | example `"PLERYGTXYOIMZZ"` |
| `$.document_type` | string | obligatoire | enum `["proforma","shipping_order","purchasing_order"]` |
| `$.exchange_rate` | string | obligatoire | Document exchange rate (used to convert the document to euros. If the document currency is euro it will be 1.0)<br>example `"1.0"` |
| `$.external_reference` | string | obligatoire | The unique external reference that was assigned during creation either by you or Pennylane. (Same attribute as `external_id` in the API v1)<br>example `"FR123"` |
| `$.filename` | string ou null | obligatoire | Name of the file attached to the document<br>example `"my_file.pdf"` |
| `$.id` | integer (int64) | obligatoire | Commercial document identifier<br>example `42` |
| `$.invoice_line_sections` | object | obligatoire | additionalProperties `false` |
| `$.invoice_line_sections.url` | string | obligatoire | URL to get the invoice line sections of the document.<br>example `"https://app.pennylane.com/api/external/v2/commercial_documents/42/invoice_line_sections"` |
| `$.invoice_lines` | object | obligatoire | additionalProperties `false` |
| `$.invoice_lines.url` | string | obligatoire | URL to get the invoice lines of the document.<br>example `"https://app.pennylane.com/api/external/v2/commercial_documents/42/invoice_lines"` |
| `$.label` | string ou null | obligatoire | example `"Commercial document label"` |
| `$.language` | string | obligatoire | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.pdf_description` | string ou null | obligatoire | example `"Document description"` |
| `$.pdf_invoice_free_text` | string | obligatoire | example `"Document details"` |
| `$.pdf_invoice_subject` | string | obligatoire | example `"Document subject"` |
| `$.public_file_url` | string ou null | obligatoire | Public URL of the document file. The URL will expire after 30 minutes.<br>example `"https://app.pennylane.com/.../file.pdf"` |
| `$.quote` | object ou null | obligatoire | additionalProperties `false` |
| `$.quote.id` | integer (int64) | obligatoire | example `201` |
| `$.quote.url` | string | obligatoire | URL to get the quote.<br>example `"https://app.pennylane.com/api/external/v2/quotes/201"` |
| `$.special_mention` | string ou null | obligatoire | Additional details<br>example `"Additional details"` |
| `$.tax` | string | obligatoire | Document taxable amount (in document currency)<br>example `"34.0"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the document has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "amount": "230.32",
  "appendices": {
    "url": "https://app.pennylane.com/api/external/v2/commercial_documents/42/appendices"
  },
  "archived_at": "2023-08-30T10:08:08.146343Z",
  "created_at": "2023-08-30T10:08:08.146343Z",
  "currency_amount": "230.32",
  "currency_amount_before_tax": "196.32",
  "currency_tax": "34.0",
  "customer": {
    "id": 42,
    "url": "https://app.pennylane.com/api/external/v2/customers/42"
  },
  "date": "2020-01-01",
  "deadline": "2020-12-31",
  "discount": {
    "value": "25"
  },
  "document_number": "PLERYGTXYOIMZZ",
  "exchange_rate": "1.0",
  "external_reference": "FR123",
  "filename": "my_file.pdf",
  "id": 42,
  "invoice_line_sections": {
    "url": "https://app.pennylane.com/api/external/v2/commercial_documents/42/invoice_line_sections"
  },
  "invoice_lines": {
    "url": "https://app.pennylane.com/api/external/v2/commercial_documents/42/invoice_lines"
  },
  "label": "Commercial document label",
  "pdf_description": "Document description",
  "pdf_invoice_free_text": "Document details",
  "pdf_invoice_subject": "Document subject",
  "public_file_url": "https://app.pennylane.com/.../file.pdf",
  "quote": {
    "id": 201,
    "url": "https://app.pennylane.com/api/external/v2/quotes/201"
  },
  "special_mention": "Additional details",
  "tax": "34.0",
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
