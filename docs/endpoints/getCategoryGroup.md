# Retrieve a category group

`GET /api/external/v2/category_groups/{id}`

This endpoint returns a specific category group.

> ℹ️
> This endpoint requires one of the following scopes: `categories:all`, `categories:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcategorygroup) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCategoryGroupWithResponse(ctx context.Context, id string, params *GetCategoryGroupParams, reqEditors ...RequestEditorFn) (*GetCategoryGroupResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | string | obligatoire | The unique identifier of the category group |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a category group

`application/json` → `réponse.JSON200` (type `*GetCategoryGroupStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.categories` | object | obligatoire | additionalProperties `false` |
| `$.categories.url` | string | obligatoire | URL to get the categories of the group.<br>example `"https://app.pennylane.com/api/external/v2/category_groups/42/categories"` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.id` | integer (int64) | obligatoire | example `989` |
| `$.kind` | string | obligatoire | enum `["treasury","profit_and_loss","structured_analytical","building"]` |
| `$.label` | string | obligatoire | example `"Types de dépenses / revenus"` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "categories": {
    "url": "https://app.pennylane.com/api/external/v2/category_groups/42/categories"
  },
  "created_at": "2023-08-30T10:08:08.146343Z",
  "id": 989,
  "label": "Types de dépenses / revenus",
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

### HTTP 404

The resource was not found

`application/json` → `réponse.JSON404` (type `*GetJournalsStatus404`).

[Champs du payload d’erreur](erreurs.md#getjournalsstatus404)

### HTTP 414

URI Too Long: the request URL exceeds the maximum length. Shorten the query string, for example by sending fewer values per query parameter and splitting them across several requests.

Aucun schéma de corps documenté.
