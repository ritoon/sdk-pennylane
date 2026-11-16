# Référence du client Go

Chaque fiche décrit l’usage, la signature Go, les paramètres HTTP, le corps de requête et les réponses par statut. Les descriptions métier et les permissions reprennent la langue du contrat Pennylane.

Les paramètres de chemin sont des arguments Go. Les paramètres de requête et les en-têtes se renseignent dans `params`. Pour les champs de type `nullable.Nullable[T]`, une valeur absente, une valeur nulle et une valeur renseignée sont distinctes. Dans les tableaux, `$` désigne le payload entier et `[]` un élément de liste. La présence d’un champ s’applique lorsque son objet parent est présent. Pour les variantes, `oneOf` impose un seul schéma, `anyOf` au moins un, et `allOf` tous.

[Exemple Go de lecture paginée](utilisation.md) · [Payloads d’erreur](endpoints/erreurs.md)

## Bank Establishments

- [`GET /api/external/v2/bank_establishments`](endpoints/getBankEstablishments.md) — List bank establishments · `GetBankEstablishmentsWithResponse`

## Bank accounts

- [`GET /api/external/v2/bank_accounts/{id}`](endpoints/getBankAccount.md) — Retrieve a bank account · `GetBankAccountWithResponse`
- [`GET /api/external/v2/bank_accounts`](endpoints/getBankAccounts.md) — List bank accounts · `GetBankAccountsWithResponse`
- [`POST /api/external/v2/bank_accounts`](endpoints/postBankAccount.md) — Create a bank account · `PostBankAccountWithResponse`

## Billing Subscriptions

- [`GET /api/external/v2/billing_subscriptions/{billing_subscription_id}/invoice_line_sections`](endpoints/getBillingSubscriptionInvoiceLineSections.md) — List the invoice line sections of a billing subscription · `GetBillingSubscriptionInvoiceLineSectionsWithResponse`
- [`GET /api/external/v2/billing_subscriptions/{billing_subscription_id}/invoice_lines`](endpoints/getBillingSubscriptionInvoiceLines.md) — List invoice lines for a billing subscription · `GetBillingSubscriptionInvoiceLinesWithResponse`
- [`GET /api/external/v2/billing_subscriptions/{id}`](endpoints/getBillingSubscription.md) — Get a billing subscription · `GetBillingSubscriptionWithResponse`
- [`GET /api/external/v2/billing_subscriptions`](endpoints/getBillingSubscriptions.md) — List billing subscriptions · `GetBillingSubscriptionsWithResponse`
- [`POST /api/external/v2/billing_subscriptions`](endpoints/postBillingSubscriptions.md) — Create a billing subscription · `PostBillingSubscriptionsWithResponse`
- [`PUT /api/external/v2/billing_subscriptions/{id}`](endpoints/putBillingSubscriptions.md) — Update a billing subscription · `PutBillingSubscriptionsWithResponse`

## Billing email templates

- [`GET /api/external/v2/billing_email_templates`](endpoints/getBillingEmailTemplates.md) — List billing email templates · `GetBillingEmailTemplatesWithResponse`

## Categories

- [`GET /api/external/v2/categories/{id}`](endpoints/getCategory.md) — Retrieve a category · `GetCategoryWithResponse`
- [`GET /api/external/v2/categories`](endpoints/getCategories.md) — List categories · `GetCategoriesWithResponse`
- [`GET /api/external/v2/category_groups`](endpoints/getCategoryGroups.md) — List category groups · `GetCategoryGroupsWithResponse`
- [`POST /api/external/v2/categories`](endpoints/postCategories.md) — Create a category · `PostCategoriesWithResponse`
- [`PUT /api/external/v2/categories/{id}`](endpoints/updateCategory.md) — Update a category · `UpdateCategoryWithResponse`

## Category Groups

- [`GET /api/external/v2/category_groups/{category_group_id}/categories`](endpoints/getCategoryGroupCategories.md) — List categories of a category group · `GetCategoryGroupCategoriesWithResponse`
- [`GET /api/external/v2/category_groups/{id}`](endpoints/getCategoryGroup.md) — Retrieve a category group · `GetCategoryGroupWithResponse`
- [`POST /api/external/v2/category_groups`](endpoints/postCategoryGroups.md) — Create a category group · `PostCategoryGroupsWithResponse`
- [`PUT /api/external/v2/category_groups/{id}`](endpoints/putCategoryGroup.md) — Update a category group · `PutCategoryGroupWithResponse`

## Changelogs

