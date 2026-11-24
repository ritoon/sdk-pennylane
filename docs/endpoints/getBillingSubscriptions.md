# List billing subscriptions

`GET /api/external/v2/billing_subscriptions`

This endpoint returns a list of subscriptions.

> ℹ️
> This endpoint requires one of the following scopes: `billing_subscriptions:all`, `billing_subscriptions:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getbillingsubscriptions) · [Index du SDK](../operations.md)

## Appel Go

```go
GetBillingSubscriptionsWithResponse(ctx context.Context, params *GetBillingSubscriptionsParams, reqEditors ...RequestEditorFn) (*GetBillingSubscriptionsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`, `start`, `customer_id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `status`: `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"status\", \"operator\": \"eq\", \"value\": \"in_progress\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"`<br>example `"id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a list of billing subscriptions

`application/json` → `réponse.JSON200` (type `*GetBillingSubscriptionsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].activated_at` | string (date-time) ou null | obligatoire | The exact timestamp for when the subscription was activated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].created_at` | string (date-time) | obligatoire | The time the subscription has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].customer.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.items[].customer_invoice_data` | object | obligatoire | additionalProperties `false` |
| `$.items[].customer_invoice_data.amount` | string | obligatoire | Invoice amount (total value of the invoice in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"230.32"` |
| `$.items[].customer_invoice_data.currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.items[].customer_invoice_data.currency_amount` | string | obligatoire | Invoice currency amount (total value of the invoice in the currency of the invoice)<br>example `"230.32"` |
| `$.items[].customer_invoice_data.currency_amount_before_tax` | string | obligatoire | Invoice currency amount before tax (total value before tax of the invoice in the currency of the invoice)<br>example `"196.32"` |
| `$.items[].customer_invoice_data.currency_tax` | string | obligatoire | Invoice taxable amount (in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"34.0"` |
| `$.items[].customer_invoice_data.customer_invoice_template` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].customer_invoice_data.customer_invoice_template.id` | integer (int64) | obligatoire |  |
| `$.items[].customer_invoice_data.discount` | object | obligatoire | additionalProperties `false` |
| `$.items[].customer_invoice_data.discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.items[].customer_invoice_data.discount.value` | string ou null | obligatoire | Discount value on the total amount before tax<br>example `"25"` |
| `$.items[].customer_invoice_data.exchange_rate` | string | obligatoire | Invoice exchange rate (used to convert the invoice to euros. If the invoice currency is euro it will be 1.0)<br>pattern `"^-?\\d+(\\.\\d+)?$"`<br>example `"100.00"` |
| `$.items[].customer_invoice_data.invoice_line_sections` | object | obligatoire | additionalProperties `false` |
| `$.items[].customer_invoice_data.invoice_line_sections.url` | string | obligatoire | URL to get the invoice line sections of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/billing_subscriptions/42/invoice_line_sections"` |
| `$.items[].customer_invoice_data.invoice_lines` | object | obligatoire | additionalProperties `false` |
| `$.items[].customer_invoice_data.invoice_lines.url` | string | obligatoire | URL to get the invoice lines of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/billing_subscriptions/42/invoice_lines"` |
| `$.items[].customer_invoice_data.label` | string ou null | obligatoire | Invoice label<br>example `"Demo label"` |
| `$.items[].customer_invoice_data.language` | string | obligatoire | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.items[].customer_invoice_data.pdf_description` | string ou null | obligatoire | example `"Invoice description"` |
| `$.items[].customer_invoice_data.pdf_invoice_free_text` | string | obligatoire | example `"Thanks for paying this invoice"` |
| `$.items[].customer_invoice_data.pdf_invoice_subject` | string | obligatoire | example `"Invoice subject"` |
| `$.items[].customer_invoice_data.special_mention` | string ou null | obligatoire | Additional details<br>example `"Additional details"` |
| `$.items[].email_settings` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].email_settings.billing_email_template` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].email_settings.billing_email_template.id` | integer (int64) | obligatoire |  |
| `$.items[].email_settings.billing_email_template.label` | string | obligatoire |  |
| `$.items[].email_settings.recipients` | array | obligatoire | Primary email recipients for invoice delivery |
| `$.items[].email_settings.recipients[]` | string (email) | par élément |  |
| `$.items[].finish` | string (date) ou null | obligatoire | The date of the last occurrence of the subscription<br>example `"2023-12-31"` |
| `$.items[].id` | integer (int64) | obligatoire | Billing subscription identifier<br>example `42` |
| `$.items[].label` | string ou null | obligatoire |  |
| `$.items[].mode` | string | obligatoire | - `awaiting_validation`: generated invoices will be in draft<br>- `finalized`: generated invoices will be finalized<br>- `email`: generated invoices will be finalized and sent by email to the recipients configured on the subscription<br><br>enum `["awaiting_validation","finalized","email"]` |
| `$.items[].next_occurrence` | string (date) ou null | obligatoire | The date for the next subscription renewal<br>example `"2023-01-01"` |
| `$.items[].payment_conditions` | oneOf | obligatoire |  |
| `$.items[].payment_conditions (oneOf 1)` | string | selon variante | enum `["upon_receipt","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$.items[].payment_conditions (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.items[].payment_method` | string | obligatoire | Payment method<br>offline means the subscription is not linked to a payment method<br>gocardless_direct-debit means at each new occurrence the client will be automatically debited thanks to GoCardless. To do so, you need a GoCardless account properly configured.<br>pro_account_sepa_core means the customer must have an email, unless it already has an enabled SEPA Core mandate or a pending SEPA Core mandate request.<br><br>enum `["offline","gocardless_direct_debit","pro_account_sepa_core"]` |
| `$.items[].prev_occurrence` | string (date) ou null | obligatoire | The date of the previous subscription renewal<br>example `"2023-01-01"` |
| `$.items[].recurring_rule` | object | obligatoire | additionalProperties `false` |
| `$.items[].recurring_rule.count` | integer ou null | obligatoire | Number of occurrences of the recurring rule<br>example `12` |
| `$.items[].recurring_rule.day` | array ou null | obligatoire |  |
| `$.items[].recurring_rule.day[]` | number | par élément |  |
| `$.items[].recurring_rule.day_of_month` | array ou null | obligatoire |  |
| `$.items[].recurring_rule.day_of_month[]` | number | par élément |  |
| `$.items[].recurring_rule.interval` | integer ou null | obligatoire |  |
| `$.items[].recurring_rule.month_of_year` | array ou null | obligatoire |  |
| `$.items[].recurring_rule.month_of_year[]` | number | par élément |  |
| `$.items[].recurring_rule.rule_type` | string | obligatoire | enum `["weekly","monthly","yearly"]`<br>example `"monthly"` |
| `$.items[].recurring_rule.until` | string ou null | obligatoire |  |
| `$.items[].recurring_rule.week_start` | number ou null | obligatoire |  |
| `$.items[].start` | string (date) ou null | obligatoire | The subscription start date<br>example `"2023-01-01"` |
| `$.items[].status` | string | obligatoire | enum `["draft","stopped","finished","pending","not_started","in_progress"]` |
| `$.items[].stopped_at` | string (date-time) ou null | obligatoire | The exact timestamp when the subscription was terminated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | The last time the subscription has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "activated_at": "2023-08-30T10:08:08.146343Z",
      "created_at": "2023-08-30T10:08:08.146343Z",
      "customer": {
        "id": 42,
        "url": "https://app.pennylane.com/api/external/v2/customers/42"
      },
      "customer_invoice_data": {
        "amount": "230.32",
        "currency_amount": "230.32",
        "currency_amount_before_tax": "196.32",
        "currency_tax": "34.0",
        "discount": {
          "value": "25"
        },
        "exchange_rate": "100.00",
        "invoice_line_sections": {
          "url": "https://app.pennylane.com/api/external/v2/billing_subscriptions/42/invoice_line_sections"
        },
        "invoice_lines": {
          "url": "https://app.pennylane.com/api/external/v2/billing_subscriptions/42/invoice_lines"
        },
        "label": "Demo label",
        "pdf_description": "Invoice description",
        "pdf_invoice_free_text": "Thanks for paying this invoice",
        "pdf_invoice_subject": "Invoice subject",
        "special_mention": "Additional details"
      },
      "finish": "2023-12-31",
      "id": 42,
      "next_occurrence": "2023-01-01",
      "prev_occurrence": "2023-01-01",
      "recurring_rule": {
        "count": 12,
        "rule_type": "monthly"
      },
      "start": "2023-01-01",
      "stopped_at": "2023-08-30T10:08:08.146343Z",
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
