# Create a customer invoice

`POST /api/external/v2/customer_invoices`

This endpoint allows you to create a draft or finalized customer
invoice or credit note


> ℹ️
> This endpoint requires the following scope: `customer_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/postcustomerinvoices) · [Index du SDK](../operations.md)

## Appel Go

```go
PostCustomerInvoicesWithResponse(ctx context.Context, params *PostCustomerInvoicesParams, body PostCustomerInvoicesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCustomerInvoicesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | anyOf | racine |  |
| `$ (anyOf 1)` | object | selon variante | title `"Draft Customer Invoice"`<br>additionalProperties `false` |
| `$ (anyOf 1).contract_reference` | string ou null | facultatif | Contract reference (e-invoice field BT-12). Maximum 50 characters.<br>maxLength `50`<br>example `"CT-2026-001"` |
| `$ (anyOf 1).currency` | string | facultatif | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$ (anyOf 1).customer_id` | integer (int64) | obligatoire | Customer identifier<br>example `42` |
| `$ (anyOf 1).customer_invoice_template_id` | integer (int64) | facultatif | The customer invoice template ID<br>example `42` |
| `$ (anyOf 1).date` | string (date) | obligatoire | Invoice date (ISO 8601)<br>example `"2023-08-30"` |
| `$ (anyOf 1).deadline` | string (date) | obligatoire | Invoice payment deadline (ISO 8601)<br>example `"2023-08-30"` |
| `$ (anyOf 1).discount` | object | facultatif | additionalProperties `false` |
| `$ (anyOf 1).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$ (anyOf 1).discount.value` | string | obligatoire | Discount value on the total amount before tax of the invoice<br>example `"25"` |
| `$ (anyOf 1).draft` | boolean | obligatoire | Indicates if the invoice is in draft status (has not been finalized)<br>enum `[true]` |
| `$ (anyOf 1).external_reference` | string | facultatif | A unique external reference you can provide to track this customer invoice. If not provided, Pennylane will generate an identifier for you.<br>example `"FR123"` |
| `$ (anyOf 1).installments` | array | facultatif | Optional list of installments for this invoice.<br>If not provided, a single installment matching the invoice total is created automatically.<br>The sum of all installments' `amount` values must equal the invoice total amount with tax.<br>Invoices issued by a billing subscription cannot have multiple installments.<br><br>> ℹ️<br>> This param requires a company plan that has access to Installments, otherwise responds with 403.<br><br>minItems `1`<br>maxItems `60` |
| `$ (anyOf 1).installments[]` | object | par élément | title `"Customer Invoice Installment"`<br>additionalProperties `false` |
| `$ (anyOf 1).installments[].amount` | string | obligatoire | Installment amount (in invoice currency)<br>pattern `"^\\d+(\\.\\d+)?$"`<br>example `"100.00"` |
| `$ (anyOf 1).installments[].deadline` | string (date) | obligatoire | Installment payment deadline (ISO 8601)<br>example `"2023-08-30"` |
| `$ (anyOf 1).installments[].payment_methods` | array | facultatif | uniqueItems `true` |
| `$ (anyOf 1).installments[].payment_methods[]` | string | par élément | Allowed payment methods to be used in installments from the API.<br>enum `["transfer","e_transfer","credit_card","sepa","cheque","cash","lcr","pro_account_pos"]` |
| `$ (anyOf 1).invoice_line_sections` | array | facultatif |  |
| `$ (anyOf 1).invoice_line_sections[]` | object | par élément | additionalProperties `false` |
| `$ (anyOf 1).invoice_line_sections[].description` | string | facultatif | Description of the invoice line section<br>example `"Description of the invoice line section"` |
| `$ (anyOf 1).invoice_line_sections[].rank` | integer | obligatoire | Defines the order in which sections will be displayed on the invoice. Each section has to have a different rank, starting from 1 and incrementing next values by 1<br>minimum `1`<br>example `1` |
| `$ (anyOf 1).invoice_line_sections[].title` | string | facultatif | Title of the invoice line section<br>example `"Invoice line section title"` |
| `$ (anyOf 1).invoice_lines` | array | obligatoire | minItems `1` |
| `$ (anyOf 1).invoice_lines[]` | oneOf | par élément |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 1)` | object | selon variante | Provide only the product_id to automatically populate label,<br>raw_currency_unit_price, unit, vat_rate, and ledger_account_id.<br>You can override any of these fields by providing them explicitly.<br><br>title `"Product-based Invoice Line"`<br>additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).custom_fields` | array | facultatif | **Experimental** — under active development, may change without notice. Do not rely on this field in production integrations. |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).custom_fields[]` | object | par élément | additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).custom_fields[].name` | string | obligatoire | The `api_name` of the custom field<br>example `"delivery_terms"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).custom_fields[].value` | anyOf | obligatoire | Format depends on the field type: plain string for `text`; numeric string for `number` (e.g. `"42.5"`); ISO 8601 date for `date` (e.g. `"2025-01-15"`); option label for `single_select`; array of option labels for `multi_select`.<br>example `"Delivered at place"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).custom_fields[].value (anyOf 1)` | string ou null | selon variante |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).custom_fields[].value (anyOf 2)` | array | selon variante |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).custom_fields[].value (anyOf 2)[]` | string | par élément |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).description` | string ou null | facultatif | The description of the invoice line |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).discount` | object | facultatif | additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).label` | string | facultatif | Line item label<br>example `"Demo label"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).ledger_account_id` | integer (int64) | facultatif | The ledger account ID<br>example `1255` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).performance_period` | oneOf | facultatif |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).performance_period (oneOf 1)` | object | selon variante | Date range during which the service or delivery was performed (ISO 8601). Used for German VAT accounting (Leistungsdatum). Only accepted for companies using German accounting, otherwise the request is rejected. Omit it to keep the current performance period, or set it to null to remove it.<br>additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).performance_period (oneOf 1).date_type` | string | obligatoire | Whether the performance period represents a single date or a range, and whether it is a delivery or service.<br>- delivery_date: single delivery date (end_date only)<br>- service_date: single service date (end_date only)<br>- delivery_period: delivery date range (start_date and end_date)<br>- service_period: service date range (start_date and end_date)<br><br>enum `["delivery_date","service_date","delivery_period","service_period"]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).performance_period (oneOf 1).end_date` | string (date) | obligatoire | End date of the performance period (ISO 8601)<br>example `"2023-07-31"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).performance_period (oneOf 1).start_date` | string (date) ou null | facultatif | Start date of the performance period (ISO 8601). Null for single-date types (`delivery_date`, `service_date`), required otherwise.<br>example `"2023-07-01"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).performance_period (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).product_id` | integer (int64) | obligatoire | The product ID. Auto-fills label, raw_currency_unit_price, unit, vat_rate, and ledger_account_id from the product.<br>example `42` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).quantity` | number | obligatoire | Line item quantity (number of items)<br>example `12` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).raw_currency_unit_price` | string | facultatif | The unit price excluding taxes (undiscounted if a discount is set). Can be set up to 6 decimals.<br>example `"33.333334"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).substance` | string | facultatif | Resolves the line's ledger account to the company's goods or services account.<br><br>When provided together with `product_id`, `product_id` auto-fill runs first. If the resulting ledger account does not match `substance`, the request returns 422.<br><br>When provided together with `ledger_account_id`, both must refer to the same account. If they do not match, the request returns 422.<br><br>enum `["goods","services"]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).unit` | string | facultatif | Line item unit (type of unit)<br>example `"piece"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 1).vat_rate` | string | facultatif | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2)` | object | selon variante | Create an invoice line by providing all required fields explicitly<br>without referencing a product.<br><br>title `"Standard Invoice Line"`<br>additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).custom_fields` | array | facultatif | **Experimental** — under active development, may change without notice. Do not rely on this field in production integrations. |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).custom_fields[]` | object | par élément | additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).custom_fields[].name` | string | obligatoire | The `api_name` of the custom field<br>example `"delivery_terms"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).custom_fields[].value` | anyOf | obligatoire | Format depends on the field type: plain string for `text`; numeric string for `number` (e.g. `"42.5"`); ISO 8601 date for `date` (e.g. `"2025-01-15"`); option label for `single_select`; array of option labels for `multi_select`.<br>example `"Delivered at place"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).custom_fields[].value (anyOf 1)` | string ou null | selon variante |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).custom_fields[].value (anyOf 2)` | array | selon variante |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).custom_fields[].value (anyOf 2)[]` | string | par élément |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).description` | string ou null | facultatif | The description of the invoice line |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).discount` | object | facultatif | additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).label` | string | obligatoire | Line item label<br>example `"Demo label"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).ledger_account_id` | integer (int64) | facultatif | The ledger account ID<br>example `1255` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).performance_period` | oneOf | facultatif |  |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).performance_period (oneOf 1)` | object | selon variante | Date range during which the service or delivery was performed (ISO 8601). Used for German VAT accounting (Leistungsdatum). Only accepted for companies using German accounting, otherwise the request is rejected. Omit it to keep the current performance period, or set it to null to remove it.<br>additionalProperties `false` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).performance_period (oneOf 1).date_type` | string | obligatoire | Whether the performance period represents a single date or a range, and whether it is a delivery or service.<br>- delivery_date: single delivery date (end_date only)<br>- service_date: single service date (end_date only)<br>- delivery_period: delivery date range (start_date and end_date)<br>- service_period: service date range (start_date and end_date)<br><br>enum `["delivery_date","service_date","delivery_period","service_period"]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).performance_period (oneOf 1).end_date` | string (date) | obligatoire | End date of the performance period (ISO 8601)<br>example `"2023-07-31"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).performance_period (oneOf 1).start_date` | string (date) ou null | facultatif | Start date of the performance period (ISO 8601). Null for single-date types (`delivery_date`, `service_date`), required otherwise.<br>example `"2023-07-01"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).performance_period (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).quantity` | number | obligatoire | Line item quantity (number of items)<br>example `12` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).raw_currency_unit_price` | string | obligatoire | The unit price excluding taxes (undiscounted if a discount is set). Can be set up to 6 decimals.<br>example `"33.333334"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).substance` | string | facultatif | Resolves the line's ledger account to the company's goods or services account.<br><br>When provided together with `ledger_account_id`, both must refer to the same account. If they do not match, the request returns 422.<br><br>enum `["goods","services"]` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).unit` | string | obligatoire | Line item unit (type of unit)<br>example `"piece"` |
| `$ (anyOf 1).invoice_lines[] (oneOf 2).vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$ (anyOf 1).label` | string ou null | facultatif | Custom label for the invoice used on accounting (ledger) entries. If not provided, Pennylane generates a label automatically.<br>maxLength `2000`<br>example `"Custom accounting label"` |
| `$ (anyOf 1).language` | string | facultatif | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$ (anyOf 1).pdf_description` | string ou null | facultatif | Invoice description. Maximum 5,000 characters.<br>example `"Invoice description"` |
| `$ (anyOf 1).pdf_invoice_free_text` | string ou null | facultatif | For example, the contact details of the person to contact<br>example `"Additional free field"` |
| `$ (anyOf 1).pdf_invoice_subject` | string ou null | facultatif | Invoice title<br>example `"Invoice title"` |
| `$ (anyOf 1).purchase_order_reference` | string ou null | facultatif | Buyer's purchase order reference (e-invoice field BT-13). Maximum 50 characters.<br>maxLength `50`<br>example `"PO-2026-001"` |
| `$ (anyOf 1).sale_type` | oneOf | facultatif | Type of sale printed on the invoice (Type de vente).<br>`goods` is a supply of goods, `services` is a supply of services,<br>and `goods_and_services` is both. `null` means the type has not been chosen.<br>This is not invoice line `substance`.<br>Omit this field to keep the invoice template default on create, or to leave<br>the invoice unchanged on update. Send `null` to clear it.<br>Finalized invoices cannot be updated with this field. |
| `$ (anyOf 1).sale_type (oneOf 1)` | string | selon variante | Type of sale printed on a customer invoice. This is not invoice line substance.<br>enum `["goods","services","goods_and_services"]` |
| `$ (anyOf 1).sale_type (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$ (anyOf 1).sales_order_reference` | string ou null | facultatif | Seller's sales order reference (e-invoice field BT-14). Maximum 50 characters.<br>maxLength `50`<br>example `"SO-2026-001"` |
| `$ (anyOf 1).special_mention` | string ou null | facultatif | Additional details. maximum 20,000 characters.<br>example `"Additional details"` |
| `$ (anyOf 2)` | object | selon variante | title `"Finalized Customer Invoice"`<br>additionalProperties `false` |
| `$ (anyOf 2).contract_reference` | string ou null | facultatif | Contract reference (e-invoice field BT-12). Maximum 50 characters.<br>maxLength `50`<br>example `"CT-2026-001"` |
| `$ (anyOf 2).currency` | string | facultatif | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$ (anyOf 2).customer_id` | integer (int64) | obligatoire | Customer identifier<br>example `42` |
| `$ (anyOf 2).customer_invoice_template_id` | integer (int64) | facultatif | The customer invoice template ID<br>example `42` |
| `$ (anyOf 2).date` | string (date) | obligatoire | Invoice date (ISO 8601)<br>example `"2023-08-30"` |
| `$ (anyOf 2).deadline` | string (date) | obligatoire | Invoice payment deadline (ISO 8601)<br>example `"2023-08-30"` |
| `$ (anyOf 2).discount` | object | facultatif | additionalProperties `false` |
| `$ (anyOf 2).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$ (anyOf 2).discount.value` | string | obligatoire | Discount value on the total amount before tax of the invoice<br>example `"25"` |
| `$ (anyOf 2).draft` | boolean | facultatif | Indicates if the invoice is in draft status (has not been finalized)<br>enum `[false]`<br>default `false` |
| `$ (anyOf 2).external_reference` | string | facultatif | A unique external reference you can provide to track this customer invoice. If not provided, Pennylane will generate an identifier for you.<br>example `"FR123"` |
| `$ (anyOf 2).installments` | array | facultatif | Optional list of installments for this invoice.<br>If not provided, a single installment matching the invoice total is created automatically.<br>The sum of all installments' `amount` values must equal the invoice total amount with tax.<br>Invoices issued by a billing subscription cannot have multiple installments.<br><br>> ℹ️<br>> This param requires a company plan that has access to Installments, otherwise responds with 403.<br><br>minItems `1`<br>maxItems `60` |
| `$ (anyOf 2).installments[]` | object | par élément | title `"Customer Invoice Installment"`<br>additionalProperties `false` |
| `$ (anyOf 2).installments[].amount` | string | obligatoire | Installment amount (in invoice currency)<br>pattern `"^\\d+(\\.\\d+)?$"`<br>example `"100.00"` |
| `$ (anyOf 2).installments[].deadline` | string (date) | obligatoire | Installment payment deadline (ISO 8601)<br>example `"2023-08-30"` |
| `$ (anyOf 2).installments[].payment_methods` | array | facultatif | uniqueItems `true` |
| `$ (anyOf 2).installments[].payment_methods[]` | string | par élément | Allowed payment methods to be used in installments from the API.<br>enum `["transfer","e_transfer","credit_card","sepa","cheque","cash","lcr","pro_account_pos"]` |
| `$ (anyOf 2).invoice_line_sections` | array | facultatif |  |
| `$ (anyOf 2).invoice_line_sections[]` | object | par élément | additionalProperties `false` |
| `$ (anyOf 2).invoice_line_sections[].description` | string | facultatif | Description of the invoice line section<br>example `"Description of the invoice line section"` |
| `$ (anyOf 2).invoice_line_sections[].rank` | integer | obligatoire | Defines the order in which sections will be displayed on the invoice. Each section has to have a different rank, starting from 1 and incrementing next values by 1<br>minimum `1`<br>example `1` |
| `$ (anyOf 2).invoice_line_sections[].title` | string | facultatif | Title of the invoice line section<br>example `"Invoice line section title"` |
| `$ (anyOf 2).invoice_lines` | array | obligatoire | minItems `1` |
| `$ (anyOf 2).invoice_lines[]` | oneOf | par élément |  |
| `$ (anyOf 2).invoice_lines[] (oneOf 1)` | object | selon variante | Provide only the product_id to automatically populate label,<br>raw_currency_unit_price, unit, vat_rate, and ledger_account_id.<br>You can override any of these fields by providing them explicitly.<br><br>title `"Product-based Invoice Line"`<br>additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).description` | string ou null | facultatif | The description of the invoice line |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).discount` | object | facultatif | additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).imputation_dates` | object ou null | facultatif | additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).imputation_dates.end_date` | string (date) | obligatoire | End date of the imputation period (ISO 8601)<br>example `"2021-06-30"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).imputation_dates.start_date` | string (date) | obligatoire | Start date of the imputation period (ISO 8601)<br>example `"2020-06-30"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).label` | string | facultatif | Line item label<br>example `"Demo label"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).ledger_account_id` | integer (int64) | facultatif | The ledger account ID<br>example `1255` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).performance_period` | oneOf | facultatif |  |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).performance_period (oneOf 1)` | object | selon variante | Date range during which the service or delivery was performed (ISO 8601). Used for German VAT accounting (Leistungsdatum). Only accepted for companies using German accounting, otherwise the request is rejected. Omit it to keep the current performance period, or set it to null to remove it.<br>additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).performance_period (oneOf 1).date_type` | string | obligatoire | Whether the performance period represents a single date or a range, and whether it is a delivery or service.<br>- delivery_date: single delivery date (end_date only)<br>- service_date: single service date (end_date only)<br>- delivery_period: delivery date range (start_date and end_date)<br>- service_period: service date range (start_date and end_date)<br><br>enum `["delivery_date","service_date","delivery_period","service_period"]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).performance_period (oneOf 1).end_date` | string (date) | obligatoire | End date of the performance period (ISO 8601)<br>example `"2023-07-31"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).performance_period (oneOf 1).start_date` | string (date) ou null | facultatif | Start date of the performance period (ISO 8601). Null for single-date types (`delivery_date`, `service_date`), required otherwise.<br>example `"2023-07-01"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).performance_period (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).product_id` | integer (int64) | obligatoire | The product ID. Auto-fills label, raw_currency_unit_price, unit, vat_rate, and ledger_account_id from the product.<br>example `42` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).quantity` | number | obligatoire | Line item quantity (number of items)<br>example `12` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).raw_currency_unit_price` | string | facultatif | The unit price excluding taxes<br>example `"33.333334"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).substance` | string | facultatif | Resolves the line's ledger account to the company's goods or services account.<br><br>When provided together with `product_id`, `product_id` auto-fill runs first. If the resulting ledger account does not match `substance`, the request returns 422.<br><br>When provided together with `ledger_account_id`, both must refer to the same account. If they do not match, the request returns 422.<br><br>enum `["goods","services"]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).unit` | string | facultatif | Line item unit (type of unit)<br>example `"piece"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 1).vat_rate` | string | facultatif | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2)` | object | selon variante | Create an invoice line by providing all required fields explicitly<br>without referencing a product.<br><br>title `"Standard Invoice Line"`<br>additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).description` | string ou null | facultatif | The description of the invoice line |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).discount` | object | facultatif | additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).imputation_dates` | object ou null | facultatif | additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).imputation_dates.end_date` | string (date) | obligatoire | End date of the imputation period (ISO 8601)<br>example `"2021-06-30"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).imputation_dates.start_date` | string (date) | obligatoire | Start date of the imputation period (ISO 8601)<br>example `"2020-06-30"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).label` | string | obligatoire | Line item label<br>example `"Demo label"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).ledger_account_id` | integer (int64) | facultatif | The ledger account ID<br>example `1255` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).performance_period` | oneOf | facultatif |  |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).performance_period (oneOf 1)` | object | selon variante | Date range during which the service or delivery was performed (ISO 8601). Used for German VAT accounting (Leistungsdatum). Only accepted for companies using German accounting, otherwise the request is rejected. Omit it to keep the current performance period, or set it to null to remove it.<br>additionalProperties `false` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).performance_period (oneOf 1).date_type` | string | obligatoire | Whether the performance period represents a single date or a range, and whether it is a delivery or service.<br>- delivery_date: single delivery date (end_date only)<br>- service_date: single service date (end_date only)<br>- delivery_period: delivery date range (start_date and end_date)<br>- service_period: service date range (start_date and end_date)<br><br>enum `["delivery_date","service_date","delivery_period","service_period"]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).performance_period (oneOf 1).end_date` | string (date) | obligatoire | End date of the performance period (ISO 8601)<br>example `"2023-07-31"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).performance_period (oneOf 1).start_date` | string (date) ou null | facultatif | Start date of the performance period (ISO 8601). Null for single-date types (`delivery_date`, `service_date`), required otherwise.<br>example `"2023-07-01"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).performance_period (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).quantity` | number | obligatoire | Line item quantity (number of items)<br>example `12` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).raw_currency_unit_price` | string | obligatoire | The unit price excluding taxes (undiscounted if a discount is set). Can be set up to 6 decimals.<br>example `"33.333334"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).substance` | string | facultatif | Resolves the line's ledger account to the company's goods or services account.<br><br>When provided together with `ledger_account_id`, both must refer to the same account. If they do not match, the request returns 422.<br><br>enum `["goods","services"]` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).unit` | string | obligatoire | Line item unit (type of unit)<br>example `"piece"` |
| `$ (anyOf 2).invoice_lines[] (oneOf 2).vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$ (anyOf 2).label` | string ou null | facultatif | Custom label for the invoice used on accounting (ledger) entries. If not provided, Pennylane generates a label automatically.<br>maxLength `2000`<br>example `"Custom accounting label"` |
| `$ (anyOf 2).language` | string | facultatif | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$ (anyOf 2).pdf_description` | string ou null | facultatif | Invoice description. Maximum 5,000 characters.<br>example `"Invoice description"` |
| `$ (anyOf 2).pdf_invoice_free_text` | string ou null | facultatif | For example, the contact details of the person to contact<br>example `"Additional free field"` |
| `$ (anyOf 2).pdf_invoice_subject` | string ou null | facultatif | Invoice title<br>example `"Invoice title"` |
| `$ (anyOf 2).purchase_order_reference` | string ou null | facultatif | Buyer's purchase order reference (e-invoice field BT-13). Maximum 50 characters.<br>maxLength `50`<br>example `"PO-2026-001"` |
| `$ (anyOf 2).sale_type` | oneOf | facultatif | Type of sale printed on the invoice (Type de vente).<br>`goods` is a supply of goods, `services` is a supply of services,<br>and `goods_and_services` is both. `null` means the type has not been chosen.<br>This is not invoice line `substance`.<br>Omit this field to keep the invoice template default on create, or to leave<br>the invoice unchanged on update. Send `null` to clear it.<br>Finalized invoices cannot be updated with this field. |
| `$ (anyOf 2).sale_type (oneOf 1)` | string | selon variante | Type of sale printed on a customer invoice. This is not invoice line substance.<br>enum `["goods","services","goods_and_services"]` |
| `$ (anyOf 2).sale_type (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$ (anyOf 2).sales_order_reference` | string ou null | facultatif | Seller's sales order reference (e-invoice field BT-14). Maximum 50 characters.<br>maxLength `50`<br>example `"SO-2026-001"` |
| `$ (anyOf 2).special_mention` | string ou null | facultatif | Additional details. maximum 20,000 characters.<br>example `"Additional details"` |
| `$ (anyOf 2).transaction_reference` | object | facultatif | By adding this field you can automatically reconcile the newly created invoice with a transaction. See documentation about [automatic payment matching](https://pennylane.readme.io/v2.0/docs/automating-payment-matching).<br>additionalProperties `false` |
| `$ (anyOf 2).transaction_reference.banking_provider` | string | obligatoire | The banking provider for the transaction<br>enum `["stripe","gocardless","bank","budgetinsight","helloasso"]`<br>example `"bank"` |
| `$ (anyOf 2).transaction_reference.provider_field_name` | string | obligatoire | Name of the field that you want to match<br>enum `["payment_id","charge_id","report_id","webid","label","order_id"]`<br>example `"label"` |
| `$ (anyOf 2).transaction_reference.provider_field_value` | string | obligatoire | Value that you want to match.<br><br>With `banking_provider: stripe` or `gocardless`, the value is matched exactly against the reference with the same `provider_field_name` synced from that provider, on non-archived transactions.<br><br>With `banking_provider: bank` and `provider_field_name: label`, every non-archived bank transaction whose label contains this value as a whole word is matched, including transactions already matched to other documents. The match is case-sensitive, and a value that does not start and end with a letter, digit or underscore never matches. Use a value that identifies a single payment, such as a card transaction ID: a merchant name or a short generic value groups unrelated documents together. Pass the literal string `invoice_number` to match on this invoice's number instead of a fixed value.<br><br>Other combinations do not match any transaction.<br><br>In all cases, matching runs again on every update of this invoice.<br>example `"invoice_number"` |

## Sortie

### HTTP 201

Renders the created finalized invoice

`application/json` → `réponse.JSON201` (type `*GetCustomerInvoiceStatus200`).

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
