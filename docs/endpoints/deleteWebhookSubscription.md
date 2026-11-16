# Delete a webhook subscription

`DELETE /api/external/v2/webhook_subscriptions/{id}`

This endpoint allows you to delete a webhook subscription by ID.

[Source Pennylane](https://pennylane.readme.io/reference/deletewebhooksubscription) · [Index du SDK](../operations.md)

## Appel Go

```go
DeleteWebhookSubscriptionWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteWebhookSubscriptionResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | ID of the webhook subscription |

Aucun corps de requête.

## Sortie

### HTTP 204

Webhook subscription deleted

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
