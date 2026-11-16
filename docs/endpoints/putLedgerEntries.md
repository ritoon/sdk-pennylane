# Update a ledger entry

`PUT /api/external/v2/ledger_entries/{id}`

Update a ledger entry


> ℹ️
> This endpoint requires the following scope: `ledger_entries:all`

[Source Pennylane](https://pennylane.readme.io/reference/putledgerentries) · [Index du SDK](../operations.md)

## Appel Go

```go
PutLedgerEntriesWithResponse(ctx context.Context, id string, params *PutLedgerEntriesParams, body PutLedgerEntriesJSONRequestBody, reqEditors ...RequestEditorFn) (*PutLedgerEntriesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | string | obligatoire | The unique identifier of the ledger entry |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.currency` | string | facultatif | Currency code of the ledger entry as per ISO 4217. This currency is applicable to all ledger_entry_lines<br>enum `["EUR","USD","GBP","AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYN","BYR","BZD","CAD","CDF","CHE","CHF","CLF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","FJD","FKP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MRU","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SLE","SOS","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TWD","TZS","UAH","UGX","UYU","UZS","VEF","VES","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMK","ZMW","ZWL"]`<br>example `"EUR"` |
| `$.date` | string (date) | facultatif | Date of the ledger entry (ISO 8601)<br>example `"2023-08-30"` |
| `$.file_attachment_id` | integer (int64) | facultatif | File attachment ID.<br><br>> ⚠️ Passing this field will soon be deprecated. We will provide an alternative using `/api/external/v2/ledger_entries/{id}/dms_files`.<br><br>example `42` |
| `$.journal_id` | integer (int64) | facultatif | The journal ID where you want to create the ledger entry<br>example `123` |
| `$.label` | string | facultatif | Label that describes the ledger entry<br>example `"Payment for Services"` |
| `$.ledger_entry_lines` | object | facultatif | Add, update, delete ledger entry lines. The entry lines must be balanced.<br>**In total**, max number of entry lines that you can create, update or delete using this endpoint is 1000 per request.<br><br>> ⚠️ The order of `ledger_entry_lines` in the response is not guaranteed to match the order of lines sent in `ledger_entry_lines.create`. To map source lines to their returned `ledger_entry_line.id`, match by content (e.g. `debit`, `credit`, `label`) rather than by position.<br><br>additionalProperties `false` |
| `$.ledger_entry_lines.create` | array | facultatif |  |
| `$.ledger_entry_lines.create[]` | object | par élément | additionalProperties `false` |
| `$.ledger_entry_lines.create[].credit` | string | obligatoire | Credit amount for the entry line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.ledger_entry_lines.create[].debit` | string | obligatoire | Debit amount for the entry line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.ledger_entry_lines.create[].label` | string | facultatif | Label that describes the entry line<br>example `"Transaction label"` |
| `$.ledger_entry_lines.create[].ledger_account_id` | integer (int64) | obligatoire | Ledger account ID<br>example `987` |
| `$.ledger_entry_lines.delete` | array | facultatif |  |
| `$.ledger_entry_lines.delete[]` | object | par élément | additionalProperties `false` |
| `$.ledger_entry_lines.delete[].id` | integer (int64) | obligatoire | ID of the entry line<br>example `42` |
| `$.ledger_entry_lines.update` | array | facultatif |  |
| `$.ledger_entry_lines.update[]` | object | par élément | additionalProperties `false` |
| `$.ledger_entry_lines.update[].credit` | string | facultatif | Credit amount for the entry line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.ledger_entry_lines.update[].debit` | string | facultatif | Debit amount for the entry line<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.ledger_entry_lines.update[].id` | integer (int64) | obligatoire | ID of the entry line<br>example `42` |
| `$.ledger_entry_lines.update[].label` | string | facultatif | Label that describes the entry line<br>example `"Transaction label"` |
| `$.ledger_entry_lines.update[].ledger_account_id` | integer (int64) | facultatif | Ledger account ID<br>example `421` |
| `$.piece_number` | string | facultatif | The piece number that was assigned during creation either by you or Pennylane.<br>example `"TENSNYLUGV"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "currency": "EUR",
  "date": "2023-08-30",
  "file_attachment_id": 42,
  "journal_id": 123,
  "label": "Payment for Services",
  "ledger_entry_lines": {
    "create": [
      {
        "credit": "100.00",
        "debit": "100.00",
        "label": "Transaction label",
        "ledger_account_id": 987
      }
    ],
    "delete": [
      {
        "id": 42
      }
    ],
    "update": [
      {
        "credit": "100.00",
        "debit": "100.00",
        "id": 42,
        "label": "Transaction label",
        "ledger_account_id": 421
      }
    ]
  },
  "piece_number": "TENSNYLUGV"
}
```

</details>

## Sortie

### HTTP 200

Returns the updated ledger entry

`application/json` → `réponse.JSON200` (type `*PostLedgerEntriesStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.attachment` | object ou null | obligatoire | The primary accounting file of this ledger entry.<br>For invoices and most document types this is auto-generated by Pennylane.<br>For miscellaneous entries (OD) it can be set via `file_attachment_id` on create/update, though that behavior will soon be deprecated.<br>Files attached through the Pennylane UI are exposed separately via `GET /api/external/v2/ledger_entries/{id}/dms_files`.<br><br>additionalProperties `false` |
| `$.attachment.filename` | string | obligatoire | example `"filename.pdf"` |
| `$.attachment.url` | string | obligatoire | URL to the uploaded file.<br>example `"http://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/filename.pdf"` |
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
| `$.date` | string (date) | obligatoire | Date of the ledger entry (ISO 8601)<br>example `"2023-08-30"` |
| `$.due_date` | string (date) ou null | obligatoire | The date the invoice is due<br>example `"2023-08-30"` |
| `$.id` | integer (int64) | obligatoire | ID of the ledger entry<br>example `1` |
| `$.invoice_number` | string ou null | obligatoire | Invoice number of the document linked to the ledger entry<br>example `"FAC-2024-001"` |
| `$.journal` | object | obligatoire | additionalProperties `false` |
| `$.journal.id` | integer (int64) | obligatoire | example `123` |
| `$.journal.url` | string | obligatoire | example `"https://app.pennylane.com/api/external/v2/journals/123"` |
| `$.journal_id` | integer (int64) | obligatoire | This attribute is **deprecated**. Please use `journal` instead. The journal ID where the ledger entry was created<br>deprecated `true`<br>example `123` |
| `$.label` | string | obligatoire | Label that describes the ledger entry<br>example `"Payment for Services"` |
| `$.ledger_attachment_filename` | string ou null | obligatoire | This attribute is deprecated. Please use `filename` from `attachment` instead.<br>deprecated `true`<br>example `"filename.pdf"` |
| `$.ledger_attachment_id` | integer (int64) ou null | obligatoire | This attribute is deprecated. Ledger attachment ID<br>deprecated `true`<br>example `42` |
| `$.ledger_entry_lines` | array | obligatoire | Array of entry lines |
| `$.ledger_entry_lines[]` | object | par élément | additionalProperties `false` |
| `$.ledger_entry_lines[].credit` | string | obligatoire | Credit amount for the entry line<br>example `"0.00"` |
| `$.ledger_entry_lines[].debit` | string | obligatoire | Debit amount for the entry line<br>example `"100.00"` |
| `$.ledger_entry_lines[].id` | integer (int64) | obligatoire | ID of the entry line<br>example `42` |
| `$.ledger_entry_lines[].label` | string | obligatoire | Label that describes the entry line<br>example `"Transaction label"` |
| `$.ledger_entry_lines[].ledger_account` | object | obligatoire | additionalProperties `false` |
| `$.ledger_entry_lines[].ledger_account.id` | integer (int64) | obligatoire | Ledger account ID<br>example `987` |
| `$.ledger_entry_lines[].ledger_account.number` | string | obligatoire | Ledger account number<br>example `"411000"` |
| `$.ledger_entry_lines[].ledger_account.url` | string | obligatoire | The URL to the ledger account<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/42"` |
| `$.ledger_entry_lines[].ledger_account_id` | integer (int64) | obligatoire | This attribute is deprecated. Please use `id` from `ledger_account` instead.<br>deprecated `true`<br>example `987` |
| `$.piece_number` | string ou null | obligatoire | example `"TENSNYLUGV"` |
| `$.status` | string ou null | obligatoire | The accounting state of the entry.<br>- `archived`: The entry has been archived.<br>- `draft`: The customer invoice is still in draft and has not been finalized.<br>- `entry`: The invoice is incomplete, some information is missing and needs to be completed.<br>- `waiting_details`: The entry is awaiting additional accounting details.<br>- `validation_needed`: The entry is sent to the accountant and needs validation.<br>- `accounting_needed`: The entry needs to be processed by the accountant.<br>- `complete`: The entry has been validated by the accountant.<br>- `null` : The entry is neither a transaction nor an invoice that needs to be processed on Pennylane<br><br>enum `["accounting_needed","archived","complete","draft","entry","validation_needed","waiting_details"]` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "attachment": {
    "filename": "filename.pdf",
    "url": "http://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/filename.pdf"
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
  "date": "2023-08-30",
  "due_date": "2023-08-30",
  "id": 1,
  "invoice_number": "FAC-2024-001",
  "journal": {
    "id": 123,
    "url": "https://app.pennylane.com/api/external/v2/journals/123"
  },
  "journal_id": 123,
  "label": "Payment for Services",
  "ledger_attachment_filename": "filename.pdf",
  "ledger_attachment_id": 42,
  "ledger_entry_lines": [
    {
      "credit": "0.00",
      "debit": "100.00",
      "id": 42,
      "label": "Transaction label",
      "ledger_account": {
        "id": 987,
        "number": "411000",
        "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/42"
      },
      "ledger_account_id": 987
    }
  ],
  "piece_number": "TENSNYLUGV"
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
