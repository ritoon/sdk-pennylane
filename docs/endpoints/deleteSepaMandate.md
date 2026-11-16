# Delete a SEPA mandate

`DELETE /api/external/v2/sepa_mandates/{id}`

This endpoint allows you to delete a specific SEPA mandate

> ℹ️
> This endpoint requires the following scope: `customer_mandates:all`

[Source Pennylane](https://pennylane.readme.io/reference/deletesepamandate) · [Index du SDK](../operations.md)

## Appel Go

```go
DeleteSepaMandateWithResponse(ctx context.Context, id int, params *DeleteSepaMandateParams, reqEditors ...RequestEditorFn) (*DeleteSepaMandateResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | ID of the SEPA mandate to delete<br>example `1234` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 204

SEPA mandate deleted

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
