# Create a journal

`POST /api/external/v2/journals`

Create a journal

> ℹ️
> This endpoint requires the following scope: `journals:all`

[Source Pennylane](https://pennylane.readme.io/reference/postjournals) · [Index du SDK](../operations.md)

## Appel Go

```go
PostJournalsWithResponse(ctx context.Context, params *PostJournalsParams, body PostJournalsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostJournalsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.code` | string | obligatoire | 2 to 5 letters that represents the journal<br>example `"RB"` |
| `$.label` | string | obligatoire | Label that describes the journal<br>example `"Journal de Reprise Balance"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "code": "RB",
  "label": "Journal de Reprise Balance"
}
```

</details>

## Sortie

### HTTP 201

Returns the created journal

`application/json` → `réponse.JSON201` (type `*GetJournalStatus200`).

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

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
