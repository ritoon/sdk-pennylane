# Référence du client Go

Les signatures indiquent les arguments métier et le type de retour. Toutes ces méthodes acceptent également des éditeurs de requête facultatifs en dernier argument. Le lien de chaque route ouvre sa fiche officielle (champs, statuts et permissions).

## Bank Establishments

- [`GET /api/external/v2/bank_establishments`](https://pennylane.readme.io/reference/getbankestablishments)

  `GetBankEstablishmentsWithResponse(ctx context.Context, params *GetBankEstablishmentsParams) (*GetBankEstablishmentsResponse, error)`

## Bank accounts

- [`GET /api/external/v2/bank_accounts`](https://pennylane.readme.io/reference/getbankaccounts)

  `GetBankAccountsWithResponse(ctx context.Context, params *GetBankAccountsParams) (*GetBankAccountsResponse, error)`

- [`POST /api/external/v2/bank_accounts`](https://pennylane.readme.io/reference/postbankaccount)

  `PostBankAccountWithResponse(ctx context.Context, params *PostBankAccountParams, body PostBankAccountJSONRequestBody) (*PostBankAccountResponse, error)`

- [`GET /api/external/v2/bank_accounts/{id}`](https://pennylane.readme.io/reference/getbankaccount)

  `GetBankAccountWithResponse(ctx context.Context, id int, params *GetBankAccountParams) (*GetBankAccountResponse, error)`

## Billing Subscriptions

- [`GET /api/external/v2/billing_subscriptions`](https://pennylane.readme.io/reference/getbillingsubscriptions)

  `GetBillingSubscriptionsWithResponse(ctx context.Context, params *GetBillingSubscriptionsParams) (*GetBillingSubscriptionsResponse, error)`

- [`POST /api/external/v2/billing_subscriptions`](https://pennylane.readme.io/reference/postbillingsubscriptions)

  `PostBillingSubscriptionsWithResponse(ctx context.Context, params *PostBillingSubscriptionsParams, body PostBillingSubscriptionsJSONRequestBody) (*PostBillingSubscriptionsResponse, error)`

- [`GET /api/external/v2/billing_subscriptions/{billing_subscription_id}/invoice_line_sections`](https://pennylane.readme.io/reference/getbillingsubscriptioninvoicelinesections)

  `GetBillingSubscriptionInvoiceLineSectionsWithResponse(ctx context.Context, billingSubscriptionId int, params *GetBillingSubscriptionInvoiceLineSectionsParams) (*GetBillingSubscriptionInvoiceLineSectionsResponse, error)`

- [`GET /api/external/v2/billing_subscriptions/{billing_subscription_id}/invoice_lines`](https://pennylane.readme.io/reference/getbillingsubscriptioninvoicelines)

  `GetBillingSubscriptionInvoiceLinesWithResponse(ctx context.Context, billingSubscriptionId int, params *GetBillingSubscriptionInvoiceLinesParams) (*GetBillingSubscriptionInvoiceLinesResponse, error)`

- [`GET /api/external/v2/billing_subscriptions/{id}`](https://pennylane.readme.io/reference/getbillingsubscription)

  `GetBillingSubscriptionWithResponse(ctx context.Context, id int, params *GetBillingSubscriptionParams) (*GetBillingSubscriptionResponse, error)`

- [`PUT /api/external/v2/billing_subscriptions/{id}`](https://pennylane.readme.io/reference/putbillingsubscriptions)

  `PutBillingSubscriptionsWithResponse(ctx context.Context, id int, params *PutBillingSubscriptionsParams, body PutBillingSubscriptionsJSONRequestBody) (*PutBillingSubscriptionsResponse, error)`

## Billing email templates

- [`GET /api/external/v2/billing_email_templates`](https://pennylane.readme.io/reference/getbillingemailtemplates)

  `GetBillingEmailTemplatesWithResponse(ctx context.Context, params *GetBillingEmailTemplatesParams) (*GetBillingEmailTemplatesResponse, error)`

## Categories

- [`GET /api/external/v2/categories`](https://pennylane.readme.io/reference/getcategories)

  `GetCategoriesWithResponse(ctx context.Context, params *GetCategoriesParams) (*GetCategoriesResponse, error)`

- [`POST /api/external/v2/categories`](https://pennylane.readme.io/reference/postcategories)

  `PostCategoriesWithResponse(ctx context.Context, params *PostCategoriesParams, body PostCategoriesJSONRequestBody) (*PostCategoriesResponse, error)`

- [`GET /api/external/v2/categories/{id}`](https://pennylane.readme.io/reference/getcategory)

  `GetCategoryWithResponse(ctx context.Context, id string, params *GetCategoryParams) (*GetCategoryResponse, error)`

- [`PUT /api/external/v2/categories/{id}`](https://pennylane.readme.io/reference/updatecategory)

  `UpdateCategoryWithResponse(ctx context.Context, id string, params *UpdateCategoryParams, body UpdateCategoryJSONRequestBody) (*UpdateCategoryResponse, error)`

- [`GET /api/external/v2/category_groups`](https://pennylane.readme.io/reference/getcategorygroups)

  `GetCategoryGroupsWithResponse(ctx context.Context, params *GetCategoryGroupsParams) (*GetCategoryGroupsResponse, error)`

## Category Groups

- [`POST /api/external/v2/category_groups`](https://pennylane.readme.io/reference/postcategorygroups)

  `PostCategoryGroupsWithResponse(ctx context.Context, params *PostCategoryGroupsParams, body PostCategoryGroupsJSONRequestBody) (*PostCategoryGroupsResponse, error)`

- [`GET /api/external/v2/category_groups/{category_group_id}/categories`](https://pennylane.readme.io/reference/getcategorygroupcategories)

  `GetCategoryGroupCategoriesWithResponse(ctx context.Context, categoryGroupId int, params *GetCategoryGroupCategoriesParams) (*GetCategoryGroupCategoriesResponse, error)`

- [`GET /api/external/v2/category_groups/{id}`](https://pennylane.readme.io/reference/getcategorygroup)

  `GetCategoryGroupWithResponse(ctx context.Context, id string, params *GetCategoryGroupParams) (*GetCategoryGroupResponse, error)`

- [`PUT /api/external/v2/category_groups/{id}`](https://pennylane.readme.io/reference/putcategorygroup)

  `PutCategoryGroupWithResponse(ctx context.Context, id string, params *PutCategoryGroupParams, body PutCategoryGroupJSONRequestBody) (*PutCategoryGroupResponse, error)`

## Changelogs

- [`GET /api/external/v2/changelogs/customer_invoices`](https://pennylane.readme.io/reference/getcustomerinvoiceschanges)

  `GetCustomerInvoicesChangesWithResponse(ctx context.Context, params *GetCustomerInvoicesChangesParams) (*GetCustomerInvoicesChangesResponse, error)`

- [`GET /api/external/v2/changelogs/customers`](https://pennylane.readme.io/reference/getcustomerchanges)

  `GetCustomerChangesWithResponse(ctx context.Context, params *GetCustomerChangesParams) (*GetCustomerChangesResponse, error)`

- [`GET /api/external/v2/changelogs/ledger_entries`](https://pennylane.readme.io/reference/getledgerentrieschanges)

  `GetLedgerEntriesChangesWithResponse(ctx context.Context, params *GetLedgerEntriesChangesParams) (*GetLedgerEntriesChangesResponse, error)`

- [`GET /api/external/v2/changelogs/ledger_entries_categories`](https://pennylane.readme.io/reference/getledgerentriescategorychanges)

  `GetLedgerEntriesCategoryChangesWithResponse(ctx context.Context, params *GetLedgerEntriesCategoryChangesParams) (*GetLedgerEntriesCategoryChangesResponse, error)`

- [`GET /api/external/v2/changelogs/ledger_entry_lines`](https://pennylane.readme.io/reference/getledgerentrylinechanges)

  `GetLedgerEntryLineChangesWithResponse(ctx context.Context, params *GetLedgerEntryLineChangesParams) (*GetLedgerEntryLineChangesResponse, error)`

- [`GET /api/external/v2/changelogs/ledger_entry_lines_categories`](https://pennylane.readme.io/reference/getledgerentrylinescategorychanges)

  `GetLedgerEntryLinesCategoryChangesWithResponse(ctx context.Context, params *GetLedgerEntryLinesCategoryChangesParams) (*GetLedgerEntryLinesCategoryChangesResponse, error)`

- [`GET /api/external/v2/changelogs/products`](https://pennylane.readme.io/reference/getproductchanges)

  `GetProductChangesWithResponse(ctx context.Context, params *GetProductChangesParams) (*GetProductChangesResponse, error)`

- [`GET /api/external/v2/changelogs/supplier_invoices`](https://pennylane.readme.io/reference/getsupplierinvoiceschanges)

  `GetSupplierInvoicesChangesWithResponse(ctx context.Context, params *GetSupplierInvoicesChangesParams) (*GetSupplierInvoicesChangesResponse, error)`

- [`GET /api/external/v2/changelogs/suppliers`](https://pennylane.readme.io/reference/getsupplierchanges)

  `GetSupplierChangesWithResponse(ctx context.Context, params *GetSupplierChangesParams) (*GetSupplierChangesResponse, error)`

- [`GET /api/external/v2/changelogs/transactions`](https://pennylane.readme.io/reference/gettransactionchanges)

  `GetTransactionChangesWithResponse(ctx context.Context, params *GetTransactionChangesParams) (*GetTransactionChangesResponse, error)`

## Commercial Documents

- [`GET /api/external/v2/commercial_documents`](https://pennylane.readme.io/reference/listcommercialdocuments)

  `ListCommercialDocumentsWithResponse(ctx context.Context, params *ListCommercialDocumentsParams) (*ListCommercialDocumentsResponse, error)`

- [`POST /api/external/v2/commercial_documents`](https://pennylane.readme.io/reference/postcommercialdocuments)

  `PostCommercialDocumentsWithResponse(ctx context.Context, params *PostCommercialDocumentsParams, body PostCommercialDocumentsJSONRequestBody) (*PostCommercialDocumentsResponse, error)`

- [`GET /api/external/v2/commercial_documents/{commercial_document_id}/appendices`](https://pennylane.readme.io/reference/getcommercialdocumentappendices)

  `GetCommercialDocumentAppendicesWithResponse(ctx context.Context, commercialDocumentId int, params *GetCommercialDocumentAppendicesParams) (*GetCommercialDocumentAppendicesResponse, error)`

- [`POST /api/external/v2/commercial_documents/{commercial_document_id}/appendices`](https://pennylane.readme.io/reference/postcommercialdocumentappendices)

  `PostCommercialDocumentAppendicesWithResponse(ctx context.Context, commercialDocumentId int, params *PostCommercialDocumentAppendicesParams, body PostCommercialDocumentAppendicesMultipartRequestBody) (*PostCommercialDocumentAppendicesResponse, error)`

- [`GET /api/external/v2/commercial_documents/{commercial_document_id}/invoice_line_sections`](https://pennylane.readme.io/reference/getcommercialdocumentinvoicelinesections)

  `GetCommercialDocumentInvoiceLineSectionsWithResponse(ctx context.Context, commercialDocumentId int, params *GetCommercialDocumentInvoiceLineSectionsParams) (*GetCommercialDocumentInvoiceLineSectionsResponse, error)`

- [`GET /api/external/v2/commercial_documents/{commercial_document_id}/invoice_lines`](https://pennylane.readme.io/reference/getcommercialdocumentinvoicelines)

  `GetCommercialDocumentInvoiceLinesWithResponse(ctx context.Context, commercialDocumentId int, params *GetCommercialDocumentInvoiceLinesParams) (*GetCommercialDocumentInvoiceLinesResponse, error)`

- [`GET /api/external/v2/commercial_documents/{id}`](https://pennylane.readme.io/reference/getcommercialdocument)

  `GetCommercialDocumentWithResponse(ctx context.Context, id int, params *GetCommercialDocumentParams) (*GetCommercialDocumentResponse, error)`

- [`PUT /api/external/v2/commercial_documents/{id}`](https://pennylane.readme.io/reference/updatecommercialdocument)

  `UpdateCommercialDocumentWithResponse(ctx context.Context, id int, params *UpdateCommercialDocumentParams, body UpdateCommercialDocumentJSONRequestBody) (*UpdateCommercialDocumentResponse, error)`

- [`POST /api/external/v2/commercial_documents/{id}/send_by_email`](https://pennylane.readme.io/reference/sendbyemailcommercialdocument)

  `SendByEmailCommercialDocumentWithResponse(ctx context.Context, id int, params *SendByEmailCommercialDocumentParams, body SendByEmailCommercialDocumentJSONRequestBody) (*SendByEmailCommercialDocumentResponse, error)`

## Company

- [`GET /api/external/v2/company/features`](https://pennylane.readme.io/reference/getcompanyfeatures)

  `GetCompanyFeaturesWithResponse(ctx context.Context, params *GetCompanyFeaturesParams) (*GetCompanyFeaturesResponse, error)`

## Customer Invoice Templates

- [`GET /api/external/v2/customer_invoice_templates`](https://pennylane.readme.io/reference/getcustomerinvoicetemplates)

  `GetCustomerInvoiceTemplatesWithResponse(ctx context.Context, params *GetCustomerInvoiceTemplatesParams) (*GetCustomerInvoiceTemplatesResponse, error)`

## Customer Invoices

- [`GET /api/external/v2/customer_invoices`](https://pennylane.readme.io/reference/getcustomerinvoices)

  `GetCustomerInvoicesWithResponse(ctx context.Context, params *GetCustomerInvoicesParams) (*GetCustomerInvoicesResponse, error)`

- [`POST /api/external/v2/customer_invoices`](https://pennylane.readme.io/reference/postcustomerinvoices)

  `PostCustomerInvoicesWithResponse(ctx context.Context, params *PostCustomerInvoicesParams, body PostCustomerInvoicesJSONRequestBody) (*PostCustomerInvoicesResponse, error)`

- [`POST /api/external/v2/customer_invoices/create_from_quote`](https://pennylane.readme.io/reference/createcustomerinvoicefromquote)

  `CreateCustomerInvoiceFromQuoteWithResponse(ctx context.Context, params *CreateCustomerInvoiceFromQuoteParams, body CreateCustomerInvoiceFromQuoteJSONRequestBody) (*CreateCustomerInvoiceFromQuoteResponse, error)`

- [`POST /api/external/v2/customer_invoices/e_invoices/imports`](https://pennylane.readme.io/reference/createcustomerinvoiceeinvoiceimport)

  `CreateCustomerInvoiceEInvoiceImportWithResponse(ctx context.Context, params *CreateCustomerInvoiceEInvoiceImportParams, body CreateCustomerInvoiceEInvoiceImportMultipartRequestBody) (*CreateCustomerInvoiceEInvoiceImportResponse, error)`

- [`POST /api/external/v2/customer_invoices/import`](https://pennylane.readme.io/reference/importcustomerinvoices)

  `ImportCustomerInvoicesWithResponse(ctx context.Context, params *ImportCustomerInvoicesParams, body ImportCustomerInvoicesJSONRequestBody) (*ImportCustomerInvoicesResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/appendices`](https://pennylane.readme.io/reference/getcustomerinvoiceappendices)

  `GetCustomerInvoiceAppendicesWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceAppendicesParams) (*GetCustomerInvoiceAppendicesResponse, error)`

- [`POST /api/external/v2/customer_invoices/{customer_invoice_id}/appendices`](https://pennylane.readme.io/reference/postcustomerinvoiceappendices)

  `PostCustomerInvoiceAppendicesWithResponse(ctx context.Context, customerInvoiceId int, params *PostCustomerInvoiceAppendicesParams, body PostCustomerInvoiceAppendicesMultipartRequestBody) (*PostCustomerInvoiceAppendicesResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/categories`](https://pennylane.readme.io/reference/getcustomerinvoicecategories)

  `GetCustomerInvoiceCategoriesWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceCategoriesParams) (*GetCustomerInvoiceCategoriesResponse, error)`

- [`PUT /api/external/v2/customer_invoices/{customer_invoice_id}/categories`](https://pennylane.readme.io/reference/putcustomerinvoicecategories)

  `PutCustomerInvoiceCategoriesWithResponse(ctx context.Context, customerInvoiceId int, params *PutCustomerInvoiceCategoriesParams, body PutCustomerInvoiceCategoriesJSONRequestBody) (*PutCustomerInvoiceCategoriesResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/comments`](https://pennylane.readme.io/reference/getcustomerinvoicecomments)

  `GetCustomerInvoiceCommentsWithResponse(ctx context.Context, customerInvoiceId int64, params *GetCustomerInvoiceCommentsParams) (*GetCustomerInvoiceCommentsResponse, error)`

- [`POST /api/external/v2/customer_invoices/{customer_invoice_id}/comments`](https://pennylane.readme.io/reference/postcustomerinvoicecomment)

  `PostCustomerInvoiceCommentWithResponse(ctx context.Context, customerInvoiceId int64, params *PostCustomerInvoiceCommentParams, body PostCustomerInvoiceCommentJSONRequestBody) (*PostCustomerInvoiceCommentResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/custom_header_fields`](https://pennylane.readme.io/reference/getcustomerinvoicecustomheaderfields)

  `GetCustomerInvoiceCustomHeaderFieldsWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceCustomHeaderFieldsParams) (*GetCustomerInvoiceCustomHeaderFieldsResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/installments`](https://pennylane.readme.io/reference/getcustomerinvoiceinstallments)

  `GetCustomerInvoiceInstallmentsWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceInstallmentsParams) (*GetCustomerInvoiceInstallmentsResponse, error)`

- [`PUT /api/external/v2/customer_invoices/{customer_invoice_id}/installments/{id}/mark_as_paid`](https://pennylane.readme.io/reference/markaspaidcustomerinvoiceinstallment)

  `MarkAsPaidCustomerInvoiceInstallmentWithResponse(ctx context.Context, customerInvoiceId int64, id int64, params *MarkAsPaidCustomerInvoiceInstallmentParams) (*MarkAsPaidCustomerInvoiceInstallmentResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/invoice_line_sections`](https://pennylane.readme.io/reference/getcustomerinvoiceinvoicelinesections)

  `GetCustomerInvoiceInvoiceLineSectionsWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceInvoiceLineSectionsParams) (*GetCustomerInvoiceInvoiceLineSectionsResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/invoice_lines`](https://pennylane.readme.io/reference/getcustomerinvoiceinvoicelines)

  `GetCustomerInvoiceInvoiceLinesWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceInvoiceLinesParams) (*GetCustomerInvoiceInvoiceLinesResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/matched_transactions`](https://pennylane.readme.io/reference/getcustomerinvoicematchedtransactions)

  `GetCustomerInvoiceMatchedTransactionsWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoiceMatchedTransactionsParams) (*GetCustomerInvoiceMatchedTransactionsResponse, error)`

- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/payments`](https://pennylane.readme.io/reference/getcustomerinvoicepayments)

  `GetCustomerInvoicePaymentsWithResponse(ctx context.Context, customerInvoiceId int, params *GetCustomerInvoicePaymentsParams) (*GetCustomerInvoicePaymentsResponse, error)`

- [`DELETE /api/external/v2/customer_invoices/{id}`](https://pennylane.readme.io/reference/deletecustomerinvoices)

  `DeleteCustomerInvoicesWithResponse(ctx context.Context, id int, params *DeleteCustomerInvoicesParams) (*DeleteCustomerInvoicesResponse, error)`

- [`GET /api/external/v2/customer_invoices/{id}`](https://pennylane.readme.io/reference/getcustomerinvoice)

  `GetCustomerInvoiceWithResponse(ctx context.Context, id int, params *GetCustomerInvoiceParams) (*GetCustomerInvoiceResponse, error)`

- [`PUT /api/external/v2/customer_invoices/{id}`](https://pennylane.readme.io/reference/updatecustomerinvoice)

  `UpdateCustomerInvoiceWithResponse(ctx context.Context, id int, params *UpdateCustomerInvoiceParams, body UpdateCustomerInvoiceJSONRequestBody) (*UpdateCustomerInvoiceResponse, error)`

- [`PUT /api/external/v2/customer_invoices/{id}/finalize`](https://pennylane.readme.io/reference/finalizecustomerinvoice)

  `FinalizeCustomerInvoiceWithResponse(ctx context.Context, id int, params *FinalizeCustomerInvoiceParams) (*FinalizeCustomerInvoiceResponse, error)`

- [`POST /api/external/v2/customer_invoices/{id}/link_credit_note`](https://pennylane.readme.io/reference/linkcreditnote)

  `LinkCreditNoteWithResponse(ctx context.Context, id int, params *LinkCreditNoteParams, body LinkCreditNoteJSONRequestBody) (*LinkCreditNoteResponse, error)`

- [`PUT /api/external/v2/customer_invoices/{id}/mark_as_paid`](https://pennylane.readme.io/reference/markaspaidcustomerinvoice)

  `MarkAsPaidCustomerInvoiceWithResponse(ctx context.Context, id int, params *MarkAsPaidCustomerInvoiceParams) (*MarkAsPaidCustomerInvoiceResponse, error)`

- [`POST /api/external/v2/customer_invoices/{id}/send_by_email`](https://pennylane.readme.io/reference/sendbyemailcustomerinvoice)

  `SendByEmailCustomerInvoiceWithResponse(ctx context.Context, id int, params *SendByEmailCustomerInvoiceParams, body SendByEmailCustomerInvoiceJSONRequestBody) (*SendByEmailCustomerInvoiceResponse, error)`

- [`POST /api/external/v2/customer_invoices/{id}/send_to_pa`](https://pennylane.readme.io/reference/sendtopacustomerinvoice)

  `SendToPaCustomerInvoiceWithResponse(ctx context.Context, id int64, params *SendToPaCustomerInvoiceParams) (*SendToPaCustomerInvoiceResponse, error)`

- [`PUT /api/external/v2/customer_invoices/{id}/update_imported`](https://pennylane.readme.io/reference/updateimportedcustomerinvoice)

  `UpdateImportedCustomerInvoiceWithResponse(ctx context.Context, id int, params *UpdateImportedCustomerInvoiceParams, body UpdateImportedCustomerInvoiceJSONRequestBody) (*UpdateImportedCustomerInvoiceResponse, error)`

## Customers

- [`POST /api/external/v2/company_customers`](https://pennylane.readme.io/reference/postcompanycustomer)

  `PostCompanyCustomerWithResponse(ctx context.Context, params *PostCompanyCustomerParams, body PostCompanyCustomerJSONRequestBody) (*PostCompanyCustomerResponse, error)`

- [`GET /api/external/v2/company_customers/{id}`](https://pennylane.readme.io/reference/getcompanycustomer)

  `GetCompanyCustomerWithResponse(ctx context.Context, id string, params *GetCompanyCustomerParams) (*GetCompanyCustomerResponse, error)`

- [`PUT /api/external/v2/company_customers/{id}`](https://pennylane.readme.io/reference/putcompanycustomer)

  `PutCompanyCustomerWithResponse(ctx context.Context, id string, params *PutCompanyCustomerParams, body PutCompanyCustomerJSONRequestBody) (*PutCompanyCustomerResponse, error)`

- [`GET /api/external/v2/customers`](https://pennylane.readme.io/reference/getcustomers)

  `GetCustomersWithResponse(ctx context.Context, params *GetCustomersParams) (*GetCustomersResponse, error)`

- [`POST /api/external/v2/customers`](https://pennylane.readme.io/reference/postcustomer)

  `PostCustomerWithResponse(ctx context.Context, params *PostCustomerParams, body PostCustomerJSONRequestBody) (*PostCustomerResponse, error)`

- [`GET /api/external/v2/customers/{customer_id}/categories`](https://pennylane.readme.io/reference/getcustomercategories)

  `GetCustomerCategoriesWithResponse(ctx context.Context, customerId int, params *GetCustomerCategoriesParams) (*GetCustomerCategoriesResponse, error)`

- [`PUT /api/external/v2/customers/{customer_id}/categories`](https://pennylane.readme.io/reference/putcustomercategories)

  `PutCustomerCategoriesWithResponse(ctx context.Context, customerId int, params *PutCustomerCategoriesParams, body PutCustomerCategoriesJSONRequestBody) (*PutCustomerCategoriesResponse, error)`

- [`GET /api/external/v2/customers/{customer_id}/contacts`](https://pennylane.readme.io/reference/getcustomercontacts)

  `GetCustomerContactsWithResponse(ctx context.Context, customerId int, params *GetCustomerContactsParams) (*GetCustomerContactsResponse, error)`

- [`POST /api/external/v2/customers/{customer_id}/contacts`](https://pennylane.readme.io/reference/postcustomercontact)

  `PostCustomerContactWithResponse(ctx context.Context, customerId int, params *PostCustomerContactParams, body PostCustomerContactJSONRequestBody) (*PostCustomerContactResponse, error)`

- [`DELETE /api/external/v2/customers/{customer_id}/contacts/{id}`](https://pennylane.readme.io/reference/deletecustomercontact)

  `DeleteCustomerContactWithResponse(ctx context.Context, customerId int, id int64, params *DeleteCustomerContactParams) (*DeleteCustomerContactResponse, error)`

- [`GET /api/external/v2/customers/{customer_id}/contacts/{id}`](https://pennylane.readme.io/reference/getcustomercontact)

  `GetCustomerContactWithResponse(ctx context.Context, customerId int, id int64, params *GetCustomerContactParams) (*GetCustomerContactResponse, error)`

- [`PUT /api/external/v2/customers/{customer_id}/contacts/{id}`](https://pennylane.readme.io/reference/putcustomercontact)

  `PutCustomerContactWithResponse(ctx context.Context, customerId int, id int64, params *PutCustomerContactParams, body PutCustomerContactJSONRequestBody) (*PutCustomerContactResponse, error)`

- [`GET /api/external/v2/customers/{customer_id}/electronic_billing_addresses`](https://pennylane.readme.io/reference/getcustomerelectronicbillingaddresses)

  `GetCustomerElectronicBillingAddressesWithResponse(ctx context.Context, customerId int64, params *GetCustomerElectronicBillingAddressesParams) (*GetCustomerElectronicBillingAddressesResponse, error)`

- [`GET /api/external/v2/customers/{id}`](https://pennylane.readme.io/reference/getcustomer)

  `GetCustomerWithResponse(ctx context.Context, id int, params *GetCustomerParams) (*GetCustomerResponse, error)`

- [`PUT /api/external/v2/customers/{id}`](https://pennylane.readme.io/reference/putcustomer)

  `PutCustomerWithResponse(ctx context.Context, id int, params *PutCustomerParams, body PutCustomerJSONRequestBody) (*PutCustomerResponse, error)`

- [`POST /api/external/v2/individual_customers`](https://pennylane.readme.io/reference/postindividualcustomer)

  `PostIndividualCustomerWithResponse(ctx context.Context, params *PostIndividualCustomerParams, body PostIndividualCustomerJSONRequestBody) (*PostIndividualCustomerResponse, error)`

- [`GET /api/external/v2/individual_customers/{id}`](https://pennylane.readme.io/reference/getindividualcustomer)

  `GetIndividualCustomerWithResponse(ctx context.Context, id string, params *GetIndividualCustomerParams) (*GetIndividualCustomerResponse, error)`

- [`PUT /api/external/v2/individual_customers/{id}`](https://pennylane.readme.io/reference/putindividualcustomer)

  `PutIndividualCustomerWithResponse(ctx context.Context, id string, params *PutIndividualCustomerParams, body PutIndividualCustomerJSONRequestBody) (*PutIndividualCustomerResponse, error)`

## Exports

- [`POST /api/external/v2/exports/analytical_general_ledgers`](https://pennylane.readme.io/reference/exportanalyticalgeneralledger)

  `ExportAnalyticalGeneralLedgerWithResponse(ctx context.Context, params *ExportAnalyticalGeneralLedgerParams, body ExportAnalyticalGeneralLedgerJSONRequestBody) (*ExportAnalyticalGeneralLedgerResponse, error)`

- [`GET /api/external/v2/exports/analytical_general_ledgers/{id}`](https://pennylane.readme.io/reference/getanalyticalgeneralledgerexport)

  `GetAnalyticalGeneralLedgerExportWithResponse(ctx context.Context, id int, params *GetAnalyticalGeneralLedgerExportParams) (*GetAnalyticalGeneralLedgerExportResponse, error)`

- [`POST /api/external/v2/exports/fecs`](https://pennylane.readme.io/reference/exportfec)

  `ExportFecWithResponse(ctx context.Context, params *ExportFecParams, body ExportFecJSONRequestBody) (*ExportFecResponse, error)`

- [`GET /api/external/v2/exports/fecs/{id}`](https://pennylane.readme.io/reference/getfecexport)

  `GetFecExportWithResponse(ctx context.Context, id int, params *GetFecExportParams) (*GetFecExportResponse, error)`

- [`POST /api/external/v2/exports/general_ledgers`](https://pennylane.readme.io/reference/exportgeneralledger)

  `ExportGeneralLedgerWithResponse(ctx context.Context, params *ExportGeneralLedgerParams, body ExportGeneralLedgerJSONRequestBody) (*ExportGeneralLedgerResponse, error)`

- [`GET /api/external/v2/exports/general_ledgers/{id}`](https://pennylane.readme.io/reference/getgeneralledgerexport)

  `GetGeneralLedgerExportWithResponse(ctx context.Context, id int, params *GetGeneralLedgerExportParams) (*GetGeneralLedgerExportResponse, error)`

## File Attachments

- [`POST /api/external/v2/file_attachments`](https://pennylane.readme.io/reference/postfileattachments)

  `PostFileAttachmentsWithResponse(ctx context.Context, params *PostFileAttachmentsParams, body PostFileAttachmentsMultipartRequestBody) (*PostFileAttachmentsResponse, error)`

## Fiscal years

- [`GET /api/external/v2/fiscal_years`](https://pennylane.readme.io/reference/company-fiscal-years)

  `CompanyFiscalYearsWithResponse(ctx context.Context, params *CompanyFiscalYearsParams) (*CompanyFiscalYearsResponse, error)`

## Journals

- [`GET /api/external/v2/journals`](https://pennylane.readme.io/reference/getjournals)

  `GetJournalsWithResponse(ctx context.Context, params *GetJournalsParams) (*GetJournalsResponse, error)`

- [`POST /api/external/v2/journals`](https://pennylane.readme.io/reference/postjournals)

  `PostJournalsWithResponse(ctx context.Context, params *PostJournalsParams, body PostJournalsJSONRequestBody) (*PostJournalsResponse, error)`

- [`GET /api/external/v2/journals/{id}`](https://pennylane.readme.io/reference/getjournal)

  `GetJournalWithResponse(ctx context.Context, id int, params *GetJournalParams) (*GetJournalResponse, error)`

## Ledger Accounts

- [`GET /api/external/v2/ledger_accounts`](https://pennylane.readme.io/reference/getledgeraccounts)

  `GetLedgerAccountsWithResponse(ctx context.Context, params *GetLedgerAccountsParams) (*GetLedgerAccountsResponse, error)`

- [`POST /api/external/v2/ledger_accounts`](https://pennylane.readme.io/reference/postledgeraccounts)

  `PostLedgerAccountsWithResponse(ctx context.Context, params *PostLedgerAccountsParams, body PostLedgerAccountsJSONRequestBody) (*PostLedgerAccountsResponse, error)`

- [`GET /api/external/v2/ledger_accounts/{id}`](https://pennylane.readme.io/reference/getledgeraccount)

  `GetLedgerAccountWithResponse(ctx context.Context, id int, params *GetLedgerAccountParams) (*GetLedgerAccountResponse, error)`

- [`PUT /api/external/v2/ledger_accounts/{id}`](https://pennylane.readme.io/reference/updateledgeraccount)

  `UpdateLedgerAccountWithResponse(ctx context.Context, id int, params *UpdateLedgerAccountParams, body UpdateLedgerAccountJSONRequestBody) (*UpdateLedgerAccountResponse, error)`

## Ledger Attachments

- [`POST /api/external/v2/ledger_attachments`](https://pennylane.readme.io/reference/postledgerattachments)

  `PostLedgerAttachmentsWithResponse(ctx context.Context, params *PostLedgerAttachmentsParams, body PostLedgerAttachmentsMultipartRequestBody) (*PostLedgerAttachmentsResponse, error)`

## Ledger Entries

- [`GET /api/external/v2/ledger_entries`](https://pennylane.readme.io/reference/getledgerentries)

  `GetLedgerEntriesWithResponse(ctx context.Context, params *GetLedgerEntriesParams) (*GetLedgerEntriesResponse, error)`

- [`POST /api/external/v2/ledger_entries`](https://pennylane.readme.io/reference/postledgerentries)

  `PostLedgerEntriesWithResponse(ctx context.Context, params *PostLedgerEntriesParams, body PostLedgerEntriesJSONRequestBody) (*PostLedgerEntriesResponse, error)`

- [`GET /api/external/v2/ledger_entries/{id}`](https://pennylane.readme.io/reference/getledgerentry)

  `GetLedgerEntryWithResponse(ctx context.Context, id int, params *GetLedgerEntryParams) (*GetLedgerEntryResponse, error)`

- [`PUT /api/external/v2/ledger_entries/{id}`](https://pennylane.readme.io/reference/putledgerentries)

  `PutLedgerEntriesWithResponse(ctx context.Context, id string, params *PutLedgerEntriesParams, body PutLedgerEntriesJSONRequestBody) (*PutLedgerEntriesResponse, error)`

- [`GET /api/external/v2/ledger_entries/{ledger_entry_id}/dms_files`](https://pennylane.readme.io/reference/getledgerentriesdmsfiles)

  `GetLedgerEntriesDmsFilesWithResponse(ctx context.Context, ledgerEntryId float32, params *GetLedgerEntriesDmsFilesParams) (*GetLedgerEntriesDmsFilesResponse, error)`

- [`GET /api/external/v2/ledger_entries/{ledger_entry_id}/ledger_entry_lines`](https://pennylane.readme.io/reference/getledgerentriesledgerentrylines)

  `GetLedgerEntriesLedgerEntryLinesWithResponse(ctx context.Context, ledgerEntryId float32, params *GetLedgerEntriesLedgerEntryLinesParams) (*GetLedgerEntriesLedgerEntryLinesResponse, error)`

## Ledger Entry Lines

- [`GET /api/external/v2/ledger_entry_lines`](https://pennylane.readme.io/reference/getledgerentrylines)

  `GetLedgerEntryLinesWithResponse(ctx context.Context, params *GetLedgerEntryLinesParams) (*GetLedgerEntryLinesResponse, error)`

- [`DELETE /api/external/v2/ledger_entry_lines/lettering`](https://pennylane.readme.io/reference/deleteledgerentrylinesunletter)

  `DeleteLedgerEntryLinesUnletterWithResponse(ctx context.Context, params *DeleteLedgerEntryLinesUnletterParams, body DeleteLedgerEntryLinesUnletterJSONRequestBody) (*DeleteLedgerEntryLinesUnletterResponse, error)`

- [`POST /api/external/v2/ledger_entry_lines/lettering`](https://pennylane.readme.io/reference/postledgerentrylinesletter)

  `PostLedgerEntryLinesLetterWithResponse(ctx context.Context, params *PostLedgerEntryLinesLetterParams, body PostLedgerEntryLinesLetterJSONRequestBody) (*PostLedgerEntryLinesLetterResponse, error)`

- [`GET /api/external/v2/ledger_entry_lines/{id}`](https://pennylane.readme.io/reference/getledgerentryline)

  `GetLedgerEntryLineWithResponse(ctx context.Context, id int, params *GetLedgerEntryLineParams) (*GetLedgerEntryLineResponse, error)`

- [`GET /api/external/v2/ledger_entry_lines/{ledger_entry_line_id}/categories`](https://pennylane.readme.io/reference/getledgerentrylinescategories)

  `GetLedgerEntryLinesCategoriesWithResponse(ctx context.Context, ledgerEntryLineId int, params *GetLedgerEntryLinesCategoriesParams) (*GetLedgerEntryLinesCategoriesResponse, error)`

- [`PUT /api/external/v2/ledger_entry_lines/{ledger_entry_line_id}/categories`](https://pennylane.readme.io/reference/putledgerentrylinescategories)

  `PutLedgerEntryLinesCategoriesWithResponse(ctx context.Context, ledgerEntryLineId float32, params *PutLedgerEntryLinesCategoriesParams, body PutLedgerEntryLinesCategoriesJSONRequestBody) (*PutLedgerEntryLinesCategoriesResponse, error)`

- [`GET /api/external/v2/ledger_entry_lines/{ledger_entry_line_id}/lettered_ledger_entry_lines`](https://pennylane.readme.io/reference/getledgerentrylinesletteredledgerentrylines)

  `GetLedgerEntryLinesLetteredLedgerEntryLinesWithResponse(ctx context.Context, ledgerEntryLineId float32, params *GetLedgerEntryLinesLetteredLedgerEntryLinesParams) (*GetLedgerEntryLinesLetteredLedgerEntryLinesResponse, error)`

## Mandates

- [`GET /api/external/v2/gocardless_mandates`](https://pennylane.readme.io/reference/getgocardlessmandates)

  `GetGocardlessMandatesWithResponse(ctx context.Context, params *GetGocardlessMandatesParams) (*GetGocardlessMandatesResponse, error)`

- [`POST /api/external/v2/gocardless_mandates/mail_requests`](https://pennylane.readme.io/reference/postgocardlessmandatemailrequests)

  `PostGocardlessMandateMailRequestsWithResponse(ctx context.Context, params *PostGocardlessMandateMailRequestsParams, body PostGocardlessMandateMailRequestsJSONRequestBody) (*PostGocardlessMandateMailRequestsResponse, error)`

- [`POST /api/external/v2/gocardless_mandates/{gocardless_mandate_id}/associations`](https://pennylane.readme.io/reference/postgocardlessmandateassociations)

  `PostGocardlessMandateAssociationsWithResponse(ctx context.Context, gocardlessMandateId int, params *PostGocardlessMandateAssociationsParams, body PostGocardlessMandateAssociationsJSONRequestBody) (*PostGocardlessMandateAssociationsResponse, error)`

- [`POST /api/external/v2/gocardless_mandates/{gocardless_mandate_id}/cancellations`](https://pennylane.readme.io/reference/postgocardlessmandatecancellations)

  `PostGocardlessMandateCancellationsWithResponse(ctx context.Context, gocardlessMandateId int, params *PostGocardlessMandateCancellationsParams) (*PostGocardlessMandateCancellationsResponse, error)`

- [`GET /api/external/v2/gocardless_mandates/{id}`](https://pennylane.readme.io/reference/getgocardlessmandate)

  `GetGocardlessMandateWithResponse(ctx context.Context, id int, params *GetGocardlessMandateParams) (*GetGocardlessMandateResponse, error)`

- [`GET /api/external/v2/pro_account/mandate_migrations`](https://pennylane.readme.io/reference/getproaccountmandatemigrations)

  `GetProAccountMandateMigrationsWithResponse(ctx context.Context, params *GetProAccountMandateMigrationsParams) (*GetProAccountMandateMigrationsResponse, error)`

- [`POST /api/external/v2/pro_account/mandate_migrations`](https://pennylane.readme.io/reference/postproaccountmandatemigrations)

  `PostProAccountMandateMigrationsWithResponse(ctx context.Context, params *PostProAccountMandateMigrationsParams, body PostProAccountMandateMigrationsJSONRequestBody) (*PostProAccountMandateMigrationsResponse, error)`

- [`POST /api/external/v2/pro_account/mandate_requests`](https://pennylane.readme.io/reference/postproaccountmandatemailrequests)

  `PostProAccountMandateMailRequestsWithResponse(ctx context.Context, params *PostProAccountMandateMailRequestsParams, body PostProAccountMandateMailRequestsJSONRequestBody) (*PostProAccountMandateMailRequestsResponse, error)`

- [`GET /api/external/v2/pro_account/mandates`](https://pennylane.readme.io/reference/getproaccountmandates)

  `GetProAccountMandatesWithResponse(ctx context.Context, params *GetProAccountMandatesParams) (*GetProAccountMandatesResponse, error)`

- [`GET /api/external/v2/sepa_mandates`](https://pennylane.readme.io/reference/getsepamandates)

  `GetSepaMandatesWithResponse(ctx context.Context, params *GetSepaMandatesParams) (*GetSepaMandatesResponse, error)`

- [`POST /api/external/v2/sepa_mandates`](https://pennylane.readme.io/reference/postsepamandates)

  `PostSepaMandatesWithResponse(ctx context.Context, params *PostSepaMandatesParams, body PostSepaMandatesJSONRequestBody) (*PostSepaMandatesResponse, error)`

- [`DELETE /api/external/v2/sepa_mandates/{id}`](https://pennylane.readme.io/reference/deletesepamandate)

  `DeleteSepaMandateWithResponse(ctx context.Context, id int, params *DeleteSepaMandateParams) (*DeleteSepaMandateResponse, error)`

- [`GET /api/external/v2/sepa_mandates/{id}`](https://pennylane.readme.io/reference/getsepamandate)

  `GetSepaMandateWithResponse(ctx context.Context, id int, params *GetSepaMandateParams) (*GetSepaMandateResponse, error)`

- [`PUT /api/external/v2/sepa_mandates/{id}`](https://pennylane.readme.io/reference/putsepamandate)

  `PutSepaMandateWithResponse(ctx context.Context, id int, params *PutSepaMandateParams, body PutSepaMandateJSONRequestBody) (*PutSepaMandateResponse, error)`

## PA Registrations

- [`GET /api/external/v2/pa_registrations`](https://pennylane.readme.io/reference/getparegistrations)

  `GetPaRegistrationsWithResponse(ctx context.Context, params *GetPaRegistrationsParams) (*GetPaRegistrationsResponse, error)`

## Products

- [`GET /api/external/v2/products`](https://pennylane.readme.io/reference/getproducts)

  `GetProductsWithResponse(ctx context.Context, params *GetProductsParams) (*GetProductsResponse, error)`

- [`POST /api/external/v2/products`](https://pennylane.readme.io/reference/postproducts-1)

  `PostProductsWithResponse(ctx context.Context, params *PostProductsParams, body PostProductsJSONRequestBody) (*PostProductsResponse, error)`

- [`GET /api/external/v2/products/{id}`](https://pennylane.readme.io/reference/getproduct)

  `GetProductWithResponse(ctx context.Context, id int, params *GetProductParams) (*GetProductResponse, error)`

- [`PUT /api/external/v2/products/{id}`](https://pennylane.readme.io/reference/putproduct)

  `PutProductWithResponse(ctx context.Context, id int, params *PutProductParams, body PutProductJSONRequestBody) (*PutProductResponse, error)`

## Purchase Requests

- [`GET /api/external/v2/purchase_requests`](https://pennylane.readme.io/reference/getpurchaserequests)

  `GetPurchaseRequestsWithResponse(ctx context.Context, params *GetPurchaseRequestsParams) (*GetPurchaseRequestsResponse, error)`

- [`POST /api/external/v2/purchase_requests/imports`](https://pennylane.readme.io/reference/createpurchaserequestimport)

  `CreatePurchaseRequestImportWithResponse(ctx context.Context, params *CreatePurchaseRequestImportParams, body CreatePurchaseRequestImportJSONRequestBody) (*CreatePurchaseRequestImportResponse, error)`

- [`GET /api/external/v2/purchase_requests/{id}`](https://pennylane.readme.io/reference/getpurchaserequest)

  `GetPurchaseRequestWithResponse(ctx context.Context, id int, params *GetPurchaseRequestParams) (*GetPurchaseRequestResponse, error)`

## Quotes

- [`GET /api/external/v2/changelogs/quotes`](https://pennylane.readme.io/reference/getquotechanges)

  `GetQuoteChangesWithResponse(ctx context.Context, params *GetQuoteChangesParams) (*GetQuoteChangesResponse, error)`

- [`GET /api/external/v2/quotes`](https://pennylane.readme.io/reference/listquotes)

  `ListQuotesWithResponse(ctx context.Context, params *ListQuotesParams) (*ListQuotesResponse, error)`

- [`POST /api/external/v2/quotes`](https://pennylane.readme.io/reference/postquotes)

  `PostQuotesWithResponse(ctx context.Context, params *PostQuotesParams, body PostQuotesJSONRequestBody) (*PostQuotesResponse, error)`

- [`GET /api/external/v2/quotes/{id}`](https://pennylane.readme.io/reference/getquote)

  `GetQuoteWithResponse(ctx context.Context, id int, params *GetQuoteParams) (*GetQuoteResponse, error)`

- [`PUT /api/external/v2/quotes/{id}`](https://pennylane.readme.io/reference/updatequote)

  `UpdateQuoteWithResponse(ctx context.Context, id int, params *UpdateQuoteParams, body UpdateQuoteJSONRequestBody) (*UpdateQuoteResponse, error)`

- [`POST /api/external/v2/quotes/{id}/send_by_email`](https://pennylane.readme.io/reference/sendbyemailquote)

  `SendByEmailQuoteWithResponse(ctx context.Context, id int, params *SendByEmailQuoteParams, body SendByEmailQuoteJSONRequestBody) (*SendByEmailQuoteResponse, error)`

- [`PUT /api/external/v2/quotes/{id}/update_status`](https://pennylane.readme.io/reference/updatestatusquote)

  `UpdateStatusQuoteWithResponse(ctx context.Context, id int, params *UpdateStatusQuoteParams, body UpdateStatusQuoteJSONRequestBody) (*UpdateStatusQuoteResponse, error)`

- [`GET /api/external/v2/quotes/{quote_id}/appendices`](https://pennylane.readme.io/reference/getquoteappendices)

  `GetQuoteAppendicesWithResponse(ctx context.Context, quoteId int, params *GetQuoteAppendicesParams) (*GetQuoteAppendicesResponse, error)`

- [`POST /api/external/v2/quotes/{quote_id}/appendices`](https://pennylane.readme.io/reference/postquoteappendices)

  `PostQuoteAppendicesWithResponse(ctx context.Context, quoteId int, params *PostQuoteAppendicesParams, body PostQuoteAppendicesMultipartRequestBody) (*PostQuoteAppendicesResponse, error)`

- [`GET /api/external/v2/quotes/{quote_id}/invoice_line_sections`](https://pennylane.readme.io/reference/getquoteinvoicelinesections)

  `GetQuoteInvoiceLineSectionsWithResponse(ctx context.Context, quoteId int, params *GetQuoteInvoiceLineSectionsParams) (*GetQuoteInvoiceLineSectionsResponse, error)`

- [`GET /api/external/v2/quotes/{quote_id}/invoice_lines`](https://pennylane.readme.io/reference/getquoteinvoicelines)

  `GetQuoteInvoiceLinesWithResponse(ctx context.Context, quoteId int, params *GetQuoteInvoiceLinesParams) (*GetQuoteInvoiceLinesResponse, error)`

## Supplier Invoices

- [`GET /api/external/v2/supplier_invoices`](https://pennylane.readme.io/reference/getsupplierinvoices)

  `GetSupplierInvoicesWithResponse(ctx context.Context, params *GetSupplierInvoicesParams) (*GetSupplierInvoicesResponse, error)`

- [`POST /api/external/v2/supplier_invoices/e_invoices/imports`](https://pennylane.readme.io/reference/createsupplierinvoiceeinvoiceimport)

  `CreateSupplierInvoiceEInvoiceImportWithResponse(ctx context.Context, params *CreateSupplierInvoiceEInvoiceImportParams, body CreateSupplierInvoiceEInvoiceImportMultipartRequestBody) (*CreateSupplierInvoiceEInvoiceImportResponse, error)`

- [`POST /api/external/v2/supplier_invoices/import`](https://pennylane.readme.io/reference/importsupplierinvoice)

  `ImportSupplierInvoiceWithResponse(ctx context.Context, params *ImportSupplierInvoiceParams, body ImportSupplierInvoiceJSONRequestBody) (*ImportSupplierInvoiceResponse, error)`

- [`GET /api/external/v2/supplier_invoices/{id}`](https://pennylane.readme.io/reference/getsupplierinvoice)

  `GetSupplierInvoiceWithResponse(ctx context.Context, id int, params *GetSupplierInvoiceParams) (*GetSupplierInvoiceResponse, error)`

- [`PUT /api/external/v2/supplier_invoices/{id}`](https://pennylane.readme.io/reference/putsupplierinvoice)

  `PutSupplierInvoiceWithResponse(ctx context.Context, id int, params *PutSupplierInvoiceParams, body PutSupplierInvoiceJSONRequestBody) (*PutSupplierInvoiceResponse, error)`

- [`PUT /api/external/v2/supplier_invoices/{id}/validate_accounting`](https://pennylane.readme.io/reference/validateaccountingsupplierinvoice)

  `ValidateAccountingSupplierInvoiceWithResponse(ctx context.Context, id int, params *ValidateAccountingSupplierInvoiceParams) (*ValidateAccountingSupplierInvoiceResponse, error)`

- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/categories`](https://pennylane.readme.io/reference/getsupplierinvoicecategories)

  `GetSupplierInvoiceCategoriesWithResponse(ctx context.Context, supplierInvoiceId int, params *GetSupplierInvoiceCategoriesParams) (*GetSupplierInvoiceCategoriesResponse, error)`

- [`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/categories`](https://pennylane.readme.io/reference/putsupplierinvoicecategories)

  `PutSupplierInvoiceCategoriesWithResponse(ctx context.Context, supplierInvoiceId int, params *PutSupplierInvoiceCategoriesParams, body PutSupplierInvoiceCategoriesJSONRequestBody) (*PutSupplierInvoiceCategoriesResponse, error)`

- [`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/e_invoice_status`](https://pennylane.readme.io/reference/putsupplierinvoiceeinvoicestatus)

  `PutSupplierInvoiceEInvoiceStatusWithResponse(ctx context.Context, supplierInvoiceId int, params *PutSupplierInvoiceEInvoiceStatusParams, body PutSupplierInvoiceEInvoiceStatusJSONRequestBody) (*PutSupplierInvoiceEInvoiceStatusResponse, error)`

- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/invoice_lines`](https://pennylane.readme.io/reference/getsupplierinvoicelines)

  `GetSupplierInvoiceLinesWithResponse(ctx context.Context, supplierInvoiceId int, params *GetSupplierInvoiceLinesParams) (*GetSupplierInvoiceLinesResponse, error)`

- [`POST /api/external/v2/supplier_invoices/{supplier_invoice_id}/linked_purchase_requests`](https://pennylane.readme.io/reference/postsupplierinvoicelinkedpurchaserequests)

  `PostSupplierInvoiceLinkedPurchaseRequestsWithResponse(ctx context.Context, supplierInvoiceId int, params *PostSupplierInvoiceLinkedPurchaseRequestsParams, body PostSupplierInvoiceLinkedPurchaseRequestsJSONRequestBody) (*PostSupplierInvoiceLinkedPurchaseRequestsResponse, error)`

- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/matched_transactions`](https://pennylane.readme.io/reference/getsupplierinvoicematchedtransactions)

  `GetSupplierInvoiceMatchedTransactionsWithResponse(ctx context.Context, supplierInvoiceId int, params *GetSupplierInvoiceMatchedTransactionsParams) (*GetSupplierInvoiceMatchedTransactionsResponse, error)`

- [`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/payment_status`](https://pennylane.readme.io/reference/updatesupplierinvoicepaymentstatus)

  `UpdateSupplierInvoicePaymentStatusWithResponse(ctx context.Context, supplierInvoiceId int, params *UpdateSupplierInvoicePaymentStatusParams, body UpdateSupplierInvoicePaymentStatusJSONRequestBody) (*UpdateSupplierInvoicePaymentStatusResponse, error)`

- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/payments`](https://pennylane.readme.io/reference/getsupplierinvoicepayments)

  `GetSupplierInvoicePaymentsWithResponse(ctx context.Context, supplierInvoiceId int, params *GetSupplierInvoicePaymentsParams) (*GetSupplierInvoicePaymentsResponse, error)`

## Suppliers

- [`GET /api/external/v2/suppliers`](https://pennylane.readme.io/reference/getsuppliers)

  `GetSuppliersWithResponse(ctx context.Context, params *GetSuppliersParams) (*GetSuppliersResponse, error)`

- [`POST /api/external/v2/suppliers`](https://pennylane.readme.io/reference/postsupplier)

  `PostSupplierWithResponse(ctx context.Context, params *PostSupplierParams, body PostSupplierJSONRequestBody) (*PostSupplierResponse, error)`

- [`GET /api/external/v2/suppliers/{id}`](https://pennylane.readme.io/reference/getsupplier)

  `GetSupplierWithResponse(ctx context.Context, id int, params *GetSupplierParams) (*GetSupplierResponse, error)`

- [`PUT /api/external/v2/suppliers/{id}`](https://pennylane.readme.io/reference/putsupplier)

  `PutSupplierWithResponse(ctx context.Context, id int, params *PutSupplierParams, body PutSupplierJSONRequestBody) (*PutSupplierResponse, error)`

- [`GET /api/external/v2/suppliers/{supplier_id}/categories`](https://pennylane.readme.io/reference/getsuppliercategories)

  `GetSupplierCategoriesWithResponse(ctx context.Context, supplierId int, params *GetSupplierCategoriesParams) (*GetSupplierCategoriesResponse, error)`

- [`PUT /api/external/v2/suppliers/{supplier_id}/categories`](https://pennylane.readme.io/reference/putsuppliercategories)

  `PutSupplierCategoriesWithResponse(ctx context.Context, supplierId int, params *PutSupplierCategoriesParams, body PutSupplierCategoriesJSONRequestBody) (*PutSupplierCategoriesResponse, error)`

## Transactions

- [`POST /api/external/v2/customer_invoices/{customer_invoice_id}/matched_transactions`](https://pennylane.readme.io/reference/postcustomerinvoicematchedtransactions)

  `PostCustomerInvoiceMatchedTransactionsWithResponse(ctx context.Context, customerInvoiceId int, params *PostCustomerInvoiceMatchedTransactionsParams, body PostCustomerInvoiceMatchedTransactionsJSONRequestBody) (*PostCustomerInvoiceMatchedTransactionsResponse, error)`

- [`DELETE /api/external/v2/customer_invoices/{customer_invoice_id}/matched_transactions/{id}`](https://pennylane.readme.io/reference/deletecustomerinvoicematchedtransactions)

  `DeleteCustomerInvoiceMatchedTransactionsWithResponse(ctx context.Context, customerInvoiceId int, id int, params *DeleteCustomerInvoiceMatchedTransactionsParams) (*DeleteCustomerInvoiceMatchedTransactionsResponse, error)`

- [`POST /api/external/v2/supplier_invoices/{supplier_invoice_id}/matched_transactions`](https://pennylane.readme.io/reference/postsupplierinvoicematchedtransactions)

  `PostSupplierInvoiceMatchedTransactionsWithResponse(ctx context.Context, supplierInvoiceId int, params *PostSupplierInvoiceMatchedTransactionsParams, body PostSupplierInvoiceMatchedTransactionsJSONRequestBody) (*PostSupplierInvoiceMatchedTransactionsResponse, error)`

- [`DELETE /api/external/v2/supplier_invoices/{supplier_invoice_id}/matched_transactions/{id}`](https://pennylane.readme.io/reference/deletesupplierinvoicematchedtransactions)

  `DeleteSupplierInvoiceMatchedTransactionsWithResponse(ctx context.Context, supplierInvoiceId int, id int, params *DeleteSupplierInvoiceMatchedTransactionsParams) (*DeleteSupplierInvoiceMatchedTransactionsResponse, error)`

- [`GET /api/external/v2/transactions`](https://pennylane.readme.io/reference/gettransactions)

  `GetTransactionsWithResponse(ctx context.Context, params *GetTransactionsParams) (*GetTransactionsResponse, error)`

- [`POST /api/external/v2/transactions`](https://pennylane.readme.io/reference/createtransaction)

  `CreateTransactionWithResponse(ctx context.Context, params *CreateTransactionParams, body CreateTransactionJSONRequestBody) (*CreateTransactionResponse, error)`

- [`GET /api/external/v2/transactions/{id}`](https://pennylane.readme.io/reference/gettransaction)

  `GetTransactionWithResponse(ctx context.Context, id int, params *GetTransactionParams) (*GetTransactionResponse, error)`

- [`PUT /api/external/v2/transactions/{id}`](https://pennylane.readme.io/reference/updatetransaction)

  `UpdateTransactionWithResponse(ctx context.Context, id int, params *UpdateTransactionParams, body UpdateTransactionJSONRequestBody) (*UpdateTransactionResponse, error)`

- [`GET /api/external/v2/transactions/{transaction_id}/categories`](https://pennylane.readme.io/reference/gettransactioncategories)

  `GetTransactionCategoriesWithResponse(ctx context.Context, transactionId int, params *GetTransactionCategoriesParams) (*GetTransactionCategoriesResponse, error)`

- [`PUT /api/external/v2/transactions/{transaction_id}/categories`](https://pennylane.readme.io/reference/puttransactioncategories)

  `PutTransactionCategoriesWithResponse(ctx context.Context, transactionId int, params *PutTransactionCategoriesParams, body PutTransactionCategoriesJSONRequestBody) (*PutTransactionCategoriesResponse, error)`

- [`GET /api/external/v2/transactions/{transaction_id}/matched_invoices`](https://pennylane.readme.io/reference/gettransactionmatchedinvoices)

  `GetTransactionMatchedInvoicesWithResponse(ctx context.Context, transactionId int, params *GetTransactionMatchedInvoicesParams) (*GetTransactionMatchedInvoicesResponse, error)`

## Trial balance

- [`GET /api/external/v2/trial_balance`](https://pennylane.readme.io/reference/gettrialbalance)

  `GetTrialBalanceWithResponse(ctx context.Context, params *GetTrialBalanceParams) (*GetTrialBalanceResponse, error)`

## Users

- [`GET /api/external/v2/me`](https://pennylane.readme.io/reference/getme)

  `GetMeWithResponse(ctx context.Context, params *GetMeParams) (*GetMeResponse, error)`

## Webhooks

- [`GET /api/external/v2/webhook_subscriptions`](https://pennylane.readme.io/reference/getwebhooksubscriptions)

  `GetWebhookSubscriptionsWithResponse(ctx context.Context, params *GetWebhookSubscriptionsParams) (*GetWebhookSubscriptionsResponse, error)`

- [`POST /api/external/v2/webhook_subscriptions`](https://pennylane.readme.io/reference/postwebhooksubscriptions)

  `PostWebhookSubscriptionsWithResponse(ctx context.Context, body PostWebhookSubscriptionsJSONRequestBody) (*PostWebhookSubscriptionsResponse, error)`

- [`DELETE /api/external/v2/webhook_subscriptions/{id}`](https://pennylane.readme.io/reference/deletewebhooksubscription)

  `DeleteWebhookSubscriptionWithResponse(ctx context.Context, id int) (*DeleteWebhookSubscriptionResponse, error)`

- [`GET /api/external/v2/webhook_subscriptions/{id}`](https://pennylane.readme.io/reference/getwebhooksubscription)

  `GetWebhookSubscriptionWithResponse(ctx context.Context, id int) (*GetWebhookSubscriptionResponse, error)`

- [`PUT /api/external/v2/webhook_subscriptions/{id}`](https://pennylane.readme.io/reference/putwebhooksubscription)

  `PutWebhookSubscriptionWithResponse(ctx context.Context, id int, body PutWebhookSubscriptionJSONRequestBody) (*PutWebhookSubscriptionResponse, error)`
