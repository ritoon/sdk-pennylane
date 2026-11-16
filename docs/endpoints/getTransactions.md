# List transactions

`GET /api/external/v2/transactions`

List transactions

> ℹ️
> This endpoint requires one of the following scopes: `transactions:readonly`, `transactions:all`

[Source Pennylane](https://pennylane.readme.io/reference/gettransactions) · [Index du SDK](../operations.md)

## Appel Go

```go
GetTransactionsWithResponse(ctx context.Context, params *GetTransactionsParams, reqEditors ...RequestEditorFn) (*GetTransactionsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`: `eq`, `not_eq`, `in`, `not_in`<br>- `bank_account_id`: `eq`, `not_eq`, `in`, `not_in`<br>- `journal_id`: `eq`, `not_eq`, `in`, `not_in`<br>- `date`: `eq`, `not_eq`, `gt`, `lt`, `lteq`, `gteq`<br>example `"[{\"field\": \"journal_id\", \"operator\": \"eq\", \"value\": \"42\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"`<br>example `"id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of transactions

`application/json` → `réponse.JSON200` (type `*GetTransactionsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].amount` | string | obligatoire | Transaction amount in euros. If the currency is euro, `currency_amount` and `amount` are identical.<br>example `"120.00"` |
| `$.items[].archived_at` | string (date-time) ou null | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].attachment_required` | boolean | obligatoire | When `true`, the transaction need to be justified and matched with an invoice<br>example `true` |
| `$.items[].bank_account` | object | obligatoire | additionalProperties `false` |
| `$.items[].bank_account.id` | integer (int64) | obligatoire | example `53` |
| `$.items[].bank_account.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/bank_accounts/53"` |
| `$.items[].categories` | array | obligatoire |  |
| `$.items[].categories[]` | object | par élément | additionalProperties `false` |
| `$.items[].categories[].analytical_code` | string ou null | obligatoire | example `"CODE123"` |
| `$.items[].categories[].category_group` | object | obligatoire | additionalProperties `false` |
| `$.items[].categories[].category_group.id` | integer (int64) | obligatoire | example `229` |
| `$.items[].categories[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].categories[].id` | integer (int64) | obligatoire | example `421` |
| `$.items[].categories[].label` | string | obligatoire | example `"HR - Salaries"` |
| `$.items[].categories[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].categories[].weight` | string | obligatoire | example `"0.25"` |
| `$.items[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].currency` | oneOf | obligatoire |  |
| `$.items[].currency (oneOf 1)` | string | selon variante | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.items[].currency (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.items[].currency_amount` | string | obligatoire | Amount in the currency of the transaction.<br>example `"120.00"` |
| `$.items[].currency_fee` | string ou null | obligatoire | Fee in the currency of the transaction.<br>example `"120.00"` |
| `$.items[].customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].customer.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].customer.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.items[].date` | string (date) | obligatoire | example `"2023-08-30"` |
| `$.items[].fee` | string ou null | obligatoire | Transaction fee in euros. If the currency is euro, `currency_fee` and `fee`<br>are identical.<br><br>example `"120.00"` |
| `$.items[].id` | integer (int64) | obligatoire | Transaction identifier<br>example `42` |
| `$.items[].interbank_code` | string ou null | obligatoire | example `"B1D"` |
| `$.items[].journal` | object | obligatoire | additionalProperties `false` |
| `$.items[].journal.id` | integer (int64) | obligatoire | example `234` |
| `$.items[].journal.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/journals/67"` |
| `$.items[].label` | string ou null | obligatoire | example `"VIR SEPA MY SUPPLIER SAS"` |
| `$.items[].matched_invoices` | object | obligatoire | additionalProperties `false` |
| `$.items[].matched_invoices.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/transactions/42/matched_invoices"` |
| `$.items[].outstanding_balance` | string ou null | obligatoire | This is the balance of the transaction after it has been<br>processed.<br><br>example `"49.3"` |
| `$.items[].pro_account_expense` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].pro_account_expense.card_masked_number` | string | obligatoire | example `"123456XXXXXX7890"` |
| `$.items[].pro_account_expense.employee` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].pro_account_expense.employee.first_name` | string | obligatoire | example `"John"` |
| `$.items[].pro_account_expense.employee.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].pro_account_expense.employee.last_name` | string | obligatoire | example `"Doe"` |
| `$.items[].supplier` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].supplier.id` | integer (int64) | obligatoire | example `42` |
| `$.items[].supplier.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/supplier/42"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "amount": "120.00",
      "archived_at": "2023-08-30T10:08:08.146343Z",
      "attachment_required": true,
      "bank_account": {
        "id": 53,
        "url": "https://app.pennylane.com/api/external/v2/bank_accounts/53"
      },
      "categories": [
        {
          "analytical_code": "CODE123",
          "category_group": {
            "id": 229
          },
          "created_at": "2023-08-30T10:08:08.146343Z",
          "id": 421,
          "label": "HR - Salaries",
          "updated_at": "2023-08-30T10:08:08.146343Z",
          "weight": "0.25"
        }
      ],
      "created_at": "2023-08-30T10:08:08.146343Z",
      "currency_amount": "120.00",
      "currency_fee": "120.00",
      "customer": {
        "id": 42,
        "url": "https://app.pennylane.com/api/external/v2/customers/42"
      },
      "date": "2023-08-30",
      "fee": "120.00",
      "id": 42,
      "interbank_code": "B1D",
      "journal": {
        "id": 234,
        "url": "https://app.pennylane.com/api/external/v2/journals/67"
      },
      "label": "VIR SEPA MY SUPPLIER SAS",
      "matched_invoices": {
        "url": "https://app.pennylane.com/api/external/v2/transactions/42/matched_invoices"
      },
      "outstanding_balance": "49.3",
      "pro_account_expense": {
        "card_masked_number": "123456XXXXXX7890",
        "employee": {
          "first_name": "John",
          "id": 42,
          "last_name": "Doe"
        }
      },
      "supplier": {
        "id": 42,
        "url": "https://app.pennylane.com/api/external/v2/supplier/42"
      },
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

### HTTP 404

The resource was not found

`application/json` → `réponse.JSON404` (type `*GetJournalsStatus404`).

[Champs du payload d’erreur](erreurs.md#getjournalsstatus404)

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.
