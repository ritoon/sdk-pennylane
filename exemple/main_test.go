package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	pennylane "github.com/ritoon/sdk-pennylane"
)

type transport func(*http.Request) (*http.Response, error)

func (f transport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestExemples(t *testing.T) {
	calls := 0
	c, err := pennylane.New(pennylane.Config{Token: "test", HTTPClient: &http.Client{Transport: transport(func(r *http.Request) (*http.Response, error) {
		calls++
		status := 200
		body := ""
		switch r.Method + " " + r.URL.Path {
		case "GET /api/external/v2/trial_balance":
			q := r.URL.Query()
			if q.Get("period_start") != "2026-01-01" || q.Get("period_end") != "2026-03-31" {
				t.Fatal(q)
			}
			switch q.Get("cursor") {
			case "":
				body = `{"items":[{"number":"401","debits":"1.20"}],"has_more":true,"next_cursor":"second"}`
			case "second":
				body = `{"items":[{"number":"512","credits":"2.30"}],"has_more":false,"next_cursor":null}`
			default:
				t.Fatal(q)
			}
		case "POST /api/external/v2/customer_invoices/18/send_by_email":
			var v pennylane.SendByEmailCustomerInvoiceJSONRequestBody
			if err := json.NewDecoder(r.Body).Decode(&v); err != nil || v.Recipients == nil || len(*v.Recipients) != 1 || (*v.Recipients)[0] != "test@example.org" {
				t.Fatal("destinataire incorrect")
			}
			status = 204
		case "POST /api/external/v2/customers":
			var v pennylane.PostCustomerJSONBody0
			if err := json.NewDecoder(r.Body).Decode(&v); err != nil || v.CustomerType != "company" || v.Name == "" || v.BillingAddress.City != "Paris" {
				t.Fatal("client incorrect")
			}
			status = 201
			body = `{"id":42,"customer_type":"company"}`
		case "POST /api/external/v2/sepa_mandates":
			var v pennylane.PostSepaMandatesJSONRequestBody
			if err := json.NewDecoder(r.Body).Decode(&v); err != nil || v.CustomerId != 42 || v.Identifier != "RUM-A-REMPLACER" || v.SignedAt.String() != "2026-01-15" {
				t.Fatal("mandat incorrect")
			}
			status = 201
			body = `{"id":71,"customer":{"id":42}}`
		default:
			t.Fatal(r.Method, r.URL.Path)
		}
		return &http.Response{StatusCode: status, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
	})}})
	if err != nil {
		t.Fatal(err)
	}
	exec := func(args []string, file string) string {
		t.Helper()
		var input []byte
		if file != "" {
			input, err = os.ReadFile("donnees/" + file)
			if err != nil {
				t.Fatal(err)
			}
		}
		var out bytes.Buffer
		if err := run(context.Background(), c, args, bytes.NewReader(input), &out); err != nil {
			t.Fatal(err)
		}
		return strings.TrimSpace(out.String())
	}
	lines := exec([]string{"balance", "2026-01-01", "2026-03-31"}, "")
	if !strings.Contains(lines, `"debits":"1.20"`) || !strings.Contains(lines, `"credits":"2.30"`) || len(strings.Split(lines, "\n")) != 2 {
		t.Fatal(lines)
	}
	exec([]string{"envoyer-facture", "18", "test@example.org"}, "")
	customer := exec([]string{"creer-client"}, "client.json")
	if customer != "42" || exec([]string{"creer-mandat", customer}, "mandat.json") != "71" {
		t.Fatal("identifiants inattendus")
	}
	if err := run(context.Background(), c, []string{"creer-mandat", "42"}, strings.NewReader(`{"customer_id":43}`), io.Discard); err == nil {
		t.Fatal("client incohérent accepté")
	}
	if calls != 5 {
		t.Fatalf("%d appels", calls)
	}
}
