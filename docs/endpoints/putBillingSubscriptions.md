# Update a billing subscription

`PUT /api/external/v2/billing_subscriptions/{id}`

Update a billing subscription

> ℹ️
> This endpoint requires the following scope: `billing_subscriptions:all`

[Source Pennylane](https://pennylane.readme.io/reference/putbillingsubscriptions) · [Index du SDK](../operations.md)

## Appel Go

```go
PutBillingSubscriptionsWithResponse(ctx context.Context, id int, params *PutBillingSubscriptionsParams, body PutBillingSubscriptionsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutBillingSubscriptionsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | The ID of the billing subscription to retrieve<br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.customer_id` | integer (int64) | facultatif | Customer identifier<br>example `42` |
| `$.customer_invoice_data` | object | facultatif | additionalProperties `false` |
| `$.customer_invoice_data.currency` | string | facultatif | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.customer_invoice_data.customer_invoice_template_id` | integer (int64) | facultatif | The customer invoice template ID<br>example `42` |
| `$.customer_invoice_data.discount` | object | facultatif | additionalProperties `false` |
| `$.customer_invoice_data.discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.customer_invoice_data.discount.value` | string | obligatoire | Discount value on the total amount before tax<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.customer_invoice_data.invoice_line_sections` | object | facultatif | Add, update, delete invoice line sections.<br>additionalProperties `false` |
| `$.customer_invoice_data.invoice_line_sections.create` | array | facultatif |  |
| `$.customer_invoice_data.invoice_line_sections.create[]` | object | par élément | additionalProperties `false` |
| `$.customer_invoice_data.invoice_line_sections.create[].description` | string | facultatif | Description of the invoice line section<br>example `"Invoice line section description"` |
| `$.customer_invoice_data.invoice_line_sections.create[].rank` | integer | obligatoire | Defines the order in which sections will be displayed on the invoice. Each section has to have a different rank, starting from 1 and incrementing next values by 1<br>minimum `1`<br>example `1` |
| `$.customer_invoice_data.invoice_line_sections.create[].title` | string | facultatif | Title of the invoice line section<br>example `"Invoice line section title"` |
| `$.customer_invoice_data.invoice_line_sections.delete` | array | facultatif |  |
| `$.customer_invoice_data.invoice_line_sections.delete[]` | object | par élément | additionalProperties `false` |
| `$.customer_invoice_data.invoice_line_sections.delete[].id` | integer (int64) | obligatoire | ID of the invoice line section<br>example `12` |
| `$.customer_invoice_data.invoice_line_sections.update` | array | facultatif |  |
| `$.customer_invoice_data.invoice_line_sections.update[]` | object | par élément | additionalProperties `false` |
| `$.customer_invoice_data.invoice_line_sections.update[].description` | string ou null | facultatif | Description of the invoice line section<br>example `"Invoice line section description"` |
| `$.customer_invoice_data.invoice_line_sections.update[].id` | integer (int64) | obligatoire | ID of the invoice line section<br>example `12` |
| `$.customer_invoice_data.invoice_line_sections.update[].rank` | integer | facultatif | Defines the order in which sections will be displayed on the invoice. Each section has to have a different rank, starting from 1 and incrementing next values by 1<br>minimum `1`<br>example `1` |
| `$.customer_invoice_data.invoice_line_sections.update[].title` | string ou null | facultatif | Title of the invoice line section<br>example `"Invoice line section title"` |
| `$.customer_invoice_data.invoice_lines` | object | facultatif | Add, update, delete invoice lines.<br>additionalProperties `false` |
| `$.customer_invoice_data.invoice_lines.create` | array | facultatif |  |
| `$.customer_invoice_data.invoice_lines.create[]` | object | par élément | additionalProperties `false` |
| `$.customer_invoice_data.invoice_lines.create[].description` | string ou null | facultatif | The description of the invoice line |
| `$.customer_invoice_data.invoice_lines.create[].discount` | object | facultatif | additionalProperties `false` |
| `$.customer_invoice_data.invoice_lines.create[].discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.customer_invoice_data.invoice_lines.create[].discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.customer_invoice_data.invoice_lines.create[].label` | string | obligatoire | Line item label<br>example `"Demo label"` |
| `$.customer_invoice_data.invoice_lines.create[].ledger_account_id` | integer (int64) | facultatif | The ledger account ID<br>example `1255` |
| `$.customer_invoice_data.invoice_lines.create[].product_id` | integer (int64) | facultatif | The product ID<br>example `42` |
| `$.customer_invoice_data.invoice_lines.create[].quantity` | number | obligatoire | Line item quantity (number of items)<br>example `12` |
| `$.customer_invoice_data.invoice_lines.create[].raw_currency_unit_price` | string | obligatoire | The unit price excluding taxes (undiscounted if a discount is set). Can be set up to 6 decimals.<br>example `"33.333334"` |
| `$.customer_invoice_data.invoice_lines.create[].section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$.customer_invoice_data.invoice_lines.create[].unit` | string | obligatoire | Line item unit (type of unit)<br>example `"piece"` |
| `$.customer_invoice_data.invoice_lines.create[].vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.customer_invoice_data.invoice_lines.delete` | array | facultatif |  |
| `$.customer_invoice_data.invoice_lines.delete[]` | object | par élément | additionalProperties `false` |
| `$.customer_invoice_data.invoice_lines.delete[].id` | integer (int64) | obligatoire | ID of the invoice line<br>example `42` |
| `$.customer_invoice_data.invoice_lines.update` | array | facultatif |  |
| `$.customer_invoice_data.invoice_lines.update[]` | object | par élément | additionalProperties `false` |
| `$.customer_invoice_data.invoice_lines.update[].description` | string ou null | facultatif | The description of the invoice line |
| `$.customer_invoice_data.invoice_lines.update[].discount` | object | facultatif | additionalProperties `false` |
| `$.customer_invoice_data.invoice_lines.update[].discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.customer_invoice_data.invoice_lines.update[].discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.customer_invoice_data.invoice_lines.update[].id` | integer (int64) | obligatoire | ID of the invoice line<br>example `42` |
| `$.customer_invoice_data.invoice_lines.update[].label` | string | facultatif | Line item label<br>example `"Demo label"` |
| `$.customer_invoice_data.invoice_lines.update[].ledger_account_id` | integer (int64) | facultatif | The ledger account ID<br>example `1255` |
| `$.customer_invoice_data.invoice_lines.update[].product_id` | integer (int64) | facultatif | The product ID<br>example `42` |
| `$.customer_invoice_data.invoice_lines.update[].quantity` | number | facultatif | Line item quantity (number of items)<br>example `12` |
| `$.customer_invoice_data.invoice_lines.update[].raw_currency_unit_price` | string | facultatif | The unit price excluding taxes (undiscounted if a discount is set). Can be set up to 6 decimals.<br>example `"33.333334"` |
| `$.customer_invoice_data.invoice_lines.update[].section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$.customer_invoice_data.invoice_lines.update[].unit` | string | facultatif | Line item unit (type of unit)<br>example `"piece"` |
| `$.customer_invoice_data.invoice_lines.update[].vat_rate` | string | facultatif | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.customer_invoice_data.language` | string | facultatif | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.customer_invoice_data.pdf_description` | string ou null | facultatif | Invoice description<br>example `"Invoice description"` |
| `$.customer_invoice_data.pdf_invoice_free_text` | string ou null | facultatif | For example, the contact details of the person to contact<br>example `"Additional free field"` |
| `$.customer_invoice_data.pdf_invoice_subject` | string ou null | facultatif | Invoice title<br>example `"Invoice title"` |
| `$.customer_invoice_data.special_mention` | string ou null | facultatif | Additional details<br>example `"Additional details"` |
| `$.label` | string ou null | facultatif |  |
| `$.mode` | object | facultatif | additionalProperties `false` |
| `$.mode (anyOf 1)` | object | selon variante | Use this mode if you need generated invoices to be sent by email.<br>title `"Email"`<br>additionalProperties `false` |
| `$.mode (anyOf 1).email_settings` | object | obligatoire | Email settings for sending invoices.<br>The total number of email addresses must not exceed 25.<br>An email template with both a subject and body content must be configured in the company settings.<br>A GoCardless subscription requires a dedicated GoCardless email template configured in the company settings.<br><br>additionalProperties `false` |
| `$.mode (anyOf 1).email_settings.recipients` | array | facultatif | Primary email recipients for invoice delivery.<br>uniqueItems `true`<br>example `["customer@example.com","accounting@example.com"]` |
| `$.mode (anyOf 1).email_settings.recipients[]` | string (email) | par élément |  |
| `$.mode (anyOf 1).type` | string | obligatoire | - `email`: generated invoices will be finalized<br>and sent by email to the recipients configured on the subscription<br><br>enum `["email"]` |
| `$.mode (anyOf 2)` | object | selon variante | title `"Other"`<br>additionalProperties `false` |
| `$.mode (anyOf 2).type` | string | obligatoire | - `awaiting_validation`: generated invoices will be in draft<br>- `finalized`: generated invoices will be finalized<br><br>enum `["awaiting_validation","finalized"]` |
| `$.payment_conditions` | string | facultatif | enum `["upon_receipt","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$.payment_method` | string | facultatif | Payment method<br>offline means the subscription is not linked to a payment method<br>gocardless_direct-debit means at each new occurrence the client will be automatically debited thanks to GoCardless. To do so, you need a GoCardless account properly configured.<br>pro_account_sepa_core means the customer must have an email, unless it already has an enabled SEPA Core mandate or a pending SEPA Core mandate request.<br><br>enum `["offline","gocardless_direct_debit","pro_account_sepa_core"]` |
| `$.recurring_rule` | allOf | facultatif |  |
| `$.recurring_rule (allOf 1)` | object | selon variante | additionalProperties `true` |
| `$.recurring_rule (allOf 1).interval` | integer | facultatif | minimum `1` |
| `$.recurring_rule (allOf 2)` | object | selon variante | additionalProperties `false` |
| `$.recurring_rule (allOf 2) (anyOf 1)` | object | selon variante | title `"Yearly recurrence"`<br>additionalProperties `false` |
| `$.recurring_rule (allOf 2) (anyOf 1).count` | integer | facultatif | Number of occurrences<br>example `12` |
| `$.recurring_rule (allOf 2) (anyOf 1).interval` | integer | facultatif | Interval between each occurrence in years.<br>minimum `1` |
| `$.recurring_rule (allOf 2) (anyOf 1).type` | string | obligatoire | enum `["yearly"]` |
| `$.recurring_rule (allOf 2) (anyOf 2)` | object | selon variante | title `"Monthly recurrence"`<br>additionalProperties `false` |
| `$.recurring_rule (allOf 2) (anyOf 2).count` | integer | facultatif | Number of occurrences<br>example `12` |
| `$.recurring_rule (allOf 2) (anyOf 2).day_of_month` | number | facultatif | minimum `1`<br>maximum `31` |
| `$.recurring_rule (allOf 2) (anyOf 2).interval` | integer | facultatif | Interval between each occurrence in months.<br>minimum `1` |
| `$.recurring_rule (allOf 2) (anyOf 2).type` | string | obligatoire | enum `["monthly"]` |
| `$.recurring_rule (allOf 2) (anyOf 3)` | object | selon variante | title `"Weekly recurrence"`<br>additionalProperties `false` |
| `$.recurring_rule (allOf 2) (anyOf 3).count` | integer | facultatif | Number of occurrences<br>example `12` |
| `$.recurring_rule (allOf 2) (anyOf 3).day_of_week` | string | facultatif | enum `["monday","tuesday","wednesday","thursday","friday","saturday","sunday"]` |
| `$.recurring_rule (allOf 2) (anyOf 3).interval` | integer | facultatif | Interval between each occurrence in weeks.<br>minimum `1` |
| `$.recurring_rule (allOf 2) (anyOf 3).type` | string | obligatoire | enum `["weekly"]` |
| `$.stop` | boolean | facultatif | To stop an `in_progress` billing subscription use value `true`.<br>To resume a `stopped` billing subscription use value `false` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "customer_id": 42,
  "customer_invoice_data": {
    "customer_invoice_template_id": 42,
    "discount": {
      "value": "100.00"
    },
    "invoice_line_sections": {
      "create": [
        {
          "description": "Invoice line section description",
          "rank": 1,
          "title": "Invoice line section title"
        }
      ],
      "delete": [
        {
          "id": 12
        }
      ],
      "update": [
        {
          "description": "Invoice line section description",
          "id": 12,
          "rank": 1,
          "title": "Invoice line section title"
        }
      ]
    },
    "invoice_lines": {
      "create": [
        {
          "discount": {
            "value": "100.00"
          },
          "label": "Demo label",
          "ledger_account_id": 1255,
          "product_id": 42,
          "quantity": 12,
          "raw_currency_unit_price": "33.333334",
          "section_rank": 1,
          "unit": "piece",
          "vat_rate": "FR_200"
        }
      ],
      "delete": [
        {
          "id": 42
        }
      ],
      "update": [
        {
          "discount": {
            "value": "100.00"
          },
          "id": 42,
          "label": "Demo label",
          "ledger_account_id": 1255,
          "product_id": 42,
          "quantity": 12,
          "raw_currency_unit_price": "33.333334",
          "section_rank": 1,
          "unit": "piece",
          "vat_rate": "FR_200"
        }
      ]
    },
    "pdf_description": "Invoice description",
    "pdf_invoice_free_text": "Additional free field",
    "pdf_invoice_subject": "Invoice title",
    "special_mention": "Additional details"
  }
}
```

</details>

## Sortie

### HTTP 200

Renders the updated billing subscription

`application/json` → `réponse.JSON200` (type `*GetBillingSubscriptionStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.activated_at` | string (date-time) ou null | obligatoire | The exact timestamp for when the subscription was activated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.created_at` | string (date-time) | obligatoire | The time the subscription has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.customer.id` | integer (int64) | obligatoire | example `42` |
| `$.customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.customer_invoice_data` | object | obligatoire | additionalProperties `false` |
| `$.customer_invoice_data.amount` | string | obligatoire | Invoice amount (total value of the invoice in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"230.32"` |
| `$.customer_invoice_data.currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.customer_invoice_data.currency_amount` | string | obligatoire | Invoice currency amount (total value of the invoice in the currency of the invoice)<br>example `"230.32"` |
| `$.customer_invoice_data.currency_amount_before_tax` | string | obligatoire | Invoice currency amount before tax (total value before tax of the invoice in the currency of the invoice)<br>example `"196.32"` |
| `$.customer_invoice_data.currency_tax` | string | obligatoire | Invoice taxable amount (in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"34.0"` |
| `$.customer_invoice_data.customer_invoice_template` | object ou null | obligatoire | additionalProperties `false` |
| `$.customer_invoice_data.customer_invoice_template.id` | integer (int64) | obligatoire |  |
| `$.customer_invoice_data.discount` | object | obligatoire | additionalProperties `false` |
| `$.customer_invoice_data.discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.customer_invoice_data.discount.value` | string ou null | obligatoire | Discount value on the total amount before tax<br>example `"25"` |
| `$.customer_invoice_data.exchange_rate` | string | obligatoire | Invoice exchange rate (used to convert the invoice to euros. If the invoice currency is euro it will be 1.0)<br>pattern `"^-?\\d+(\\.\\d+)?$"`<br>example `"100.00"` |
| `$.customer_invoice_data.invoice_line_sections` | object | obligatoire | additionalProperties `false` |
| `$.customer_invoice_data.invoice_line_sections.url` | string | obligatoire | URL to get the invoice line sections of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/billing_subscriptions/42/invoice_line_sections"` |
| `$.customer_invoice_data.invoice_lines` | object | obligatoire | additionalProperties `false` |
| `$.customer_invoice_data.invoice_lines.url` | string | obligatoire | URL to get the invoice lines of the invoice.<br>example `"https://app.pennylane.com/api/external/v2/billing_subscriptions/42/invoice_lines"` |
| `$.customer_invoice_data.label` | string ou null | obligatoire | Invoice label<br>example `"Demo label"` |
| `$.customer_invoice_data.language` | string | obligatoire | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.customer_invoice_data.pdf_description` | string ou null | obligatoire | example `"Invoice description"` |
| `$.customer_invoice_data.pdf_invoice_free_text` | string | obligatoire | example `"Thanks for paying this invoice"` |
| `$.customer_invoice_data.pdf_invoice_subject` | string | obligatoire | example `"Invoice subject"` |
| `$.customer_invoice_data.special_mention` | string ou null | obligatoire | Additional details<br>example `"Additional details"` |
| `$.email_settings` | object ou null | obligatoire | additionalProperties `false` |
| `$.email_settings.billing_email_template` | object ou null | obligatoire | additionalProperties `false` |
| `$.email_settings.billing_email_template.id` | integer (int64) | obligatoire |  |
| `$.email_settings.billing_email_template.label` | string | obligatoire |  |
| `$.email_settings.recipients` | array | obligatoire | Primary email recipients for invoice delivery |
| `$.email_settings.recipients[]` | string (email) | par élément |  |
| `$.finish` | string (date) ou null | obligatoire | The date of the last occurrence of the subscription<br>example `"2023-12-31"` |
| `$.id` | integer (int64) | obligatoire | Billing subscription identifier<br>example `42` |
| `$.label` | string ou null | obligatoire |  |
| `$.mode` | string | obligatoire | - `awaiting_validation`: generated invoices will be in draft<br>- `finalized`: generated invoices will be finalized<br>- `email`: generated invoices will be finalized and sent by email to the recipients configured on the subscription<br><br>enum `["awaiting_validation","finalized","email"]` |
| `$.next_occurrence` | string (date) ou null | obligatoire | The date for the next subscription renewal<br>example `"2023-01-01"` |
| `$.payment_conditions` | oneOf | obligatoire |  |
| `$.payment_conditions (oneOf 1)` | string | selon variante | enum `["upon_receipt","7_days","15_days","30_days","30_days_end_of_month","45_days","45_days_end_of_month","60_days"]` |
| `$.payment_conditions (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.payment_method` | string | obligatoire | Payment method<br>offline means the subscription is not linked to a payment method<br>gocardless_direct-debit means at each new occurrence the client will be automatically debited thanks to GoCardless. To do so, you need a GoCardless account properly configured.<br>pro_account_sepa_core means the customer must have an email, unless it already has an enabled SEPA Core mandate or a pending SEPA Core mandate request.<br><br>enum `["offline","gocardless_direct_debit","pro_account_sepa_core"]` |
| `$.prev_occurrence` | string (date) ou null | obligatoire | The date of the previous subscription renewal<br>example `"2023-01-01"` |
| `$.recurring_rule` | object | obligatoire | additionalProperties `false` |
| `$.recurring_rule.count` | integer ou null | obligatoire | Number of occurrences of the recurring rule<br>example `12` |
| `$.recurring_rule.day` | array ou null | obligatoire |  |
| `$.recurring_rule.day[]` | number | par élément |  |
| `$.recurring_rule.day_of_month` | array ou null | obligatoire |  |
| `$.recurring_rule.day_of_month[]` | number | par élément |  |
| `$.recurring_rule.interval` | integer ou null | obligatoire |  |
| `$.recurring_rule.month_of_year` | array ou null | obligatoire |  |
| `$.recurring_rule.month_of_year[]` | number | par élément |  |
| `$.recurring_rule.rule_type` | string | obligatoire | enum `["weekly","monthly","yearly"]`<br>example `"monthly"` |
| `$.recurring_rule.until` | string ou null | obligatoire |  |
| `$.recurring_rule.week_start` | number ou null | obligatoire |  |
| `$.start` | string (date) ou null | obligatoire | The subscription start date<br>example `"2023-01-01"` |
| `$.status` | string | obligatoire | enum `["draft","stopped","finished","pending","not_started","in_progress"]` |
| `$.stopped_at` | string (date-time) ou null | obligatoire | The exact timestamp when the subscription was terminated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the subscription has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
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
