# List products

`GET /api/external/v2/products`

List products

> ℹ️
> This endpoint requires one of the following scopes: `products:all`, `products:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getproducts) · [Index du SDK](../operations.md)

## Appel Go

```go
GetProductsWithResponse(ctx context.Context, params *GetProductsParams, reqEditors ...RequestEditorFn) (*GetProductsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id` : `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `label` : `eq`, `in`<br>- `reference` : `eq`, `in`<br>- `external_reference` : `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"id\", \"operator\": \"eq\", \"value\": \"42\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of Products

`application/json` → `réponse.JSON200` (type `*GetProductsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].archived_at` | string (date-time) ou null | obligatoire | The time the product has been archived<br>example `"2023-09-30T10:08:08.146343Z"` |
| `$.items[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.items[].custom_fields` | array | obligatoire |  |
| `$.items[].custom_fields[]` | object | par élément | additionalProperties `false` |
| `$.items[].custom_fields[].name` | string | obligatoire | example `"color"` |
| `$.items[].custom_fields[].value` | anyOf | obligatoire | Select fields use the option label. Multi-select fields use an array of option labels.<br>example `"red"` |
| `$.items[].custom_fields[].value (anyOf 1)` | string ou null | selon variante |  |
| `$.items[].custom_fields[].value (anyOf 2)` | array | selon variante |  |
| `$.items[].custom_fields[].value (anyOf 2)[]` | string | par élément |  |
| `$.items[].description` | string | obligatoire | Product description<br>example `"This is product 1"` |
| `$.items[].ean_code` | string ou null | obligatoire | Product EAN code<br>example `"3017620422003"` |
| `$.items[].external_reference` | string | obligatoire | The unique external reference assigned to this Product, assigned on creation either by you or Pennylane. (Same attribute as `source_id` in the API v1)<br>example `"0e67fc3c-c632-4feb-ad34-e18ed5fbf66a"` |
| `$.items[].id` | integer (int64) | obligatoire | example `1` |
| `$.items[].label` | string | obligatoire | Product label<br>example `"Product 1"` |
| `$.items[].ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].ledger_account.id` | integer (int64) | obligatoire |  |
| `$.items[].price` | string | obligatoire | example `"13.6"` |
| `$.items[].price_before_tax` | string | obligatoire | Product price without taxes<br>example `"12.5"` |
| `$.items[].reference` | string ou null | obligatoire | Product reference<br>example `"REF-123"` |
| `$.items[].substance` | oneOf | obligatoire | Derived from the product's ledger account. Returns `goods` or `services`, or `null` when no ledger account is set. |
| `$.items[].substance (oneOf 1)` | string | selon variante | enum `["goods","services"]` |
| `$.items[].substance (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.items[].unit` | string | obligatoire | Product unit<br>example `"piece"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100"]`<br>example `"FR_200"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "archived_at": "2023-09-30T10:08:08.146343Z",
      "created_at": "2023-08-30T10:08:08.146343Z",
      "custom_fields": [
        {
          "name": "color",
          "value": "red"
        }
      ],
      "description": "This is product 1",
      "ean_code": "3017620422003",
      "external_reference": "0e67fc3c-c632-4feb-ad34-e18ed5fbf66a",
      "id": 1,
      "label": "Product 1",
      "price": "13.6",
      "price_before_tax": "12.5",
      "reference": "REF-123",
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
