# List installments for a customer invoice

`GET /api/external/v2/customer_invoices/{customer_invoice_id}/installments`

List installments for a customer invoice

> ℹ️
> This endpoint requires one of the following scopes: `customer_invoices:all`, `customer_invoices:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcustomerinvoiceinstallments) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCustomerInvoiceInstallmentsWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceInstallmentsParams, reqEditors ...RequestEditorFn) (*GetCustomerInvoiceInstallmentsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_invoice_id` | `customerInvoiceId` | path | integer | obligatoire | <br>example `1234` |
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order. By default uses deadline in ascending order.<br>Example : `deadline` will sort by ascending order, `-deadline` will sort by descending order.<br>Available fields :<br>- `deadline`<br>example `"deadline"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns the list of installments for a customer invoice

`application/json` → `réponse.JSON200` (type `*GetCustomerInvoiceInstallmentsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire | additionalProperties `false` |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].amount` | string ou null | obligatoire | Installment amount (total value of the installment in euros. If the currency is `EUR`, `currency_amount` and `amount` are identical)<br>example `"230.32"` |
| `$.items[].currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.items[].currency_amount` | string ou null | obligatoire | Installment currency amount (total value of the installment in the currency of the invoice)<br>example `"230.32"` |
| `$.items[].deadline` | string (date) ou null | obligatoire | Installment payment deadline (ISO 8601)<br>example `"2023-08-30"` |
| `$.items[].id` | integer (int64) | obligatoire | Installment id<br>example `444` |
| `$.items[].payment_methods` | array ou null | obligatoire | uniqueItems `true` |
| `$.items[].payment_methods[]` | string | par élément | enum `["transfer","e_transfer","credit_card","sepa","cheque","cash","lcr","pro_account_sepa_core","pro_account_pos","pro_account_payment_link","pro_account_sepa_b2b"]` |
| `$.items[].remaining_amount_with_tax` | string ou null | obligatoire | The remaining amount with VAT to pay for the installment to be considered paid<br>example `"20.0"` |
| `$.items[].status` | string | obligatoire | enum `["upcoming","late","paid","waiting_funds","cancelled"]` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "amount": "230.32",
      "currency_amount": "230.32",
      "deadline": "2023-08-30",
      "id": 444,
      "remaining_amount_with_tax": "20.0"
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