- [`GET /api/external/v2/changelogs/customer_invoices`](endpoints/getCustomerInvoicesChanges.md) — Get customer invoices changes events · `GetCustomerInvoicesChangesWithResponse`
- [`GET /api/external/v2/changelogs/customers`](endpoints/getCustomerChanges.md) — Get customer changes events · `GetCustomerChangesWithResponse`
- [`GET /api/external/v2/changelogs/ledger_entries_categories`](endpoints/getLedgerEntriesCategoryChanges.md) — Get ledger entry category change events · `GetLedgerEntriesCategoryChangesWithResponse`
- [`GET /api/external/v2/changelogs/ledger_entries`](endpoints/getLedgerEntriesChanges.md) — Get ledger entry change events · `GetLedgerEntriesChangesWithResponse`
- [`GET /api/external/v2/changelogs/ledger_entry_lines_categories`](endpoints/getLedgerEntryLinesCategoryChanges.md) — Get ledger entry line category change events · `GetLedgerEntryLinesCategoryChangesWithResponse`
- [`GET /api/external/v2/changelogs/ledger_entry_lines`](endpoints/getLedgerEntryLineChanges.md) — Get ledger entry line change events · `GetLedgerEntryLineChangesWithResponse`
- [`GET /api/external/v2/changelogs/products`](endpoints/getProductChanges.md) — Get product change events · `GetProductChangesWithResponse`
- [`GET /api/external/v2/changelogs/supplier_invoices`](endpoints/getSupplierInvoicesChanges.md) — Get supplier invoices changes events · `GetSupplierInvoicesChangesWithResponse`
- [`GET /api/external/v2/changelogs/suppliers`](endpoints/getSupplierChanges.md) — Get supplier changes events · `GetSupplierChangesWithResponse`
- [`GET /api/external/v2/changelogs/transactions`](endpoints/getTransactionChanges.md) — Get transaction change events · `GetTransactionChangesWithResponse`

## Commercial Documents

- [`GET /api/external/v2/commercial_documents/{commercial_document_id}/appendices`](endpoints/getCommercialDocumentAppendices.md) — List appendices of a commercial document · `GetCommercialDocumentAppendicesWithResponse`
- [`GET /api/external/v2/commercial_documents/{commercial_document_id}/invoice_line_sections`](endpoints/getCommercialDocumentInvoiceLineSections.md) — List invoice line sections for a commercial document · `GetCommercialDocumentInvoiceLineSectionsWithResponse`
- [`GET /api/external/v2/commercial_documents/{commercial_document_id}/invoice_lines`](endpoints/getCommercialDocumentInvoiceLines.md) — List invoice lines for a commercial document · `GetCommercialDocumentInvoiceLinesWithResponse`
- [`GET /api/external/v2/commercial_documents/{id}`](endpoints/getCommercialDocument.md) — Retrieve a commercial document · `GetCommercialDocumentWithResponse`
- [`GET /api/external/v2/commercial_documents`](endpoints/listCommercialDocuments.md) — List commercial documents · `ListCommercialDocumentsWithResponse`
- [`POST /api/external/v2/commercial_documents/{commercial_document_id}/appendices`](endpoints/postCommercialDocumentAppendices.md) — Upload an appendix for a commercial document · `PostCommercialDocumentAppendicesWithResponse`
- [`POST /api/external/v2/commercial_documents/{id}/send_by_email`](endpoints/sendByEmailCommercialDocument.md) — Send a commercial document by email · `SendByEmailCommercialDocumentWithResponse`
- [`POST /api/external/v2/commercial_documents`](endpoints/postCommercialDocuments.md) — Create a commercial document · `PostCommercialDocumentsWithResponse`
- [`PUT /api/external/v2/commercial_documents/{id}`](endpoints/updateCommercialDocument.md) — Update a commercial document · `UpdateCommercialDocumentWithResponse`

## Company

- [`GET /api/external/v2/company/features`](endpoints/getCompanyFeatures.md) — Company features · `GetCompanyFeaturesWithResponse`

## Customer Invoice Templates

- [`GET /api/external/v2/customer_invoice_templates`](endpoints/getCustomerInvoiceTemplates.md) — List customer invoice templates · `GetCustomerInvoiceTemplatesWithResponse`

## Customer Invoices

