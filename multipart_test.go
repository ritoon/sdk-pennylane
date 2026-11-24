package pennylane

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/oapi-codegen/runtime/types"
)

func TestMultipartOptionalFields(t *testing.T) {
	var file types.File
	file.InitFromBytes([]byte("data"), "facture.unknown-extension")
	filename, send := "renamed.pdf", false
	contentType, body, err := multipartBody(file, map[string]any{
		"filename": &filename, "send_to_pa": &send,
		"invoice_options": map[string]any{"customer_id": 42},
		"installments":    []any{map[string]any{"amount": "10.00"}},
		"absent_string":   (*string)(nil), "absent_bool": (*bool)(nil), "absent_json": nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	r, err := http.NewRequest(http.MethodPost, "https://example.test", body)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Content-Type", contentType)
	if err := r.ParseMultipartForm(4096); err != nil {
		t.Fatal(err)
	}
	defer r.MultipartForm.RemoveAll()
	f, h, err := r.FormFile("file")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil || string(data) != "data" || h.Header.Get("Content-Type") != "application/octet-stream" {
		t.Fatal("incorrect file encoding")
	}
	if len(r.MultipartForm.Value) != 4 || r.FormValue("filename") != filename || r.FormValue("send_to_pa") != "false" {
		t.Fatalf("unexpected fields: %v", r.MultipartForm.Value)
	}
	for key, want := range map[string]string{"invoice_options": `{"customer_id":42}`, "installments": `[{"amount":"10.00"}]`} {
		if r.FormValue(key) != want {
			t.Errorf("%s = %q", key, r.FormValue(key))
		}
	}
	// JSON parts must carry their content type, not just valid JSON text.
	_, body, err = multipartBody(file, map[string]any{"invoice_options": map[string]int{"customer_id": 42}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(body.String(), "Content-Type: application/json") {
		t.Fatal("missing JSON part content type")
	}
	if _, _, err := multipartBody(file, map[string]any{"invalid": make(chan int)}); err == nil {
		t.Fatal("accepted unserializable JSON")
	}
}

func TestMultipartWrappersRejectInvalidFile(t *testing.T) {
	c := mockClient(t, 200, "{}", func(*http.Request) { t.Fatal("invalid upload reached transport") })
	ctx := context.Background()
	tests := map[string]func() error{
		"commercial appendix": func() error {
			_, e := c.PostCommercialDocumentAppendicesWithResponse(ctx, 73, nil, PostCommercialDocumentAppendicesMultipartRequestBody{})
			return e
		},
		"customer import": func() error {
			_, e := c.CreateCustomerInvoiceEInvoiceImportWithResponse(ctx, nil, CreateCustomerInvoiceEInvoiceImportMultipartRequestBody{})
			return e
		},
		"customer appendix": func() error {
			_, e := c.PostCustomerInvoiceAppendicesWithResponse(ctx, 73, nil, PostCustomerInvoiceAppendicesMultipartRequestBody{})
			return e
		},
		"file attachment": func() error {
			_, e := c.PostFileAttachmentsWithResponse(ctx, nil, PostFileAttachmentsMultipartRequestBody{})
			return e
		},
		"ledger attachment": func() error {
			_, e := c.PostLedgerAttachmentsWithResponse(ctx, nil, PostLedgerAttachmentsMultipartRequestBody{})
			return e
		},
		"quote appendix": func() error {
			_, e := c.PostQuoteAppendicesWithResponse(ctx, 73, nil, PostQuoteAppendicesMultipartRequestBody{})
			return e
		},
		"supplier import": func() error {
			_, e := c.CreateSupplierInvoiceEInvoiceImportWithResponse(ctx, nil, CreateSupplierInvoiceEInvoiceImportMultipartRequestBody{})
			return e
		},
	}
	for name, call := range tests {
		t.Run(name, func(t *testing.T) {
			if err := call(); err == nil {
				t.Fatal("accepted missing file")
			}
		})
	}
}

func TestImportJSONOptions(t *testing.T) {
	var body CreateCustomerInvoiceEInvoiceImportMultipartRequestBody
	if err := json.Unmarshal([]byte(`{"invoice_options":{"customer_id":42},"installments":[{"amount":"10.00","deadline":"2026-11-24"}],"send_to_pa":false}`), &body); err != nil {
		t.Fatal(err)
	}
	body.File.InitFromBytes([]byte("pdf"), "invoice.pdf")
	c := mockClient(t, 400, "{}", func(r *http.Request) {
		if err := r.ParseMultipartForm(4096); err != nil {
			t.Fatal(err)
		}
		defer r.MultipartForm.RemoveAll()
		if r.FormValue("invoice_options") != `{"customer_id":42}` || !strings.Contains(r.FormValue("installments"), `"amount":"10.00"`) || r.FormValue("send_to_pa") != "false" {
			t.Fatal("lost import options")
		}
	})
	if _, err := c.CreateCustomerInvoiceEInvoiceImportWithResponse(context.Background(), nil, body); err != nil {
		t.Fatal(err)
	}
}
