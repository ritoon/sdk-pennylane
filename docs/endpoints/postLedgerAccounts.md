# Create a ledger account

`POST /api/external/v2/ledger_accounts`

Create a ledger account


> ℹ️
> This endpoint requires the following scope: `ledger_accounts:all`

[Source Pennylane](https://pennylane.readme.io/reference/postledgeraccounts) · [Index du SDK](../operations.md)

## Appel Go

```go
PostLedgerAccountsWithResponse(ctx context.Context, params *PostLedgerAccountsParams, body PostLedgerAccountsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostLedgerAccountsResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `application/json`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.country_alpha2` | string | facultatif | enum `["AT","BE","BG","CY","CZ","DE","DK","EE","ES","FI","FR","GR","HR","HU","IE","IT","LT","LU","LV","MT","NL","PL","PT","RO","SE","SI","SK","GB","MC","CH","AD","MU","NO","XK","any"]` |
| `$.label` | string | obligatoire | Ledger Account's label<br>example `"Secondary Account"` |
| `$.number` | string | obligatoire | Ledger Account's number.<br>If the number starts with 401 (supplier) or 411 (customer) a corresponding supplier or company customer will also be created.<br><br>pattern `"^\\S+$"`<br>example `"512"` |
| `$.vat_rate` | string | facultatif | enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_100","SK_190","SK_200","SK_230","SK_50","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","DE_exempt_construction_BU0046","DE_190_construction_BU0526","DE_190_construction_BU0546","exempt","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","extracom","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","ES_isp_purchases_210","ES_isp_purchases_100","any"]`<br>example `"FR_200"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "label": "Secondary Account",
  "number": "512",
  "vat_rate": "FR_200"
}
```

</details>

## Sortie

### HTTP 201

Returns the created ledger account

`application/json` → `réponse.JSON201` (type `*GetLedgerAccountStatus200`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.country_alpha2` | string | obligatoire | Ledger Account's country code (alpha2)<br>example `"FR"` |
| `$.created_at` | string (date-time) | obligatoire | The time the ledger account has been created<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.enabled` | boolean | obligatoire | example `true` |
| `$.id` | integer (int64) | obligatoire | example `124` |
| `$.label` | string | obligatoire | example `"Secondary Account"` |
| `$.letterable` | boolean | obligatoire | example `false` |
| `$.number` | string | obligatoire | example `"512"` |
| `$.type` | string | obligatoire | example `"custom"` |
| `$.updated_at` | string (date-time) | obligatoire | The last time the ledger account has been updated<br>example `"2023-08-30T10:08:08.146343Z"` |
| `$.vat_rate` | string | obligatoire | Ledger Account's VAT rate in percentage<br>example `"any"` |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "country_alpha2": "FR",
  "created_at": "2023-08-30T10:08:08.146343Z",
  "enabled": true,
  "id": 124,
  "label": "Secondary Account",
  "letterable": false,
  "number": "512",
  "type": "custom",
  "updated_at": "2023-08-30T10:08:08.146343Z",
  "vat_rate": "any"
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
