# Add a comment to a customer invoice

`POST /api/external/v2/customer_invoices/{customer_invoice_id}/comments`

Add a comment to a customer invoice. The comment appears in the
**Comments** section of the invoice in Pennylane, authored by the user
who granted the access token. Access tokens that are not issued to a
user are rejected with a 403.

The content is stored as plain text. Mentions and rich text are not
supported. A content made only of Unicode whitespace, such as no-break
spaces, is rejected with a 422.

This request is not idempotent: each call adds a new comment, so
retrying after a timeout can create a duplicate.


> ℹ️
> This endpoint requires the following scope: `customer_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/postcustomerinvoicecomment) · [Index du SDK](../operations.md)

## Appel Go

```go
PostCustomerInvoiceCommentWithResponse(ctx context.Context, customerInvoiceId int64, params *PostCustomerInvoiceCommentParams, body PostCustomerInvoiceCommentJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCustomerInvoiceCommentResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_invoice_id` | `customerInvoiceId` | path | integer (int64) | obligatoire | The unique identifier of the customer invoice.<br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.content` | allOf | obligatoire |  |
| `$.content (allOf 1)` | string | selon variante | pattern `"^[^\\x00]*$"` |
| `$.content (allOf 2)` | string | selon variante | Plain text of the comment. Must contain at least one non-whitespace character.<br>maxLength `10000`<br>pattern `"\\S"`<br>example `"Reminder sent to the customer on October 14."` |

## Sortie

### HTTP 201

The created comment

`application/json` → `réponse.JSON201` (type `*PostCustomerInvoiceCommentStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.content` | string | obligatoire | Plain text of the comment.<br>example `"Reminder sent to the customer on October 14."` |
| `$.created_at` | string (date-time) | obligatoire | The time the comment was added<br>example `"2026-10-14T16:00:00.000000Z"` |
| `$.id` | integer (int64) | obligatoire | Comment identifier<br>example `9001` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the comment has been updated, for example when its author edits it in Pennylane.<br>example `"2026-10-14T16:00:00.000000Z"` |
| `$.user` | object ou null | obligatoire | Author of the comment. `null` when the comment was not written by a user.<br>additionalProperties `false` |
| `$.user.first_name` | string | obligatoire | First name of the author<br>example `"Jane"` |
| `$.user.last_name` | string | obligatoire | Last name of the author<br>example `"Doe"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "content": "Reminder sent to the customer on October 14.",
  "created_at": "2026-10-14T16:00:00.000000Z",
  "id": 9001,
  "updated_at": "2026-10-14T16:00:00.000000Z",
  "user": {
    "first_name": "Jane",
    "last_name": "Doe"
  }
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
