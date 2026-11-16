# Import a supplier invoice with a file attached

`POST /api/external/v2/supplier_invoices/import`

This endpoint allows you to import a supplier invoice with a file
attached.


> ℹ️
> This endpoint requires the following scope: `supplier_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/importsupplierinvoice) · [Index du SDK](../operations.md)

## Appel Go

```go
ImportSupplierInvoiceWithResponse(ctx context.Context, params *ImportSupplierInvoiceParams, body ImportSupplierInvoiceJSONRequestBody, reqEditors ...RequestEditorFn) (*ImportSupplierInvoiceResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.amount` | string | facultatif | Invoice amount in euros (total value of the invoice in euros). If the currency is euro, `currency_amount` and `amount` are identical.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.currency` | string | facultatif | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.currency_amount` | string | obligatoire | Invoice currency amount (total value of the invoice in the currency of the invoice)<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.currency_amount_before_tax` | string | obligatoire | Invoice currency amount before tax (total value before tax of the invoice in the currency of the invoice)<br>pattern `"^-?\\d+(\\.\\d+)?$"`<br>example `"100.00"` |
| `$.currency_tax` | string | obligatoire | Invoice taxable amount (in invoice currency)<br>pattern `"^-?\\d+(\\.\\d+)?$"`<br>example `"100.00"` |
| `$.date` | string (date) | obligatoire | The date of the invoice (ISO 8601)<br>example `"2021-01-30"` |
| `$.deadline` | string (date) | obligatoire | Invoice payment deadline (ISO 8601)<br>example `"2021-01-30"` |
| `$.external_reference` | string | facultatif | A unique external reference you can provide to track this supplier. If not provided, Pennylane will generate an identifier for you.<br>pattern `"^[^\\x00]*$"`<br>example `"FR123"` |
| `$.file_attachment_id` | integer (int64) | obligatoire | ID of a previously uploaded file attachment. The attachment must already have a file.<br>If the file is not available yet, the API returns 409; retry after a few seconds.<br><br>example `123` |
| `$.import_as_incomplete` | boolean | facultatif | This will mark the invoice as incomplete.<br>default `false` |
| `$.invoice_lines` | array | obligatoire | minItems `1` |
| `$.invoice_lines[]` | object | par élément | additionalProperties `false` |
| `$.invoice_lines[].amount` | string | facultatif | Invoice line amount in euros. If the currency is euro, `currency_amount` and `amount` are identical.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.invoice_lines[].currency_amount` | string | obligatoire | Invoice line total amount in the currency of the invoice.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.invoice_lines[].currency_tax` | string | obligatoire | Invoice line taxable amount (in invoice currency)<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.invoice_lines[].description` | string | facultatif | Invoice line description<br><br>> ⚠️ This attribute is deprecated and will be removed soon.<br><br>pattern `"^[^\\x00]*$"`<br>deprecated `true`<br>example `"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor"` |
| `$.invoice_lines[].imputation_dates` | object | facultatif | additionalProperties `false` |
| `$.invoice_lines[].imputation_dates.end_date` | string (date) | obligatoire | End date of the imputation period (ISO 8601)<br>example `"2021-06-30"` |
| `$.invoice_lines[].imputation_dates.start_date` | string (date) | obligatoire | Start date of the imputation period (ISO 8601)<br>example `"2020-06-30"` |
| `$.invoice_lines[].label` | string | facultatif | Invoice line label<br><br>> ⚠️ This attribute is deprecated and will be removed soon.<br><br>pattern `"^[^\\x00]*$"`<br>deprecated `true`<br>example `"Demo label"` |
| `$.invoice_lines[].ledger_account_id` | integer (int64) | facultatif | The ledger account ID<br>example `1255` |
| `$.invoice_lines[].ledger_entry_line` | object | facultatif | additionalProperties `false` |
| `$.invoice_lines[].ledger_entry_line.label` | string | obligatoire | Ledger entry line label.<br>Ledger entry lines with identical labels, tax, and ledger accounts will be automatically compacted into a single line.<br><br>minLength `1`<br>maxLength `255`<br>pattern `"^[^\\x00]*$"`<br>example `"Demo label"` |
| `$.invoice_lines[].performance_period` | object ou null | facultatif | Date range during which the service or delivery was performed (ISO 8601). Used for German VAT accounting (Leistungsdatum).<br>additionalProperties `false` |
| `$.invoice_lines[].performance_period.date_type` | string | obligatoire | Whether the performance period represents a single date or a range, and whether it is a delivery or service.<br>- delivery_date: single delivery date (end_date only)<br>- service_date: single service date (end_date only)<br>- delivery_period: delivery date range (start_date and end_date)<br>- service_period: service date range (start_date and end_date)<br><br>enum `["delivery_date","service_date","delivery_period","service_period"]` |
| `$.invoice_lines[].performance_period.end_date` | string (date) | obligatoire | End date of the performance period (ISO 8601)<br>example `"2023-07-31"` |
| `$.invoice_lines[].performance_period.start_date` | string (date) ou null | facultatif | Start date of the performance period (ISO 8601)<br>example `"2023-07-01"` |
| `$.invoice_lines[].tax` | string | facultatif | Invoice line taxable amount (in euros). If the currency is euro, `currency_tax` and `tax` are identical.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.invoice_lines[].vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.invoice_number` | string | facultatif | The invoice number<br>pattern `"^[^\\x00]*$"`<br>example `"F20230001"` |
| `$.label` | string ou null | facultatif | Custom label for the invoice used on accounting (ledger) entries. If not provided, Pennylane generates a label automatically.<br>maxLength `2000`<br>pattern `"^[^\\x00]*$"`<br>example `"Custom accounting label"` |
| `$.supplier_id` | integer (int64) | obligatoire | The ID of the supplier to import the invoice for<br>example `456` |
| `$.tax` | string | facultatif | Invoice taxable amount (in euros). If the currency is euro, `currency_tax` and `tax` are identical.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.transaction_reference` | object | facultatif | By adding this field you can automatically reconcile the newly imported invoice with a transaction. See documentation about [automatic payment matching](https://pennylane.readme.io/v2.0/docs/automating-payment-matching).<br>additionalProperties `false` |
| `$.transaction_reference.banking_provider` | string | obligatoire | The banking provider for the transaction<br>enum `["stripe","gocardless","bank","budgetinsight","helloasso"]`<br>example `"bank"` |
| `$.transaction_reference.provider_field_name` | string | obligatoire | Name of the field that you want to match<br>enum `["payment_id","charge_id","report_id","webid","label","order_id"]`<br>example `"label"` |
| `$.transaction_reference.provider_field_value` | string | obligatoire | Value that you want to match.<br><br>With `banking_provider: stripe` or `gocardless`, the value is matched exactly against the reference with the same `provider_field_name` synced from that provider, on non-archived transactions.<br><br>With `banking_provider: bank` and `provider_field_name: label`, every non-archived bank transaction whose label contains this value as a whole word is matched, including transactions already matched to other documents. The match is case-sensitive, and a value that does not start and end with a letter, digit or underscore never matches. Use a value that identifies a single payment, such as a card transaction ID: a merchant name or a short generic value groups unrelated documents together. Pass the literal string `invoice_number` to match on this invoice's number instead of a fixed value.<br><br>Other combinations do not match any transaction.<br><br>In all cases, matching runs again on every update of this invoice.<br>pattern `"^[^\\x00]*$"`<br>example `"invoice_number"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "amount": "100.00",
  "currency_amount": "100.00",
  "currency_amount_before_tax": "100.00",
  "currency_tax": "100.00",
  "date": "2021-01-30",
  "deadline": "2021-01-30",
  "external_reference": "FR123",
  "file_attachment_id": 123,
  "invoice_lines": [
    {
      "amount": "100.00",
      "currency_amount": "100.00",
      "currency_tax": "100.00",
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
      "imputation_dates": {
        "end_date": "2021-06-30",
        "start_date": "2020-06-30"
      },
      "label": "Demo label",
      "ledger_account_id": 1255,
      "ledger_entry_line": {
        "label": "Demo label"
      },
      "performance_period": {
        "end_date": "2023-07-31",
        "start_date": "2023-07-01"
      },
      "tax": "100.00",
      "vat_rate": "FR_200"
    }
  ],
  "invoice_number": "F20230001",
  "label": "Custom accounting label",
  "supplier_id": 456,
  "tax": "100.00",
  "transaction_reference": {
    "banking_provider": "bank",
    "provider_field_name": "label",
    "provider_field_value": "invoice_number"
  }
}
```

</details>

## Sortie

### HTTP 201

The imported supplier invoice

`application/json` → `réponse.JSON201` (type `*GetSupplierInvoiceStatus200`).

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
| `$.currency_tax` | string | obligatoire | Invoice taxable amount (in invoice currency)<br>example `"34.0"` |
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
| `$.performance_date` | string (date) ou null | obligatoire | End date of the performance period of the first invoice line that has one (ISO 8601). Used for German VAT accounting (Leistungsdatum). Null if the company does not use German accounting, or if no invoice line has a performance period.<br>example `"2023-07-31"` |
| `$.public_file_url` | string ou null | obligatoire | Public URL of the invoice file. The URL will expire after 30 minutes.<br>example `"https://app.pennylane.com/public/invoice/pdf?encrypted_id=bzjoVJe...3D%3D"` |
| `$.reconciled` | boolean | obligatoire | Whether the invoice has been reconciled or not<br>example `false` |
| `$.remaining_amount_with_tax` | string ou null | obligatoire | The remaining amount with VAT to pay for the invoice to be considered paid<br>example `"20.0"` |
| `$.remaining_amount_without_tax` | string ou null | obligatoire | The remaining amount without VAT to pay for the invoice to be considered paid<br>example `"16.0"` |
| `$.supplier` | object ou null | obligatoire | additionalProperties `false` |
| `$.supplier.id` | integer (int64) | obligatoire | example `456` |
| `$.supplier.url` | string | obligatoire | URL to get the supplier.<br>example `"https://app.pennylane.com/api/external/v2/suppliers/42"` |
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
  "performance_date": "2023-07-31",
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

### HTTP 409

Conflict with the current state of the target resource, such as when trying to create a resource that already exists.

`application/json` → `réponse.JSON409` (type `*SendByEmailCommercialDocumentStatus409`).

[Champs du payload d’erreur](erreurs.md#sendbyemailcommercialdocumentstatus409)

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
