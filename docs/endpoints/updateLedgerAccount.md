# Update a ledger account

`PUT /api/external/v2/ledger_accounts/{id}`

Update a ledger account

> ℹ️
> This endpoint requires the following scope: `ledger_accounts:all`

[Source Pennylane](https://pennylane.readme.io/reference/updateledgeraccount) · [Index du SDK](../operations.md)

## Appel Go

```go
UpdateLedgerAccountWithResponse(ctx context.Context, id int, params *UpdateLedgerAccountParams, body UpdateLedgerAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLedgerAccountResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | Ledger Account ID<br>example `234` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.enabled` | boolean | facultatif | Whether the ledger account is enabled.<br>The change applies to every ledger account sharing the same number.<br><br>Side effects:<br>- Enabling a personal ledger account of a German capital company is refused with a 422.<br>- Enabling a 401 or 411 ledger account creates the missing supplier or customer.<br><br>example `false` |
| `$.label` | string | facultatif | Label that describes the ledger account<br>example `"Capital"` |
| `$.letterable` | boolean | facultatif | Whether the ledger entries of this ledger account are letterable<br>example `true` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "enabled": false,
  "label": "Capital",
  "letterable": true
}
```

</details>

## Sortie

### HTTP 200

Returns the updated ledger account

`application/json` → `réponse.JSON200` (type `*GetLedgerAccountStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.country_alpha2` | string | obligatoire | Ledger Account's country code (alpha2)<br>example `"FR"` |
| `$.created_at` | string (date-time) | obligatoire | The time the ledger account has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.enabled` | boolean | obligatoire | example `true` |
| `$.id` | integer (int64) | obligatoire | example `124` |
| `$.label` | string | obligatoire | example `"Secondary Account"` |
| `$.letterable` | boolean | obligatoire | example `false` |
| `$.number` | string | obligatoire | example `"512"` |
| `$.type` | string | obligatoire | example `"custom"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the ledger account has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.vat_rate` | string | obligatoire | Ledger Account's VAT rate in percentage<br>example `"any"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "country_alpha2": "FR",
  "created_at": "2023-08-30T10:08:08.146343Z",
  "enabled": true,
  "id": 124,
  "label": "Secondary Account",
  "letterable": false,
  "number": "512",
  "type": "custom",
  "updated_at": "2023-08-30T10:08:08.146343Z",
  "vat_rate": "any"
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

### HTTP 422

Unprocessable content

`application/json` → `réponse.JSON422` (type `*PostJournalsStatus422`).

[Champs du payload d’erreur](erreurs.md#postjournalsstatus422)