- [`DELETE /api/external/v2/customer_invoices/{id}`](endpoints/deleteCustomerInvoices.md) — Delete draft invoice · `DeleteCustomerInvoicesWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/appendices`](endpoints/getCustomerInvoiceAppendices.md) — List appendices of a customer invoice · `GetCustomerInvoiceAppendicesWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/categories`](endpoints/getCustomerInvoiceCategories.md) — List categories of a customer invoice · `GetCustomerInvoiceCategoriesWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/comments`](endpoints/getCustomerInvoiceComments.md) — List comments of a customer invoice · `GetCustomerInvoiceCommentsWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/custom_header_fields`](endpoints/getCustomerInvoiceCustomHeaderFields.md) — List custom header fields for a customer invoice · `GetCustomerInvoiceCustomHeaderFieldsWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/installments`](endpoints/getCustomerInvoiceInstallments.md) — List installments for a customer invoice · `GetCustomerInvoiceInstallmentsWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/invoice_line_sections`](endpoints/getCustomerInvoiceInvoiceLineSections.md) — List invoice line sections for a customer invoice · `GetCustomerInvoiceInvoiceLineSectionsWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/invoice_lines`](endpoints/getCustomerInvoiceInvoiceLines.md) — List invoice lines for a customer invoice · `GetCustomerInvoiceInvoiceLinesWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/matched_transactions`](endpoints/getCustomerInvoiceMatchedTransactions.md) — List matched transactions for a customer invoice · `GetCustomerInvoiceMatchedTransactionsWithResponse`
- [`GET /api/external/v2/customer_invoices/{customer_invoice_id}/payments`](endpoints/getCustomerInvoicePayments.md) — List payments for a customer invoice · `GetCustomerInvoicePaymentsWithResponse`
- [`GET /api/external/v2/customer_invoices/{id}`](endpoints/getCustomerInvoice.md) — Retrieve a customer invoice · `GetCustomerInvoiceWithResponse`
- [`GET /api/external/v2/customer_invoices`](endpoints/getCustomerInvoices.md) — List customer invoices · `GetCustomerInvoicesWithResponse`
- [`POST /api/external/v2/customer_invoices/create_from_quote`](endpoints/createCustomerInvoiceFromQuote.md) — Create a customer invoice from a quote · `CreateCustomerInvoiceFromQuoteWithResponse`
- [`POST /api/external/v2/customer_invoices/e_invoices/imports`](endpoints/createCustomerInvoiceEInvoiceImport.md) — Import a customer e-invoice · `CreateCustomerInvoiceEInvoiceImportWithResponse`
- [`POST /api/external/v2/customer_invoices/import`](endpoints/importCustomerInvoices.md) — Import an invoice with file attached · `ImportCustomerInvoicesWithResponse`
- [`POST /api/external/v2/customer_invoices/{customer_invoice_id}/appendices`](endpoints/postCustomerInvoiceAppendices.md) — Upload an appendix for a customer invoice · `PostCustomerInvoiceAppendicesWithResponse`
- [`POST /api/external/v2/customer_invoices/{customer_invoice_id}/comments`](endpoints/postCustomerInvoiceComment.md) — Add a comment to a customer invoice · `PostCustomerInvoiceCommentWithResponse`
- [`POST /api/external/v2/customer_invoices/{id}/link_credit_note`](endpoints/linkCreditNote.md) — Link a credit note to a customer invoice · `LinkCreditNoteWithResponse`
- [`POST /api/external/v2/customer_invoices/{id}/send_by_email`](endpoints/sendByEmailCustomerInvoice.md) — Send a customer invoice by email · `SendByEmailCustomerInvoiceWithResponse`
- [`POST /api/external/v2/customer_invoices/{id}/send_to_pa`](endpoints/sendToPaCustomerInvoice.md) — Send a customer e-invoice to PA · `SendToPaCustomerInvoiceWithResponse`
- [`POST /api/external/v2/customer_invoices`](endpoints/postCustomerInvoices.md) — Create a customer invoice · `PostCustomerInvoicesWithResponse`
- [`PUT /api/external/v2/customer_invoices/{customer_invoice_id}/categories`](endpoints/putCustomerInvoiceCategories.md) — Categorize a customer invoice · `PutCustomerInvoiceCategoriesWithResponse`
- [`PUT /api/external/v2/customer_invoices/{customer_invoice_id}/installments/{id}/mark_as_paid`](endpoints/markAsPaidCustomerInvoiceInstallment.md) — Mark a customer invoice installment as paid · `MarkAsPaidCustomerInvoiceInstallmentWithResponse`
- [`PUT /api/external/v2/customer_invoices/{id}/finalize`](endpoints/finalizeCustomerInvoice.md) — Turn the draft invoice into a finalized invoice. · `FinalizeCustomerInvoiceWithResponse`
- [`PUT /api/external/v2/customer_invoices/{id}/mark_as_paid`](endpoints/markAsPaidCustomerInvoice.md) — Mark a customer invoice as paid · `MarkAsPaidCustomerInvoiceWithResponse`
- [`PUT /api/external/v2/customer_invoices/{id}/update_imported`](endpoints/updateImportedCustomerInvoice.md) — Update an Imported customer invoice · `UpdateImportedCustomerInvoiceWithResponse`
- [`PUT /api/external/v2/customer_invoices/{id}`](endpoints/updateCustomerInvoice.md) — Update a customer invoice · `UpdateCustomerInvoiceWithResponse`

