# Create a transaction

`POST /api/external/v2/transactions`

Create a banking transaction

> ℹ️
> This endpoint requires the following scope: `transactions:all`

[Source Pennylane](https://pennylane.readme.io/reference/createtransaction) · [Index du SDK](../operations.md)

## Appel Go

```go
CreateTransactionWithResponse(ctx context.Context, params *CreateTransactionParams, body CreateTransactionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateTransactionResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.amount` | string | obligatoire | Transaction amount<br>example `"120.00"` |
| `$.bank_account_id` | integer (int64) | obligatoire | The bank account where the transaction is registered<br>example `42` |
| `$.date` | string (date) | obligatoire | Transaction date<br>example `"2025-07-30"` |
| `$.fee` | string | facultatif | Transaction fee<br>example `"12.00"` |
| `$.label` | string | obligatoire | Transaction label<br>example `"test label"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "amount": "120.00",
  "bank_account_id": 42,
  "date": "2025-07-30",
  "fee": "12.00",
  "label": "test label"
}
```

</details>

## Sortie

### HTTP 201

Transaction created

`application/json` → `réponse.JSON201` (type `*GetTransactionStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.amount` | string | obligatoire | Transaction amount in euros. If the currency is euro, `currency_amount` and `amount` are identical.<br>example `"120.00"` |
| `$.archived_at` | string (date-time) ou null | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.attachment_required` | boolean | obligatoire | When `true`, the transaction need to be justified and matched with an invoice<br>example `true` |
| `$.bank_account` | object | obligatoire | additionalProperties `false` |
| `$.bank_account.id` | integer (int64) | obligatoire | example `53` |
| `$.bank_account.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/bank_accounts/53"` |
| `$.categories` | array | obligatoire |  |
| `$.categories[]` | object | par élément | additionalProperties `false` |
| `$.categories[].analytical_code` | string ou null | obligatoire | example `"CODE123"` |
| `$.categories[].category_group` | object | obligatoire | additionalProperties `false` |
| `$.categories[].category_group.id` | integer (int64) | obligatoire | example `229` |
| `$.categories[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.categories[].id` | integer (int64) | obligatoire | example `421` |
| `$.categories[].label` | string | obligatoire | example `"HR - Salaries"` |
| `$.categories[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.categories[].weight` | string | obligatoire | example `"0.25"` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.currency` | oneOf | obligatoire |  |
| `$.currency (oneOf 1)` | string | selon variante | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.currency (oneOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.currency_amount` | string | obligatoire | Amount in the currency of the transaction.<br>example `"120.00"` |
| `$.currency_fee` | string ou null | obligatoire | Fee in the currency of the transaction.<br>example `"120.00"` |
| `$.customer` | object ou null | obligatoire | additionalProperties `false` |
| `$.customer.id` | integer (int64) | obligatoire | example `42` |
| `$.customer.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.date` | string (date) | obligatoire | example `"2023-08-30"` |
| `$.fee` | string ou null | obligatoire | Transaction fee in euros. If the currency is euro, `currency_fee` and `fee`<br>are identical.<br><br>example `"120.00"` |
| `$.id` | integer (int64) | obligatoire | Transaction identifier<br>example `42` |
| `$.interbank_code` | string ou null | obligatoire | example `"B1D"` |
| `$.journal` | object | obligatoire | additionalProperties `false` |
| `$.journal.id` | integer (int64) | obligatoire | example `234` |
| `$.journal.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/journals/67"` |
| `$.label` | string ou null | obligatoire | example `"VIR SEPA MY SUPPLIER SAS"` |
| `$.matched_invoices` | object | obligatoire | additionalProperties `false` |
| `$.matched_invoices.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/transactions/42/matched_invoices"` |
| `$.outstanding_balance` | string ou null | obligatoire | This is the balance of the transaction after it has been<br>processed.<br><br>example `"49.3"` |
| `$.pro_account_expense` | object ou null | obligatoire | additionalProperties `false` |
| `$.pro_account_expense.card_masked_number` | string | obligatoire | example `"123456XXXXXX7890"` |
| `$.pro_account_expense.employee` | object ou null | obligatoire | additionalProperties `false` |
| `$.pro_account_expense.employee.first_name` | string | obligatoire | example `"John"` |
| `$.pro_account_expense.employee.id` | integer (int64) | obligatoire | example `42` |
| `$.pro_account_expense.employee.last_name` | string | obligatoire | example `"Doe"` |
| `$.supplier` | object ou null | obligatoire | additionalProperties `false` |
| `$.supplier.id` | integer (int64) | obligatoire | example `42` |
| `$.supplier.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/supplier/42"` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
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
