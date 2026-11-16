# List commercial documents

`GET /api/external/v2/commercial_documents`

This endpoint lists commercial documents.

> ℹ️
> This endpoint requires one of the following scopes: `commercial_documents:all`, `commercial_documents:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/listcommercialdocuments) · [Index du SDK](../operations.md)

## Appel Go

```go
ListCommercialDocumentsWithResponse(ctx context.Context, params *ListCommercialDocumentsParams, reqEditors ...RequestEditorFn) (*ListCommercialDocumentsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and operators:<br>- `id`, `customer_id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `document_type`: `eq`, `not_eq`, `in`, `not_in`<br><br>Available document_types:<br>- proforma<br>- shipping_order<br>- purchasing_order<br>example `"[{\"field\": \"id\", \"operator\": \"eq\", \"value\": \"42\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of Commercial Documents

`application/json` → `réponse.JSON200` (type `*ListCommercialDocumentsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].amount` | string | obligatoire | Document amount (total value of the document in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"230.32"` |
| `$.items[].appendices` | object | obligatoire | additionalProperties `false` |
| `$.items[].appendices.url` | string | obligatoire | URL to get the appendices of the document.<br>example `"https://app.pennylane.com/api/external/v2/commercial_documents/42/appendices"` |
| `$.items[].archived_at` | string (date-time) ou null | obligatoire | The time the document has been archived<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].commercial_document_template` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].commercial_document_template.id` | integer (int64) | obligatoire |  |
| `$.items[].created_at` | string (date-time) | obligatoire | The time the document has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.items[].currency_amount` | string | obligatoire | Document currency amount (total value of the document in the currency of the document)<br>example `"230.32"` |
| `$.items[].currency_amount_before_tax` | string | obligatoire | Document currency amount before tax (total value before tax of the document in the currency of the document)<br>example `"196.32"` |
| `$.items[].currency_tax` | string | obligatoire | Document taxable amount (in document currency)<br>example `"34.0"` |
| `$.items[].customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].customer.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.items[].date` | string (date) ou null | obligatoire | Document issue date (ISO 8601)<br>example `"2020-01-01"` |
| `$.items[].deadline` | string (date) ou null | obligatoire | Document deadline (ISO 8601)<br>example `"2020-12-31"` |
| `$.items[].discount` | object | obligatoire | additionalProperties `false` |
| `$.items[].discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.items[].discount.value` | string ou null | obligatoire | Discount value on the total amount before tax of the line<br>example `"25"` |
| `$.items[].document_number` | string | obligatoire | example `"PLERYGTXYOIMZZ"` |
| `$.items[].document_type` | string | obligatoire | enum `["proforma","shipping_order","purchasing_order"]` |
| `$.items[].exchange_rate` | string | obligatoire | Document exchange rate (used to convert the document to euros. If the document currency is euro it will be 1.0)<br>example `"1.0"` |
| `$.items[].external_reference` | string | obligatoire | The unique external reference that was assigned during creation either by you or Pennylane. (Same attribute as `external_id` in the API v1)<br>example `"FR123"` |
| `$.items[].filename` | string ou null | obligatoire | Name of the file attached to the document<br>example `"my_file.pdf"` |
| `$.items[].id` | integer (int64) | obligatoire | Commercial document identifier<br>example `42` |
| `$.items[].invoice_line_sections` | object | obligatoire | additionalProperties `false` |
| `$.items[].invoice_line_sections.url` | string | obligatoire | URL to get the invoice line sections of the document.<br>example `"https://app.pennylane.com/api/external/v2/commercial_documents/42/invoice_line_sections"` |
| `$.items[].invoice_lines` | object | obligatoire | additionalProperties `false` |
| `$.items[].invoice_lines.url` | string | obligatoire | URL to get the invoice lines of the document.<br>example `"https://app.pennylane.com/api/external/v2/commercial_documents/42/invoice_lines"` |
| `$.items[].label` | string ou null | obligatoire | example `"Commercial document label"` |
| `$.items[].language` | string | obligatoire | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.items[].pdf_description` | string ou null | obligatoire | example `"Document description"` |
| `$.items[].pdf_invoice_free_text` | string | obligatoire | example `"Document details"` |
| `$.items[].pdf_invoice_subject` | string | obligatoire | example `"Document subject"` |
| `$.items[].public_file_url` | string ou null | obligatoire | Public URL of the document file. The URL will expire after 30 minutes.<br>example `"https://app.pennylane.com/.../file.pdf"` |
| `$.items[].quote` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].quote.id` | integer (int64) | obligatoire | example `201` |
| `$.items[].quote.url` | string | obligatoire | URL to get the quote.<br>example `"https://app.pennylane.com/api/external/v2/quotes/201"` |
| `$.items[].special_mention` | string ou null | obligatoire | Additional details<br>example `"Additional details"` |
| `$.items[].tax` | string | obligatoire | Document taxable amount (in document currency)<br>example `"34.0"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | The last time the document has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
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