## Customers

- [`DELETE /api/external/v2/customers/{customer_id}/contacts/{id}`](endpoints/deleteCustomerContact.md) — Delete a contact of a customer · `DeleteCustomerContactWithResponse`
- [`GET /api/external/v2/company_customers/{id}`](endpoints/getCompanyCustomer.md) — Retrieve a company customer · `GetCompanyCustomerWithResponse`
- [`GET /api/external/v2/customers/{customer_id}/categories`](endpoints/getCustomerCategories.md) — List categories of a customer · `GetCustomerCategoriesWithResponse`
- [`GET /api/external/v2/customers/{customer_id}/contacts/{id}`](endpoints/getCustomerContact.md) — Retrieve a contact of a customer · `GetCustomerContactWithResponse`
- [`GET /api/external/v2/customers/{customer_id}/contacts`](endpoints/getCustomerContacts.md) — List contacts of a customer · `GetCustomerContactsWithResponse`
- [`GET /api/external/v2/customers/{customer_id}/electronic_billing_addresses`](endpoints/getCustomerElectronicBillingAddresses.md) — List eligible electronic billing addresses for a customer · `GetCustomerElectronicBillingAddressesWithResponse`
- [`GET /api/external/v2/customers/{id}`](endpoints/getCustomer.md) — Retrieve a customer · `GetCustomerWithResponse`
- [`GET /api/external/v2/customers`](endpoints/getCustomers.md) — List customers (company and individual) · `GetCustomersWithResponse`
- [`GET /api/external/v2/individual_customers/{id}`](endpoints/getIndividualCustomer.md) — Retrieve an individual customer · `GetIndividualCustomerWithResponse`
- [`POST /api/external/v2/company_customers`](endpoints/postCompanyCustomer.md) — Create a company customer · `PostCompanyCustomerWithResponse`
- [`POST /api/external/v2/customers/{customer_id}/contacts`](endpoints/postCustomerContact.md) — Create a contact of a customer · `PostCustomerContactWithResponse`
- [`POST /api/external/v2/customers`](endpoints/postCustomer.md) — Create a customer (company or individual) · `PostCustomerWithResponse`
- [`POST /api/external/v2/individual_customers`](endpoints/postIndividualCustomer.md) — Create an individual customer · `PostIndividualCustomerWithResponse`
- [`PUT /api/external/v2/company_customers/{id}`](endpoints/putCompanyCustomer.md) — Update a company customer · `PutCompanyCustomerWithResponse`
- [`PUT /api/external/v2/customers/{customer_id}/categories`](endpoints/putCustomerCategories.md) — Categorize a customer · `PutCustomerCategoriesWithResponse`
- [`PUT /api/external/v2/customers/{customer_id}/contacts/{id}`](endpoints/putCustomerContact.md) — Update a contact of a customer · `PutCustomerContactWithResponse`
- [`PUT /api/external/v2/customers/{id}`](endpoints/putCustomer.md) — Update a customer (company or individual) · `PutCustomerWithResponse`
- [`PUT /api/external/v2/individual_customers/{id}`](endpoints/putIndividualCustomer.md) — Update an individual customer · `PutIndividualCustomerWithResponse`

## Exports

- [`GET /api/external/v2/exports/analytical_general_ledgers/{id}`](endpoints/getAnalyticalGeneralLedgerExport.md) — Retrieve an Analytical General Ledger export · `GetAnalyticalGeneralLedgerExportWithResponse`
- [`GET /api/external/v2/exports/fecs/{id}`](endpoints/getFecExport.md) — Retrieve a FEC export · `GetFecExportWithResponse`
- [`GET /api/external/v2/exports/general_ledgers/{id}`](endpoints/getGeneralLedgerExport.md) — Retrieve a General Ledger export · `GetGeneralLedgerExportWithResponse`
- [`POST /api/external/v2/exports/analytical_general_ledgers`](endpoints/exportAnalyticalGeneralLedger.md) — Create an Analytical General Ledger export · `ExportAnalyticalGeneralLedgerWithResponse`
- [`POST /api/external/v2/exports/fecs`](endpoints/exportFec.md) — Create a FEC export · `ExportFecWithResponse`
- [`POST /api/external/v2/exports/general_ledgers`](endpoints/exportGeneralLedger.md) — Create a General Ledger export · `ExportGeneralLedgerWithResponse`

