# List customer invoices

`GET /api/external/v2/customer_invoices`

List customer invoices and credit notes

> ℹ️
> This endpoint requires one of the following scopes: `customer_invoices:all`, `customer_invoices:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcustomerinvoices) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCustomerInvoicesWithResponse(ctx context.Context, params *GetCustomerInvoicesParams, reqEditors ...RequestEditorFn) (*GetCustomerInvoicesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`, `date`, `customer_id`, `billing_subscription_id`, `quote_id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `invoice_number`: `eq`, `not_eq`, `in`, `not_in`<br>- `draft`: `eq` (boolean)<br>- `credit_note`: `eq` (boolean)<br>- `external_reference`: `eq`<br>- `flow_id`: `eq`, `not_eq`, `in`, `not_in`<br>- `category_id`: `in`<br>example `"[{\"field\": \"date\", \"operator\": \"gteq\", \"value\": \"2024-01-01\"}, {\"field\": \"credit_note\", \"operator\": \"eq\", \"value\": \"true\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`, `date`<br>default `"-id"`<br>example `"id"` |
| `include` | `params.Include` | query | string | facultatif | You can choose to include additional related resources in the response.<br>When specified, related resources will be returned in a separate `included` section.<br>Available includes: `invoice_lines`<br><br>⚠️ **Warning**: This feature is currently experimental and may change or be removed in future releases.<br>We recommend using it with caution in production environments.<br>example `"invoice_lines"` |
| `include_invoices_created_from_ledger_entries` | `params.IncludeInvoicesCreatedFromLedgerEntries` | query | boolean | facultatif | Set to `true` to also return customer invoices created from imported accounting entries<br>(FEC, journal entry imports or accounting software migrations).<br>By default, only customer invoices created in Pennylane (in the app, through the API or from a quote) are returned.<br>default `false`<br>example `true` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of Customer Invoices and/or credit notes

