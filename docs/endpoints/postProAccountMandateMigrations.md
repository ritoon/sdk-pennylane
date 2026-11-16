# Migrate a mandate to Pro Account

`POST /api/external/v2/pro_account/mandate_migrations`

This endpoint allows you to migrate a mandate to a Pro Account.
Only mandates with status 'available' are eligible for migration.

**Requirements:**
- Company must have a Pro Account (returns 404 if not)
- Company must have an enabled merchant profile (returns 403 if not)


> ℹ️
> This endpoint requires the following scope: `customer_mandates:all`

[Source Pennylane](https://pennylane.readme.io/reference/postproaccountmandatemigrations) · [Index du SDK](../operations.md)

## Appel Go

```go
PostProAccountMandateMigrationsWithResponse(ctx context.Context, params *PostProAccountMandateMigrationsParams, body PostProAccountMandateMigrationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostProAccountMandateMigrationsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.early_execution_date_permitted` | boolean | facultatif | Whether to permit early execution date for the migration.<br>Defaults to false if not specified.<br><br>example `false` |
| `$.mandate_id` | integer (int64) | obligatoire | The ID of the mandate to migrate.<br>example `12345` |
| `$.mandate_type` | string | obligatoire | The type of the mandate to migrate.<br>Use "Mandate" for GoCardless mandates or "SepaMandate" for SEPA mandates.<br><br>enum `["SepaMandate","Mandate"]`<br>example `"Mandate"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "early_execution_date_permitted": false,
  "mandate_id": 12345,
  "mandate_type": "Mandate"
}
```

</details>

## Sortie

### HTTP 201

Mandate migration initiated successfully

`application/json` → `réponse.JSON201` (type `*PostProAccountMandateMigrationsStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.mandate_migration` | object | obligatoire | additionalProperties `false` |
| `$.mandate_migration.created_at` | string (date-time) | obligatoire | example `"2023-08-07T14:23:12.000Z"` |
| `$.mandate_migration.customer` | object | obligatoire | additionalProperties `false` |
| `$.mandate_migration.customer.id` | integer (int64) | obligatoire | example `42` |
| `$.mandate_migration.customer.pro_account_mandate` | object | obligatoire | Link to the pro account mandate for this customer.<br>additionalProperties `false` |
| `$.mandate_migration.customer.pro_account_mandate.url` | string | obligatoire | URL to list pro account mandates for this customer.<br>example `"https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=..."` |
| `$.mandate_migration.customer.url` | string | obligatoire | URL to get the customer.<br>example `"https://app.pennylane.com/api/external/v2/customers/42"` |
| `$.mandate_migration.direct_debit_method` | string ou null | obligatoire | The direct debit method used by this mandate.<br>- `GoCardless`: GoCardless mandate<br>- `XML`: SEPA mandate<br>- `ProAccount`: Pro Account mandate<br><br>enum `["GoCardless","XML","ProAccount"]`<br>example `"GoCardless"` |
| `$.mandate_migration.error_message` | string ou null | obligatoire | Error message if the migration failed<br>example `null` |
| `$.mandate_migration.id` | integer (int64) | obligatoire | example `42` |
| `$.mandate_migration.mandate` | object | obligatoire | A polymorphic mandate object can either be a SepaMandate or a Gocardless mandate.<br>additionalProperties `false` |
| `$.mandate_migration.mandate.id` | integer (int64) | obligatoire | example `42` |
| `$.mandate_migration.mandate.type` | string | obligatoire | example `"SepaMandate"` |
| `$.mandate_migration.migrated_at` | string (date-time) ou null | obligatoire | Date when the mandate was successfully migrated<br>example `null` |
| `$.mandate_migration.migration_started_at` | string (date-time) ou null | obligatoire | Date when the migration was started<br>example `null` |
| `$.mandate_migration.signed_at` | string (date-time) ou null | obligatoire | Date when the mandate was signed<br>example `"2023-08-07T14:23:12.000Z"` |
| `$.mandate_migration.status` | string | obligatoire | example `"pending"` |
| `$.mandate_migration.updated_at` | string (date-time) | obligatoire | example `"2023-08-07T14:23:12.000Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "mandate_migration": {
    "created_at": "2023-08-07T14:23:12.000Z",
    "customer": {
      "id": 42,
      "pro_account_mandate": {
        "url": "https://app.pennylane.com/api/external/v2/pro_account/mandates?filter=..."
      },
      "url": "https://app.pennylane.com/api/external/v2/customers/42"
    },
    "direct_debit_method": "GoCardless",
    "error_message": null,
    "id": 42,
    "mandate": {
      "id": 42,
      "type": "SepaMandate"
    },
    "migrated_at": null,
    "migration_started_at": null,
    "signed_at": "2023-08-07T14:23:12.000Z",
    "status": "pending",
    "updated_at": "2023-08-07T14:23:12.000Z"
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
