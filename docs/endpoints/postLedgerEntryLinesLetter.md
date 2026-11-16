# Letter ledger entry lines

`POST /api/external/v2/ledger_entry_lines/lettering`

This endpoint lets you letter ledger entry lines together. All
received entry lines will be lettered together. If a passed entry line is
already lettered, then the lettering will be applied to its associated lettered
entry lines as well.

> ℹ️
> This endpoint requires the following scope: `ledger_entries:all`

[Source Pennylane](https://pennylane.readme.io/reference/postledgerentrylinesletter) · [Index du SDK](../operations.md)

## Appel Go

```go
PostLedgerEntryLinesLetterWithResponse(ctx context.Context, params *PostLedgerEntryLinesLetterParams, body PostLedgerEntryLinesLetterJSONRequestBody, reqEditors ...RequestEditorFn) (*PostLedgerEntryLinesLetterResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.ledger_entry_lines` | array | obligatoire | The list of ledger entry lines you want to letter together.<br><br>You can provide ledger entry lines already lettered and they don't have to be part of the same lettering.<br><br>All received entry lines will be lettered together. If a passed entry line is already lettered, then the lettering will be applied to its lettered entry lines as well. For example:<br>Ledger entry lines `A` and `B` are already lettered together. `C` in not lettered.<br>When requesting `[A, C]` to be lettered, the final lettering will be `[A, B, C]`.<br><br>minItems `2` |
| `$.ledger_entry_lines[]` | object | par élément | additionalProperties `false` |
| `$.ledger_entry_lines[].id` | integer (int64) | obligatoire | Id of the ledger entry line<br>example `3455` |
| `$.unbalanced_lettering_strategy` | string | obligatoire | - `none`: the API won't let you create an unbalanced lettering<br>and will respond with an error.<br>- `partial`: a potentially unbalanced lettering will be created.<br><br>enum `["none","partial"]` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "ledger_entry_lines": [
    {
      "id": 3455
    }
  ]
}
```

</details>

## Sortie

### HTTP 200

Returns all the ledger entry lines of the new lettering

`application/json` → `réponse.JSON200` (type `*`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | array | racine | An array containing all the ledger entry lines of the new lettering |
| `$[]` | object | par élément | additionalProperties `false` |
| `$[].id` | integer (int64) | obligatoire | Id of the ledger entry line<br>example `1` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
[
  {
    "id": 1
  }
]
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

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
