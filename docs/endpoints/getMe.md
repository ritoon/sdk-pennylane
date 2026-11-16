# User Profile

`GET /api/external/v2/me`

This endpoint returns information about the company and the user associated to the token.

[Source Pennylane](https://pennylane.readme.io/reference/getme) · [Index du SDK](../operations.md)

## Appel Go

```go
GetMeWithResponse(ctx context.Context, params *GetMeParams, reqEditors ...RequestEditorFn) (*GetMeResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns the user and the company

`application/json` → `réponse.JSON200` (type `*GetMeStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.company` | object | obligatoire | additionalProperties `false` |
| `$.company.accounting_logic` | string | obligatoire | Chart of accounts framework used by the company.<br>enum `["FR_PCG","FR_AGRICULTURE","FR_NON_PROFIT","FR_NOTARY","DE_SKR03","DE_SKR04","ES_PGC"]`<br>example `"FR_PCG"` |
| `$.company.id` | integer (int64) | obligatoire | example `123456` |
| `$.company.name` | string | obligatoire | example `"Pennylane"` |
| `$.company.reg_no` | string | obligatoire | example `"123456789"` |
| `$.scopes` | array | obligatoire | OAuth scopes associated with the access token<br>example `["customer_invoices","suppliers"]` |
| `$.scopes[]` | string | par élément |  |
| `$.user` | object ou null | obligatoire | additionalProperties `false` |
| `$.user.email` | string (email) | obligatoire | example `"jdoe@pennylane.com"` |
| `$.user.first_name` | string | obligatoire | example `"John"` |
| `$.user.id` | integer (int64) | obligatoire | example `12345` |
| `$.user.last_name` | string | obligatoire | example `"Doe"` |
| `$.user.locale` | string | obligatoire | enum `["fr","en","de"]`<br>example `"fr"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "company": {
    "accounting_logic": "FR_PCG",
    "id": 123456,
    "name": "Pennylane",
    "reg_no": "123456789"
  },
  "scopes": [
    "customer_invoices",
    "suppliers"
  ],
  "user": {
    "email": "jdoe@pennylane.com",
    "first_name": "John",
    "id": 12345,
    "last_name": "Doe",
    "locale": "fr"
  }
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
