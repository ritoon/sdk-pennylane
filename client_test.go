package pennylane

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestConfigRedactsToken(t *testing.T) {
	const token = "test-private-token"
	var cfg Config
	if err := json.Unmarshal([]byte(`{"Token":"test-private-token","CompanyID":42}`), &cfg); err != nil {
		t.Fatal(err)
	}
	if cfg.Token != token {
		t.Fatal("JSON input did not populate the token")
	}
	for _, format := range []string{"%v", "%+v", "%#v", "%s", "%q"} {
		for _, value := range []any{cfg, &cfg, struct{ Config Config }{cfg}} {
			output := fmt.Sprintf(format, value)
			if strings.Contains(output, token) || !strings.Contains(output, "[REDACTED]") {
				t.Errorf("token was not redacted with %s", format)
			}
		}
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}
	if _, present := fields["Token"]; present || string(fields["CompanyID"]) != "42" {
		t.Fatal("JSON should omit the token and preserve the company ID")
	}
	var logs bytes.Buffer
	for _, handler := range []slog.Handler{slog.NewTextHandler(&logs, nil), slog.NewJSONHandler(&logs, nil)} {
		slog.New(handler).Info("Pennylane", "config", cfg)
	}
	if strings.Contains(logs.String(), token) {
		t.Fatal("structured logs contain the token")
	}
}

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

func TestNewRejectsInvalidConfig(t *testing.T) {
	cases := map[string]Config{
		"blank":            {Token: "   "},
		"embedded space":   {Token: "first second"},
		"embedded newline": {Token: "first\nsecond"},
		"negative company": {Token: "test-token", CompanyID: -1},
	}
	for name, cfg := range cases {
		t.Run(name, func(t *testing.T) {
			if c, err := New(cfg); err == nil || c != nil {
				t.Fatal("expected config rejection")
			}
		})
	}
}

func TestCompanyTokenAndTransportError(t *testing.T) {
	failure := errors.New("transport unavailable")
	c, err := New(Config{Token: "  test-token  ", HTTPClient: &http.Client{Transport: transportFunc(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") != "Bearer test-token" || r.Header.Get("X-Company-Id") != "" {
			t.Fatal("unexpected company token headers")
		}
		return nil, failure
	})}})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Ping(context.Background()); !errors.Is(err, failure) {
		t.Fatalf("expected transport error, got %v", err)
	}
}

func TestInvalidMultipartFilenameStopsBeforeRequest(t *testing.T) {
	c := mockClient(t, 200, "{}", func(*http.Request) { t.Fatal("unexpected HTTP request") })
	for _, name := range []string{"", "invoice\r\n.pdf"} {
		var body PostFileAttachmentsMultipartRequestBody
		body.File.InitFromBytes([]byte("data"), name)
		if _, err := c.PostFileAttachmentsWithResponse(context.Background(), nil, body); err == nil {
			t.Fatal("expected filename error")
		}
	}
}
