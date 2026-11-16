# Update a contact of a customer

`PUT /api/external/v2/customers/{customer_id}/contacts/{id}`

Update a contact of a customer.
Only the submitted fields are updated, the others are left unchanged.
Contacts are a separate resource: updating one does not change the invoice
recipients of the customer.


> ℹ️
> This endpoint requires the following scope: `customers:all`

[Source Pennylane](https://pennylane.readme.io/reference/putcustomercontact) · [Index du SDK](../operations.md)

## Appel Go

```go
PutCustomerContactWithResponse(ctx context.Context, customerId int, id int64, params *PutCustomerContactParams, body PutCustomerContactJSONRequestBody, reqEditors ...RequestEditorFn) (*PutCustomerContactResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `customer_id` | `customerId` | path | integer | obligatoire | <br>example `42` |
| `id` | `id` | path | integer (int64) | obligatoire | <br>example `421` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.email` | string (email) | facultatif | maxLength `2000`<br>example `"jane.doe@example.com"` |
| `$.first_name` | string | facultatif | maxLength `2000`<br>pattern `"\\S"`<br>example `"Jane"` |
| `$.last_name` | string | facultatif | maxLength `2000`<br>pattern `"\\S"`<br>example `"Doe"` |
| `$.mobile_number` | string | facultatif | Mobile number of the contact.<br>Send an empty string to clear it.<br><br>maxLength `1744`<br>example `"+33612345678"` |
| `$.role` | string | facultatif | Role of the contact within the customer organization.<br>Send an empty string to clear it.<br><br>maxLength `2000`<br>example `"accountant"` |
| `$.telephone_number` | string | facultatif | Landline number of the contact.<br>Send an empty string to clear it.<br><br>maxLength `1744`<br>example `"+33123456789"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "email": "jane.doe@example.com",
  "first_name": "Jane",
  "last_name": "Doe",
  "mobile_number": "+33612345678",
  "role": "accountant",
  "telephone_number": "+33123456789"
}
```

</details>

## Sortie

### HTTP 200

Returns the updated contact

`application/json` → `réponse.JSON200` (type `*PutCustomerContactStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.email` | string | obligatoire | example `"jane.doe@example.com"` |
| `$.first_name` | string | obligatoire | example `"Jane"` |
| `$.id` | integer (int64) | obligatoire | example `421` |
| `$.last_name` | string | obligatoire | example `"Doe"` |
| `$.mobile_number` | string | obligatoire |  |
| `$.role` | string | obligatoire | example `"accountant"` |
| `$.telephone_number` | string | obligatoire |  |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "created_at": "2023-08-30T10:08:08.146343Z",
  "email": "jane.doe@example.com",
  "first_name": "Jane",
  "id": 421,
  "last_name": "Doe",
  "role": "accountant",
  "updated_at": "2023-08-30T10:08:08.146343Z"
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
