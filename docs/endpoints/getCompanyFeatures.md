# Company features

`GET /api/external/v2/company/features`

Effective availability of the features that are gated for this company,
so that they can be checked before calling the endpoints they restrict.
A feature is `true` only when every gate is open, including the billing plan
and the progressive rollout, so it can be `false` on a plan that includes it.
Keys are added here as new gated features are introduced.

[Source Pennylane](https://pennylane.readme.io/reference/getcompanyfeatures) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCompanyFeaturesWithResponse(ctx context.Context, params *GetCompanyFeaturesParams, reqEditors ...RequestEditorFn) (*GetCompanyFeaturesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns the features of the company

`application/json` → `réponse.JSON200` (type `*GetCompanyFeaturesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.custom_fields` | boolean | obligatoire | Whether the company can set custom fields on customer invoices, quotes, and products.<br>True only when the billing plan includes custom field definitions.<br><br>example `true` |
| `$.installments` | boolean | obligatoire | Whether the company can pay a customer invoice in several installments.<br>When `false`, sending `installments` on the creation of a customer invoice either<br>responds with `403` or is ignored and creates a single installment, so check this<br>field rather than relying on the response.<br><br>example `true` |
| `$.products_and_services_catalog` | boolean | obligatoire | Whether the company can use the fields of the advanced products and services catalog, such as<br>`ean_code` on product creation and update. True only when the billing plan includes the purchase<br>product catalog, the advanced products and services catalog module is enabled, and the rollout is open.<br>When `false`, sending these fields is refused with `403`.<br><br>example `true` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "custom_fields": true,
  "installments": true,
  "products_and_services_catalog": true
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
