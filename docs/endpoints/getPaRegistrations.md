# List PA Registrations

`GET /api/external/v2/pa_registrations`

Returns all PA (Plateforme Agrée) registrations for the company,
including activation status and exchange direction. Use this to determine whether the company has completed
PA onboarding. Records with a `null` `siret` represent the SIREN-level
(head office), other records represent establishments.

> ⚠️
> By default, addresses of **every** status are returned. Only `activated` addresses are relevant to
> receive invoices, so we recommend filtering on the `activated` status:
> `filter=[{"field":"status","operator":"eq","value":"activated"}]`.
> In a future version, only `activated` addresses will be returned, and the `exchange_direction` property
> will be removed.


> ℹ️
> This endpoint requires the following scope: `pa_registrations:readonly`

[Source Pennylane](https://pennylane.readme.io/reference/getparegistrations) · [Index du SDK](../operations.md)

## Appel Go

```go
GetPaRegistrationsWithResponse(ctx context.Context, params *GetPaRegistrationsParams, reqEditors ...RequestEditorFn) (*GetPaRegistrationsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |
| `filter` | `params.Filter` | query | string | facultatif | You can choose to filter items on specific fields.<br>Available fields and operators:<br>- `id`: `lt`, `lteq`, `gt`, `gteq`, `eq`, `not_eq`, `in`, `not_in`<br>- `siret`: `eq`, `not_eq`, `in`, `not_in`<br>- `status`: `eq`, `not_eq`, `in`, `not_in`<br><br>We recommend filtering on the `activated` status to only get the addresses that are live in the<br>PA directory.<br>example `"[{\"field\": \"status\", \"operator\": \"eq\", \"value\": \"activated\"}]"` |
| `sort` | `params.Sort` | query | string | facultatif | You can choose to sort items on specific attributes.<br>Sort field may be prefixed with `-` for descending order.<br>Example: `id` will sort by ascending order, `-id` will sort by descending order.<br>Available fields: `id`<br>default `"-id"`<br>example `"id"` |

Aucun corps de requête.

## Sortie

### HTTP 200

List of PA registrations for the company

`application/json` → `réponse.JSON200` (type `*GetPaRegistrationsStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.emission_activated` | boolean | obligatoire | Whether the company has activated invoice emission through Pennylane.<br>`true` when the company is enabled to send invoices via Pennylane's emission feature,<br>regardless of whether a PPF routing address exists.<br>Use this alongside `items` to determine the full PA registration state.<br><br>example `false` |
| `$.has_more` | boolean | obligatoire | Indicates whether additional results are available beyond this set.<br>Use this flag to determine if another request is needed.<br><br>example `true` |
| `$.items` | array | obligatoire |  |
| `$.items[]` | object | par élément | additionalProperties `false` |
| `$.items[].created_at` | string (date-time) | obligatoire | example `"2026-02-23T10:53:00.000Z"` |
| `$.items[].effective_until` | string (date) ou null | obligatoire | Last date the address is effective, `null` when there is no end date.<br>example `"2027-01-01"` |
| `$.items[].exchange_direction` | string | obligatoire | Exchange direction of the address.<br><br>> ⚠️ This attribute is deprecated and will be removed in a future version.<br><br>enum `["emission","reception","emission_and_reception"]`<br>deprecated `true`<br>example `"emission_and_reception"` |
| `$.items[].id` | integer (int64) | obligatoire | example `1` |
| `$.items[].routing_code` | string ou null | obligatoire | Optional routing code of the address, `null` when the address has none.<br>example `"ROUTING-01"` |
| `$.items[].siren` | string | obligatoire | example `"993369859"` |
| `$.items[].siret` | string ou null | obligatoire | Establishment SIRET when the row is establishment-level; `null` when the row is SIREN-level (head office).<br><br>example `"98493148500010"` |
| `$.items[].status` | string | obligatoire | Registration status of a PA (Plateforme Agrée) address.<br><br>> ⚠️ All statuses except `activated` are deprecated. In a future version, only `activated` addresses will be returned.<br>- `activated`: live and confirmed in the PA directory; the address is registered.<br>- `provisioned`: [DEPRECATED] please do not use, records with this status will be deleted in a future version.<br>- `draft`: [DEPRECATED] the SME user started registration but did not complete it and explicitly saved the address as draft.<br>- `pending`: [DEPRECATED] consent completed and registration submitted to the PA directory, waiting for confirmation.<br>- `conflict`: [DEPRECATED] the directory callback indicates the address conflicts (e.g. already claimed elsewhere).<br>- `ineligible`: [DEPRECATED] the directory callback indicates the address cannot be registered.<br>- `lost`: [DEPRECATED] the address was lost to another PA — either after a portability transfer, or when another PA registered the same address on the PPF first.<br>- `archived`: [DEPRECATED] please do not use, records with this status will be deleted in a future version.<br><br>enum `["pending","activated","provisioned","draft","archived","conflict","ineligible","lost"]` |
| `$.items[].suffix` | string ou null | obligatoire | Optional suffix of the address, `null` when the address has none.<br>example `"invoices"` |
| `$.items[].updated_at` | string (date-time) | obligatoire | example `"2026-02-23T10:53:00.000Z"` |
| `$.next_cursor` | string ou null | obligatoire | Cursor to retrieve the next set of results.<br>Include this value in the cursor parameter of your next request to fetch subsequent items.<br>A `null` `next_cursor` in the response indicates no further results.<br><br>example `"dXBkYXRlZF9hdDoxNjc0MTIzNDU2"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "emission_activated": false,
  "has_more": true,
  "items": [
    {
      "created_at": "2026-02-23T10:53:00.000Z",
      "effective_until": "2027-01-01",
      "exchange_direction": "emission_and_reception",
      "id": 1,
      "routing_code": "ROUTING-01",
      "siren": "993369859",
      "siret": "98493148500010",
      "suffix": "invoices",
      "updated_at": "2026-02-23T10:53:00.000Z"
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