`application/json` → `réponse.JSON200` (type `*GetCustomerInvoicesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | allOf | racine |  |
| `$ (allOf 1)` | object | selon variante |  |
| `$ (allOf 1).has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$ (allOf 1).items` | array | obligatoire | Array of customer invoices |
| `$ (allOf 1).items[]` | object | par élément | additionalProperties `false` |
| `$ (allOf 1).items[].amount` | string | obligatoire | Invoice amount (total value of the invoice in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"230.32"` |
| `$ (allOf 1).items[].appendices` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].appendices.url` | string | obligatoire | URL to get the appendices of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/appendices"` |
| `$ (allOf 1).items[].archived_at` | string (date-time) ou null | obligatoire | The time the invoice has been archived<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$ (allOf 1).items[].billing_subscription` | object ou null | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].billing_subscription.id` | integer (int64) | obligatoire |  |
| `$ (allOf 1).items[].categories` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].categories.url` | string | obligatoire | URL to get the categories of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/categories"` |
| `$ (allOf 1).items[].contract_reference` | string ou null | obligatoire | Contract reference (e-invoice field BT-12). Maximum 50 characters.<br>example `"CT-2026-001"` |
| `$ (allOf 1).items[].created_at` | string (date-time) | obligatoire | The time the invoice has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$ (allOf 1).items[].credited_invoice` | object ou null | obligatoire | The credited invoice if the invoice is a credit note.<br>additionalProperties `false` |
| `$ (allOf 1).items[].credited_invoice.id` | integer (int64) | obligatoire | example `40` |
| `$ (allOf 1).items[].credited_invoice.url` | string | obligatoire | URL to get the credited invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/40"` |
| `$ (allOf 1).items[].currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$ (allOf 1).items[].currency_amount` | string | obligatoire | Invoice currency amount (total value of the invoice in the currency of the invoice)<br>example `"230.32"` |
| `$ (allOf 1).items[].currency_amount_before_tax` | string | obligatoire | Invoice currency amount before tax (total value before tax of the invoice in the currency of the invoice)<br>example `"196.32"` |
| `$ (allOf 1).items[].currency_tax` | string | obligatoire | Invoice taxable amount (in invoice currency)<br>example `"34.0"` |
| `$ (allOf 1).items[].custom_header_fields` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].custom_header_fields.url` | string | obligatoire | URL to get the custom header fields of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/custom_header_fields"` |
| `$ (allOf 1).items[].customer` | object ou null | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].customer.id` | integer (int64) | obligatoire | example `42` |
| `$ (allOf 1).items[].customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$ (allOf 1).items[].customer_invoice_template` | object ou null | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].customer_invoice_template.id` | integer (int64) | obligatoire |  |
| `$ (allOf 1).items[].date` | string (date) ou null | obligatoire | Invoice issue date (ISO 8601)<br>example `"2023-08-30"` |
| `$ (allOf 1).items[].deadline` | string (date) ou null | obligatoire | Invoice payment deadline (ISO 8601)<br>example `"2020-09-02"` |
| `$ (allOf 1).items[].discount` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$ (allOf 1).items[].discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$ (allOf 1).items[].draft` | boolean | obligatoire | Indicates if the invoice is in draft (has not been finalized)<br>example `false` |
| `$ (allOf 1).items[].e_invoicing` | object ou null | obligatoire | E-invoicing lifecycle information managed by the PA.<br>Contains the current status and associated reason when applicable.<br>This field is only present for invoices sent through e-invoicing.<br><br>additionalProperties `false` |
| `$ (allOf 1).items[].e_invoicing.flow` | object ou null | obligatoire | The AFNOR flow associated with this e-invoice.<br>additionalProperties `false` |
| `$ (allOf 1).items[].e_invoicing.flow.id` | string (uuid) | obligatoire |  |
| `$ (allOf 1).items[].e_invoicing.lifecycle_events` | array ou null | obligatoire | > ⚠️ **Warning**: This feature is in beta.<br>> Until further rollout, this field will be `null` for all invoices.<br>> If you plan to rely on it, please get in touch with us first.<br>Chronologically ordered e-invoicing lifecycle events received from the PA.<br>Some events are optional and only present when the receiving platform reports them. |
| `$ (allOf 1).items[].e_invoicing.lifecycle_events[]` | object | par élément | additionalProperties `false` |
| `$ (allOf 1).items[].e_invoicing.lifecycle_events[].timestamp` | string (date-time) | obligatoire |  |
| `$ (allOf 1).items[].e_invoicing.lifecycle_events[].type` | string | obligatoire | enum `["accepted","available_to_buyer","emitted_by_seller_platform","in_dispute","partially_collected","received_by_buyer_platform","refused","rejected","sent","submitted"]` |
| `$ (allOf 1).items[].e_invoicing.reason` | string ou null | obligatoire |  |
| `$ (allOf 1).items[].e_invoicing.status` | string | obligatoire | enum `["accepted","approved","collected","in_dispute","partially_collected","refused","rejected","sent","submitted"]` |
| `$ (allOf 1).items[].electronic_billing_address` | string ou null | obligatoire | The electronic billing address used for the e-invoicing routing of this invoice.<br>The address comes from the e-invoicing directory, for example `123456789` or `123456789_00012`.<br><br>It is null when:<br>- the invoice is imported without a conversion to an e-invoice.<br>- the company could not send e-invoices when the invoice was finalized.<br>- the customer has no electronic billing address and no SIREN, such as an individual.<br>- the invoice is a draft.<br>- the imported e-invoice file is still being processed.<br><br>It is set later when:<br>- the draft invoice is finalized.<br>- the imported e-invoice file is processed.<br><br>example `"123456789_00012"` |
| `$ (allOf 1).items[].exchange_rate` | string | obligatoire | Invoice exchange rate (used to convert the invoice to euros. If the invoice currency is euro it will be 1.0)<br>example `"1.0"` |
| `$ (allOf 1).items[].external_reference` | string | obligatoire | The unique external reference that was assigned during creation either by you or Pennylane.  (Same attribute as `external_id` in the API v1)<br>example `"FR123"` |
| `$ (allOf 1).items[].factur_x` | boolean | obligatoire | Whether the attached invoice file is a Factur-X document (a PDF/A-3 with an embedded structured XML). |
| `$ (allOf 1).items[].filename` | string ou null | obligatoire | Name of the file attached to the invoice. May be null immediately after creation if the PDF has not been generated yet; fetch the invoice again to get the filename once it is available.<br>example `"my_file.pdf"` |
| `$ (allOf 1).items[].id` | integer (int64) | obligatoire | Invoice identifier<br>example `42` |
| `$ (allOf 1).items[].installments` | object ou null | obligatoire | Link to the installments of the invoice.<br>additionalProperties `false` |
| `$ (allOf 1).items[].installments.url` | string | obligatoire | URL to get the installments of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/installments"` |
| `$ (allOf 1).items[].invoice_line_sections` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].invoice_line_sections.url` | string | obligatoire | URL to get the invoice line sections of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/invoice_line_sections"` |
| `$ (allOf 1).items[].invoice_lines` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].invoice_lines.url` | string | obligatoire | URL to get the invoice lines of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/invoice_lines"` |
| `$ (allOf 1).items[].invoice_number` | string | obligatoire | example `"F20230001"` |
| `$ (allOf 1).items[].label` | string ou null | obligatoire | example `"Invoice label"` |
| `$ (allOf 1).items[].language` | string | obligatoire | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$ (allOf 1).items[].ledger_entry` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].ledger_entry.id` | integer (int64) | obligatoire | Ledger entry identifier<br>example `42002` |
| `$ (allOf 1).items[].matched_transactions` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].matched_transactions.url` | string | obligatoire | URL to get the transactions of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/matched_transactions"` |
| `$ (allOf 1).items[].paid` | boolean | obligatoire | Invoice paid status (set to True if the invoice is paid)<br>example `false` |
| `$ (allOf 1).items[].payment_method_ledger_account` | object ou null | obligatoire | Ledger account used for reconciliation by payment method.<br>When present, the receivable was transferred to this account.<br><br>additionalProperties `false` |
| `$ (allOf 1).items[].payment_method_ledger_account.id` | integer (int64) | obligatoire | example `1255` |
| `$ (allOf 1).items[].payment_method_ledger_account.url` | string | obligatoire | URL to get the ledger account.<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/1255"` |
| `$ (allOf 1).items[].payments` | object | obligatoire | additionalProperties `false` |
| `$ (allOf 1).items[].payments.url` | string | obligatoire | URL to get the payments of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/payments"` |
| `$ (allOf 1).items[].pdf_description` | string ou null | obligatoire | example `"Invoice description"` |
| `$ (allOf 1).items[].pdf_invoice_free_text` | string | obligatoire | example `"Thanks for paying this invoice"` |
| `$ (allOf 1).items[].pdf_invoice_subject` | string | obligatoire | example `"Invoice subject"` |
| `$ (allOf 1).items[].performance_date` | string (date) ou null | obligatoire | End date of the performance period of the first invoice line that has one (ISO 8601). Used for German VAT accounting (Leistungsdatum). Null if the company does not use German accounting, or if no invoice line has a performance period.<br>example `"2023-07-31"` |
| `$ (allOf 1).items[].public_file_url` | string ou null | obligatoire | Public URL of the invoice file. The URL will expire after 30 minutes. May be null if the PDF has not been generated yet (for example right after creation or finalization); fetch the invoice again to get the URL once it is available.<br>example `"https://app.pennylane.com/public/invoice/pdf?encrypted_id=bzjoVJe...3D%3D"` |
| `$ (allOf 1).items[].purchase_order_reference` | string ou null | obligatoire | Buyer's purchase order reference (e-invoice field BT-13). Maximum 50 characters.<br>example `"PO-2026-001"` |
| `$ (allOf 1).items[].quote` | object ou null | obligatoire | The quote at the origin of the invoice<br>additionalProperties `false` |
| `$ (allOf 1).items[].quote.id` | integer (int64) | obligatoire | example `42` |
| `$ (allOf 1).items[].remaining_amount_with_tax` | string ou null | obligatoire | The remaining amount with VAT to pay for the invoice to be considered paid<br>example `"20.0"` |
| `$ (allOf 1).items[].remaining_amount_without_tax` | string ou null | obligatoire | The remaining amount without VAT to pay for the invoice to be considered paid<br>example `"16.0"` |
| `$ (allOf 1).items[].sale_type` | oneOf | obligatoire | Type of sale printed on the invoice (Type de vente).<br>`goods` is a supply of goods, `services` is a supply of services,<br>and `goods_and_services` is both. `null` means the type has not been chosen.<br>This is not invoice line `substance`. |
| `$ (allOf 1).items[].sale_type (oneOf 1)` | string | selon variante | Type of sale printed on a customer invoice. This is not invoice line substance.<br>enum `["goods","services","goods_and_services"]` |
| `$ (allOf 1).items[].sale_type (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$ (allOf 1).items[].sales_order_reference` | string ou null | obligatoire | Seller's sales order reference (e-invoice field BT-14). Maximum 50 characters.<br>example `"SO-2026-001"` |
| `$ (allOf 1).items[].schematron_validation_status` | string ou null | obligatoire | Result of Schematron structural validation for Factur-X invoices.<br>Only relevant when `factur_x` is true, null otherwise.<br><br>enum `["pending","valid","invalid"]` |
| `$ (allOf 1).items[].special_mention` | string ou null | obligatoire | Additional details<br>example `"Additional details"` |
| `$ (allOf 1).items[].status` | string | obligatoire | enum `["archived","incomplete","cancelled","paid","partially_paid","partially_cancelled","upcoming","late","draft","credit_note","proforma","shipping_order","purchasing_order","estimate_pending","estimate_accepted","estimate_invoiced","estimate_denied"]` |
| `$ (allOf 1).items[].tax` | string | obligatoire | Invoice taxable amount (in euros). If the currency is euro, `currency_tax` and `tax` are identical.<br>example `"34.0"` |
| `$ (allOf 1).items[].transaction_reference` | object ou null | obligatoire | This reconciles the invoice with a transaction. See documentation about [automatic payment matching](https://pennylane.readme.io/v2.0/docs/automating-payment-matching).<br>additionalProperties `false` |
| `$ (allOf 1).items[].transaction_reference.banking_provider` | string | obligatoire | The banking provider for the transaction<br>example `"bank"` |
| `$ (allOf 1).items[].transaction_reference.provider_field_name` | string | obligatoire | Name of the field that you want to match<br>example `"label"` |
| `$ (allOf 1).items[].transaction_reference.provider_field_value` | string | obligatoire | Value that you want to match<br>example `"invoice_number"` |
| `$ (allOf 1).items[].updated_at` | string (date-time) | obligatoire | The last time the invoice has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$ (allOf 1).next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `$ (allOf 2)` | object | selon variante | Optional included resources. Only present when the `include` parameter is used.<br>Each property corresponds to a requested inclusion and contains data keyed by parent invoice ID. |
| `$ (allOf 2).included` | object | facultatif | Container for all included resources |
| `$ (allOf 2).included.invoice_lines` | object | facultatif | Invoice lines grouped by parent invoice ID<br>additionalProperties `{"additionalProperties":false,"description":"Paginated collection of invoice lines for a specific invoice","properties":{"has_more":{"description":"Whether more invoice lines exist beyond this set","example":false,"type":"boolean"},"items":{"description":"Array of invoice lines (up to 20 items)","items":{"additionalProperties":false,"properties":{"amount":{"description":"The total amount of the invoice lines in euros including taxes and deducting discounts","example":"50.4","type":"string"},"created_at":{"example":"2023-08-30T10:08:08.146343Z","format":"date-time","type":"string"},"currency_amount":{"description":"The total amount of the invoice lines in the document's currency including taxes and deducting discounts.\nIf the currency is euro, amount and currency_amount are identical.\n","example":"50.4","type":"string"},"currency_amount_before_tax":{"description":"Total amount before tax in currency","example":"30","type":"string"},"currency_tax":{"description":"Total tax amount in currency","example":"10","type":"string"},"custom_fields":{"description":"**Experimental** — under active development, may change without notice. Do not rely on this field in production integrations.","items":{"additionalProperties":false,"properties":{"name":{"description":"The `api_name` of the custom field","example":"delivery_terms","type":"string"},"value":{"anyOf":[{"nullable":true,"type":"string"},{"items":{"type":"string"},"type":"array"}],"description":"Format depends on the field type: plain string for `text`; numeric string for `number` (e.g. `\"42.5\"`); ISO 8601 date for `date` (e.g. `\"2025-01-15\"`); option label for `single_select`; array of option labels for `multi_select`.","example":"Delivered at place"}},"required":["name","value"],"type":"object"},"type":"array"},"description":{"description":"Invoice line description","example":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor","type":"string"},"discount":{"additionalProperties":false,"properties":{"type":{"description":"Discount type.\n- absolute if it is an amount\n- relative if it is a percentage\n","enum":["absolute","relative"],"type":"string"},"value":{"description":"Discount value on the total amount before tax of the invoice","example":"25","type":"string"}},"required":["type","value"],"type":"object"},"ean_code":{"description":"EAN code (8 or 13 digits) set on the line itself. `null` when the line has none; a product's EAN code is on the product.","example":"4006381333931","nullable":true,"type":"string"},"id":{"description":"Invoice line id","example":444,"type":"number"},"imputation_dates":{"additionalProperties":false,"nullable":true,"properties":{"end_date":{"description":"End date of the imputation period (ISO 8601)","example":"2021-06-30","format":"date","type":"string"},"start_date":{"description":"Start date of the imputation period (ISO 8601)","example":"2020-06-30","format":"date","type":"string"}},"required":["start_date","end_date"],"type":"object"},"label":{"description":"Invoice line label","example":"Demo label","type":"string"},"ledger_account":{"additionalProperties":false,"description":"Ledger account of the invoice line.","properties":{"id":{"example":1255,"format":"int64","type":"integer"},"number":{"description":"Ledger account number","example":"706000","type":"string"},"url":{"description":"URL to get the ledger account.","example":"https://app.pennylane.com/api/external/v2/ledger_accounts/1255","type":"string"}},"required":["id","number","url"],"type":"object"},"performance_period":{"additionalProperties":false,"description":"Date range during which the service or delivery was performed (ISO 8601). Used for German VAT accounting (Leistungsdatum).","nullable":true,"properties":{"date_type":{"oneOf":[{"description":"Whether the performance period represents a single date or a range, and whether it is a delivery or service.\n- delivery_date: single delivery date (end_date only)\n- service_date: single service date (end_date only)\n- delivery_period: delivery date range (start_date and end_date)\n- service_period: service date range (start_date and end_date)\n","enum":["delivery_date","service_date","delivery_period","service_period"],"type":"string"},{"enum":[null],"nullable":true,"type":"string"}]},"end_date":{"description":"End date of the performance period (ISO 8601)","example":"2023-07-31","format":"date","type":"string"},"start_date":{"description":"Start date of the performance period (ISO 8601). Null for single-date types.","example":"2023-07-01","format":"date","nullable":true,"type":"string"}},"required":["start_date","end_date","date_type"],"type":"object"},"product":{"additionalProperties":false,"nullable":true,"properties":{"id":{"description":"Product id","example":3049,"type":"number"},"url":{"description":"URL to get the product.","example":"https://app.pennylane.com/api/external/v2/products/42","type":"string"}},"required":["id","url"],"type":"object"},"quantity":{"description":"Invoice line item quantity (number of items)","example":"12","type":"string"},"raw_currency_unit_price":{"description":"Unit price (excluding tax)","example":"5","type":"string"},"section_rank":{"description":"Has to correspond to the rank number of a line items section in which the line item should be","example":1,"nullable":true,"type":"integer"},"substance":{"description":"Derived from the line's ledger account. Returns `goods` or `services` when unambiguous, or `null` otherwise (e.g. mixed partial-invoice lines).\n","oneOf":[{"enum":["goods","services"],"type":"string"},{"enum":[null],"nullable":true,"type":"string"}]},"tax":{"description":"Total tax amount in euros","example":"10","type":"string"},"unit":{"description":"The kind of unit which applies to the amount","example":"piece","nullable":true,"type":"string"},"updated_at":{"example":"2023-08-30T10:08:08.146343Z","format":"date-time","type":"string"},"vat_rate":{"description":"Product VAT rate. A 20% VAT in France is FR_200.","enum":["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"],"example":"FR_200","type":"string"}},"required":["id","label","unit","quantity","amount","currency_amount","description","product","ean_code","ledger_account","vat_rate","currency_amount_before_tax","currency_tax","tax","raw_currency_unit_price","discount","section_rank","imputation_dates","performance_period","created_at","updated_at","substance","custom_fields"],"title":"Customer Invoice Line","type":"object"},"maxItems":20,"type":"array"},"next_url":{"description":"URL to fetch the next page of invoice lines (null if no more pages)","example":"https://app.pennylane.com/api/external/v2/customer_invoices/42/invoice_lines?cursor=eyJpZCI6...","nullable":true,"type":"string"}},"required":["items","has_more","next_url"],"type":"object"}` |

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
