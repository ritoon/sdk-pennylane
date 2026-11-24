# Update a supplier

`PUT /api/external/v2/suppliers/{id}`

This endpoint returns the updated supplier.

> ℹ️
> This endpoint requires the following scope: `suppliers:all`

[Source Pennylane](https://pennylane.readme.io/reference/putsupplier) · [Index du SDK](../operations.md)

## Appel Go

```go
PutSupplierWithResponse(ctx context.Context, id int, params *PutSupplierParams, body PutSupplierJSONRequestBody, reqEditors ...RequestEditorFn) (*PutSupplierResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `id` | `id` | path | integer | obligatoire | <br>example `1` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.emails` | array | facultatif | example `["hello@example.org"]` |
| `$.emails[]` | string | par élément |  |
| `$.establishment_no` | string ou null | facultatif | Supplier identification number (SIRET).<br>- 14-digit number combining SIREN (9 digits) and NIC (5 digits)<br>- Only applicable for French companies<br><br>pattern `"^\\d{14}$"`<br>example `"82762938500014"` |
| `$.external_reference` | string | facultatif | The unique external reference that was assigned during creation either by you or Pennylane.<br>example `"FR123"` |
| `$.iban` | string | facultatif | example `"FR3330002005500000157841Z25"` |
| `$.name` | string | facultatif | example `"Pennylane"` |
| `$.postal_address` | object | facultatif | additionalProperties `false` |
| `$.postal_address.address` | string | obligatoire | example `"8 rue de la paix"` |
| `$.postal_address.city` | string | obligatoire | example `"Paris"` |
| `$.postal_address.country_alpha2` | string | obligatoire | example `"FR"` |
| `$.postal_address.postal_code` | string | obligatoire | example `"75002"` |
| `$.reg_no` | string | facultatif | Supplier registration number (SIREN).<br>- 9-digit number<br>- Only applicable for French companies<br><br>pattern `"^\\d{9}$"`<br>example `"827629385"` |
| `$.supplier_due_date_delay` | integer ou null | facultatif | example `30` |
| `$.supplier_due_date_rule` | string ou null | facultatif | enum `["days","days_or_end_of_month"]` |
| `$.supplier_payment_method` | string ou null | facultatif | enum `["automatic_transfer","manual_transfer","automatic_debiting","bill_of_exchange","check","cash","card","other"]` |
| `$.vat_number` | string | facultatif | example `"FR32123456789"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "emails": [
    "hello@example.org"
  ],
  "establishment_no": "82762938500014",
  "external_reference": "FR123",
  "iban": "FR3330002005500000157841Z25",
  "name": "Pennylane",
  "postal_address": {
    "address": "8 rue de la paix",
    "city": "Paris",
    "country_alpha2": "FR",
    "postal_code": "75002"
  },
  "reg_no": "827629385",
  "supplier_due_date_delay": 30,
  "vat_number": "FR32123456789"
}
```

</details>

## Sortie

### HTTP 200

Returns the updated supplier

`application/json` → `réponse.JSON200` (type `*GetSupplierStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.bic` | string ou null | obligatoire | The Bank Identifier Code (BIC) of the supplier's bank. `null` when the supplier has no BIC for its current IBAN.<br>example `"BNPAFRPPXXX"` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.emails` | array | obligatoire | example `["hello@example.org"]` |
| `$.emails[]` | string | par élément |  |
| `$.establishment_no` | string ou null | obligatoire | Supplier identification number (SIRET).<br>- 14-digit number combining SIREN (9 digits) and NIC (5 digits)<br>- Only applicable for French companies<br><br>example `"82762938500014"` |
| `$.external_reference` | string | obligatoire | The unique external reference that was assigned during creation either by you or Pennylane<br>example `"FR123"` |
| `$.iban` | string | obligatoire | The IBAN of the supplier<br>example `"FRXXXXXXXXXXXXXXXXXXXXXXXXX"` |
| `$.id` | integer (int64) | obligatoire | example `42` |
| `$.ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$.ledger_account.id` | integer (int64) | obligatoire |  |
| `$.name` | string | obligatoire | example `"Pennylane"` |
| `$.postal_address` | object | obligatoire | additionalProperties `false` |
| `$.postal_address.address` | string | obligatoire |  |
| `$.postal_address.city` | string | obligatoire |  |
| `$.postal_address.country_alpha2` | string | obligatoire |  |
| `$.postal_address.postal_code` | string | obligatoire |  |
| `$.reg_no` | string ou null | obligatoire | Supplier registration number (SIREN).<br>- 9-digit number<br>- Only applicable for French companies<br><br>example `"827629385"` |
| `$.supplier_due_date_delay` | integer ou null | obligatoire | example `30` |
| `$.supplier_due_date_rule` | string ou null | obligatoire | enum `["days","days_or_end_of_month"]` |
| `$.supplier_payment_method` | string ou null | obligatoire | enum `["automatic_transfer","manual_transfer","automatic_debiting","bill_of_exchange","check","cash","card","other"]` |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.vat_number` | string | obligatoire | example `"FR32123456789"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "bic": "BNPAFRPPXXX",
  "created_at": "2023-08-30T10:08:08.146343Z",
  "emails": [
    "hello@example.org"
  ],
  "establishment_no": "82762938500014",
  "external_reference": "FR123",
  "iban": "FRXXXXXXXXXXXXXXXXXXXXXXXXX",
  "id": 42,
  "name": "Pennylane",
  "reg_no": "827629385",
  "supplier_due_date_delay": 30,
  "updated_at": "2023-08-30T10:08:08.146343Z",
  "vat_number": "FR32123456789"
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