## File Attachments

- [`POST /api/external/v2/file_attachments`](endpoints/postFileAttachments.md) — Upload a file · `PostFileAttachmentsWithResponse`

## Fiscal years

- [`GET /api/external/v2/fiscal_years`](endpoints/company-fiscal-years.md) — List Company's Fiscal Years · `CompanyFiscalYearsWithResponse`

## Journals

- [`GET /api/external/v2/journals/{id}`](endpoints/getJournal.md) — Retrieve a journal · `GetJournalWithResponse`
- [`GET /api/external/v2/journals`](endpoints/getJournals.md) — List journals · `GetJournalsWithResponse`
- [`POST /api/external/v2/journals`](endpoints/postJournals.md) — Create a journal · `PostJournalsWithResponse`

## Ledger Accounts

- [`GET /api/external/v2/ledger_accounts/{id}`](endpoints/getLedgerAccount.md) — Get a ledger account · `GetLedgerAccountWithResponse`
- [`GET /api/external/v2/ledger_accounts`](endpoints/getLedgerAccounts.md) — List Ledger Accounts · `GetLedgerAccountsWithResponse`
- [`POST /api/external/v2/ledger_accounts`](endpoints/postLedgerAccounts.md) — Create a ledger account · `PostLedgerAccountsWithResponse`
- [`PUT /api/external/v2/ledger_accounts/{id}`](endpoints/updateLedgerAccount.md) — Update a ledger account · `UpdateLedgerAccountWithResponse`

## Ledger Attachments

- [`POST /api/external/v2/ledger_attachments`](endpoints/postLedgerAttachments.md) — Upload a file · `PostLedgerAttachmentsWithResponse`

## Ledger Entries

- [`GET /api/external/v2/ledger_entries/{id}`](endpoints/getLedgerEntry.md) — Retrieve a Ledger entry · `GetLedgerEntryWithResponse`
- [`GET /api/external/v2/ledger_entries/{ledger_entry_id}/dms_files`](endpoints/getLedgerEntriesDmsFiles.md) — List DMS files of a Ledger Entry · `GetLedgerEntriesDmsFilesWithResponse`
- [`GET /api/external/v2/ledger_entries/{ledger_entry_id}/ledger_entry_lines`](endpoints/getLedgerEntriesLedgerEntryLines.md) — List ledger entry lines of a Ledger Entry · `GetLedgerEntriesLedgerEntryLinesWithResponse`
- [`GET /api/external/v2/ledger_entries`](endpoints/getLedgerEntries.md) — List Ledger Entries · `GetLedgerEntriesWithResponse`
- [`POST /api/external/v2/ledger_entries`](endpoints/postLedgerEntries.md) — Create a ledger entry · `PostLedgerEntriesWithResponse`
- [`PUT /api/external/v2/ledger_entries/{id}`](endpoints/putLedgerEntries.md) — Update a ledger entry · `PutLedgerEntriesWithResponse`

## Ledger Entry Lines

- [`DELETE /api/external/v2/ledger_entry_lines/lettering`](endpoints/deleteLedgerEntryLinesUnletter.md) — Unletter ledger entry lines · `DeleteLedgerEntryLinesUnletterWithResponse`
- [`GET /api/external/v2/ledger_entry_lines/{id}`](endpoints/getLedgerEntryLine.md) — Retrieve a Ledger entry line · `GetLedgerEntryLineWithResponse`
- [`GET /api/external/v2/ledger_entry_lines/{ledger_entry_line_id}/categories`](endpoints/getLedgerEntryLinesCategories.md) — List categories of a Ledger Entry line · `GetLedgerEntryLinesCategoriesWithResponse`
- [`GET /api/external/v2/ledger_entry_lines/{ledger_entry_line_id}/lettered_ledger_entry_lines`](endpoints/getLedgerEntryLinesLetteredLedgerEntryLines.md) — List ledger entry lines lettered to a given ledger entry line · `GetLedgerEntryLinesLetteredLedgerEntryLinesWithResponse`
- [`GET /api/external/v2/ledger_entry_lines`](endpoints/getLedgerEntryLines.md) — List ledger entry lines · `GetLedgerEntryLinesWithResponse`
- [`POST /api/external/v2/ledger_entry_lines/lettering`](endpoints/postLedgerEntryLinesLetter.md) — Letter ledger entry lines · `PostLedgerEntryLinesLetterWithResponse`
- [`PUT /api/external/v2/ledger_entry_lines/{ledger_entry_line_id}/categories`](endpoints/putLedgerEntryLinesCategories.md) — Link Analytical Categories to a Ledger Entry line · `PutLedgerEntryLinesCategoriesWithResponse`

