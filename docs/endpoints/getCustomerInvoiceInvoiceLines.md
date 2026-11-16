# List invoice lines for a customer invoice

`GET /api/external/v2/customer_invoices/{customer_invoice_id}/invoice_lines`

List invoice lines for a customer invoice

> ℹ️
> This endpoint requires one of the following scopes: `customer_invoices:all`, `customer_invoices:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcustomerinvoiceinvoicelines) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCustomerInvoiceInvoiceLinesWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceInvoiceLinesParams, reqEditors ...RequestEditorFn) (*GetCustomerInvoiceInvoiceLinesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_invoice_id` | `customerInvoiceId` | path | integer | obligatoire | <br>example `1234` |
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`, `rank`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns the list of invoice lines for a customer invoice

`application/json` → `réponse.JSON200` (type `*GetCustomerInvoiceInvoiceLinesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire | additionalProperties `false` |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].amount` | string | obligatoire | The total amount of the invoice lines in euros including taxes and deducting discounts<br>example `"50.4"` |
| `$.items[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].currency_amount` | string | obligatoire | The total amount of the invoice lines in the document's currency including taxes and deducting discounts.<br>If the currency is euro, amount and currency_amount are identical.<br><br>example `"50.4"` |
| `$.items[].currency_amount_before_tax` | string | obligatoire | Total amount before tax in currency<br>example `"30"` |
| `$.items[].currency_tax` | string | obligatoire | Total tax amount in currency<br>example `"10"` |
| `$.items[].custom_fields` | array | obligatoire |  |
| `$.items[].custom_fields[]` | object | par élément | additionalProperties `false` |
| `$.items[].custom_fields[].name` | string | obligatoire | example `"color"` |
| `$.items[].custom_fields[].value` | anyOf | obligatoire | Select fields use the option label. Multi-select fields use an array of option labels.<br>example `"red"` |
| `$.items[].custom_fields[].value (anyOf 1)` | string ou null | selon variante |  |
| `$.items[].custom_fields[].value (anyOf 2)` | array | selon variante |  |
| `$.items[].custom_fields[].value (anyOf 2)[]` | string | par élément |  |
| `$.items[].description` | string | obligatoire | Invoice line description<br>example `"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor"` |
| `$.items[].discount` | object | obligatoire | additionalProperties `false` |
| `$.items[].discount.type` | string | obligatoire | Discount type.<br>- absolute if it is an amount<br>- relative if it is a percentage<br><br>enum `["absolute","relative"]` |
| `$.items[].discount.value` | string | obligatoire | Discount value on the total amount before tax of the invoice<br>example `"25"` |
| `$.items[].id` | integer (int64) | obligatoire | Invoice line id<br>example `444` |
| `$.items[].imputation_dates` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].imputation_dates.end_date` | string (date) | obligatoire | End date of the imputation period (ISO 8601)<br>example `"2021-06-30"` |
| `$.items[].imputation_dates.start_date` | string (date) | obligatoire | Start date of the imputation period (ISO 8601)<br>example `"2020-06-30"` |
| `$.items[].label` | string | obligatoire | Invoice line label<br>example `"Demo label"` |
| `$.items[].performance_period` | object ou null | obligatoire | Date range during which the service or delivery was performed (ISO 8601). Used for German VAT accounting (Leistungsdatum).<br>additionalProperties `false` |
| `$.items[].performance_period.date_type` | oneOf | obligatoire |  |
| `$.items[].performance_period.date_type (oneOf 1)` | string | selon variante | Whether the performance period represents a single date or a range, and whether it is a delivery or service.<br>- delivery_date: single delivery date (end_date only)<br>- service_date: single service date (end_date only)<br>- delivery_period: delivery date range (start_date and end_date)<br>- service_period: service date range (start_date and end_date)<br><br>enum `["delivery_date","service_date","delivery_period","service_period"]` |
| `$.items[].performance_period.date_type (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.items[].performance_period.end_date` | string (date) | obligatoire | End date of the performance period (ISO 8601)<br>example `"2023-07-31"` |
| `$.items[].performance_period.start_date` | string (date) ou null | obligatoire | Start date of the performance period (ISO 8601). Null for single-date types.<br>example `"2023-07-01"` |
| `$.items[].product` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].product.id` | integer (int64) | obligatoire | Product id<br>example `3049` |
| `$.items[].product.url` | string | obligatoire | URL to get the product.<br>example `"https://app.pennylane.com/api/external/v2/products/42"` |
| `$.items[].quantity` | string | obligatoire | Invoice line item quantity (number of items)<br>example `"12"` |
| `$.items[].raw_currency_unit_price` | string | obligatoire | Unit price (excluding tax)<br>example `"5"` |
| `$.items[].section_rank` | integer ou null | obligatoire | Has to correspond to the rank number of a line items section in which the line item should be<br>example `1` |
| `$.items[].substance` | oneOf | obligatoire | Derived from the line's ledger account. Returns `goods` or `services` when unambiguous, or `null` otherwise (e.g. mixed partial-invoice lines). |
| `$.items[].substance (oneOf 1)` | string | selon variante | enum `["goods","services"]` |
| `$.items[].substance (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.items[].tax` | string | obligatoire | Total tax amount in euros<br>example `"10"` |
| `$.items[].unit` | string ou null | obligatoire | The kind of unit which applies to the amount<br>example `"piece"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "amount": "50.4",
      "created_at": "2023-08-30T10:08:08.146343Z",
      "currency_amount": "50.4",
      "currency_amount_before_tax": "30",
      "currency_tax": "10",
      "custom_fields": [
        {
          "name": "color",
          "value": "red"
        }
      ],
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
      "discount": {
        "value": "25"
      },
      "id": 444,
      "imputation_dates": {
        "end_date": "2021-06-30",
        "start_date": "2020-06-30"
      },
      "label": "Demo label",
      "performance_period": {
        "end_date": "2023-07-31",
        "start_date": "2023-07-01"
      },
      "product": {
        "id": 3049,
        "url": "https://app.pennylane.com/api/external/v2/products/42"
      },
      "quantity": "12",
      "raw_currency_unit_price": "5",
      "section_rank": 1,
      "tax": "10",
      "unit": "piece",
      "updated_at": "2023-08-30T10:08:08.146343Z",
      "vat_rate": "FR_200"
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
