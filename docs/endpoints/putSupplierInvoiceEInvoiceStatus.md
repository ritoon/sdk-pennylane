# Update e-invoice status for a supplier invoice

`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/e_invoice_status`

Applies an electronic invoicing lifecycle transition to a supplier invoice received via the PA.

Invoices arrive with a `null` e-invoicing status, which means implicitly approved — no action is required to accept them.

Allowed transitions:
- **Dispute** (`status: "disputed"`): allowed when the current e-invoicing status is `null`, `approved`, or `waiting_for_validation`, the invoice is not archived, and it has no linked payments. Requires a `reason`.
- **Refuse** (`status: "refused"`): allowed when the invoice is not archived and has no linked payments. Requires a `reason`. This is a terminal action — the invoice will be archived and cannot be disputed or refused again.
- **Undispute** (`status: "approved"`): allowed only when the current e-invoicing status is `disputed`. Removes the dispute and restores the invoice to approved. No reason required.

Returns 422 if the transition is not allowed given the current status (e.g. undisputing a non-disputed invoice, or disputing/refusing a collected invoice).


> ℹ️
> This endpoint requires the following scope: `supplier_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/putsupplierinvoiceeinvoicestatus) · [Index du SDK](../operations.md)

## Appel Go

```go
PutSupplierInvoiceEInvoiceStatusWithResponse(ctx context.Context, supplierInvoiceId int, params *PutSupplierInvoiceEInvoiceStatusParams, body PutSupplierInvoiceEInvoiceStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*PutSupplierInvoiceEInvoiceStatusResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `supplier_invoice_id` | `supplierInvoiceId` | path | integer | obligatoire | <br>example `12345` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | oneOf | racine |  |
| `$ (oneOf 1)` | object | selon variante | title `"Dispute"`<br>additionalProperties `false` |
| `$ (oneOf 1).reason` | string | obligatoire | enum `["incorrect_vat_rate","incorrect_unit_prices","incorrect_billed_quantity","incorrect_billed_item","defective_delivered_item","delivery_issue","bank_details_error","incorrect_payment_terms","missing_legal_notice","missing_contractual_reference","recipient_error"]` |
| `$ (oneOf 1).status` | string | obligatoire | enum `["disputed"]` |
| `$ (oneOf 2)` | object | selon variante | title `"Refuse"`<br>additionalProperties `false` |
| `$ (oneOf 2).reason` | string | obligatoire | enum `["incorrect_vat_rate","contract_completed","duplicate_invoice","recipient_error","incorrect_prices","non_compliant_invoice"]` |
| `$ (oneOf 2).status` | string | obligatoire | enum `["refused"]` |
| `$ (oneOf 3)` | object | selon variante | title `"Undispute"`<br>additionalProperties `false` |
| `$ (oneOf 3).status` | string | obligatoire | enum `["approved"]` |

## Sortie

### HTTP 200

The updated supplier invoice

`application/json` → `réponse.JSON200` (type `*PutSupplierInvoiceEInvoiceStatusStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.accounting_status` | string | obligatoire | The accounting state of the invoice.<br>- `draft`: The invoice is not yet sent to the accountant.<br>- `archived`: The invoice has been archived.<br>- `entry`: The invoice is incomplete. Some information is missing on the invoice and needs to be completed by SME.<br>- `validation_needed`: The invoice is sent to the accountant and needs validation.<br>- `complete`: The invoice has been validated by the accountant.<br><br>enum `["draft","archived","entry","validation_needed","complete"]` |
| `$.amount` | string | obligatoire | Invoice amount (total value of the invoice in euros). If the currency is euro, `currency_amount` and `amount` are identical.<br>example `"230.32"` |
| `$.archived_at` | string (date-time) ou null | obligatoire | The time the invoice has been archived<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.categories` | object | obligatoire | additionalProperties `false` |
| `$.categories.url` | string | obligatoire | URL to get the categories of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/supplier_invoices/42/categories"` |
| `$.created_at` | string (date-time) | obligatoire | The time the invoice has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.currency_amount` | string | obligatoire | Invoice currency amount (total value of the invoice in the currency of the invoice)<br>example `"230.32"` |
| `$.currency_amount_before_tax` | string | obligatoire | Invoice currency amount before tax (total value before tax of the invoice in the currency of the invoice)<br>example `"196.32"` |
| `$.currency_tax` | string | obligatoire | Invoice taxable amount (in euros). If the currency is euro, `currency_tax` and `tax` are identical.<br>example `"34.0"` |
| `$.date` | string (date) ou null | obligatoire | Invoice issue date (ISO 8601)<br>example `"2023-08-30"` |
| `$.deadline` | string (date) ou null | obligatoire | Invoice payment deadline (ISO 8601)<br>example `"2023-09-30"` |
| `$.e_invoicing` | object ou null | obligatoire | E-invoicing lifecycle information managed by the PA.<br>Contains the current status and associated reason when applicable.<br>This field is only present for invoices received through e-invoicing.<br><br>additionalProperties `false` |
| `$.e_invoicing.electronic_billing_address` | string ou null | obligatoire | The electronic billing address (Peppol address) that received this<br>e-invoice. Use it to isolate invoices routed to a specific entity.<br><br>example `"0009:123456789"` |
| `$.e_invoicing.flow` | object ou null | obligatoire | The AFNOR flow associated with this e-invoice.<br>additionalProperties `false` |
| `$.e_invoicing.flow.id` | string (uuid) | obligatoire |  |
| `$.e_invoicing.reason` | anyOf | obligatoire |  |
| `$.e_invoicing.reason (anyOf 1)` | string | selon variante | enum `["incorrect_vat_rate","incorrect_unit_prices","incorrect_billed_quantity","incorrect_billed_item","defective_delivered_item","delivery_issue","bank_details_error","incorrect_payment_terms","missing_legal_notice","missing_contractual_reference","recipient_error","contract_completed","duplicate_invoice","incorrect_prices","non_compliant_invoice"]` |
| `$.e_invoicing.reason (anyOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.e_invoicing.source_file_url` | string ou null | obligatoire | URL to download the original e-invoice file as received from the PA<br>(CII, UBL, or Factur-X). Pennylane may replace the main invoice file<br>with a generated PDF; this field preserves the untouched file received<br>from the PA.<br><br>example `"https://app.pennylane.com/rails/active_storage/blobs/redirect/abc123/invoice.xml"` |
| `$.e_invoicing.status` | anyOf | obligatoire |  |
| `$.e_invoicing.status (anyOf 1)` | string | selon variante | enum `["waiting_for_validation","approved","rejected","disputed","refused","collected","partially_collected"]` |
| `$.e_invoicing.status (anyOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.exchange_rate` | string | obligatoire | Invoice exchange rate (used to convert the invoice to euros. If the invoice currency is euro it will be 1.0)<br>example `"1.0"` |
| `$.external_reference` | string | obligatoire | The unique external reference that was assigned during creation either by you or Pennylane<br>example `"FR123"` |
| `$.filename` | string ou null | obligatoire | Name of the file attached to the invoice<br>example `"my_file.pdf"` |
| `$.id` | integer (int64) | obligatoire | The ID of the supplier invoice<br>example `123` |
| `$.import_source` | object ou null | obligatoire | Origin file information for invoices imported via a file connector<br>(Dropbox or Google Drive). Null for invoices created by other means.<br><br>additionalProperties `false` |
| `$.import_source.path` | string | obligatoire | The file path of the invoice in the connector's storage.<br>example `"/ACCOUNTING INVOICES/2024/invoice.pdf"` |
| `$.import_source.source` | string | obligatoire | The connector type that imported this invoice.<br>enum `["dropbox","gdrive"]`<br>example `"dropbox"` |
| `$.invoice_lines` | object | obligatoire | additionalProperties `false` |
| `$.invoice_lines.url` | string | obligatoire | URL to get the invoice lines of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/supplier_invoices/42/invoice_lines"` |
| `$.invoice_number` | string | obligatoire | example `"F20230001"` |
| `$.label` | string ou null | obligatoire | example `"Demo label"` |
| `$.ledger_entry` | object | obligatoire | additionalProperties `false` |
| `$.ledger_entry.id` | integer (int64) | obligatoire | Ledger entry identifier<br>example `42003` |
| `$.matched_transactions` | object | obligatoire | additionalProperties `false` |
| `$.matched_transactions.url` | string | obligatoire | URL to get the transactions of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/supplier_invoices/42/matched_transactions"` |
| `$.paid` | boolean | obligatoire | Invoice paid status (set to true if the invoice is paid)<br>example `false` |
| `$.payment_method_ledger_account` | object ou null | obligatoire | Ledger account used for reconciliation by payment method.<br>When present, the receivable/payable was transferred to this account.<br><br>additionalProperties `false` |
| `$.payment_method_ledger_account.id` | integer (int64) | obligatoire | example `1255` |
| `$.payment_method_ledger_account.url` | string | obligatoire | URL to get the ledger account.<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/1255"` |
| `$.payment_status` | string | obligatoire | enum `["to_be_processed","to_be_paid","partially_paid","payment_error","payment_scheduled","payment_in_progress","payment_emitted","payment_found","paid_offline","fully_paid"]` |
| `$.payments` | object | obligatoire | additionalProperties `false` |
| `$.payments.url` | string | obligatoire | URL to get the payments of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/supplier_invoices/42/payments"` |
| `$.public_file_url` | string ou null | obligatoire | Public URL of the invoice file. The URL will expire after 30 minutes.<br>example `"https://app.pennylane.com/public/invoice/pdf?encrypted_id=bzjoVJe...3D%3D"` |
| `$.reconciled` | boolean | obligatoire | Whether the invoice has been reconciled or not<br>example `false` |
| `$.remaining_amount_with_tax` | string ou null | obligatoire | The remaining amount with VAT to pay for the invoice to be considered paid<br>example `"20.0"` |
| `$.remaining_amount_without_tax` | string ou null | obligatoire | The remaining amount without VAT and without the Spanish IRPF income tax withholding, if any, to pay for the invoice to be considered paid<br>example `"16.0"` |
| `$.supplier` | object ou null | obligatoire | additionalProperties `false` |
| `$.supplier.id` | integer (int64) | obligatoire | example `456` |
| `$.supplier.url` | string | obligatoire | URL to get the supplier.<br>example `"https://app.pennylane.com/api/external/v2/suppliers/42"` |
| `$.tax` | string | obligatoire | Invoice taxable amount (in invoice currency)<br>example `"34.0"` |
| `$.transaction_reference` | object ou null | obligatoire | This reconciles the invoice with a transaction. See documentation about [automatic payment matching](https://pennylane.readme.io/v2.0/docs/automating-payment-matching).<br>additionalProperties `false` |
| `$.transaction_reference.banking_provider` | string | obligatoire | The banking provider for the transaction<br>example `"bank"` |
| `$.transaction_reference.provider_field_name` | string | obligatoire | Name of the field that you want to match<br>example `"label"` |
| `$.transaction_reference.provider_field_value` | string | obligatoire | Value that you want to match<br>example `"invoice_number"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the invoice has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "amount": "230.32",
  "archived_at": "2023-08-30T10:08:08.146343Z",
  "categories": {
    "url": "https://app.pennylane.com/api/external/v2/supplier_invoices/42/categories"
  },
  "created_at": "2023-08-30T10:08:08.146343Z",
  "currency_amount": "230.32",
  "currency_amount_before_tax": "196.32",
  "currency_tax": "34.0",
  "date": "2023-08-30",
  "deadline": "2023-09-30",
  "e_invoicing": {
    "electronic_billing_address": "0009:123456789",
    "source_file_url": "https://app.pennylane.com/rails/active_storage/blobs/redirect/abc123/invoice.xml"
  },
  "exchange_rate": "1.0",
  "external_reference": "FR123",
  "filename": "my_file.pdf",
  "id": 123,
  "import_source": {
    "path": "/ACCOUNTING INVOICES/2024/invoice.pdf",
    "source": "dropbox"
  },
  "invoice_lines": {
    "url": "https://app.pennylane.com/api/external/v2/supplier_invoices/42/invoice_lines"
  },
  "invoice_number": "F20230001",
  "label": "Demo label",
  "ledger_entry": {
    "id": 42003
  },
  "matched_transactions": {
    "url": "https://app.pennylane.com/api/external/v2/supplier_invoices/42/matched_transactions"
  },
  "paid": false,
  "payment_method_ledger_account": {
    "id": 1255,
    "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/1255"
  },
  "payments": {
    "url": "https://app.pennylane.com/api/external/v2/supplier_invoices/42/payments"
  },
  "public_file_url": "https://app.pennylane.com/public/invoice/pdf?encrypted_id=bzjoVJe...3D%3D",
  "reconciled": false,
  "remaining_amount_with_tax": "20.0",
  "remaining_amount_without_tax": "16.0",
  "supplier": {
    "id": 456,
    "url": "https://app.pennylane.com/api/external/v2/suppliers/42"
  },
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
