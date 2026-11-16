# List webhook subscriptions

`GET /api/external/v2/webhook_subscriptions`

Returns all webhook subscriptions for the authenticated token, paginated.
Note that the secret is not included in the response.

[Source Pennylane](https://pennylane.readme.io/reference/getwebhooksubscriptions) · [Index du SDK](../operations.md)

## Appel Go

```go
GetWebhookSubscriptionsWithResponse(ctx context.Context, params *GetWebhookSubscriptionsParams, reqEditors ...RequestEditorFn) (*GetWebhookSubscriptionsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Opaque cursor from a previous response to retrieve the next page |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per page<br>default `20`<br>minimum `1`<br>maximum `100` |

Aucun corps de requête.

## Sortie

### HTTP 200

A paginated list of webhook subscriptions

`application/json` → `réponse.JSON200` (type `*GetWebhookSubscriptionsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].callback_url` | string (uri) | obligatoire | HTTPS URL where webhook events are sent<br>example `"https://example.com/webhooks"` |
| `$.items[].consecutive_failures` | integer | obligatoire | Number of consecutive delivery failures since the last successful delivery<br>example `3` |
| `$.items[].created_at` | string (date-time) | obligatoire | Creation date of the webhook subscription<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.items[].disabled_at` | string (date-time) ou null | obligatoire | Timestamp when the subscription was disabled<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.items[].disabled_reason` | string ou null | obligatoire | Reason the subscription was disabled.<br>- `permanent_error`: disabled due to a permanent delivery error (HTTP 401, 403, 404, or 410).<br>- `repeated_failure`: disabled after repeated delivery failures over a sustained period.<br>- `manual`: disabled manually via the API.<br><br>enum `["permanent_error","repeated_failure","manual"]` |
| `$.items[].enabled` | boolean | obligatoire | Indicates whether the webhook subscription is active<br>example `true` |
| `$.items[].events` | array | obligatoire | Array of event types to subscribe to.<br>- `customer_invoice.e_invoicing_status_updated`: the event is triggered when a customer invoice e-invoicing status is updated.<br>- `dms_file.created`: the event is triggered when a dms file is created.<br>- `supplier_invoice.e_invoicing_received`: the event is triggered when a supplier e-invoice is received via the Pennylane PA (Plateforme Agrée).<br><br>minItems `1`<br>example `["dms_file.created","customer_invoice.e_invoicing_status_updated","supplier_invoice.e_invoicing_received"]` |
| `$.items[].events[]` | string | par élément | enum `["customer_invoice.e_invoicing_status_updated","dms_file.created","supplier_invoice.e_invoicing_received"]` |
| `$.items[].id` | integer | obligatoire | ID of the webhook subscription<br>example `12345` |
| `$.items[].last_failure_at` | string (date-time) ou null | obligatoire | Timestamp of the last failed delivery attempt<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.items[].last_failure_status` | integer ou null | obligatoire | HTTP status code of the last failed delivery attempt<br>example `500` |
| `$.items[].updated_at` | string (date-time) | obligatoire | Last update date of the webhook subscription<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "items": [
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
  ]
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
