# Upload an appendix for a commercial document

`POST /api/external/v2/commercial_documents/{commercial_document_id}/appendices`

Upload a file that will be an appendix attached to a commercial document.

Note that this will not upload a file into the DMS (GED).


> ℹ️
> This endpoint requires the following scope: `commercial_documents:all`

[Source Pennylane](https://pennylane.readme.io/reference/postcommercialdocumentappendices) · [Index du SDK](../operations.md)

## Appel Go

```go
PostCommercialDocumentAppendicesWithResponse(ctx context.Context, commercialDocumentId int, params *PostCommercialDocumentAppendicesParams, body PostCommercialDocumentAppendicesMultipartRequestBody, editors ...RequestEditorFn) (*PostCommercialDocumentAppendicesResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `commercial_document_id` | `commercialDocumentId` | path | integer | obligatoire | <br>example `1234` |
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `multipart/form-data`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.file` | string (binary) | obligatoire | The appendix file you want to upload.<br>Allowed content types:<br>  - image/png<br>  - image/jpeg<br>  - image/tiff<br>  - image/bmp<br>  - image/gif<br>  - application/pdf |

## Sortie

### HTTP 201

Returns the created appendix

`application/json` → `réponse.JSON201` (type `*PostCommercialDocumentAppendicesStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.filename` | string | obligatoire |  |
| `$.id` | integer (int64) | obligatoire |  |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.url` | string | obligatoire | URL to the uploaded file.<br>example `"https://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/Invoice42.pdf"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "created_at": "2023-08-30T10:08:08.146343Z",
  "updated_at": "2023-08-30T10:08:08.146343Z",
  "url": "https://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/Invoice42.pdf"
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
