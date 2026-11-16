# Cancel a Gocardless mandate

`POST /api/external/v2/gocardless_mandates/{gocardless_mandate_id}/cancellations`

Cancels a specific Gocardless mandate by ID. The mandate must be in a cancellable state, having one of the following statuses: `pending_submission`, `submitted` or `active`.

> ℹ️
> This endpoint requires the following scope: `customer_mandates:all`

[Source Pennylane](https://pennylane.readme.io/reference/postgocardlessmandatecancellations) · [Index du SDK](../operations.md)

## Appel Go

```go
PostGocardlessMandateCancellationsWithResponse(ctx context.Context, gocardlessMandateId int, params *PostGocardlessMandateCancellationsParams, reqEditors ...RequestEditorFn) (*PostGocardlessMandateCancellationsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `gocardless_mandate_id` | `gocardlessMandateId` | path | integer | obligatoire | ID of the Gocardless mandate to cancel |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 204

Mandate cancelled successfully. No content returned.

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