## Mandates

- [`DELETE /api/external/v2/sepa_mandates/{id}`](endpoints/deleteSepaMandate.md) — Delete a SEPA mandate · `DeleteSepaMandateWithResponse`
- [`GET /api/external/v2/gocardless_mandates/{id}`](endpoints/getGocardlessMandate.md) — Get a Gocardless mandate · `GetGocardlessMandateWithResponse`
- [`GET /api/external/v2/gocardless_mandates`](endpoints/getGocardlessMandates.md) — List gocardless mandates · `GetGocardlessMandatesWithResponse`
- [`GET /api/external/v2/pro_account/mandate_migrations`](endpoints/getProAccountMandateMigrations.md) — List mandate migration candidates · `GetProAccountMandateMigrationsWithResponse`
- [`GET /api/external/v2/pro_account/mandates`](endpoints/getProAccountMandates.md) — List Pro Account payment mandates · `GetProAccountMandatesWithResponse`
- [`GET /api/external/v2/sepa_mandates/{id}`](endpoints/getSepaMandate.md) — Get a SEPA mandate · `GetSepaMandateWithResponse`
- [`GET /api/external/v2/sepa_mandates`](endpoints/getSepaMandates.md) — List SEPA mandates · `GetSepaMandatesWithResponse`
- [`POST /api/external/v2/gocardless_mandates/mail_requests`](endpoints/postGocardlessMandateMailRequests.md) — Send a GoCardless mandate email request · `PostGocardlessMandateMailRequestsWithResponse`
- [`POST /api/external/v2/gocardless_mandates/{gocardless_mandate_id}/associations`](endpoints/postGocardlessMandateAssociations.md) — Associate a GoCardless mandate to a customer · `PostGocardlessMandateAssociationsWithResponse`
- [`POST /api/external/v2/gocardless_mandates/{gocardless_mandate_id}/cancellations`](endpoints/postGocardlessMandateCancellations.md) — Cancel a Gocardless mandate · `PostGocardlessMandateCancellationsWithResponse`
- [`POST /api/external/v2/pro_account/mandate_migrations`](endpoints/postProAccountMandateMigrations.md) — Migrate a mandate to Pro Account · `PostProAccountMandateMigrationsWithResponse`
- [`POST /api/external/v2/pro_account/mandate_requests`](endpoints/postProAccountMandateMailRequests.md) — Send a Pro Account SEPA mandate request · `PostProAccountMandateMailRequestsWithResponse`
- [`POST /api/external/v2/sepa_mandates`](endpoints/postSepaMandates.md) — Create a SEPA mandate · `PostSepaMandatesWithResponse`
- [`PUT /api/external/v2/sepa_mandates/{id}`](endpoints/putSepaMandate.md) — Update a SEPA mandate · `PutSepaMandateWithResponse`

## PA Registrations

- [`GET /api/external/v2/pa_registrations`](endpoints/getPaRegistrations.md) — List PA Registrations · `GetPaRegistrationsWithResponse`

## Products

- [`GET /api/external/v2/products/{id}`](endpoints/getProduct.md) — Retrieve a product · `GetProductWithResponse`
- [`GET /api/external/v2/products`](endpoints/getProducts.md) — List products · `GetProductsWithResponse`
- [`POST /api/external/v2/products`](endpoints/postProducts.md) — Create a product · `PostProductsWithResponse`
- [`PUT /api/external/v2/products/{id}`](endpoints/putProduct.md) — Update a product · `PutProductWithResponse`

## Purchase Requests

- [`GET /api/external/v2/purchase_requests/{id}`](endpoints/getPurchaseRequest.md) — Retrieve a purchase request · `GetPurchaseRequestWithResponse`
- [`GET /api/external/v2/purchase_requests`](endpoints/getPurchaseRequests.md) — List purchase requests · `GetPurchaseRequestsWithResponse`
- [`POST /api/external/v2/purchase_requests/imports`](endpoints/createPurchaseRequestImport.md) — Import a purchase order. · `CreatePurchaseRequestImportWithResponse`

## Quotes

