# Upload a file

`POST /api/external/v2/file_attachments`

Upload a file to attach to any resource that provides a `file_attachment_id`.

The maximum allowed file size is 100MB.
Note that this will not upload a file into the DMS (GED).


> ℹ️
> This endpoint requires the following scope: `file_attachments:all`

[Source Pennylane](https://pennylane.readme.io/reference/postfileattachments) · [Index du SDK](../operations.md)

## Appel Go

```go
PostFileAttachmentsWithResponse(ctx context.Context, params *PostFileAttachmentsParams, body PostFileAttachmentsMultipartRequestBody, editors ...RequestEditorFn) (*PostFileAttachmentsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `multipart/form-data`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.file` | string (binary) | obligatoire | The file you want to upload.<br>Allowed content types:<br>  - image/png<br>  - image/jpeg<br>  - image/tiff<br>  - image/bmp<br>  - image/gif<br>  - application/pdf |
| `$.filename` | string | facultatif | Name of the file. If not provided, the default value will be the uploaded file name.<br>maxLength `255`<br>pattern `"^[^\\x00]*$"`<br>example `"Invoice42.pdf"` |

## Sortie

### HTTP 201

Returns the created attachment

`application/json` → `réponse.JSON201` (type `*PostFileAttachmentsStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.created_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.filename` | string | obligatoire |  |
| `$.id` | integer (int64) | obligatoire |  |
| `$.updated_at` | string (date-time) | obligatoire | example `"2023-08-30T10:08:08.146343Z"` |
| `$.url` | string ou null | obligatoire | URL to the uploaded file.<br>This URL expires after 30 minutes.<br><br>example `"http://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/Invoice42.pdf"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "created_at": "2023-08-30T10:08:08.146343Z",
  "updated_at": "2023-08-30T10:08:08.146343Z",
  "url": "http://www.pennylane.com/rails/active_storage/blobs/redirect/eyJfc..==--26d6e7391dd728fae2cde59bd3fbe4dd11e20a95/Invoice42.pdf"
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
