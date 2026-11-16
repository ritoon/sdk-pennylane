# Retrieve a purchase request

`GET /api/external/v2/purchase_requests/{id}`

Retrieve a purchase request

> ℹ️
> This endpoint requires one of the following scopes: `purchase_requests:all`, `purchase_requests:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getpurchaserequest) · [Index du SDK](../operations.md)

## Appel Go

```go
GetPurchaseRequestWithResponse(ctx context.Context, id int, params *GetPurchaseRequestParams, reqEditors ...RequestEditorFn) (*GetPurchaseRequestResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | <br>example `1` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

The requested Purchase Request

`application/json` → `réponse.JSON200` (type `*GetPurchaseRequestStatus200`).

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
