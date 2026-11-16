# List eligible electronic billing addresses for a customer

`GET /api/external/v2/customers/{customer_id}/electronic_billing_addresses`

List the electronic billing addresses that can be selected for a company customer.

> ℹ️
> This endpoint requires one of the following scopes: `customers:all`, `customers:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getcustomerelectronicbillingaddresses) · [Index du SDK](../operations.md)

## Appel Go

```go
GetCustomerElectronicBillingAddressesWithResponse(ctx context.Context, customerId int64, params *GetCustomerElectronicBillingAddressesParams, reqEditors ...RequestEditorFn) (*GetCustomerElectronicBillingAddressesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_id` | `customerId` | path | integer (int64) | obligatoire | <br>example `42` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

Aucun corps de requête.

## Sortie

### HTTP 200

The list of eligible electronic billing addresses

`application/json` → `réponse.JSON200` (type `*GetCustomerElectronicBillingAddressesStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.has_more` | boolean | obligatoire | Always `false`. This endpoint returns<br>every eligible address in a single response.<br><br>example `false` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].electronic_billing_address` | string | obligatoire | example `"123456789_00012"` |
| `$.next_cursor` | string ou null | obligatoire | Always `null`. This endpoint returns<br>every eligible address in a single response. |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "has_more": false,
  "items": [
    {
      "electronic_billing_address": "123456789_00012"
    }
  ]
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

### HTTP 503

The service is temporarily unavailable

`application/json` → `réponse.JSON503` (type `*ImportCustomerInvoicesStatus503`).

[Champs du payload d’erreur](erreurs.md#importcustomerinvoicesstatus503)
