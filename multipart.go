package pennylane

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"path/filepath"
	"strings"

	"github.com/oapi-codegen/runtime/types"
)

func multipartBody(file types.File, fields map[string]any) (string, *bytes.Buffer, error) {
	name := file.Filename()
	if name == "" || strings.ContainsAny(name, "\r\n") {
		return "", nil, fmt.Errorf("invalid multipart filename")
	}
	source, err := file.Reader()
	if err != nil {
		return "", nil, err
	}
	defer source.Close()
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	mimeType := mime.TypeByExtension(filepath.Ext(name))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	part, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": {mime.FormatMediaType("form-data", map[string]string{"name": "file", "filename": name})},
		"Content-Type":        {mimeType},
	})
	if err != nil {
		return "", nil, err
	}
	if _, err = io.Copy(part, source); err != nil {
		return "", nil, err
	}
	for key, value := range fields {
		switch v := value.(type) {
		case *string:
			if v != nil {
				err = writer.WriteField(key, *v)
			}
		case *bool:
			if v != nil {
				err = writer.WriteField(key, fmt.Sprint(*v))
			}
		default:
			b, e := json.Marshal(value)
			if e != nil {
				return "", nil, e
			}
			if string(b) == "null" {
				continue
			}
			part, err = writer.CreatePart(textproto.MIMEHeader{"Content-Disposition": {mime.FormatMediaType("form-data", map[string]string{"name": key})}, "Content-Type": {"application/json"}})
			if err == nil {
				_, err = part.Write(b)
			}
		}
		if err != nil {
			return "", nil, err
		}
	}
	err = writer.Close()
	return writer.FormDataContentType(), buf, err
}

func (c *Client) PostCommercialDocumentAppendicesWithResponse(ctx context.Context, commercialDocumentId int, params *PostCommercialDocumentAppendicesParams, body PostCommercialDocumentAppendicesMultipartRequestBody, editors ...RequestEditorFn) (*PostCommercialDocumentAppendicesResponse, error) {
	contentType, reader, err := multipartBody(body.File, nil)
	if err != nil {
		return nil, err
	}
	return c.PostCommercialDocumentAppendicesWithBodyWithResponse(ctx, commercialDocumentId, params, contentType, reader, editors...)
}

func (c *Client) CreateCustomerInvoiceEInvoiceImportWithResponse(ctx context.Context, params *CreateCustomerInvoiceEInvoiceImportParams, body CreateCustomerInvoiceEInvoiceImportMultipartRequestBody, editors ...RequestEditorFn) (*CreateCustomerInvoiceEInvoiceImportResponse, error) {
	contentType, reader, err := multipartBody(body.File, map[string]any{"installments": body.Installments, "invoice_options": body.InvoiceOptions, "override_invoice_lines": body.OverrideInvoiceLines, "send_to_pa": body.SendToPa})
	if err != nil {
		return nil, err
	}
	return c.CreateCustomerInvoiceEInvoiceImportWithBodyWithResponse(ctx, params, contentType, reader, editors...)
}

func (c *Client) PostCustomerInvoiceAppendicesWithResponse(ctx context.Context, customerInvoiceId int, params *PostCustomerInvoiceAppendicesParams, body PostCustomerInvoiceAppendicesMultipartRequestBody, editors ...RequestEditorFn) (*PostCustomerInvoiceAppendicesResponse, error) {
	contentType, reader, err := multipartBody(body.File, nil)
	if err != nil {
		return nil, err
	}
	return c.PostCustomerInvoiceAppendicesWithBodyWithResponse(ctx, customerInvoiceId, params, contentType, reader, editors...)
}

func (c *Client) PostFileAttachmentsWithResponse(ctx context.Context, params *PostFileAttachmentsParams, body PostFileAttachmentsMultipartRequestBody, editors ...RequestEditorFn) (*PostFileAttachmentsResponse, error) {
	contentType, reader, err := multipartBody(body.File, map[string]any{"filename": body.Filename})
	if err != nil {
		return nil, err
	}
	return c.PostFileAttachmentsWithBodyWithResponse(ctx, params, contentType, reader, editors...)
}

func (c *Client) PostLedgerAttachmentsWithResponse(ctx context.Context, params *PostLedgerAttachmentsParams, body PostLedgerAttachmentsMultipartRequestBody, editors ...RequestEditorFn) (*PostLedgerAttachmentsResponse, error) {
	contentType, reader, err := multipartBody(body.File, map[string]any{"filename": body.Filename})
	if err != nil {
		return nil, err
	}
	return c.PostLedgerAttachmentsWithBodyWithResponse(ctx, params, contentType, reader, editors...)
}

func (c *Client) PostQuoteAppendicesWithResponse(ctx context.Context, quoteId int, params *PostQuoteAppendicesParams, body PostQuoteAppendicesMultipartRequestBody, editors ...RequestEditorFn) (*PostQuoteAppendicesResponse, error) {
	contentType, reader, err := multipartBody(body.File, nil)
	if err != nil {
		return nil, err
	}
	return c.PostQuoteAppendicesWithBodyWithResponse(ctx, quoteId, params, contentType, reader, editors...)
}

func (c *Client) CreateSupplierInvoiceEInvoiceImportWithResponse(ctx context.Context, params *CreateSupplierInvoiceEInvoiceImportParams, body CreateSupplierInvoiceEInvoiceImportMultipartRequestBody, editors ...RequestEditorFn) (*CreateSupplierInvoiceEInvoiceImportResponse, error) {
	contentType, reader, err := multipartBody(body.File, map[string]any{"invoice_options": body.InvoiceOptions, "override_invoice_lines": body.OverrideInvoiceLines})
	if err != nil {
		return nil, err
	}
	return c.CreateSupplierInvoiceEInvoiceImportWithBodyWithResponse(ctx, params, contentType, reader, editors...)
}
