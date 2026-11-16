# Pages officielles Pennylane

Répertoire des guides, fiches de référence et changements du portail Company API consulté le 16 novembre 2026. Les liens ouvrent le Markdown publié par Pennylane. [Index source](https://pennylane.readme.io/llms.txt)


## Guides: getting started

- [API Overview](https://pennylane.readme.io/docs/api-overview.md)
- [Introduction to the Pennylane API](https://pennylane.readme.io/docs/getting-started.md)
- [Choose the Right API for Your Integration](https://pennylane.readme.io/docs/what-apis-are-available.md)
- [Data Sharing or API: Which Should You Use?](https://pennylane.readme.io/docs/should-i-use-data-sharing-or-the-api-1.md)
- Authentication
  - [Create a Company API Token](https://pennylane.readme.io/docs/generating-my-api-token.md)
  - [Create a Firm API Token](https://pennylane.readme.io/docs/firm-api.md)
  - [Implement OAuth 2.0](https://pennylane.readme.io/docs/oauth-20-walkthrough.md)
  - [Understand Scopes](https://pennylane.readme.io/docs/v2-scopes.md)
- Setup & Configuration
  - [Error Handling & Status Codes](https://pennylane.readme.io/docs/error-handling-status-codes.md)
- [Webhooks](https://pennylane.readme.io/docs/webhooks.md)
  - [Quickstart](https://pennylane.readme.io/docs/quickstart.md)
  - [List of webhook events](https://pennylane.readme.io/docs/list-of-events.md)
  - [Subscribing and Managing Subscriptions ](https://pennylane.readme.io/docs/subscribing-and-managing-subscriptions-new.md)
  - [Subscriber types & company scoping](https://pennylane.readme.io/docs/authorization-new.md)
  - [Receiving & verifying](https://pennylane.readme.io/docs/authorization.md)
  - [Failures & Lifecycle](https://pennylane.readme.io/docs/failures-lifecycle.md)
  - [Rules & limits](https://pennylane.readme.io/docs/rules-limits.md)

## Guides: tutorials

- Use Cases
  - [Supported Use Cases](https://pennylane.readme.io/docs/supported-use-cases.md)
  - Point of Sale Integration
    - [Connect Your Point of Sale (POS) to Pennylane](https://pennylane.readme.io/docs/how-to-build-a-pos-integration.md)
  - Accounting Management
    - [Accounting Reporting 📊](https://pennylane.readme.io/docs/accounting-reporting.md)
    - [Create Ledger Entries via API](https://pennylane.readme.io/docs/create-a-ledger-entry.md)
  - [Invoicing Management](https://pennylane.readme.io/docs/invoicing-use-cases.md)
    - [Import a Supplier Invoice via API](https://pennylane.readme.io/docs/supplier-invoicing.md)
    - [Import a Customer Invoice via API](https://pennylane.readme.io/docs/customer-invoicing.md)
    - [Create a Customer Invoice via API](https://pennylane.readme.io/docs/create-a-customer-invoice-use-case.md)
    - [Manage a Customer Invoice Payment Schedule via the Company API](https://pennylane.readme.io/docs/manage-a-customer-invoice-payment-schedule-via-the-company-api.md)
- Feature Guides
  - [Filter API Data](https://pennylane.readme.io/docs/setting-up-filters.md)
    - [Filter Customer Invoices](https://pennylane.readme.io/docs/filter-my-customer-invoices.md)
    - [Filter Supplier Invoices](https://pennylane.readme.io/docs/filter-my-supplier-invoices.md)
    - [Filter Draft Invoices](https://pennylane.readme.io/docs/filter-my-draft-invoices.md)
    - [Filter Credit Notes](https://pennylane.readme.io/docs/filter-my-credit-notes.md)
    - [Filter Customers](https://pennylane.readme.io/docs/filter-my-customers.md)
    - [Filter Suppliers](https://pennylane.readme.io/docs/filter-my-suppliers.md)
    - [Filter Ledger Accounts](https://pennylane.readme.io/docs/filter-my-ledger-accounts.md)
    - [Filter Ledger Entries](https://pennylane.readme.io/docs/filter-my-ledger-entries.md)
    - [Filter Ledger Entry Lines](https://pennylane.readme.io/docs/filter-my-ledger-entry-lines.md)
    - [Filter Categories](https://pennylane.readme.io/docs/filter-my-categories.md)
    - [Filter Products](https://pennylane.readme.io/docs/filter-my-products.md)
  - [Automate Payment Matching](https://pennylane.readme.io/docs/automating-payment-matching.md)
  - [Use Cursor-Based Pagination](https://pennylane.readme.io/docs/using-cursor-based-pagination.md)
  - [Track Data Changes with the API](https://pennylane.readme.io/docs/tracking-data-changes-with-pennylane-api.md)
  - [Create a Downpayment Invoice](https://pennylane.readme.io/docs/create-a-downpayment-invoice-facture-dacompte.md)
  - [Send Documents by Email](https://pennylane.readme.io/docs/sending-documents-by-email.md)

## Guides: CONCEPTS & BEST PRACTICES

- Concepts
  - [String Character Limits](https://pennylane.readme.io/docs/string-characters-limit.md)
- Best Practices
  - [Payment vs Matched Transaction](https://pennylane.readme.io/docs/what-is-the-difference-between-a-payment-and-a-matched-transaction.md)
  - [Import vs Create a Customer Invoice](https://pennylane.readme.io/docs/import-vs-create-a-customer-invoice-whats-the-difference.md)
  - [Understanding VAT Rates on Ledger Accounts](https://pennylane.readme.io/docs/handle-vat-rates-by-ledger-accounts.md)
    - [How to use the customer update endpoints](https://pennylane.readme.io/docs/how-to-update-individual-and-company-customers.md)

## Guides: ADVANCED GUIDES

- [Rate Limiting in API v2](https://pennylane.readme.io/docs/rate-limiting-1.md)

## Guides: MIGRATION & RELEASE NOTES

- [Migrate from API v1 to v2](https://pennylane.readme.io/docs/api-v2-vs-v1.md)
- [API Changelog](https://pennylane.readme.io/docs/stay-updated-with-our-api-changes.md)
- [Public Roadmap](https://pennylane.readme.io/docs/api-public-roadmap.md)
- [2026 API changes Migration Guide](https://pennylane.readme.io/docs/2026-api-changes-guide.md)

## Guides: RESOURCES

- [Glossary](https://pennylane.readme.io/docs/glossary.md)

## Guides: Need support?

- [How to reach out to us](https://pennylane.readme.io/docs/how-to-reach-out-to-us.md)

## API Reference: Accounting

- Journals
  - [List journals](https://pennylane.readme.io/reference/getjournals.md)
  - [Retrieve a journal](https://pennylane.readme.io/reference/getjournal.md)
  - [Create a journal](https://pennylane.readme.io/reference/postjournals.md)
- Ledger Accounts
  - [List Ledger Accounts](https://pennylane.readme.io/reference/getledgeraccounts.md)
  - [Get a ledger account](https://pennylane.readme.io/reference/getledgeraccount.md)
  - [Create a ledger account](https://pennylane.readme.io/reference/postledgeraccounts.md)
  - [Update a ledger account](https://pennylane.readme.io/reference/updateledgeraccount.md)
- Ledger Attachments
  - [Upload a file](https://pennylane.readme.io/reference/postledgerattachments.md)
- Ledger Entries
  - [List Ledger Entries](https://pennylane.readme.io/reference/getledgerentries.md)
  - [List ledger entry lines of a Ledger Entry](https://pennylane.readme.io/reference/getledgerentriesledgerentrylines.md)
  - [Retrieve a Ledger entry](https://pennylane.readme.io/reference/getledgerentry.md)
  - [Create a ledger entry](https://pennylane.readme.io/reference/postledgerentries.md)
  - [Update a ledger entry](https://pennylane.readme.io/reference/putledgerentries.md)
  - [List DMS files of a Ledger Entry](https://pennylane.readme.io/reference/getledgerentriesdmsfiles.md)
- Ledger Entry Lines
  - [List ledger entry lines](https://pennylane.readme.io/reference/getledgerentrylines.md)
  - [Retrieve a Ledger entry line](https://pennylane.readme.io/reference/getledgerentryline.md)
  - [List ledger entry lines lettered to a given ledger entry line](https://pennylane.readme.io/reference/getledgerentrylinesletteredledgerentrylines.md)
  - [List categories of a Ledger Entry line](https://pennylane.readme.io/reference/getledgerentrylinescategories.md)
  - [Link Analytical Categories to a Ledger Entry line](https://pennylane.readme.io/reference/putledgerentrylinescategories.md)
  - [Letter ledger entry lines](https://pennylane.readme.io/reference/postledgerentrylinesletter.md)
  - [Unletter ledger entry lines](https://pennylane.readme.io/reference/deleteledgerentrylinesunletter.md)
- Trial balance
  - [Get the trial balance](https://pennylane.readme.io/reference/gettrialbalance.md)
- Fiscal years
  - [List Company's Fiscal Years](https://pennylane.readme.io/reference/company-fiscal-years.md)

## API Reference: Accounting exports

- Exports
  - [Create an Analytical General Ledger export](https://pennylane.readme.io/reference/exportanalyticalgeneralledger.md)
  - [Retrieve an Analytical General Ledger export](https://pennylane.readme.io/reference/getanalyticalgeneralledgerexport.md)
  - [Create a General Ledger export](https://pennylane.readme.io/reference/exportgeneralledger.md)
  - [Retrieve a General Ledger export](https://pennylane.readme.io/reference/getgeneralledgerexport.md)
  - [Create a FEC export](https://pennylane.readme.io/reference/exportfec.md)
  - [Retrieve a FEC export](https://pennylane.readme.io/reference/getfecexport.md)

## API Reference: Analytics

- Category Groups
  - [List category groups](https://pennylane.readme.io/reference/getcategorygroups.md)
  - [Retrieve a category group](https://pennylane.readme.io/reference/getcategorygroup.md)
- Categories
  - [List categories of a category group](https://pennylane.readme.io/reference/getcategorygroupcategories.md)
  - [List categories](https://pennylane.readme.io/reference/getcategories.md)
  - [Retrieve a category](https://pennylane.readme.io/reference/getcategory.md)
  - [Create a category](https://pennylane.readme.io/reference/postcategories.md)
  - [Update a category](https://pennylane.readme.io/reference/updatecategory.md)

## API Reference: Billing Subscriptions

- Billing Subscriptions
  - [List billing subscriptions](https://pennylane.readme.io/reference/getbillingsubscriptions.md)
  - [Get a billing subscription](https://pennylane.readme.io/reference/getbillingsubscription.md)
  - [Create a billing subscription](https://pennylane.readme.io/reference/postbillingsubscriptions.md)
  - [Update a billing subscription](https://pennylane.readme.io/reference/putbillingsubscriptions.md)
  - [List invoice lines for a billing subscription](https://pennylane.readme.io/reference/getbillingsubscriptioninvoicelines.md)
  - [List the invoice line sections of a billing subscription](https://pennylane.readme.io/reference/getbillingsubscriptioninvoicelinesections.md)

## API Reference: Changelogs

- Changelogs
  - [Get customer invoices changes events](https://pennylane.readme.io/reference/getcustomerinvoiceschanges.md)
  - [Get supplier invoices changes events](https://pennylane.readme.io/reference/getsupplierinvoiceschanges.md)
  - [Get customer changes events](https://pennylane.readme.io/reference/getcustomerchanges.md)
  - [Get supplier changes events](https://pennylane.readme.io/reference/getsupplierchanges.md)
  - [Get product change events](https://pennylane.readme.io/reference/getproductchanges.md)
  - [Get ledger entry line change events](https://pennylane.readme.io/reference/getledgerentrylinechanges.md)
  - [Get transaction change events](https://pennylane.readme.io/reference/gettransactionchanges.md)
  - [Get quotes changes events](https://pennylane.readme.io/reference/getquotechanges.md)
  - [Get ledger entry change events](https://pennylane.readme.io/reference/getledgerentrieschanges.md)
  - [Get ledger entry line category change events](https://pennylane.readme.io/reference/getledgerentrylinescategorychanges.md)
  - [Get ledger entry category change events](https://pennylane.readme.io/reference/getledgerentriescategorychanges.md)

## API Reference: COMMERCIAL DOCUMENTS

- Commercial Documents
  - [List commercial documents](https://pennylane.readme.io/reference/listcommercialdocuments.md)
  - [Update a commercial document](https://pennylane.readme.io/reference/updatecommercialdocument.md)
  - [Retrieve a commercial document](https://pennylane.readme.io/reference/getcommercialdocument.md)
  - [List invoice line sections for a commercial document](https://pennylane.readme.io/reference/getcommercialdocumentinvoicelinesections.md)
  - [List invoice lines for a commercial document](https://pennylane.readme.io/reference/getcommercialdocumentinvoicelines.md)
  - [List appendices of a commercial document](https://pennylane.readme.io/reference/getcommercialdocumentappendices.md)
  - [Create a commercial document](https://pennylane.readme.io/reference/postcommercialdocuments.md)
  - [Upload an appendix for a commercial document](https://pennylane.readme.io/reference/postcommercialdocumentappendices.md)
  - [Send a commercial document by email](https://pennylane.readme.io/reference/sendbyemailcommercialdocument.md)
- Category Groups
  - [Create a category group](https://pennylane.readme.io/reference/postcategorygroups.md)
  - [Update a category group](https://pennylane.readme.io/reference/putcategorygroup.md)

## API Reference: COMPANY

- Company
  - [Company features](https://pennylane.readme.io/reference/getcompanyfeatures.md)

## API Reference: Customer Invoices

- Customer Invoices
  - [List customer invoices](https://pennylane.readme.io/reference/getcustomerinvoices.md)
  - [List invoice line sections for a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoiceinvoicelinesections.md)
  - [List invoice lines for a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoiceinvoicelines.md)
  - [List payments for a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoicepayments.md)
  - [List matched transactions for a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoicematchedtransactions.md)
  - [List appendices of a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoiceappendices.md)
  - [List categories of a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoicecategories.md)
  - [Retrieve a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoice.md)
  - [Upload an appendix for a customer invoice](https://pennylane.readme.io/reference/postcustomerinvoiceappendices.md)
  - [Send a customer invoice by email](https://pennylane.readme.io/reference/sendbyemailcustomerinvoice.md)
  - [Create a customer invoice](https://pennylane.readme.io/reference/postcustomerinvoices.md)
  - [Import an invoice with file attached](https://pennylane.readme.io/reference/importcustomerinvoices.md)
  - [Link a credit note to a customer invoice](https://pennylane.readme.io/reference/linkcreditnote.md)
  - [Create a customer invoice from a quote](https://pennylane.readme.io/reference/createcustomerinvoicefromquote.md)
  - [Import a customer e-invoice](https://pennylane.readme.io/reference/createcustomerinvoiceeinvoiceimport.md)
  - [Send a customer e-invoice to PA](https://pennylane.readme.io/reference/sendtopacustomerinvoice.md)
  - [Update a customer invoice](https://pennylane.readme.io/reference/updatecustomerinvoice.md)
  - [Mark a customer invoice as paid](https://pennylane.readme.io/reference/markaspaidcustomerinvoice.md)
  - [Categorize a customer invoice](https://pennylane.readme.io/reference/putcustomerinvoicecategories.md)
  - [Update an Imported customer invoice](https://pennylane.readme.io/reference/updateimportedcustomerinvoice.md)
  - [Turn the draft invoice into a finalized invoice.](https://pennylane.readme.io/reference/finalizecustomerinvoice.md)
  - [Delete draft invoice](https://pennylane.readme.io/reference/deletecustomerinvoices.md)
  - [List custom header fields for a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoicecustomheaderfields.md)
- Products
  - [List products](https://pennylane.readme.io/reference/getproducts.md)
  - [Retrieve a product](https://pennylane.readme.io/reference/getproduct.md)
  - [Create a product](https://pennylane.readme.io/reference/postproducts-1.md)
  - [Update a product](https://pennylane.readme.io/reference/putproduct.md)
- Customer Invoice Templates
  - [List customer invoice templates](https://pennylane.readme.io/reference/getcustomerinvoicetemplates.md)
- Customers
  - [Delete a contact of a customer](https://pennylane.readme.io/reference/deletecustomercontact.md)
  - [Update a contact of a customer](https://pennylane.readme.io/reference/putcustomercontact.md)
  - [Create a contact of a customer](https://pennylane.readme.io/reference/postcustomercontact.md)
  - [Retrieve a contact of a customer](https://pennylane.readme.io/reference/getcustomercontact.md)
  - [List contacts of a customer](https://pennylane.readme.io/reference/getcustomercontacts.md)
  - [Retrieve a company customer](https://pennylane.readme.io/reference/getcompanycustomer.md)
  - [Retrieve an individual customer](https://pennylane.readme.io/reference/getindividualcustomer.md)
  - [List customers (company and individual)](https://pennylane.readme.io/reference/getcustomers.md)
  - [Retrieve a customer](https://pennylane.readme.io/reference/getcustomer.md)
  - [List eligible electronic billing addresses for a customer](https://pennylane.readme.io/reference/getcustomerelectronicbillingaddresses.md)
  - [Create a customer (company or individual)](https://pennylane.readme.io/reference/postcustomer.md)
  - [Create a company customer](https://pennylane.readme.io/reference/postcompanycustomer.md)
  - [Create an individual customer](https://pennylane.readme.io/reference/postindividualcustomer.md)
  - [Update a customer (company or individual)](https://pennylane.readme.io/reference/putcustomer.md)
  - [Update a company customer](https://pennylane.readme.io/reference/putcompanycustomer.md)
  - [Update an individual customer](https://pennylane.readme.io/reference/putindividualcustomer.md)
  - [List categories of a customer](https://pennylane.readme.io/reference/getcustomercategories.md)
  - [Categorize a customer](https://pennylane.readme.io/reference/putcustomercategories.md)

## API Reference: File upload

- File Attachments
  - [Upload a file](https://pennylane.readme.io/reference/postfileattachments.md)

## API Reference: Mandates

- Mandates
  - [Create a SEPA mandate](https://pennylane.readme.io/reference/postsepamandates.md)
  - [List SEPA mandates](https://pennylane.readme.io/reference/getsepamandates.md)
  - [Get a SEPA mandate](https://pennylane.readme.io/reference/getsepamandate.md)
  - [List mandate migration candidates](https://pennylane.readme.io/reference/getproaccountmandatemigrations.md)
    - [Send a Pro Account SEPA mandate request](https://pennylane.readme.io/reference/postproaccountmandatemailrequests.md)
  - [Migrate a mandate to Pro Account](https://pennylane.readme.io/reference/postproaccountmandatemigrations.md)
  - [Update a SEPA mandate](https://pennylane.readme.io/reference/putsepamandate.md)
    - [List Pro Account payment mandates](https://pennylane.readme.io/reference/getproaccountmandates.md)
  - [Delete a SEPA mandate](https://pennylane.readme.io/reference/deletesepamandate.md)
  - [List gocardless mandates](https://pennylane.readme.io/reference/getgocardlessmandates.md)
  - [Get a Gocardless mandate](https://pennylane.readme.io/reference/getgocardlessmandate.md)
  - [Send a GoCardless mandate email request](https://pennylane.readme.io/reference/postgocardlessmandatemailrequests.md)
  - [Associate a GoCardless mandate to a customer](https://pennylane.readme.io/reference/postgocardlessmandateassociations.md)
  - [Cancel a Gocardless mandate](https://pennylane.readme.io/reference/postgocardlessmandatecancellations.md)

## API Reference: QUOTES

- Quotes
  - [List quotes](https://pennylane.readme.io/reference/listquotes.md)
  - [Retrieve a quote](https://pennylane.readme.io/reference/getquote.md)
  - [List invoice line sections for a quote](https://pennylane.readme.io/reference/getquoteinvoicelinesections.md)
  - [List invoice lines for a quote](https://pennylane.readme.io/reference/getquoteinvoicelines.md)
  - [List appendices of a quote](https://pennylane.readme.io/reference/getquoteappendices.md)
  - [Create a quote](https://pennylane.readme.io/reference/postquotes.md)
  - [Upload an appendix for a quote](https://pennylane.readme.io/reference/postquoteappendices.md)
  - [Send a quote by email](https://pennylane.readme.io/reference/sendbyemailquote.md)
  - [Update a quote](https://pennylane.readme.io/reference/updatequote.md)
  - [Update status of a quote](https://pennylane.readme.io/reference/updatestatusquote.md)

## API Reference: Supplier invoices

- Supplier Invoices
  - [List invoice lines for a supplier invoice](https://pennylane.readme.io/reference/getsupplierinvoicelines.md)
  - [List supplier invoices](https://pennylane.readme.io/reference/getsupplierinvoices.md)
  - [Retrieve a supplier invoice](https://pennylane.readme.io/reference/getsupplierinvoice.md)
  - [List categories of a supplier invoice](https://pennylane.readme.io/reference/getsupplierinvoicecategories.md)
  - [List payments for a supplier invoice](https://pennylane.readme.io/reference/getsupplierinvoicepayments.md)
  - [List matched transactions for a supplier invoice](https://pennylane.readme.io/reference/getsupplierinvoicematchedtransactions.md)
  - [Import a supplier invoice with a file attached](https://pennylane.readme.io/reference/importsupplierinvoice.md)
  - [Link a purchase request to a supplier invoice](https://pennylane.readme.io/reference/postsupplierinvoicelinkedpurchaserequests.md)
  - [Import a supplier e-invoice](https://pennylane.readme.io/reference/createsupplierinvoiceeinvoiceimport.md)
  - [Update a supplier invoice](https://pennylane.readme.io/reference/putsupplierinvoice.md)
  - [Categorize a supplier invoice](https://pennylane.readme.io/reference/putsupplierinvoicecategories.md)
  - [Update a supplier invoice payment status](https://pennylane.readme.io/reference/updatesupplierinvoicepaymentstatus.md)
  - [Validate the accounting of a supplier invoice](https://pennylane.readme.io/reference/validateaccountingsupplierinvoice.md)
  - [Update e-invoice status for a supplier invoice](https://pennylane.readme.io/reference/putsupplierinvoiceeinvoicestatus.md)
- Purchase Requests
  - [List purchase requests](https://pennylane.readme.io/reference/getpurchaserequests.md)
  - [Retrieve a purchase request](https://pennylane.readme.io/reference/getpurchaserequest.md)
  - [Import a purchase order.](https://pennylane.readme.io/reference/createpurchaserequestimport.md)
- Suppliers
  - [List suppliers](https://pennylane.readme.io/reference/getsuppliers.md)
  - [Retrieve a supplier](https://pennylane.readme.io/reference/getsupplier.md)
  - [Create a Supplier](https://pennylane.readme.io/reference/postsupplier.md)
  - [Update a supplier](https://pennylane.readme.io/reference/putsupplier.md)
  - [List categories of a supplier](https://pennylane.readme.io/reference/getsuppliercategories.md)
  - [Categorize a supplier](https://pennylane.readme.io/reference/putsuppliercategories.md)

## API Reference: Transactions

- Bank accounts
  - [List bank establishments](https://pennylane.readme.io/reference/getbankestablishments.md)
  - [List bank accounts](https://pennylane.readme.io/reference/getbankaccounts.md)
  - [Retrieve a bank account](https://pennylane.readme.io/reference/getbankaccount.md)
  - [Create a bank account](https://pennylane.readme.io/reference/postbankaccount.md)
- Transactions
  - [List transactions](https://pennylane.readme.io/reference/gettransactions.md)
  - [Retrieve a transaction](https://pennylane.readme.io/reference/gettransaction.md)
  - [List invoices matched to a bank transaction](https://pennylane.readme.io/reference/gettransactionmatchedinvoices.md)
  - [List categories of a bank transaction](https://pennylane.readme.io/reference/gettransactioncategories.md)
  - [Create a transaction](https://pennylane.readme.io/reference/createtransaction.md)
  - [Match a transaction to a customer invoice](https://pennylane.readme.io/reference/postcustomerinvoicematchedtransactions.md)
  - [Match a transaction to a supplier invoice](https://pennylane.readme.io/reference/postsupplierinvoicematchedtransactions.md)
  - [Categorize a bank transaction](https://pennylane.readme.io/reference/puttransactioncategories.md)
  - [Update a transaction](https://pennylane.readme.io/reference/updatetransaction.md)
  - [Unmatch a transaction to a customer invoice](https://pennylane.readme.io/reference/deletecustomerinvoicematchedtransactions.md)
  - [Unmatch a transaction to a supplier invoice](https://pennylane.readme.io/reference/deletesupplierinvoicematchedtransactions.md)

## API Reference: Users

- Users
  - [User Profile](https://pennylane.readme.io/reference/getme.md)

## API Reference: PA Registrations

- PA Registrations
  - [List PA Registrations](https://pennylane.readme.io/reference/getparegistrations.md)

## API Reference: Webhooks

- Webhook Subscriptions
  - [Get a webhook subscription](https://pennylane.readme.io/reference/getwebhooksubscription.md)
  - [Update a webhook subscription](https://pennylane.readme.io/reference/putwebhooksubscription.md)
  - [Delete a webhook subscription](https://pennylane.readme.io/reference/deletewebhooksubscription.md)
  - [List webhook subscriptions](https://pennylane.readme.io/reference/getwebhooksubscriptions.md)
  - [Create a webhook subscription](https://pennylane.readme.io/reference/postwebhooksubscriptions.md)
- Webhook Events
  - [supplier_invoice.e_invoicing_received](https://pennylane.readme.io/reference/supplier_invoicee_invoicing_received.md)
  - [customer_invoice.e_invoicing_status_updated](https://pennylane.readme.io/reference/customer_invoicee_invoicing_status_updated.md)
  - [dms_file.created](https://pennylane.readme.io/reference/dms_filecreated.md)

## API Reference: Company V2

- Customer Invoices
  - [List installments for a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoiceinstallments.md)
  - [Mark a customer invoice installment as paid](https://pennylane.readme.io/reference/markaspaidcustomerinvoiceinstallment.md)
  - [List comments of a customer invoice](https://pennylane.readme.io/reference/getcustomerinvoicecomments.md)
  - [Add a comment to a customer invoice](https://pennylane.readme.io/reference/postcustomerinvoicecomment.md)
- Billing email templates
  - [List billing email templates](https://pennylane.readme.io/reference/getbillingemailtemplates.md)

## Pages

- [API Contract terms](https://pennylane.readme.io/page/api-contract-terms.md)

## Changelog

- [Enabling and disabling ledger accounts is now available in the Company API v2](https://pennylane.readme.io/changelog/enabling-and-disabling-ledger-accounts-is-now-available-in-the-company-api-v2.md)
- [Multi-installment invoices are now available in the Company API v2](https://pennylane.readme.io/changelog/multi-installment-invoices-are-now-available-in-the-company-api-v2.md)
- [New `electronic_billing_address` field on the invoice import](https://pennylane.readme.io/changelog/new-electronic_billing_address-field-on-the-invoice-import.md)
- [Fix: public_file_url on customer invoices is now null until the PDF is ready](https://pennylane.readme.io/changelog/fix-public_file_url-on-customer-invoices-is-now-null-until-the-pdf-is-ready.md)
- [Removal: Reseller API](https://pennylane.readme.io/changelog/removal-reseller-api.md)
- [Inline legal notices when converting imported customer invoices to e-invoices](https://pennylane.readme.io/changelog/inline-legal-notices-when-converting-imported-customer-invoices-to-e-invoices.md)
- [New `contract_reference` field on customer invoices](https://pennylane.readme.io/changelog/new-contract_reference-field-on-customer-invoices.md)
- [New `public_administration` field on Customers endpoints](https://pennylane.readme.io/changelog/new-public_administration-field-on-customers-endpoints.md)
- [New endpoints: create and update a customer, and list electronic billing addresses](https://pennylane.readme.io/changelog/new-endpoints-create-and-update-a-customer-and-list-electronic-billing-addresses.md)
- [Supplier invoices: expose and filter by electronic billing address](https://pennylane.readme.io/changelog/supplier-invoices-expose-and-filter-by-electronic-billing-address.md)
