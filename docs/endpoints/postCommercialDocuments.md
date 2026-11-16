# Create a commercial document

`POST /api/external/v2/commercial_documents`

This endpoint allows you to create a commercial document (proforma,
shipping order or purchasing order).

The document number is assigned by Pennylane, based on the numbering
configured for this type of document.
`GET /api/external/v2/numberings` lists whether `proforma`, `shipping_order`,
or `purchasing_order` is configured.


> ℹ️
> This endpoint requires the following scope: `commercial_documents:all`

[Source Pennylane](https://pennylane.readme.io/reference/postcommercialdocuments) · [Index du SDK](../operations.md)

## Appel Go

```go
PostCommercialDocumentsWithResponse(ctx context.Context, params *PostCommercialDocumentsParams, body PostCommercialDocumentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCommercialDocumentsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | title `"Commercial Document Request"`<br>additionalProperties `false` |
| `$.commercial_document_template_id` | integer (int64) | facultatif | The commercial document template ID<br>example `42` |
| `$.currency` | string | facultatif | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.customer_id` | integer (int64) | obligatoire | Customer identifier<br>example `42` |
| `$.date` | string (date) | obligatoire | Document issue date (ISO 8601)<br>example `"2023-08-30"` |
| `$.deadline` | string (date) | obligatoire | Document validity deadline (ISO 8601)<br>example `"2023-09-02"` |
| `$.discount` | object | facultatif | additionalProperties `false` |
| `$.discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.discount.value` | string | obligatoire | Discount value on the total amount before tax of the document<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.document_type` | string | obligatoire | enum `["proforma","shipping_order","purchasing_order"]` |
| `$.external_reference` | string | facultatif | A unique external reference you can provide to track this document. If not provided, Pennylane will generate an identifier for you.<br>example `"FR123"` |
| `$.invoice_line_sections` | array | facultatif |  |
| `$.invoice_line_sections[]` | object | par élément | additionalProperties `false` |
| `$.invoice_line_sections[].description` | string | facultatif | Description of the invoice line section<br>example `"Consulting for August"` |
| `$.invoice_line_sections[].rank` | integer | obligatoire | Defines the order in which sections will be displayed on the document. Each section has to have a different rank, starting from 1 and incrementing next values by 1<br>minimum `1`<br>example `1` |
| `$.invoice_line_sections[].title` | string | facultatif | Title of the invoice line section<br>example `"Consulting Services"` |
| `$.invoice_lines` | array | obligatoire | minItems `1` |
| `$.invoice_lines[]` | oneOf | par élément |  |
| `$.invoice_lines[] (oneOf 1)` | object | selon variante | Provide only the product_id to automatically populate label,<br>raw_currency_unit_price, unit and vat_rate.<br>You can override any of these fields by providing them explicitly.<br><br>title `"Product-based Invoice Line"`<br>additionalProperties `false` |
| `$.invoice_lines[] (oneOf 1).description` | string ou null | facultatif | The description of the invoice line<br>example `"Consulting services for August"` |
| `$.invoice_lines[] (oneOf 1).discount` | object | facultatif | additionalProperties `false` |
| `$.invoice_lines[] (oneOf 1).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.invoice_lines[] (oneOf 1).discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.invoice_lines[] (oneOf 1).label` | string | facultatif | Line item label<br>example `"Consulting fees"` |
| `$.invoice_lines[] (oneOf 1).product_id` | integer (int64) | obligatoire | The product ID. Auto-fills label, raw_currency_unit_price, unit and vat_rate from the product.<br>example `42` |
| `$.invoice_lines[] (oneOf 1).quantity` | number | obligatoire | Line item quantity (number of items)<br>example `2` |
| `$.invoice_lines[] (oneOf 1).raw_currency_unit_price` | string | facultatif | The unit price excluding taxes (undiscounted if a discount is set). Can be set up to 6 decimals.<br>example `"33.333334"` |
| `$.invoice_lines[] (oneOf 1).section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$.invoice_lines[] (oneOf 1).unit` | string | facultatif | Line item unit (type of unit)<br>example `"piece"` |
| `$.invoice_lines[] (oneOf 1).vat_rate` | string | facultatif | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.invoice_lines[] (oneOf 2)` | object | selon variante | Create an invoice line by providing all required fields explicitly<br>without referencing a product.<br><br>title `"Standard Invoice Line"`<br>additionalProperties `false` |
| `$.invoice_lines[] (oneOf 2).description` | string ou null | facultatif | The description of the invoice line<br>example `"Consulting services for August"` |
| `$.invoice_lines[] (oneOf 2).discount` | object | facultatif | additionalProperties `false` |
| `$.invoice_lines[] (oneOf 2).discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.invoice_lines[] (oneOf 2).discount.value` | string | obligatoire | Discount value on the total amount before tax of the line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.invoice_lines[] (oneOf 2).label` | string | obligatoire | Line item label<br>example `"Consulting fees"` |
| `$.invoice_lines[] (oneOf 2).quantity` | number | obligatoire | Line item quantity (number of items)<br>example `2` |
| `$.invoice_lines[] (oneOf 2).raw_currency_unit_price` | string | obligatoire | The unit price excluding taxes (undiscounted if a discount is set). Can be set up to 6 decimals.<br>example `"33.333334"` |
| `$.invoice_lines[] (oneOf 2).section_rank` | integer | facultatif | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$.invoice_lines[] (oneOf 2).unit` | string | obligatoire | Line item unit (type of unit)<br>example `"piece"` |
| `$.invoice_lines[] (oneOf 2).vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.language` | string | facultatif | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.pdf_description` | string ou null | facultatif | Document description<br>example `"Document description"` |
| `$.pdf_invoice_free_text` | string ou null | facultatif | For example, the contact details of the person to contact<br>example `"Additional free field"` |
| `$.pdf_invoice_subject` | string ou null | facultatif | Document title<br>example `"Document subject"` |
| `$.special_mention` | string ou null | facultatif | Additional details. Maximum 2,000 characters.<br>example `"Additional details"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "commercial_document_template_id": 42,
  "customer_id": 42,
  "date": "2023-08-30",
  "deadline": "2023-09-02",
  "discount": {
    "value": "100.00"
  },
  "external_reference": "FR123",
  "invoice_line_sections": [
    {
      "description": "Consulting for August",
      "rank": 1,
      "title": "Consulting Services"
    }
  ],
  "pdf_description": "Document description",
  "pdf_invoice_free_text": "Additional free field",
  "pdf_invoice_subject": "Document subject",
  "special_mention": "Additional details"
}
```

</details>

## Sortie

### HTTP 201

Renders the created commercial document

`application/json` → `réponse.JSON201` (type `*UpdateCommercialDocumentStatus200`).

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

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
