package pennylane

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

type transportFunc func(*http.Request) (*http.Response, error)

func (f transportFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func mockClient(t *testing.T, status int, body string, inspect func(*http.Request)) *Client {
	t.Helper()
	c, err := New(Config{Token: "test-token", CompanyID: 42, HTTPClient: &http.Client{Transport: transportFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.Scheme != "https" || r.URL.Host != "app.pennylane.com" || r.Header.Get("Authorization") != "Bearer test-token" {
			t.Fatal("unexpected destination or authorization")
		}
		inspect(r)
		return &http.Response{StatusCode: status, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
	})}})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestPing(t *testing.T) {
	for _, tc := range []struct {
		status int
		body   string
		ok     bool
	}{
		{200, `{"company":{"id":42},"user":null,"scopes":[]}`, true}, {200, "{}", false}, {200, "invalid", false}, {401, "{}", false}, {403, "{}", false},
	} {
		c := mockClient(t, tc.status, tc.body, func(r *http.Request) {
			if r.Method != "GET" || r.URL.Path != "/api/external/v2/me" || r.Header.Get("X-Company-Id") != "42" {
				t.Fatal("unexpected profile request")
			}
		})
		if err := c.Ping(context.Background()); (err == nil) != tc.ok {
			t.Errorf("status %d, body %q: %v", tc.status, tc.body, err)
		}
	}
	if _, err := New(Config{}); err == nil {
		t.Fatal("accepted missing token")
	}
}

func TestJournals(t *testing.T) {
	cursor := "a+b/c"
	company := "123"
	c := mockClient(t, 200, `{"items":[{"id":7,"code":"VE","label":"Ventes","type":"sales"}],"has_more":false,"next_cursor":null}`, func(r *http.Request) {
		if r.URL.Query().Get("cursor") != cursor || r.Header.Get("X-Company-Id") != company {
			t.Fatal("unexpected journal parameters")
		}
	})
	resp, err := c.GetJournalsWithResponse(context.Background(), &GetJournalsParams{Cursor: &cursor, XCompanyId: &company})
	if err != nil || resp.JSON200 == nil || len(resp.JSON200.Items) != 1 || resp.JSON200.Items[0].Id != 7 || !resp.JSON200.NextCursor.IsNull() {
		t.Fatalf("journal decoding: %v", err)
	}
}

func TestWebhookJSON(t *testing.T) {
	enabled := false
	c := mockClient(t, 403, `{"message":"Forbidden"}`, func(r *http.Request) {
		var payload map[string]any
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatal(err)
		}
		if r.Method != "POST" || r.Header.Get("X-Company-Id") != "" || payload["enabled"] != false || payload["callback_url"] != "https://example.org/hook" {
			t.Fatal("unexpected webhook request")
		}
	})
	resp, err := c.PostWebhookSubscriptionsWithResponse(context.Background(), PostWebhookSubscriptionsJSONRequestBody{CallbackUrl: "https://example.org/hook", Enabled: &enabled})
	if err != nil || resp.StatusCode() != 403 {
		t.Fatalf("webhook status: %v", err)
	}
}

func TestMultipartImport(t *testing.T) {
	send := false
	body := CreateCustomerInvoiceEInvoiceImportMultipartRequestBody{SendToPa: &send}
	body.File.InitFromBytes([]byte("%PDF-1.7"), "facture.pdf")
	c := mockClient(t, 400, "{}", func(r *http.Request) {
		if err := r.ParseMultipartForm(4096); err != nil {
			t.Fatal(err)
		}
		defer r.MultipartForm.RemoveAll()
		f, h, err := r.FormFile("file")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		b, err := io.ReadAll(f)
		if err != nil || string(b) != "%PDF-1.7" || h.Filename != "facture.pdf" || h.Header.Get("Content-Type") != "application/pdf" || r.FormValue("send_to_pa") != "false" {
			t.Fatal("incorrect multipart encoding")
		}
		if len(r.MultipartForm.Value) != 1 {
			t.Fatal("unexpected optional fields")
		}
	})
	if _, err := c.CreateCustomerInvoiceEInvoiceImportWithResponse(context.Background(), nil, body); err != nil {
		t.Fatal(err)
	}
}

func TestContractCoverage(t *testing.T) {
	data, err := os.ReadFile("internal/spec/company.json")
	if err != nil {
		t.Fatal(err)
	}
	var spec struct {
		Paths map[string]map[string]struct {
			OperationID string `json:"operationId"`
		}
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	client := reflect.TypeOf((*Client)(nil))
	count := 0
	for path, ops := range spec.Paths {
		for verb, op := range ops {
			name := ""
			for _, part := range strings.Split(op.OperationID, "-") {
				name += strings.ToUpper(part[:1]) + part[1:]
			}
			if _, ok := client.MethodByName(name + "WithResponse"); !ok {
				t.Errorf("%s %s missing typed method %s", verb, path, name)
			}
			count++
		}
	}
	if count != 186 {
		t.Errorf("expected 186 contract operations, got %d", count)
	}
}
