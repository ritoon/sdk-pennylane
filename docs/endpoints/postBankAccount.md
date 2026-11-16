# Create a bank account

`POST /api/external/v2/bank_accounts`

Create a bank account

> ℹ️
> This endpoint requires the following scope: `bank_accounts:all`

[Source Pennylane](https://pennylane.readme.io/reference/postbankaccount) · [Index du SDK](../operations.md)

## Appel Go

```go
PostBankAccountWithResponse(ctx context.Context, params *PostBankAccountParams, body PostBankAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*PostBankAccountResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.account_type` | string | facultatif | - `current`: Deprecated. This value will be removed in a future version. Please use `checking` instead.<br><br>enum `["current","card","savings","shares","loan","life_insurance","other","checking"]` |
| `$.bank_establishment_id` | integer (int64) | facultatif | ID of the bank establishment. If not provided, `Other` establishment will be used.<br>example `25` |
| `$.bic` | string | facultatif | Bank Identifier Code (BIC)<br>example `"BNPAFRPPXXX"` |
| `$.currency` | string | facultatif | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.iban` | string | facultatif | International Bank Account Number (IBAN)<br>example `"FR1420041010050500013M02606"` |
| `$.name` | string | obligatoire | The name of the bank account<br>example `"Main Account"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "bank_establishment_id": 25,
  "bic": "BNPAFRPPXXX",
  "iban": "FR1420041010050500013M02606",
  "name": "Main Account"
}
```

</details>

## Sortie

### HTTP 201

Returns the created bank account

`application/json` → `réponse.JSON201` (type `*GetBankAccountStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.balance` | string | obligatoire | example `"100.15"` |
| `$.bank_establishment` | object | obligatoire | additionalProperties `false` |
| `$.bank_establishment.id` | integer (int64) | obligatoire | example `42` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.currency` | string | obligatoire | enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>default `"EUR"` |
| `$.id` | integer (int64) | obligatoire | example `42` |
| `$.journal` | object ou null | obligatoire | additionalProperties `false` |
| `$.journal.id` | integer (int64) | obligatoire | example `42` |
| `$.journal.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/journals/7"` |
| `$.ledger_account` | object | obligatoire | additionalProperties `false` |
| `$.ledger_account.id` | integer (int64) | obligatoire | example `42` |
| `$.ledger_account.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/ledger_accounts/8"` |
| `$.name` | string | obligatoire | example `"Main account"` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "balance": "100.15",
  "bank_establishment": {
    "id": 42
  },
  "created_at": "2023-08-30T10:08:08.146343Z",
  "id": 42,
  "journal": {
    "id": 42,
    "url": "https://app.pennylane.com/api/external/v2/journals/7"
  },
  "ledger_account": {
    "id": 42,
    "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/8"
  },
  "name": "Main account",
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
