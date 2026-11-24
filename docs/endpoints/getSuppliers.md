# List suppliers

`GET /api/external/v2/suppliers`

List suppliers

> ℹ️
> This endpoint requires one of the following scopes: `suppliers:all`, `suppliers:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getsuppliers) · [Index du SDK](../operations.md)

## Appel Go

```go
GetSuppliersWithResponse(ctx context.Context, params *GetSuppliersParams, reqEditors ...RequestEditorFn) (*GetSuppliersResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `cursor` | `params.Cursor` | query | string | facultatif | Cursor for pagination. Use this to fetch the next set of results.<br>The cursor is an opaque string returned in the previous response's metadata.<br>Leave empty for the first request.<br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |
| `limit` | `params.Limit` | query | integer | facultatif | Number of items to return per request.<br>Defaults to 20 if not specified.<br>Must be between 1 and 100.<br>minimum `1`<br>maximum `100`<br>example `20` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and values:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `ledger_account_id`: `eq`, `not_eq`, `in`, `not_in`<br>- `name`: `start_with`<br>- `external_reference`: `eq`, `not_eq`, `in`, `not_in`<br>- `emails`: `in`<br>example `"[{\"field\": \"ledger_account_id\", \"operator\": \"eq\", \"value\": \"123\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes<br>Sort field may be prefixed with `-` for descending order.<br>Example : `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields : `id`<br>default `"-id"` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

Returns a list of suppliers

`application/json` → `réponse.JSON200` (type `*GetSuppliersStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire | The list of items returned |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].bic` | string ou null | obligatoire | The Bank Identifier Code (BIC) of the supplier's bank. `null` when the supplier has no BIC for its current IBAN.<br>example `"BNPAFRPPXXX"` |
| `$.items[].created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].emails` | array | obligatoire | example `["hello@example.org"]` |
| `$.items[].emails[]` | string | par élément |  |
| `$.items[].establishment_no` | string ou null | obligatoire | Supplier identification number (SIRET).<br>- 14-digit number combining SIREN (9 digits) and NIC (5 digits)<br>- Only applicable for French companies<br><br>example `"82762938500014"` |
| `$.items[].external_reference` | string | obligatoire | The unique external reference that was assigned during creation either by you or Pennylane<br>example `"FR123"` |
| `$.items[].iban` | string | obligatoire | The IBAN of the supplier<br>example `"FRXXXXXXXXXXXXXXXXXXXXXXXXX"` |
| `$.items[].id` | integer (int64) | obligatoire | example `42` |
| `$.items[].ledger_account` | object ou null | obligatoire | additionalProperties `false` |
| `$.items[].ledger_account.id` | integer (int64) | obligatoire |  |
| `$.items[].name` | string | obligatoire | example `"Pennylane"` |
| `$.items[].postal_address` | object | obligatoire | additionalProperties `false` |
| `$.items[].postal_address.address` | string | obligatoire |  |
| `$.items[].postal_address.city` | string | obligatoire |  |
| `$.items[].postal_address.country_alpha2` | string | obligatoire |  |
| `$.items[].postal_address.postal_code` | string | obligatoire |  |
| `$.items[].reg_no` | string ou null | obligatoire | Supplier registration number (SIREN).<br>- 9-digit number<br>- Only applicable for French companies<br><br>example `"827629385"` |
| `$.items[].supplier_due_date_delay` | integer ou null | obligatoire | example `30` |
| `$.items[].supplier_due_date_rule` | string ou null | obligatoire | enum `["days","days_or_end_of_month"]` |
| `$.items[].supplier_payment_method` | string ou null | obligatoire | enum `["automatic_transfer","manual_transfer","automatic_debiting","bill_of_exchange","check","cash","card","other"]` |
| `$.items[].updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.items[].vat_number` | string | obligatoire | example `"FR32123456789"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": true,
  "items": [
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
