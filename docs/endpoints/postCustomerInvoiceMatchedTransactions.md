# Match a transaction to a customer invoice

`POST /api/external/v2/customer_invoices/{customer_invoice_id}/matched_transactions`

This endpoint allows you to match a transaction to a customer invoice. It is not applicable for draft invoices.

You can match one transaction with one customer invoice at a time. To match multiple transactions to a customer invoice, you need to call this endpoint multiple times.
It's possible to match a transaction to multiple customer invoices too.


> ℹ️
> This endpoint requires the following scope: `customer_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/postcustomerinvoicematchedtransactions) · [Index du SDK](../operations.md)

## Appel Go

```go
PostCustomerInvoiceMatchedTransactionsWithResponse(ctx context.Context, customerInvoiceId int, params *PostCustomerInvoiceMatchedTransactionsParams, body PostCustomerInvoiceMatchedTransactionsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCustomerInvoiceMatchedTransactionsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_invoice_id` | `customerInvoiceId` | path | integer | obligatoire |  |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.transaction_id` | integer (int64) | obligatoire | ID of the transaction<br>example `42` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "transaction_id": 42
}
```

</details>

## Sortie

### HTTP 204

Transaction matched successfully

Aucun schéma de corps documenté.

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
