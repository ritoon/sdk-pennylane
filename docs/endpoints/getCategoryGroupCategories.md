# List categories of a category group

`GET /api/external/v2/category_groups/{category_group_id}/categories`

List categories of a category group

> ℹ️
> This endpoint requires one of the following scopes: `categories:all`, `categories:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcategorygroupcategories) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCategoryGroupCategoriesWithResponse(ctx context.Context, categoryGroupId int, params *GetCategoryGroupCategoriesParams, reqEditors ...RequestEditorFn) (*GetCategoryGroupCategoriesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `category_group_id` | `categoryGroupId` | path | integer | obligatoire | The unique identifier of the category group<br>example `42` |
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

A list of categories of a category group

`application/json` → `réponse.JSON200` (type `*GetCategoryGroupCategoriesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].analytical_code` | string ou null | obligatoire | example `"CODE123"` |
| `$.items[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].direction` | anyOf | obligatoire |  |
| `$.items[].direction (anyOf 1)` | string | selon variante | Only applicable for _treasury_ categories<br>enum `["cash_in","cash_out"]` |
| `$.items[].direction (anyOf 2)` | string ou null | selon variante | enum `[null]` |
| `$.items[].id` | integer (int64) | obligatoire | example `42` |
| `$.items[].label` | string | obligatoire | example `"Alimentaire"` |
| `$.items[].ledger_account` | object ou null | obligatoire | The ledger account the category is wired to, or `null` when none is set.<br>additionalProperties `false` |
| `$.items[].ledger_account.id` | integer (int64) | obligatoire | example `124` |
| `$.items[].ledger_account.url` | string | obligatoire | The URL to the ledger account<br>example `"https://app.pennylane.com/api/external/v2/ledger_accounts/124"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
    {
      "analytical_code": "CODE123",
      "created_at": "2023-08-30T10:08:08.146343Z",
      "id": 42,
      "label": "Alimentaire",
      "ledger_account": {
        "id": 124,
        "url": "https://app.pennylane.com/api/external/v2/ledger_accounts/124"
      },
      "updated_at": "2023-08-30T10:08:08.146343Z"
    }
  ],
  "next_cursor": "dXBkYXRlZF9hdDoxNjc0MTIzNDU2"
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