- [`GET /api/external/v2/changelogs/quotes`](endpoints/getQuoteChanges.md) — Get quotes changes events · `GetQuoteChangesWithResponse`
- [`GET /api/external/v2/quotes/{id}`](endpoints/getQuote.md) — Retrieve a quote · `GetQuoteWithResponse`
- [`GET /api/external/v2/quotes/{quote_id}/appendices`](endpoints/getQuoteAppendices.md) — List appendices of a quote · `GetQuoteAppendicesWithResponse`
- [`GET /api/external/v2/quotes/{quote_id}/invoice_line_sections`](endpoints/getQuoteInvoiceLineSections.md) — List invoice line sections for a quote · `GetQuoteInvoiceLineSectionsWithResponse`
- [`GET /api/external/v2/quotes/{quote_id}/invoice_lines`](endpoints/getQuoteInvoiceLines.md) — List invoice lines for a quote · `GetQuoteInvoiceLinesWithResponse`
- [`GET /api/external/v2/quotes`](endpoints/listQuotes.md) — List quotes · `ListQuotesWithResponse`
- [`POST /api/external/v2/quotes/{id}/send_by_email`](endpoints/sendByEmailQuote.md) — Send a quote by email · `SendByEmailQuoteWithResponse`
- [`POST /api/external/v2/quotes/{quote_id}/appendices`](endpoints/postQuoteAppendices.md) — Upload an appendix for a quote · `PostQuoteAppendicesWithResponse`
- [`POST /api/external/v2/quotes`](endpoints/postQuotes.md) — Create a quote · `PostQuotesWithResponse`
- [`PUT /api/external/v2/quotes/{id}/update_status`](endpoints/updateStatusQuote.md) — Update status of a quote · `UpdateStatusQuoteWithResponse`
- [`PUT /api/external/v2/quotes/{id}`](endpoints/updateQuote.md) — Update a quote · `UpdateQuoteWithResponse`

## Supplier Invoices

