# Retrieve a journal

`GET /api/external/v2/journals/{id}`

Retrieve a journal

> ℹ️
> This endpoint requires one of the following scopes: `journals:readonly`, `journals:all`

[Source Pennylane](https://pennylane.readme.io/reference/getjournal) · [Index du SDK](../operations.md)

## Appel Go

```go
GetJournalWithResponse(ctx context.Context, id int, params *GetJournalParams, reqEditors ...RequestEditorFn) (*GetJournalResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | The unique identifier of the journal<br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a journal

`application/json` → `réponse.JSON200` (type `*GetJournalStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.code` | string | obligatoire |  |
| `$.id` | integer (int64) | obligatoire |  |
| `$.label` | string | obligatoire |  |
| `$.type` | string | obligatoire |  |

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
