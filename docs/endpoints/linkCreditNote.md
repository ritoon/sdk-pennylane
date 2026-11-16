# Link a credit note to a customer invoice

`POST /api/external/v2/customer_invoices/{id}/link_credit_note`

Link a credit note to a customer invoice

> ℹ️
> This endpoint requires the following scope: `customer_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/linkcreditnote) · [Index du SDK](../operations.md)

## Appel Go

```go
LinkCreditNoteWithResponse(ctx context.Context, id int, params *LinkCreditNoteParams, body LinkCreditNoteJSONRequestBody, reqEditors ...RequestEditorFn) (*LinkCreditNoteResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | The unique identifier of the customer invoice<br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.credit_note_id` | integer (int64) | obligatoire | The credit note ID<br>example `42` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "credit_note_id": 42
}
```

</details>

## Sortie

### HTTP 200

The linked credit note

`application/json` → `réponse.JSON200` (type `*GetCustomerInvoiceStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.amount` | string | obligatoire | Invoice amount (total value of the invoice in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"230.32"` |
| `$.appendices` | object | obligatoire | additionalProperties `false` |
| `$.appendices.url` | string | obligatoire | URL to get the appendices of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/appendices"` |
| `$.archived_at` | string (date-time) ou null | obligatoire | The time the invoice has been archived<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.billing_subscription` | object ou null | obligatoire | additionalProperties `false` |
| `$.billing_subscription.id` | integer (int64) | obligatoire |  |
| `$.categories` | object | obligatoire | additionalProperties `false` |
| `$.categories.url` | string | obligatoire | URL to get the categories of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/categories"` |
| `$.contract_reference` | string ou null | obligatoire | Contract reference (e-invoice field BT-12). Maximum 50 characters.<br>example `"CT-2026-001"` |
| `$.created_at` | string (date-time) | obligatoire | The time the invoice has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.credited_invoice` | object ou null | obligatoire | The credited invoice if the invoice is a credit note.<br>additionalProperties `false` |
| `$.credited_invoice.id` | integer (int64) | obligatoire | example `40` |
| `$.credited_invoice.url` | string | obligatoire | URL to get the credited invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/40"` |
| `$.currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.currency_amount` | string | obligatoire | Invoice currency amount (total value of the invoice in the currency of the invoice)<br>example `"230.32"` |
| `$.currency_amount_before_tax` | string | obligatoire | Invoice currency amount before tax (total value before tax of the invoice in the currency of the invoice)<br>example `"196.32"` |
| `$.currency_tax` | string | obligatoire | Invoice taxable amount (in invoice currency)<br>example `"34.0"` |
| `$.custom_header_fields` | object | obligatoire | additionalProperties `false` |
| `$.custom_header_fields.url` | string | obligatoire | URL to get the custom header fields of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/custom_header_fields"` |
| `$.customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.customer.id` | integer (int64) | obligatoire | example `42` |
| `$.customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.customer_invoice_template` | object ou null | obligatoire | additionalProperties `false` |
| `$.customer_invoice_template.id` | integer (int64) | obligatoire |  |
| `$.date` | string (date) ou null | obligatoire | Invoice issue date (ISO 8601)<br>example `"2023-08-30"` |
| `$.deadline` | string (date) ou null | obligatoire | Invoice payment deadline (ISO 8601)<br>example `"2020-09-02"` |
| `$.discount` | object | obligatoire | additionalProperties `false` |
| `$.discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.draft` | boolean | obligatoire | Indicates if the invoice is in draft (has not been finalized)<br>example `false` |
| `$.e_invoicing` | object ou null | obligatoire | E-invoicing lifecycle information managed by the PA.<br>Contains the current status and associated reason when applicable.<br>This field is only present for invoices sent through e-invoicing.<br><br>additionalProperties `false` |
| `$.e_invoicing.flow` | object ou null | obligatoire | The AFNOR flow associated with this e-invoice.<br>additionalProperties `false` |
| `$.e_invoicing.flow.id` | string (uuid) | obligatoire |  |
| `$.e_invoicing.lifecycle_events` | array ou null | obligatoire | > ⚠️ **Warning**: This feature is in beta.<br>> Until further rollout, this field will be `null` for all invoices.<br>> If you plan to rely on it, please get in touch with us first.<br>Chronologically ordered e-invoicing lifecycle events received from the PA.<br>Some events are optional and only present when the receiving platform reports them. |
| `$.e_invoicing.lifecycle_events[]` | object | par élément | additionalProperties `false` |
| `$.e_invoicing.lifecycle_events[].timestamp` | string (date-time) | obligatoire |  |
| `$.e_invoicing.lifecycle_events[].type` | string | obligatoire | enum `["accepted","available_to_buyer","emitted_by_seller_platform","in_dispute","partially_collected","received_by_buyer_platform","refused","rejected","sent","submitted"]` |
| `$.e_invoicing.reason` | string ou null | obligatoire |  |
| `$.e_invoicing.status` | string | obligatoire | enum `["accepted","approved","collected","in_dispute","partially_collected","refused","rejected","sent","submitted"]` |
| `$.electronic_billing_address` | string ou null | obligatoire | The electronic billing address used for the e-invoicing routing of this invoice.<br>The address comes from the e-invoicing directory, for example `123456789` or `123456789_00012`.<br><br>It is null when:<br>- the invoice is imported without a conversion to an e-invoice.<br>- the company could not send e-invoices when the invoice was finalized.<br>- the customer has no electronic billing address and no SIREN, such as an individual.<br>- the invoice is a draft.<br>- the imported e-invoice file is still being processed.<br><br>It is set later when:<br>- the draft invoice is finalized.<br>- the imported e-invoice file is processed.<br><br>example `"123456789_00012"` |
| `$.exchange_rate` | string | obligatoire | Invoice exchange rate (used to convert the invoice to euros. If the invoice currency is euro it will be 1.0)<br>example `"1.0"` |
| `$.external_reference` | string | obligatoire | The unique external reference that was assigned during creation either by you or Pennylane.  (Same attribute as `external_id` in the API v1)<br>example `"FR123"` |
| `$.factur_x` | boolean | obligatoire | Whether the attached invoice file is a Factur-X document (a PDF/A-3 with an embedded structured XML). |
| `$.filename` | string ou null | obligatoire | Name of the file attached to the invoice. May be null immediately after creation if the PDF has not been generated yet; fetch the invoice again to get the filename once it is available.<br>example `"my_file.pdf"` |
| `$.id` | integer (int64) | obligatoire | Invoice identifier<br>example `42` |
| `$.installments` | object ou null | obligatoire | Link to the installments of the invoice.<br>additionalProperties `false` |
| `$.installments.url` | string | obligatoire | URL to get the installments of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/installments"` |
| `$.invoice_line_sections` | object | obligatoire | additionalProperties `false` |
| `$.invoice_line_sections.url` | string | obligatoire | URL to get the invoice line sections of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/invoice_line_sections"` |
| `$.invoice_lines` | object | obligatoire | additionalProperties `false` |
| `$.invoice_lines.url` | string | obligatoire | URL to get the invoice lines of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/invoice_lines"` |
| `$.invoice_number` | string | obligatoire | example `"F20230001"` |
| `$.label` | string ou null | obligatoire | example `"Invoice label"` |
| `$.language` | string | obligatoire | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.ledger_entry` | object | obligatoire | additionalProperties `false` |
| `$.ledger_entry.id` | integer (int64) | obligatoire | Ledger entry identifier<br>example `42002` |
| `$.matched_transactions` | object | obligatoire | additionalProperties `false` |
| `$.matched_transactions.url` | string | obligatoire | URL to get the transactions of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/matched_transactions"` |
| `$.paid` | boolean | obligatoire | Invoice paid status (set to True if the invoice is paid)<br>example `false` |
| `$.payment_method_ledger_account` | object ou null | obligatoire | Ledger account used for reconciliation by payment method.<br>When present, the receivable was transferred to this account.<br><br>additionalProperties `false` |
| `$.payment_method_ledger_account.id` | integer (int64) | obligatoire | example `1255` |
| `$.payment_method_ledger_account.url` | string | obligatoire | URL to get the ledger account.<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/1255"` |
| `$.payments` | object | obligatoire | additionalProperties `false` |
| `$.payments.url` | string | obligatoire | URL to get the payments of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/customer_invoices/42/payments"` |
| `$.pdf_description` | string ou null | obligatoire | example `"Invoice description"` |
| `$.pdf_invoice_free_text` | string | obligatoire | example `"Thanks for paying this invoice"` |
| `$.pdf_invoice_subject` | string | obligatoire | example `"Invoice subject"` |
| `$.performance_date` | string (date) ou null | obligatoire | End date of the performance period of the first invoice line that has one (ISO 8601). Used for German VAT accounting (Leistungsdatum). Null if the company does not use German accounting, or if no invoice line has a performance period.<br>example `"2023-07-31"` |
| `$.public_file_url` | string ou null | obligatoire | Public URL of the invoice file. The URL will expire after 30 minutes. May be null if the PDF has not been generated yet (for example right after creation or finalization); fetch the invoice again to get the URL once it is available.<br>example `"https://app.pennylane.com/public/invoice/pdf?encrypted_id=bzjoVJe...3D%3D"` |
| `$.purchase_order_reference` | string ou null | obligatoire | Buyer's purchase order reference (e-invoice field BT-13). Maximum 50 characters.<br>example `"PO-2026-001"` |
| `$.quote` | object ou null | obligatoire | The quote at the origin of the invoice<br>additionalProperties `false` |
| `$.quote.id` | integer (int64) | obligatoire | example `42` |
| `$.remaining_amount_with_tax` | string ou null | obligatoire | The remaining amount with VAT to pay for the invoice to be considered paid<br>example `"20.0"` |
| `$.remaining_amount_without_tax` | string ou null | obligatoire | The remaining amount without VAT to pay for the invoice to be considered paid<br>example `"16.0"` |
| `$.sale_type` | oneOf | obligatoire | Type of sale printed on the invoice (Type de vente).<br>`goods` is a supply of goods, `services` is a supply of services,<br>and `goods_and_services` is both. `null` means the type has not been chosen.<br>This is not invoice line `substance`. |
| `$.sale_type (oneOf 1)` | string | selon variante | Type of sale printed on a customer invoice. This is not invoice line substance.<br>enum `["goods","services","goods_and_services"]` |
| `$.sale_type (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.sales_order_reference` | string ou null | obligatoire | Seller's sales order reference (e-invoice field BT-14). Maximum 50 characters.<br>example `"SO-2026-001"` |
| `$.schematron_validation_status` | string ou null | obligatoire | Result of Schematron structural validation for Factur-X invoices.<br>Only relevant when `factur_x` is true, null otherwise.<br><br>enum `["pending","valid","invalid"]` |
| `$.special_mention` | string ou null | obligatoire | Additional details<br>example `"Additional details"` |
| `$.status` | string | obligatoire | enum `["archived","incomplete","cancelled","paid","partially_paid","partially_cancelled","upcoming","late","draft","credit_note","proforma","shipping_order","purchasing_order","estimate_pending","estimate_accepted","estimate_invoiced","estimate_denied"]` |
| `$.tax` | string | obligatoire | Invoice taxable amount (in euros). If the currency is euro, `currency_tax` and `tax` are identical.<br>example `"34.0"` |
| `$.transaction_reference` | object ou null | obligatoire | This reconciles the invoice with a transaction. See documentation about [automatic payment matching](https://pennylane.readme.io/v2.0/docs/automating-payment-matching).<br>additionalProperties `false` |
| `$.transaction_reference.banking_provider` | string | obligatoire | The banking provider for the transaction<br>example `"bank"` |
| `$.transaction_reference.provider_field_name` | string | obligatoire | Name of the field that you want to match<br>example `"label"` |
| `$.transaction_reference.provider_field_value` | string | obligatoire | Value that you want to match<br>example `"invoice_number"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the invoice has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "amount": "230.32",
  "appendices": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/appendices"
  },
  "archived_at": "2023-08-30T10:08:08.146343Z",
  "categories": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/categories"
  },
  "contract_reference": "CT-2026-001",
  "created_at": "2023-08-30T10:08:08.146343Z",
  "credited_invoice": {
    "id": 40,
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/40"
  },
  "currency_amount": "230.32",
  "currency_amount_before_tax": "196.32",
  "currency_tax": "34.0",
  "custom_header_fields": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/custom_header_fields"
  },
  "customer": {
    "id": 42,
    "url": "https://app.pennylane.com/api/external/v2/customers/42"
  },
  "date": "2023-08-30",
  "deadline": "2020-09-02",
  "discount": {
    "value": "100.00"
  },
  "draft": false,
  "electronic_billing_address": "123456789_00012",
  "exchange_rate": "1.0",
  "external_reference": "FR123",
  "filename": "my_file.pdf",
  "id": 42,
  "installments": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/installments"
  },
  "invoice_line_sections": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/invoice_line_sections"
  },
  "invoice_lines": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/invoice_lines"
  },
  "invoice_number": "F20230001",
  "label": "Invoice label",
  "ledger_entry": {
    "id": 42002
  },
  "matched_transactions": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/matched_transactions"
  },
  "paid": false,
  "payment_method_ledger_account": {
    "id": 1255,
    "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/1255"
  },
  "payments": {
    "url": "https://app.pennylane.com/api/external/v2/customer_invoices/42/payments"
  },
  "pdf_description": "Invoice description",
  "pdf_invoice_free_text": "Thanks for paying this invoice",
  "pdf_invoice_subject": "Invoice subject",
  "performance_date": "2023-07-31",
  "public_file_url": "https://app.pennylane.com/public/invoice/pdf?encrypted_id=bzjoVJe...3D%3D",
  "purchase_order_reference": "PO-2026-001",
  "quote": {
    "id": 42
  },
  "remaining_amount_with_tax": "20.0",
  "remaining_amount_without_tax": "16.0",
  "sales_order_reference": "SO-2026-001",
  "special_mention": "Additional details",
  "tax": "34.0",
  "transaction_reference": {
    "banking_provider": "bank",
    "provider_field_name": "label",
    "provider_field_value": "invoice_number"
  },
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

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
