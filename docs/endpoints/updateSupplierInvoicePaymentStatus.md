# Update a supplier invoice payment status

`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/payment_status`

This endpoint allows you to update the payment status of a supplier
invoice.

> ⚠️ **Changing a paid invoice to `to_be_paid` removes its reconciliation.**
>
> When the invoice is fully paid and reconciled, in an open fiscal year and outside any frozen period, setting `payment_status` to `to_be_paid`:
> - unmatches all the bank transactions matched to the invoice;
> - cancels its SEPA XML transfers that were not found on the bank account yet.
>
> If the invoice is part of a grouped payment (one transaction matched to several invoices),
> this breaks the reconciliation of every invoice in the group, not only of this invoice.
> If the invoice was paid in a SEPA XML batch with other invoices,
> the transfers of every invoice in the batch that were not found on the bank account yet are canceled too.
> Check the matched transactions with
> [List matched transactions for a supplier invoice](https://pennylane.readme.io/reference/getsupplierinvoicematchedtransactions)
> before you change the status.


> ℹ️
> This endpoint requires the following scope: `supplier_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/updatesupplierinvoicepaymentstatus) · [Index du SDK](../operations.md)

## Appel Go

```go
UpdateSupplierInvoicePaymentStatusWithResponse(ctx context.Context, supplierInvoiceId int, params *UpdateSupplierInvoicePaymentStatusParams, body UpdateSupplierInvoicePaymentStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSupplierInvoicePaymentStatusResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `supplier_invoice_id` | `supplierInvoiceId` | path | integer | obligatoire | Supplier invoice identifier<br>example `123` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.payment_status` | string | obligatoire | enum `["paid","to_be_paid"]` |

## Sortie

### HTTP 204

The supplier invoice payment status was successfully updated

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
