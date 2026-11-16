# Import a customer e-invoice

`POST /api/external/v2/customer_invoices/e_invoices/imports`

Import a customer invoice from an e-invoice file.
The file must be a Factur-X PDF, a standalone UBL XML invoice, or a standalone CII XML invoice.
Optionally provide `invoice_options` to pre-fill customer and line-level data.
Invoice line `e_invoice_line_id` must match Factur-X BT-126 (LineID) or the UBL/CII line identifier.

> ⚠️ **UBL and CII XML support is in alpha.**
> A 201 response confirms the file was accepted and parsed, but invoices imported via
> standalone UBL or CII XML have two known limitations:
> - The uploaded file is **not displayed** in the app — no document preview is generated yet for these formats.
> - The invoice is **not recognized as an electronic invoice** within Pennylane.
>
> Factur-X PDF remains the recommended format for production use.


> ℹ️
> This endpoint requires the following scope: `customer_invoices:all`

[Source Pennylane](https://pennylane.readme.io/reference/createcustomerinvoiceeinvoiceimport) · [Index du SDK](../operations.md)

## Appel Go

```go
CreateCustomerInvoiceEInvoiceImportWithResponse(ctx context.Context, params *CreateCustomerInvoiceEInvoiceImportParams, body CreateCustomerInvoiceEInvoiceImportMultipartRequestBody, editors ...RequestEditorFn) (*CreateCustomerInvoiceEInvoiceImportResponse, error)
```

## Entrée

| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |
|---|---|---|---|---|---|
| `X-Company-Id` | `params.XCompanyId` | header | string | facultatif | Numeric ID of the company to act on. Required when authenticating with a firm token. Ignored when the token is already bound to a company.<br>example `"12345"` |

**Corps obligatoire** · `multipart/form-data`

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.file` | string (binary) | obligatoire | The customer invoice file to be imported.<br>Must be a Factur-X PDF, a UBL XML invoice, or a CII XML invoice. |
| `$.installments` | array | facultatif | Optional list of installments for this invoice.<br>If not provided, a single installment matching the invoice total is created automatically.<br>The sum of all installments' `amount` values must equal the invoice total amount with tax.<br><br>> ℹ️<br>> This endpoint requires a company plan that has access to Installments, otherwise responds with 403.<br><br>minItems `1`<br>maxItems `60` |
| `$.installments[]` | object | par élément | title `"Customer Invoice Installment"`<br>additionalProperties `false` |
| `$.installments[].amount` | string | obligatoire | Installment amount (in invoice currency)<br>pattern `"^\\d+(\\.\\d+)?$"`<br>example `"100.00"` |
| `$.installments[].deadline` | string (date) | obligatoire | Installment payment deadline (ISO 8601)<br>example `"2023-08-30"` |
| `$.installments[].payment_methods` | array | facultatif | uniqueItems `true` |
| `$.installments[].payment_methods[]` | string | par élément | Allowed payment methods to be used in installments from the API.<br>enum `["transfer","e_transfer","credit_card","sepa","cheque","cash","lcr","pro_account_pos"]` |
| `$.invoice_options` | object | facultatif | Optional JSON payload to enrich the imported customer invoice.<br>Invoice line `e_invoice_line_id` must match Factur-X BT-126 (LineID).<br><br>additionalProperties `false` |
| `$.invoice_options.customer_id` | integer (int64) | facultatif | Customer id for the imported invoice.<br>example `12` |
| `$.invoice_options.invoice_lines` | array | facultatif |  |
| `$.invoice_options.invoice_lines[]` | object | par élément | additionalProperties `false` |
| `$.invoice_options.invoice_lines[].e_invoice_line_id` | string | obligatoire | Factur-X line identifier (BT-126 LineID). |
| `$.invoice_options.invoice_lines[].ledger_account_id` | integer (int64) ou null | facultatif | Ledger account id for this line.<br>example `123` |
| `$.invoice_options.invoice_lines[].product_id` | integer (int64) ou null | facultatif | Product id for this line.<br>example `456` |
| `$.override_invoice_lines` | array | facultatif | Complete set of accounting lines that **replaces** the ones produced by the<br>Factur-X parser. When provided, the parser's lines are discarded and the invoice is<br>built from these lines instead; the document total is still read from the file for a<br>coherence check. If the lines do not reconcile with the Factur-X total, the override<br>is ignored and the parsed Factur-X lines are kept.<br><br>Mutually exclusive with `invoice_options.invoice_lines` (sending both returns 400).<br><br>minItems `1`<br>maxItems `500` |
| `$.override_invoice_lines[]` | object | par élément | A complete accounting line used to override the Factur-X invoice lines.<br>Lines absent from the XML (e.g. a down-payment deduction, a reverse-charge line)<br>are expressed directly here — a TTC deduction is a normal line with a negative<br>`currency_amount`.<br><br>additionalProperties `false` |
| `$.override_invoice_lines[].currency_amount` | string | obligatoire | Line total amount excluding tax, in the invoice currency. May be negative.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.override_invoice_lines[].currency_tax` | string | obligatoire | Line VAT amount, in the invoice currency.<br>pattern `"^(-?\\d+(\\.\\d+)?)?$"`<br>example `"100.00"` |
| `$.override_invoice_lines[].ledger_account_id` | integer (int64) | obligatoire | The ledger account id (counterpart) for this line.<br>example `706` |
| `$.override_invoice_lines[].product_id` | integer (int64) ou null | facultatif | The product id for this line.<br>example `42` |
| `$.override_invoice_lines[].quantity` | number | facultatif | Line item quantity.<br>example `1` |
| `$.override_invoice_lines[].raw_currency_unit_price` | string | facultatif | The unit price excluding taxes. Can be set up to 6 decimals.<br>pattern `"^-?\\d+(\\.\\d{1,6})?$"`<br>example `"50.00"` |
| `$.override_invoice_lines[].vat_rate` | string | obligatoire | Product VAT rate. A 20% VAT in France is FR_200.<br>enum `["FR_1_05","FR_1_75","FR_09","FR_21","FR_40","FR_50","FR_55","FR_60","FR_65","FR_85","FR_92","FR_100","FR_130","FR_15_385","FR_160","FR_196","FR_200","AD_10","AD_45","AD_95","AT_100","AT_130","AT_190","AT_200","BE_60","BE_120","BE_210","BG_90","BG_200","CH_25","CH_37","CH_77","CH_26","CH_38","CH_81","CY_30","CY_50","CY_80","CY_90","CY_190","CZ_100","CZ_120","CZ_150","CZ_210","DE_70","DE_190","DK_250","EE_90","EE_200","EE_220","EE_240","ES_40","ES_70","ES_100","ES_210","FI_100","FI_135","FI_140","FI_240","FI_255","GB_50","GB_200","GR_60","GR_130","GR_170","GR_240","GR_40","HR_50","HR_130","HR_250","HU_50","HU_180","HU_270","IE_48","IE_90","IE_135","IE_210","IE_230","IT_40","IT_50","IT_100","IT_220","LT_50","LT_90","LT_120","LT_210","LU_30","LU_70","LU_80","LU_120","LU_130","LU_140","LU_160","LU_170","LV_50","LV_120","LV_210","MC_09","MC_21","MC_55","MC_85","MC_100","MC_200","MT_50","MT_70","MT_120","MT_180","MU_150","NL_90","NL_210","PL_50","PL_80","PL_230","PT_60","PT_130","PT_160","PT_180","PT_220","PT_230","RO_50","RO_90","RO_110","RO_190","RO_210","SE_60","SE_120","SE_250","SI_50","SI_95","SI_220","SK_50","SK_100","SK_190","SK_200","SK_230","NO_120","NO_150","NO_250","DE_70_BU0002","DE_190_BU0003","DE_70_BU0008","DE_190_BU0009","DE_intracom_70_BU0012","DE_intracom_190_BU0013","DE_190_electronic_BU0226","exempt","DE_intracom_exempt_BU0011","DE_intracom_exempt_BU0047","DE_intracom_exempt_BU0233","DE_exempt_BU0001","DE_exempt_BU0181","DE_exempt_BU0490","intracom_21","intracom_55","intracom_85","intracom_100","crossborder","ES_intracom_sales_goods","ES_intracom_sales_services","DE_intracom_70_BU0018","DE_intracom_190_BU0019","DE_intracom_190_BU0506","extracom","DE_extracom_BU0173","DE_extracom_BU0191","DE_extracom_190_BU0511","ES_isp_sales","ES_suplido","FR_85_construction","FR_100_construction","FR_200_construction","DE_exempt_construction_BU0046","DE_70_construction_BU0091","DE_190_construction_BU0094","DE_190_construction_BU0526","DE_190_construction_BU0546","ES_isp_purchases_210","ES_isp_purchases_100","mixed"]`<br>example `"FR_200"` |
| `$.send_to_pa` | boolean | facultatif | If true, the invoice will be automatically sent to the PA after parsing.<br>The send is optimistic: it happens without waiting for schematron validation.<br>If the invoice is invalid, the PA may reject it later (visible via e_invoicing.status). |

Partie `installments` encodée en `application/json`.

Partie `invoice_options` encodée en `application/json`.

Partie `override_invoice_lines` encodée en `application/json`.

## Sortie

### HTTP 201

Customer e-invoice imported successfully

`application/json` → `réponse.JSON201` (type `*CreateCustomerInvoiceEInvoiceImportStatus201`).

| Champ | Type | Présence | Description, contraintes et exemple |
|---|---|---|---|
| `$` | object | racine | additionalProperties `false` |
| `$.id` | integer (int64) | obligatoire | example `42` |
| `$.url` | string | obligatoire | The URL of the imported customer e-invoice.<br>This can be used to access the document later. |
| `$.warnings` | array | obligatoire | Non-blocking issues found during the import. The invoice was created,<br>but part of the data could not be applied. Empty when there is no issue.<br><br>example `[]` |
| `$.warnings[]` | string | par élément |  |

<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>

```json
{
  "id": 42,
  "warnings": []
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