- [`GET /api/external/v2/supplier_invoices/{id}`](endpoints/getSupplierInvoice.md) — Retrieve a supplier invoice · `GetSupplierInvoiceWithResponse`
- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/categories`](endpoints/getSupplierInvoiceCategories.md) — List categories of a supplier invoice · `GetSupplierInvoiceCategoriesWithResponse`
- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/invoice_lines`](endpoints/getSupplierInvoiceLines.md) — List invoice lines for a supplier invoice · `GetSupplierInvoiceLinesWithResponse`
- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/matched_transactions`](endpoints/getSupplierInvoiceMatchedTransactions.md) — List matched transactions for a supplier invoice · `GetSupplierInvoiceMatchedTransactionsWithResponse`
- [`GET /api/external/v2/supplier_invoices/{supplier_invoice_id}/payments`](endpoints/getSupplierInvoicePayments.md) — List payments for a supplier invoice · `GetSupplierInvoicePaymentsWithResponse`
- [`GET /api/external/v2/supplier_invoices`](endpoints/getSupplierInvoices.md) — List supplier invoices · `GetSupplierInvoicesWithResponse`
- [`POST /api/external/v2/supplier_invoices/e_invoices/imports`](endpoints/createSupplierInvoiceEInvoiceImport.md) — Import a supplier e-invoice · `CreateSupplierInvoiceEInvoiceImportWithResponse`
- [`POST /api/external/v2/supplier_invoices/import`](endpoints/importSupplierInvoice.md) — Import a supplier invoice with a file attached · `ImportSupplierInvoiceWithResponse`
- [`POST /api/external/v2/supplier_invoices/{supplier_invoice_id}/linked_purchase_requests`](endpoints/postSupplierInvoiceLinkedPurchaseRequests.md) — Link a purchase request to a supplier invoice · `PostSupplierInvoiceLinkedPurchaseRequestsWithResponse`
- [`PUT /api/external/v2/supplier_invoices/{id}/validate_accounting`](endpoints/ValidateAccountingSupplierInvoice.md) — Validate the accounting of a supplier invoice · `ValidateAccountingSupplierInvoiceWithResponse`
- [`PUT /api/external/v2/supplier_invoices/{id}`](endpoints/putSupplierInvoice.md) — Update a supplier invoice · `PutSupplierInvoiceWithResponse`
- [`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/categories`](endpoints/putSupplierInvoiceCategories.md) — Categorize a supplier invoice · `PutSupplierInvoiceCategoriesWithResponse`
- [`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/e_invoice_status`](endpoints/putSupplierInvoiceEInvoiceStatus.md) — Update e-invoice status for a supplier invoice · `PutSupplierInvoiceEInvoiceStatusWithResponse`
- [`PUT /api/external/v2/supplier_invoices/{supplier_invoice_id}/payment_status`](endpoints/updateSupplierInvoicePaymentStatus.md) — Update a supplier invoice payment status · `UpdateSupplierInvoicePaymentStatusWithResponse`

## Suppliers

- [`GET /api/external/v2/suppliers/{id}`](endpoints/getSupplier.md) — Retrieve a supplier · `GetSupplierWithResponse`
- [`GET /api/external/v2/suppliers/{supplier_id}/categories`](endpoints/getSupplierCategories.md) — List categories of a supplier · `GetSupplierCategoriesWithResponse`
- [`GET /api/external/v2/suppliers`](endpoints/getSuppliers.md) — List suppliers · `GetSuppliersWithResponse`
- [`POST /api/external/v2/suppliers`](endpoints/postSupplier.md) — Create a Supplier · `PostSupplierWithResponse`
- [`PUT /api/external/v2/suppliers/{id}`](endpoints/putSupplier.md) — Update a supplier · `PutSupplierWithResponse`
- [`PUT /api/external/v2/suppliers/{supplier_id}/categories`](endpoints/putSupplierCategories.md) — Categorize a supplier · `PutSupplierCategoriesWithResponse`

## Transactions

- [`DELETE /api/external/v2/customer_invoices/{customer_invoice_id}/matched_transactions/{id}`](endpoints/deleteCustomerInvoiceMatchedTransactions.md) — Unmatch a transaction to a customer invoice · `DeleteCustomerInvoiceMatchedTransactionsWithResponse`
- [`DELETE /api/external/v2/supplier_invoices/{supplier_invoice_id}/matched_transactions/{id}`](endpoints/deleteSupplierInvoiceMatchedTransactions.md) — Unmatch a transaction to a supplier invoice · `DeleteSupplierInvoiceMatchedTransactionsWithResponse`
- [`GET /api/external/v2/transactions/{id}`](endpoints/getTransaction.md) — Retrieve a transaction · `GetTransactionWithResponse`
- [`GET /api/external/v2/transactions/{transaction_id}/categories`](endpoints/getTransactionCategories.md) — List categories of a bank transaction · `GetTransactionCategoriesWithResponse`
- [`GET /api/external/v2/transactions/{transaction_id}/matched_invoices`](endpoints/getTransactionMatchedInvoices.md) — List invoices matched to a bank transaction · `GetTransactionMatchedInvoicesWithResponse`
- [`GET /api/external/v2/transactions`](endpoints/getTransactions.md) — List transactions · `GetTransactionsWithResponse`
- [`POST /api/external/v2/customer_invoices/{customer_invoice_id}/matched_transactions`](endpoints/postCustomerInvoiceMatchedTransactions.md) — Match a transaction to a customer invoice · `PostCustomerInvoiceMatchedTransactionsWithResponse`
- [`POST /api/external/v2/supplier_invoices/{supplier_invoice_id}/matched_transactions`](endpoints/postSupplierInvoiceMatchedTransactions.md) — Match a transaction to a supplier invoice · `PostSupplierInvoiceMatchedTransactionsWithResponse`
- [`POST /api/external/v2/transactions`](endpoints/createTransaction.md) — Create a transaction · `CreateTransactionWithResponse`
- [`PUT /api/external/v2/transactions/{id}`](endpoints/updateTransaction.md) — Update a transaction · `UpdateTransactionWithResponse`
- [`PUT /api/external/v2/transactions/{transaction_id}/categories`](endpoints/putTransactionCategories.md) — Categorize a bank transaction · `PutTransactionCategoriesWithResponse`

## Trial balance

- [`GET /api/external/v2/trial_balance`](endpoints/getTrialBalance.md) — Get the trial balance · `GetTrialBalanceWithResponse`

## Users

- [`GET /api/external/v2/me`](endpoints/getMe.md) — User Profile · `GetMeWithResponse`

## Webhooks

- [`DELETE /api/external/v2/webhook_subscriptions/{id}`](endpoints/deleteWebhookSubscription.md) — Delete a webhook subscription · `DeleteWebhookSubscriptionWithResponse`
- [`GET /api/external/v2/webhook_subscriptions/{id}`](endpoints/getWebhookSubscription.md) — Get a webhook subscription · `GetWebhookSubscriptionWithResponse`
- [`GET /api/external/v2/webhook_subscriptions`](endpoints/getWebhookSubscriptions.md) — List webhook subscriptions · `GetWebhookSubscriptionsWithResponse`
- [`POST /api/external/v2/webhook_subscriptions`](endpoints/postWebhookSubscriptions.md) — Create a webhook subscription · `PostWebhookSubscriptionsWithResponse`
- [`PUT /api/external/v2/webhook_subscriptions/{id}`](endpoints/putWebhookSubscription.md) — Update a webhook subscription · `PutWebhookSubscriptionWithResponse`
