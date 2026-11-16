# Send a customer e-invoice to PA

`POST /api/external/v2/customer_invoices/{id}/send_to_pa`

Send a customer e-invoice to the Partner Dematerialization Platform (PA).

After an e-invoice import, the file may still be processing. If you call this
endpoint before processing is finished, the API returns 409. Retry after a few
seconds. A 204 response means the invoice is being sent to the PA.


> ℹ️
> This endpoint requires the following scope: `customer_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/sendtopacustomerinvoice) · [Index du SDK](../operations.md)

## Appel Go

```go
SendToPaCustomerInvoiceWithResponse(ctx context.Context, id int64, params *SendToPaCustomerInvoiceParams, reqEditors ...RequestEditorFn) (*SendToPaCustomerInvoiceResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer (int64) | obligatoire | The unique identifier of the customer e-invoice to send.<br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 204

Customer e-invoice is being sent to the Partner Dematerialization Platform (PA)

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

### HTTP 409

Conflict with the current state of the target resource, such as when trying to create a resource that already exists.

`application/json` → `réponse.JSON409` (type `*SendByEmailCommercialDocumentStatus409`).

[Champs du payload d’erreur](erreurs.md#sendbyemailcommercialdocumentstatus409)

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
