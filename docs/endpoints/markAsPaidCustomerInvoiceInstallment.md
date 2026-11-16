# Mark a customer invoice installment as paid

`PUT /api/external/v2/customer_invoices/{customer_invoice_id}/installments/{id}/mark_as_paid`

Mark a customer invoice installment as paid. If all installments of a customer invoice are paid, the customer invoice will be marked as paid. No automatic reconciliation will
be done once the installment or customer invoice is marked as paid.

> ℹ️
> This endpoint requires a company plan that has access to Installments, otherwise responds with 403.


> ℹ️
> This endpoint requires the following scope: `customer_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/markaspaidcustomerinvoiceinstallment) · [Index du SDK](../operations.md)

## Appel Go

```go
MarkAsPaidCustomerInvoiceInstallmentWithResponse(ctx context.Context, customerInvoiceId int64, id int64, params *MarkAsPaidCustomerInvoiceInstallmentParams, reqEditors ...RequestEditorFn) (*MarkAsPaidCustomerInvoiceInstallmentResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_invoice_id` | `customerInvoiceId` | path | integer (int64) | obligatoire | The unique identifier of the customer invoice which the installment belongs to.<br>example `1234` |
| `id` | `id` | path | integer (int64) | obligatoire | The unique identifier of the installment to mark as paid.<br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 204

Customer invoice installment marked as paid

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
