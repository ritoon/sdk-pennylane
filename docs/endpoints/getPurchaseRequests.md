# List purchase requests

`GET /api/external/v2/purchase_requests`

List purchase requests

> ℹ️
> This endpoint requires one of the following scopes: `purchase_requests:all`, `purchase_requests:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getpurchaserequests) · [Index du SDK](../operations.md)

## Appel Go

```go
GetPurchaseRequestsWithResponse(ctx context.Context, params *GetPurchaseRequestsParams, reqEditors ...RequestEditorFn) (*GetPurchaseRequestsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `reviewed_by_id`: `eq`, `not_eq`, `in`, `not_in`<br>- `user_id`: `eq`, `not_eq`, `in`, `not_in`<br>- `supplier_id`: `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"user_id\", \"operator\": \"eq\", \"value\": \"123\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a list of purchase requests

`application/json` → `réponse.JSON200` (type `*GetPurchaseRequestsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire | The list of items returned |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].amount` | non spécifié (string) | obligatoire | example `"1234.56"` |
| `$.items[].created_at` | string (date-time) | obligatoire | The time the purchase request has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.items[].currency_amount` | string | obligatoire | Purchase Request currency amount (total value of the purchase request in the currency of the purchase request)<br>example `"1234.56"` |
| `$.items[].currency_amount_before_tax` | string | obligatoire | Purchase Request currency amount before tax (total value before tax of the purchase request in the currency of the purchase request)<br>example `"1200.23"` |
| `$.items[].currency_tax` | string | obligatoire | Purchase request taxable amount (in euros. If the currency is euro, `currency_amount` and `amount` are identical)<br>example `"34.33"` |
| `$.items[].delivery_address` | object | obligatoire | additionalProperties `false` |
| `$.items[].delivery_address.address` | string ou null | obligatoire | example `"8 rue de la paix"` |
| `$.items[].delivery_address.city` | string ou null | obligatoire | example `"Paris"` |
| `$.items[].delivery_address.country_alpha2` | string ou null | obligatoire | example `"FR"` |
| `$.items[].delivery_address.postal_code` | string ou null | obligatoire | example `"75002"` |
| `$.items[].estimated_delivery_date` | string (date) ou null | obligatoire | Estimated delivery date for the purchase request<br>example `"2023-12-31"` |
| `$.items[].exchange_rate` | string | obligatoire | Purchase request exchange rate (used to convert the purchase request to euros. If the purchase request currency is euro it will be 1.0)<br>example `"1.0"` |
| `$.items[].id` | integer (int64) | obligatoire | Purchase request identifier<br>example `1` |
| `$.items[].linked_invoices` | object | obligatoire | additionalProperties `false` |
| `$.items[].linked_invoices.items` | array | obligatoire | List of linked supplier invoices |
| `$.items[].linked_invoices.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].linked_invoices.items[].id` | integer (int64) | obligatoire | example `42` |
| `$.items[].linked_invoices.items[].url` | string | obligatoire | URL to get the supplier invoice.<br>example `"https://app.pennylane.com/api/external/v2/supplier_invoices/42"` |
| `$.items[].purchase_order` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].purchase_order.filename` | string | obligatoire | example `"Bon-de-commande_n°2025-1.pdf"` |
| `$.items[].purchase_order.url` | string | obligatoire | example `"https://www.pennylane.com/Bon-de-commande_n%C2%B02025-1.pdf"` |
| `$.items[].purchase_order_number` | string ou null | obligatoire | Purchase order number<br>example `"PR20230001"` |
| `$.items[].reason` | string | obligatoire | Reason of the request<br>example `"I need a computer to be able to create purchase requests"` |
| `$.items[].status` | string | obligatoire | enum `["to_be_validated","approved","rejected","invoiced"]` |
| `$.items[].supplier` | object | obligatoire | additionalProperties `false` |
| `$.items[].supplier.id` | integer (int64) | obligatoire | example `2` |
| `$.items[].supplier.url` | string (uri) | obligatoire | URL to get the supplier.<br>example `"https://app.pennylane.com/api/external/v2/suppliers/2"` |
| `$.items[].tax` | string | obligatoire | Purchase request taxable amount (in purchase request currency)<br>example `"34.33"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | The last time the purchase request has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
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

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.
