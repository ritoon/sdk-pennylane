# List billing email templates

`GET /api/external/v2/billing_email_templates`

List the company's billing email templates that can be used to send a document.

`kind` decides which document accepts the template:
- `billing_email_invoice`: a customer invoice or a proforma
- `billing_email_quote`: a quote, a shipping order, or a purchasing order

Payment reminders, GoCardless subscription emails, and PDF footer templates are not included.
These are not the PDF layout templates listed by `customer_invoice_templates`.

Liquid variables in `subject` and `content` are left unrendered.


> ℹ️
> This endpoint requires one of the following scopes: `customer_invoices:all`, `customer_invoices:readonly`, `quotes:all`, `quotes:readonly`, `commercial_documents:all`, `commercial_documents:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getbillingemailtemplates) · [Index du SDK](../operations.md)

## Appel Go

```go
GetBillingEmailTemplatesWithResponse(ctx context.Context, params *GetBillingEmailTemplatesParams, reqEditors ...RequestEditorFn) (*GetBillingEmailTemplatesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and operators:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `kind`: `eq`, `not_eq`, `in`, `not_in`<br>- `locale`: `eq`, `not_eq`, `in`, `not_in`<br>example `"[{\"field\": \"kind\", \"operator\": \"eq\", \"value\": \"billing_email_invoice\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"`<br>example `"id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of billing email templates

`application/json` → `réponse.JSON200` (type `*GetBillingEmailTemplatesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].bcc_recipients` | array | obligatoire | Blind copy recipients stored on the template. Empty when the template does not use them.<br>example `[]` |
| `$.items[].bcc_recipients[]` | string (email) | par élément |  |
| `$.items[].cc_recipients` | array | obligatoire | Copy recipients stored on the template. Empty when the template does not use them.<br>example `["accounting@company.com"]` |
| `$.items[].cc_recipients[]` | string (email) | par élément |  |
| `$.items[].content` | string | obligatoire | Email body. Liquid variables are left unrendered.<br>example `"Hello, please find your invoice attached."` |
| `$.items[].default` | boolean | obligatoire | Whether this template is selected automatically when sending a document of this kind and locale.<br>example `true` |
| `$.items[].id` | integer (int64) | obligatoire | example `42` |
| `$.items[].kind` | string | obligatoire | Which document a billing email template is used to send.<br>- `billing_email_invoice`: a customer invoice or a proforma<br>- `billing_email_quote`: a quote, a shipping order, or a purchasing order<br><br>enum `["billing_email_invoice","billing_email_quote"]` |
| `$.items[].label` | string ou null | obligatoire | example `"Invoice email"` |
| `$.items[].locale` | string | obligatoire | The local default is based on the thirdparty (customer / supplier ) billing_language . In case the customer is not present, we will default to the company locale.<br>enum `["fr_FR","en_GB","de_DE","es_ES"]` |
| `$.items[].subject` | string ou null | obligatoire | Email subject. Liquid variables are left unrendered.<br>example `"Your invoice {{ numero }}"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "bcc_recipients": [],
      "cc_recipients": [
        "accounting@company.com"
      ],
      "content": "Hello, please find your invoice attached.",
      "default": true,
      "id": 42,
      "label": "Invoice email",
      "subject": "Your invoice {{ numero }}"
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
