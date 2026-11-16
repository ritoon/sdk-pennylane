# Send a quote by email

`POST /api/external/v2/quotes/{id}/send_by_email`

This endpoint allows you to send a quote by email to your customer.
This requires that the PDF file for that document has been generated
(this process can take a few minutes), so if you just created
the quote in our system, we may return a 409 error. You should
retry the request in a few minutes - if you receive a 204 response, that means
that the email is on its way. For more information about email sending, please
read \[this guide\](https://pennylane.readme.io/v2.0/docs/sending-documents-by-email).

The subject and body come from the company's default billing email template
for the customer's billing language, unless `email_template_id` is set.


> ℹ️
> This endpoint requires the following scope: `quotes:all`

[Source Pennylane](https://pennylane.readme.io/reference/sendbyemailquote) · [Index du SDK](../operations.md)

## Appel Go

```go
SendByEmailQuoteWithResponse(ctx context.Context, id int, params *SendByEmailQuoteParams, body SendByEmailQuoteJSONRequestBody, reqEditors ...RequestEditorFn) (*SendByEmailQuoteResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | <br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps facultatif** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.email_template_id` | integer (int64) | facultatif | Billing email template to use for this send.<br><br>It must belong to the company and have kind `billing_email_quote`.<br>Its locale does not have to match the customer's billing language.<br>Subject, body, cc and bcc then come from that template.<br><br>When omitted, the company's default template for the customer's billing language is used.<br>In that case only the subject and body are applied.<br><br>example `42` |
| `$.recipients` | array | facultatif | Email recipients.<br><br>If empty, the email will be sent to the recipient email addresses specified for the customer of this quote.<br><br>uniqueItems `true`<br>example `["audit@company.com"]` |
| `$.recipients[]` | string (email) | par élément |  |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "email_template_id": 42,
  "recipients": [
    "audit@company.com"
  ]
}
```

</details>

## Sortie

### HTTP 204

Quote is being sent by email

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
