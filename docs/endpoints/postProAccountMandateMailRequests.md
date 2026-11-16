# Send a Pro Account SEPA mandate request

`POST /api/external/v2/pro_account/mandate_requests`

This endpoint allows you to send a mandate request for a Pro Account
SEPA Direct Debit mandate to a customer.

Requirements:
- Company must have a Pro Account (returns 404 if not)
- Company must have an enabled merchant profile (returns 403 if not)


> ℹ️
> This endpoint requires the following scope: `customer_mandates:all`

[Source Pennylane](https://pennylane.readme.io/reference/postproaccountmandatemailrequests) · [Index du SDK](../operations.md)

## Appel Go

```go
PostProAccountMandateMailRequestsWithResponse(ctx context.Context, params *PostProAccountMandateMailRequestsParams, body PostProAccountMandateMailRequestsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostProAccountMandateMailRequestsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.customer_id` | integer (int64) | obligatoire | Customer identifier<br>example `42` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "customer_id": 42
}
```

</details>

## Sortie

### HTTP 201

Mandate request successfully created

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
