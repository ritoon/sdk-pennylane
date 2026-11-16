# Import a purchase order.

`POST /api/external/v2/purchase_requests/imports`

Import a purchase order. This will create a purchase request with an existing purchase order attached.
The purchase request will be **automatically validated**.


> ℹ️
> This endpoint requires the following scope: `purchase_requests:all`

[Source Pennylane](https://pennylane.readme.io/reference/createpurchaserequestimport) · [Index du SDK](../operations.md)

## Appel Go

```go
CreatePurchaseRequestImportWithResponse(ctx context.Context, params *CreatePurchaseRequestImportParams, body CreatePurchaseRequestImportJSONRequestBody, reqEditors ...RequestEditorFn) (*CreatePurchaseRequestImportResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.amount` | string | facultatif | Purchase request amount in euros (total value of the purchase request in euros). If the currency is euro, `currency_amount` and `amount` are identical.<br>example `"120.00"` |
| `$.currency` | string | facultatif | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.currency_amount` | string | obligatoire | Purchase request currency amount (total value of the purchase request in the currency of the purchase request)<br>example `"120.00"` |
| `$.currency_amount_before_tax` | string | obligatoire | Purchase request currency amount before tax (total value before tax of the purchase request in the currency of the purchase request)<br>example `"100.00"` |
| `$.currency_tax` | string | obligatoire | Purchase request taxable amount (in purchase request currency)<br>example `"20.00"` |
| `$.delivery_address` | object | obligatoire | additionalProperties `false` |
| `$.delivery_address.address` | string | obligatoire | Delivery address<br>example `"8 rue de la paix"` |
| `$.delivery_address.city` | string | obligatoire | City of the delivery address<br>example `"Paris"` |
| `$.delivery_address.country_alpha2` | string | obligatoire | Country of the delivery address (ISO 3166-1 alpha-2)<br>example `"FR"` |
| `$.delivery_address.postal_code` | string | obligatoire | Postal code of the delivery address<br>example `"75002"` |
| `$.estimated_delivery_date` | string (date) | facultatif | Estimated delivery date (ISO 8601)<br>example `"2023-08-30"` |
| `$.file_attachment_id` | integer (int64) | obligatoire | File attachment id. This file attachment must contain the purchase order to import.<br>example `42` |
| `$.purchase_order_number` | string | obligatoire | Purchase order number<br>example `"PR20230001"` |
| `$.purchase_request_lines` | array | obligatoire | minItems `1` |
| `$.purchase_request_lines[]` | object | par élément | additionalProperties `false` |
| `$.purchase_request_lines[].amount` | string | facultatif | Purchase request line amount in euros. If the currency is euro, `currency_amount` and `amount` are identical.<br>example `"120.00"` |
| `$.purchase_request_lines[].currency_amount` | string | obligatoire | Purchase request line total amount in the currency of the purchase request.<br>example `"120.00"` |
| `$.purchase_request_lines[].currency_tax` | string | obligatoire | Purchase request line taxable amount (in purchase request currency)<br>example `"20.00"` |
| `$.purchase_request_lines[].description` | string ou null | facultatif | The description of the purchase request line |
| `$.purchase_request_lines[].label` | string | obligatoire | Line item label<br>example `"Demo label"` |
| `$.purchase_request_lines[].quantity` | number | obligatoire | Line item quantity (number of items)<br>example `12` |
| `$.purchase_request_lines[].tax` | string | facultatif | Purchase request line taxable amount (in euros). If the currency is euro, currency_tax and tax are identical.<br>example `"20.00"` |
| `$.purchase_request_lines[].unit` | string | obligatoire | enum `["piece","hour","day","month","kilogram","m2","m3","ton","mg","no_unit"]` |
| `$.purchase_request_lines[].unit_price` | string | obligatoire | The unit price excluding taxes<br>example `"8.34"` |
| `$.purchase_request_lines[].vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.reason` | string | obligatoire | Reason of the purchase request<br>example `"I need a computer to be able to create purchase requests"` |
| `$.supplier_id` | integer (int64) | obligatoire | Supplier identifier<br>example `42` |
| `$.tax` | string | facultatif | Purchase request taxable amount (in euros). If the currency is euro, `currency_tax` and `tax` are identical.<br>example `"20.00"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "amount": "120.00",
  "currency_amount": "120.00",
  "currency_amount_before_tax": "100.00",
  "currency_tax": "20.00",
  "delivery_address": {
    "address": "8 rue de la paix",
    "city": "Paris",
    "country_alpha2": "FR",
    "postal_code": "75002"
  },
  "estimated_delivery_date": "2023-08-30",
  "file_attachment_id": 42,
  "purchase_order_number": "PR20230001",
  "purchase_request_lines": [
    {
      "amount": "120.00",
      "currency_amount": "120.00",
      "currency_tax": "20.00",
      "label": "Demo label",
      "quantity": 12,
      "tax": "20.00",
      "unit_price": "8.34",
      "vat_rate": "FR_200"
    }
  ],
  "reason": "I need a computer to be able to create purchase requests",
  "supplier_id": 42,
  "tax": "20.00"
}
```

</details>

## Sortie

### HTTP 201

The imported Purchase Request

`application/json` → `réponse.JSON201` (type `*GetPurchaseRequestStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.amount` | non spécifié (string) | obligatoire | example `"1234.56"` |
| `$.created_at` | string (date-time) | obligatoire | The time the purchase request has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.currency_amount` | string | obligatoire | Purchase Request currency amount (total value of the purchase request in the currency of the purchase request)<br>example `"1234.56"` |
| `$.currency_amount_before_tax` | string | obligatoire | Purchase Request currency amount before tax (total value before tax of the purchase request in the currency of the purchase request)<br>example `"1200.23"` |
| `$.currency_tax` | string | obligatoire | Purchase request taxable amount (in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"34.33"` |
| `$.delivery_address` | object | obligatoire | additionalProperties `false` |
| `$.delivery_address.address` | string ou null | obligatoire | example `"8 rue de la paix"` |
| `$.delivery_address.city` | string ou null | obligatoire | example `"Paris"` |
| `$.delivery_address.country_alpha2` | string ou null | obligatoire | example `"FR"` |
| `$.delivery_address.postal_code` | string ou null | obligatoire | example `"75002"` |
| `$.estimated_delivery_date` | string (date) ou null | obligatoire | Estimated delivery date for the purchase request<br>example `"2023-12-31"` |
| `$.exchange_rate` | string | obligatoire | Purchase request exchange rate (used to convert the purchase request to euros. If the purchase request currency is euro it will be 1.0)<br>example `"1.0"` |
| `$.id` | integer (int64) | obligatoire | Purchase request identifier<br>example `1` |
| `$.linked_invoices` | object | obligatoire | additionalProperties `false` |
| `$.linked_invoices.items` | array | obligatoire | List of linked supplier invoices |
| `$.linked_invoices.items[]` | object | par élément | additionalProperties `false` |
| `$.linked_invoices.items[].id` | integer (int64) | obligatoire | example `42` |
| `$.linked_invoices.items[].url` | string | obligatoire | URL to get the supplier invoice.<br>example `"https://app.pennylane.com/api/external/v2/supplier_invoices/42"` |
| `$.purchase_order` | object ou null | obligatoire | additionalProperties `false` |
| `$.purchase_order.filename` | string | obligatoire | example `"Bon-de-commande_n°2025-1.pdf"` |
| `$.purchase_order.url` | string | obligatoire | example `"https://www.pennylane.com/Bon-de-commande_n%C2%B02025-1.pdf"` |
| `$.purchase_order_number` | string ou null | obligatoire | Purchase order number<br>example `"PR20230001"` |
| `$.reason` | string | obligatoire | Reason of the request<br>example `"I need a computer to be able to create purchase requests"` |
| `$.status` | string | obligatoire | enum `["to_be_validated","approved","rejected","invoiced"]` |
| `$.supplier` | object | obligatoire | additionalProperties `false` |
| `$.supplier.id` | integer (int64) | obligatoire | example `2` |
| `$.supplier.url` | string (uri) | obligatoire | URL to get the supplier.<br>example `"https://app.pennylane.com/api/external/v2/suppliers/2"` |
| `$.tax` | string | obligatoire | Purchase request taxable amount (in purchase request currency)<br>example `"34.33"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the purchase request has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "amount": "1234.56",
  "created_at": "2023-08-30T10:08:08.146343Z",
  "currency_amount": "1234.56",
  "currency_amount_before_tax": "1200.23",
  "currency_tax": "34.33",
  "delivery_address": {
    "address": "8 rue de la paix",
    "city": "Paris",
    "country_alpha2": "FR",
    "postal_code": "75002"
  },
  "estimated_delivery_date": "2023-12-31",
  "exchange_rate": "1.0",
  "id": 1,
  "linked_invoices": {
    "items": [
      {
        "id": 42,
        "url": "https://app.pennylane.com/api/external/v2/supplier_invoices/42"
      }
    ]
  },
  "purchase_order": {
    "filename": "Bon-de-commande_n°2025-1.pdf",
    "url": "https://www.pennylane.com/Bon-de-commande_n%C2%B02025-1.pdf"
  },
  "purchase_order_number": "PR20230001",
  "reason": "I need a computer to be able to create purchase requests",
  "supplier": {
    "id": 2,
    "url": "https://app.pennylane.com/api/external/v2/suppliers/2"
  },
  "tax": "34.33",
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
