# Get a webhook subscription

`GET /api/external/v2/webhook_subscriptions/{id}`

This endpoint allows you to retrieve a webhook subscription by ID.
Note that the secret is not included in the response.

[Source Pennylane](https://pennylane.readme.io/reference/getwebhooksubscription) · [Index du SDK](../operations.md)

## Appel Go

```go
GetWebhookSubscriptionWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetWebhookSubscriptionResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | ID of the webhook subscription |

Aucun corps de requête.

## Sortie

### HTTP 200

A webhook subscription

`application/json` → `réponse.JSON200` (type `*GetWebhookSubscriptionStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.callback_url` | string (uri) | obligatoire | HTTPS URL where webhook events are sent<br>example `"https://example.com/webhooks"` |
| `$.consecutive_failures` | integer | obligatoire | Number of consecutive delivery failures since the last successful delivery<br>example `3` |
| `$.created_at` | string (date-time) | obligatoire | Creation date of the webhook subscription<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.disabled_at` | string (date-time) ou null | obligatoire | Timestamp when the subscription was disabled<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.disabled_reason` | string ou null | obligatoire | Reason the subscription was disabled.<br>- `permanent_error`: disabled due to a permanent delivery error (HTTP 401, 403, 404, or 410).<br>- `repeated_failure`: disabled after repeated delivery failures over a sustained period.<br>- `manual`: disabled manually via the API.<br><br>enum `["permanent_error","repeated_failure","manual"]` |
| `$.enabled` | boolean | obligatoire | Indicates whether the webhook subscription is active<br>example `true` |
| `$.events` | array | obligatoire | Array of event types to subscribe to.<br>- `customer_invoice.e_invoicing_status_updated`: the event is triggered when a customer invoice e-invoicing status is updated.<br>- `dms_file.created`: the event is triggered when a dms file is created.<br>- `supplier_invoice.e_invoicing_received`: the event is triggered when a supplier e-invoice is received via the Pennylane PA (Plateforme Agrée).<br><br>minItems `1`<br>example `["dms_file.created","customer_invoice.e_invoicing_status_updated","supplier_invoice.e_invoicing_received"]` |
| `$.events[]` | string | par élément | enum `["customer_invoice.e_invoicing_status_updated","dms_file.created","supplier_invoice.e_invoicing_received"]` |
| `$.id` | integer | obligatoire | ID of the webhook subscription<br>example `12345` |
| `$.last_failure_at` | string (date-time) ou null | obligatoire | Timestamp of the last failed delivery attempt<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.last_failure_status` | integer ou null | obligatoire | HTTP status code of the last failed delivery attempt<br>example `500` |
| `$.updated_at` | string (date-time) | obligatoire | Last update date of the webhook subscription<br>example `"2023-08-07T14:23:12.000Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "callback_url": "https://example.com/webhooks",
  "consecutive_failures": 3,
  "created_at": "2023-08-07T14:23:12.000Z",
  "disabled_at": "2023-08-07T14:23:12.000Z",
  "enabled": true,
  "events": [
    "dms_file.created",
    "customer_invoice.e_invoicing_status_updated",
    "supplier_invoice.e_invoicing_received"
  ],
  "id": 12345,
  "last_failure_at": "2023-08-07T14:23:12.000Z",
  "last_failure_status": 500,
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

### HTTP 404

The resource was not found

`application/json` → `réponse.JSON404` (type `*GetJournalsStatus404`).

[Champs du payload d’erreur](erreurs.md#getjournalsstatus404)

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.
