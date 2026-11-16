# Create a webhook subscription

`POST /api/external/v2/webhook_subscriptions`

Creates a webhook subscription to receive real-time notifications for events occurring within your company or companies.

**Authentication & Scope**

Subscription scope:
- **Company-scoped token** (developer token or OAuth token bound to a specific company): the subscription is scoped to that single company.
- **App-bound token** (OAuth token without a company context): the subscription covers all companies accessible by the OAuth application.

**Limits**
Up to 10 webhook subscriptions are allowed per company (for company-scoped tokens) or per subscriber (for app-bound tokens).

**Secret**
The secret will be auto-generated.

> 🔒 The secret is **only returned in the creation response** and cannot be retrieved afterwards. Make sure to store it securely.

[Source Pennylane](https://pennylane.readme.io/reference/postwebhooksubscriptions) · [Index du SDK](../operations.md)

## Appel Go

```go
PostWebhookSubscriptionsWithResponse(ctx context.Context, body PostWebhookSubscriptionsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostWebhookSubscriptionsResponse, error)
```

## Entrée

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.callback_url` | string (uri) | obligatoire | HTTPS URL where webhook events will be sent<br>example `"https://example.com/webhooks"` |
| `$.enabled` | boolean | facultatif | Indicates whether the webhook subscription is active<br>default `true`<br>example `false` |
| `$.events` | array | obligatoire | Array of event types to subscribe to.<br>- `customer_invoice.e_invoicing_status_updated`: the event is triggered when a customer invoice e-invoicing status is updated.<br>- `dms_file.created`: the event is triggered when a dms file is created.<br>- `supplier_invoice.e_invoicing_received`: the event is triggered when a supplier e-invoice is received via the Pennylane PA (Plateforme Agrée).<br><br>minItems `1`<br>example `["dms_file.created","customer_invoice.e_invoicing_status_updated","supplier_invoice.e_invoicing_received"]` |
| `$.events[]` | string | par élément | enum `["customer_invoice.e_invoicing_status_updated","dms_file.created","supplier_invoice.e_invoicing_received"]` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "callback_url": "https://example.com/webhooks",
  "enabled": false,
  "events": [
    "dms_file.created",
    "customer_invoice.e_invoicing_status_updated",
    "supplier_invoice.e_invoicing_received"
  ]
}
```

</details>

## Sortie

### HTTP 201

Renders the created webhook subscription

`application/json` → `réponse.JSON201` (type `*PostWebhookSubscriptionsStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.callback_url` | string (uri) | obligatoire | HTTPS URL where webhook events are sent<br>example `"https://example.com/webhooks"` |
| `$.created_at` | string (date-time) | obligatoire | Creation date of the webhook subscription<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.enabled` | boolean | obligatoire | Indicates whether the webhook subscription is active<br>example `true` |
| `$.events` | array | obligatoire | Array of event types to subscribe to.<br>- `customer_invoice.e_invoicing_status_updated`: the event is triggered when a customer invoice e-invoicing status is updated.<br>- `dms_file.created`: the event is triggered when a dms file is created.<br>- `supplier_invoice.e_invoicing_received`: the event is triggered when a supplier e-invoice is received via the Pennylane PA (Plateforme Agrée).<br><br>minItems `1`<br>example `["dms_file.created","customer_invoice.e_invoicing_status_updated","supplier_invoice.e_invoicing_received"]` |
| `$.events[]` | string | par élément | enum `["customer_invoice.e_invoicing_status_updated","dms_file.created","supplier_invoice.e_invoicing_received"]` |
| `$.id` | integer | obligatoire | ID of the webhook subscription<br>example `12345` |
| `$.secret` | string | obligatoire | Automatically generated secret for HMAC signature verification. This is only returned once<br>on creation. Store it securely as it cannot be retrieved afterwards. |
| `$.updated_at` | string (date-time) | obligatoire | Last update date of the webhook subscription<br>example `"2023-08-07T14:23:12.000Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "callback_url": "https://example.com/webhooks",
  "created_at": "2023-08-07T14:23:12.000Z",
  "enabled": true,
  "events": [
    "dms_file.created",
    "customer_invoice.e_invoicing_status_updated",
    "supplier_invoice.e_invoicing_received"
  ],
  "id": 12345,
  "updated_at": "2023-08-07T14:23:12.000Z"
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

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
