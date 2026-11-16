# Create a category group

`POST /api/external/v2/category_groups`

Create a category group

> ℹ️
> This endpoint requires the following scope: `categories:all`

[Source Pennylane](https://pennylane.readme.io/reference/postcategorygroups) · [Index du SDK](../operations.md)

## Appel Go

```go
PostCategoryGroupsWithResponse(ctx context.Context, params *PostCategoryGroupsParams, body PostCategoryGroupsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCategoryGroupsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.kind` | string | facultatif | enum `["profit_and_loss","treasury","building"]`<br>example `"profit_and_loss"` |
| `$.label` | string | obligatoire | maxLength `2000`<br>example `"Types de depenses / revenus"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "kind": "profit_and_loss",
  "label": "Types de depenses / revenus"
}
```

</details>

## Sortie

### HTTP 201

Returns the created category group

`application/json` → `réponse.JSON201` (type `*PostCategoryGroupsStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.categories` | object | obligatoire | additionalProperties `false` |
| `$.categories.url` | string | obligatoire | URL to get the categories of the group.<br>example `"https://app.pennylane.com/api/external/v2/category_groups/42/categories"` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.id` | integer (int64) | obligatoire | example `989` |
| `$.label` | string | obligatoire | example `"Types de depenses / revenus"` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "categories": {
    "url": "https://app.pennylane.com/api/external/v2/category_groups/42/categories"
  },
  "created_at": "2023-08-30T10:08:08.146343Z",
  "id": 989,
  "label": "Types de depenses / revenus",
  "updated_at": "2023-08-30T10:08:08.146343Z"
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
